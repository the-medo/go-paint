package ui

import (
	"fyne.io/fyne/v2"
	"go-paint/apptype"
	"go-paint/pxcanvas"
	"go-paint/swatch"
)

type AppInit struct {
	PaintCanvas *pxcanvas.PxCanvas
	PaintWindow fyne.Window
	State       *apptype.State
	Swatches    []*swatch.Swatch
}
