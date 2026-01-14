package app

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) OnBeforeClose(ctx context.Context) (prevent bool) {
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Input Stats",
		Message: "确定要退出 Input Stats 吗？",
		Buttons: []string{"确定", "取消"},
	})
	if err != nil {
		return false
	}
	return dialog != "Yes"
}
