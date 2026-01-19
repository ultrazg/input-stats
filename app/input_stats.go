package app

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"unsafe"
)

var (
	user32            = syscall.NewLazyDLL("user32.dll")
	setWindowsHookEx  = user32.NewProc("SetWindowsHookExW")
	callNextHookEx    = user32.NewProc("CallNextHookEx")
	unhookWindowsHook = user32.NewProc("UnhookWindowsHookEx")
	getMessage        = user32.NewProc("GetMessageW")
)

const (
	WH_KEYBOARD_LL = 13
	WH_MOUSE_LL    = 14

	WM_KEYDOWN    = 0x0100
	WM_KEYUP      = 0x0101
	WM_SYSKEYDOWN = 0x0104

	WM_LBUTTONDOWN = 0x0201
	WM_RBUTTONDOWN = 0x0204
	WM_MOUSEMOVE   = 0x0200
	WM_MOUSEWHEEL  = 0x020A

	VK_LSHIFT   = 0xA0
	VK_RSHIFT   = 0xA1
	VK_LCONTROL = 0xA2
	VK_RCONTROL = 0xA3
	VK_LMENU    = 0xA4
	VK_RMENU    = 0xA5
	VK_LWIN     = 0x5B
	VK_RWIN     = 0x5C
)

type KBDLLHOOKSTRUCT struct {
	VkCode uint32
}

type POINT struct {
	X int32
	Y int32
}

type MSLLHOOKSTRUCT struct {
	Pt        POINT
	MouseData uint32
}

var (
	totalKeyCount uint64

	keyStats      = make(map[string]uint64)
	modifierStats = make(map[string]uint64)
	comboStats    = make(map[string]uint64)

	mouseLeftClick  uint64
	mouseRightClick uint64
	mouseWheelDelta int64
	mouseMovePixels float64

	statsLock sync.Mutex

	ctrlDown  bool
	shiftDown bool
	altDown   bool
	keyDown   = make(map[uint32]bool)

	curX, curY           int32
	lastCalcX, lastCalcY int32
	hasSample            bool
)

func isModifier(vk uint32) bool {
	switch vk {
	case VK_LCONTROL, VK_RCONTROL,
		VK_LSHIFT, VK_RSHIFT,
		VK_LMENU, VK_RMENU,
		VK_LWIN, VK_RWIN:
		return true
	}
	return false
}

func keyName(vk uint32) string {
	switch vk {
	case VK_LCONTROL:
		return "Ctrl"
	case VK_RCONTROL:
		return "Ctrl"
	case VK_LSHIFT:
		return "Shift"
	case VK_RSHIFT:
		return "Shift"
	case VK_LMENU:
		return "Alt"
	case VK_RMENU:
		return "Alt"
	case VK_LWIN:
		return "Win"
	case VK_RWIN:
		return "Win"
	case 0x20:
		return "Space"
	case 0x0D:
		return "Enter"
	case 0x09:
		return "Tab"
	case 0x08:
		return "Backspace"
	case 0x1B:
		return "Esc"
	}

	if vk >= 'A' && vk <= 'Z' {
		return string(rune(vk))
	}
	if vk >= '0' && vk <= '9' {
		return string(rune(vk))
	}
	if vk >= 0x70 && vk <= 0x7B {
		return fmt.Sprintf("F%d", vk-0x6F)
	}

	return ""
}

func keyboardProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 {
		k := (*KBDLLHOOKSTRUCT)(unsafe.Pointer(lParam))
		vk := k.VkCode
		name := keyName(vk)

		switch wParam {

		case WM_KEYDOWN, WM_SYSKEYDOWN:
			statsLock.Lock()
			if keyDown[vk] {
				statsLock.Unlock()
				goto NEXT
			}
			keyDown[vk] = true
			statsLock.Unlock()

			atomic.AddUint64(&totalKeyCount, 1)

			if isModifier(vk) {
				if name != "" {
					statsLock.Lock()
					modifierStats[name]++
					statsLock.Unlock()
				}
				switch vk {
				case VK_LCONTROL, VK_RCONTROL:
					ctrlDown = true
				case VK_LSHIFT, VK_RSHIFT:
					shiftDown = true
				case VK_LMENU, VK_RMENU:
					altDown = true
				}
				goto NEXT
			}

			if name != "" {
				statsLock.Lock()
				keyStats[name]++
				statsLock.Unlock()
			}

			if (ctrlDown || shiftDown || altDown) && name != "" {
				combo := ""
				if ctrlDown {
					combo += "Ctrl+"
				}
				if shiftDown {
					combo += "Shift+"
				}
				if altDown {
					combo += "Alt+"
				}
				combo += name

				statsLock.Lock()
				comboStats[combo]++
				statsLock.Unlock()
			}

		case WM_KEYUP:
			statsLock.Lock()
			keyDown[vk] = false
			statsLock.Unlock()

			switch vk {
			case VK_LCONTROL, VK_RCONTROL:
				ctrlDown = false
			case VK_LSHIFT, VK_RSHIFT:
				shiftDown = false
			case VK_LMENU, VK_RMENU:
				altDown = false
			}
		}
	}

