package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"go-paint/apptype"
	"go-paint/pxcanvas"
	"go-paint/swatch"
	"go-paint/ui"
	"image/color"
)

func main() {
	paintApp := app.New()
	paintWindow := paintApp.NewWindow("Go Paint!")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	paintCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(600, 600),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       20,
		PxCols:       20,
		PxSize:       30,
	}

	paintCanvas := pxcanvas.NewPxCanvas(&state, paintCanvasConfig)

	appInit := ui.AppInit{
		PaintCanvas: paintCanvas,
		PaintWindow: paintWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PaintWindow.ShowAndRun()
}
