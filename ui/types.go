package ui

import (
	"fyne.io/fyne/v2"
	"go-paint/apptype"
	"go-paint/swatch"
)

type AppInit struct {
	PaintWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
