package utils

import (
	"image"
	"log"

	"go-gui/pkg/shared/handlers"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func Run(w *app.Window) error {
	// UI state
	var fetchButton widget.Clickable
	var currentImage image.Image

	// Theme for material widgets
	th := material.NewTheme()

	var ops op.Ops
	for {
		switch e := w.Event().(type) {
		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)

			// Handle button click
			if fetchButton.Clicked(gtx) {
				go func() {
					img, err := handlers.HandleButtonClick()
					if err != nil {
						log.Printf("Error handling button click: %v", err)
					} else {
						currentImage = img
					}

				}()
			}

			// Layout
			layout.Flex{
				Axis:    layout.Vertical,
				Spacing: layout.SpaceStart,
			}.Layout(gtx,
				// Button at top
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &fetchButton, "Fetch Image")
						return btn.Layout(gtx)
					})
				}),

				// Image display area
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(16)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if currentImage == nil {
							// Show placeholder
							return layout.Dimensions{Size: gtx.Constraints.Min}
						}
						return DrawImage(gtx, currentImage)
					})
				}),
			)

			e.Frame(gtx.Ops)
		}
	}
}
