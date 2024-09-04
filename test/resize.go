package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"image"
	"image/color"
)

func main() {
	var clickable widget.Clickable

	var th = theme.NewTheme()
	card := widgets.NewCard(th)
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	resize := component.Resize{
		Axis:  layout.Horizontal,
		Ratio: 0.5,
	}

	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		resize.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).SetIcon(resource.DeleteIcon, 0).Layout(gtx)
					})
				})
			},
			func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).SetIcon(resource.DeleteIcon, 0).Layout(gtx)
							})
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).SetIcon(resource.DeleteIcon, 0).Layout(gtx)
							})
						}),
					)
				})
			},
			func(gtx layout.Context) layout.Dimensions {
				rect := image.Rectangle{
					Max: image.Point{
						X: (gtx.Dp(unit.Dp(4))),
						Y: (gtx.Constraints.Max.Y),
					},
				}
				paint.FillShape(gtx.Ops, color.NRGBA{A: 200}, clip.Rect(rect).Op())
				return layout.Dimensions{Size: rect.Max}
			},
		)

	})
	win.Run()
}
