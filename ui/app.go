package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// NewApp creates and returns a new application.
func NewApp(onAppInputCapture func(*tcell.EventKey) *tcell.EventKey) (app *tview.Application) {
	app = tview.NewApplication().
		EnableMouse(true).
		SetInputCapture(onAppInputCapture)

	return app
}
