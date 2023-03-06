package ui

import (
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"go-paint/util"
	"image"
	"image/png"
	"os"
	"strconv"
)

func saveFileDialog(app *AppInit) {
	dialog.ShowFileSave(func(uri fyne.URIWriteCloser, err error) {
		if uri == nil {
			return
		} else {
			err := png.Encode(uri, app.PaintCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PaintWindow)
				return
			}
			app.State.SetFilePath(uri.URI().Path())
		}
	}, app.PaintWindow)
}

func BuildSaveAsMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save as...", func() {
		saveFileDialog(app)
	})
}

func BuildSaveMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Save", func() {
		if app.State.FilePath == "" {
			saveFileDialog(app)
			return
		} else {
			tryClose := func(fh *os.File) {
				err := fh.Close()
				if err != nil {
					dialog.ShowError(err, app.PaintWindow)
				}
			}

			fh, err := os.Create(app.State.FilePath)
			defer tryClose(fh)

			if err != nil {
				dialog.ShowError(err, app.PaintWindow)
				return
			}

			err = png.Encode(fh, app.PaintCanvas.PixelData)
			if err != nil {
				dialog.ShowError(err, app.PaintWindow)
				return
			}

		}
	})
}

func BuildNewMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("New", func() {
		sizeValidator := func(s string) error {
			size, err := strconv.Atoi(s)
			if err != nil {
				return errors.New("Must be a positive integer")
			}
			if size <= 0 {
				return errors.New("Must be > 0")
			}
			return nil
		}

		widthEntry := widget.NewEntry()
		widthEntry.Validator = sizeValidator

		heightEntry := widget.NewEntry()
		heightEntry.Validator = sizeValidator

		widthFormEntry := widget.NewFormItem("Width", widthEntry)
		heightFormEntry := widget.NewFormItem("Height", heightEntry)

		formItems := []*widget.FormItem{widthFormEntry, heightFormEntry}

		dialog.ShowForm("New Image", "Create", "Cancel", formItems, func(ok bool) {
			if ok {
				pixelWidth := 0
				pixelHeight := 0
				if widthEntry.Validate() != nil {
					dialog.ShowError(errors.New("Invalid width"), app.PaintWindow)
				} else {
					pixelWidth, _ = strconv.Atoi(widthEntry.Text)
				}

				if heightEntry.Validate() != nil {
					dialog.ShowError(errors.New("Invalid height"), app.PaintWindow)
				} else {
					pixelHeight, _ = strconv.Atoi(heightEntry.Text)
				}

				app.PaintCanvas.NewDrawing(pixelWidth, pixelHeight)
			}
		}, app.PaintWindow)
	})
}

func BuildOpenMenu(app *AppInit) *fyne.MenuItem {
	return fyne.NewMenuItem("Open...", func() {
		dialog.ShowFileOpen(func(uri fyne.URIReadCloser, err error) {
			if uri == nil {
				return
			} else {
				image, _, err := image.Decode(uri)
				if err != nil {
					dialog.ShowError(err, app.PaintWindow)
					return
				}
				app.PaintCanvas.LoadImage(image)
				app.State.SetFilePath(uri.URI().Path())
				imgColors := util.GetImageColors(image)
				i := 0
				for c := range imgColors {
					if i == len(app.Swatches) {
						break
					}

					app.Swatches[i].SetColor(c)
					i++
				}
			}
		}, app.PaintWindow)
	})
}

func BuildMenus(app *AppInit) *fyne.Menu {
	return fyne.NewMenu("File",
		BuildNewMenu(app),
		BuildOpenMenu(app),
		BuildSaveMenu(app),
		BuildSaveAsMenu(app),
	)
}

func SetupMenus(app *AppInit) {
	app.PaintWindow.SetMainMenu(fyne.NewMainMenu(BuildMenus(app)))
}
