package main

import (
	"fyne.io/fyne/v2/app"
	"go-paint/apptype"
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

	appInit := ui.AppInit{
		PaintWindow: paintWindow,
		State:       &state,
		Swatches:    make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PaintWindow.ShowAndRun()
}
