package app

import (
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
	systray.Quit()
	runtime.Quit(a.Ctx)
	os.Exit(0)
}