NEXT:
	ret, _, _ := callNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
	return ret
}

func mouseProc(nCode int, wParam uintptr, lParam uintptr) uintptr {
	if nCode >= 0 {
		m := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))

		switch wParam {
		case WM_LBUTTONDOWN:
			atomic.AddUint64(&mouseLeftClick, 1)

		case WM_RBUTTONDOWN:
			atomic.AddUint64(&mouseRightClick, 1)

		case WM_MOUSEMOVE:
			curX = m.Pt.X
			curY = m.Pt.Y
			hasSample = true

		case WM_MOUSEWHEEL:
			delta := int16(m.MouseData >> 16)
			if delta < 0 {
				delta = -delta
			}
			atomic.AddInt64(&mouseWheelDelta, int64(delta))
		}
	}

	ret, _, _ := callNextHookEx.Call(0, uintptr(nCode), wParam, lParam)
	return ret
}

func ListenStats() {
	hKey, _, _ := setWindowsHookEx.Call(
		WH_KEYBOARD_LL,
		syscall.NewCallback(keyboardProc),
		0,
		0,
	)
	hMouse, _, _ := setWindowsHookEx.Call(
		WH_MOUSE_LL,
		syscall.NewCallback(mouseProc),
		0,
		0,
	)
	defer unhookWindowsHook.Call(hKey)
	defer unhookWindowsHook.Call(hMouse)

	go func() {
		ticker := time.NewTicker(50 * time.Millisecond)
		for range ticker.C {
			if !hasSample {
				continue
			}
			dx := float64(curX - lastCalcX)
			dy := float64(curY - lastCalcY)
			if dx != 0 || dy != 0 {
				mouseMovePixels += math.Sqrt(dx*dx + dy*dy)
				lastCalcX, lastCalcY = curX, curY
			}
		}
	}()

	// go func() {
	// 	for {
	// 		time.Sleep(5 * time.Second)
	// 		statsLock.Lock()
	// 		fmt.Println("==== Modifier Stats ====")
	// 		for k, v := range modifierStats {
	// 			fmt.Println(k, v)
	// 		}
	// 		fmt.Println("==== Key Stats ====")
	// 		for k, v := range keyStats {
	// 			fmt.Println(k, v)
	// 		}
	// 		fmt.Println("==== Combo Stats ====")
	// 		for k, v := range comboStats {
	// 			fmt.Println(k, v)
	// 		}
	// 		fmt.Println("==== Mouse Stats ====")
	// 		fmt.Println("Left:", mouseLeftClick, "Right:", mouseRightClick)
	// 		fmt.Printf("Move: %.2f px\n", mouseMovePixels)
	// 		fmt.Println("Wheel:", mouseWheelDelta)
	// 		statsLock.Unlock()
	// 	}
	// }()

	var msg struct{}
	for {
		getMessage.Call(uintptr(unsafe.Pointer(&msg)), 0, 0, 0)
	}
}

func (a *App) GetInputStats() StatsSnapshot {
	statsLock.Lock()
	defer statsLock.Unlock()

	modifier := make(map[string]uint64)
	key := make(map[string]uint64)
	combo := make(map[string]uint64)

	for k, v := range modifierStats {
		modifier[k] = v
	}
	for k, v := range keyStats {
		key[k] = v
	}
	for k, v := range comboStats {
		combo[k] = v
	}

	return StatsSnapshot{
		ModifierStats:   modifier,
		KeyStats:        key,
		ComboStats:      combo,
		MouseLeftClick:  mouseLeftClick,
		MouseRightClick: mouseRightClick,
		MouseMovePixels: mouseMovePixels,
		MouseWheel:      mouseWheelDelta,
	}
}
