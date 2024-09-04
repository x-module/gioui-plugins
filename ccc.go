package main

import (
	"image"
	"image/color"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Tooltip 示例"),
			app.Size(unit.Dp(400), unit.Dp(300)),
		)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	var ops op.Ops
	th := material.NewTheme(gofont.Collection())
	var button widget.Clickable
	showTooltip := false
	var tooltipArea image.Rectangle
	var tooltipTimer *time.Timer

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			pointer.Rect(image.Rectangle{Max: gtx.Constraints.Max}).Add(gtx.Ops)
			pointer.InputOp{Tag: &button, Types: pointer.Enter | pointer.Leave}.Add(gtx.Ops)

			for _, ev := range button.History() {
				if ev.Type == pointer.Enter {
					if tooltipTimer == nil {
						tooltipTimer = time.AfterFunc(500*time.Millisecond, func() {
							showTooltip = true
							w.Invalidate()
						})
					}
				} else if ev.Type == pointer.Leave {
					if tooltipTimer != nil {
						tooltipTimer.Stop()
						tooltipTimer = nil
					}
					showTooltip = false
				}
			}

			layout.Flex{}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &button, "Hover over me")
						btn.Background = color.NRGBA{R: 255, A: 255}
						dims := btn.Layout(gtx)

						// Update tooltip area based on the button dimensions.
						tooltipArea = image.Rectangle{
							Min: gtx.Constraints.Min,
							Max: gtx.Constraints.Min.Add(dims.Size),
						}
						return dims
					})
				}),
			)

			if showTooltip {
				op.Offset(layout.FPt(tooltipArea.Min)).Add(gtx.Ops)
				paint.FillShape(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 0, A: 150},
					clip.RRect{Rect: image.Rectangle{
						Max: image.Point{X: 160, Y: 40},
					}}.Op(gtx.Ops),
				)
				lbl := material.Label(th, unit.Sp(16), "Tooltip text goes here")
				lbl.Color = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
				lbl.Alignment = text.Middle
				lbl.Layout(gtx)
			}

			e.Frame(gtx.Ops)
		}
	}
}
