package app

import (
	"fmt"
	"os"

	"github.com/energye/systray"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func CreateTray(a *App, icon []byte) (start, end func()) {
	return systray.RunWithExternalLoop(func() {
		systray.SetIcon(icon)
		systray.SetTooltip("Input Stats")

		systray.SetOnClick(func(menu systray.IMenu) {
			if runtime.WindowIsMinimised(a.Ctx) {
				runtime.WindowShow(a.Ctx)
			}
		})

		mQuit := systray.AddMenuItem("退出", "退出 Input Stats")
		mQuit.Click(func() {
			a.ExitApp()
		})
	}, nil)
}

func (a *App) ExitApp() {
	if !a.OnBeforeClose(a.Ctx) {
		systray.Quit()
		runtime.Quit(a.Ctx)
		os.Exit(0)
	}
}

func createMenuItem(a *App, menu MenuItem) {
	switch menu.Type {
	case "item":
		var m *systray.MenuItem
		m = systray.AddMenuItem(menu.Title, menu.Tooltip)
		m.Click(func() {
			fmt.Printf("%s", menu.Event)
			go runtime.EventsEmit(a.Ctx, menu.Event)
		})

		if menu.Disabled {
			m.Disable()
		}
	case "separator":
		systray.AddSeparator()
	}
}

func updateTrayMenus(a *App, menus []MenuItem) {
	systray.ResetMenu()
	for _, menu := range menus {
		createMenuItem(a, menu)
	}
}

func (a *App) UpdateTrayMenus(menus []MenuItem) {
	updateTrayMenus(a, menus)
}
