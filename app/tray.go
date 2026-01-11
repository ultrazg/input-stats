package app

import (
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
	}, nil)
}
