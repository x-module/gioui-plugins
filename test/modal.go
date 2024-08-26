package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable
	var clickable1 widget.Clickable
	th := theme.NewTheme()
	var model = widgets.NewModal(th)

	win := window.NewInitialize()
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		if clickable.Clicked(gtx) {
			model.Display(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.Label(th, "Hello, World!").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.DefaultButton(th, &clickable1, "click me", unit.Dp(100)).Layout(gtx)
					}),
				)
			})
		}
		layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(400))
						return widgets.Label(th, "&clickable, nil, 0,  unit.Dp(100)").Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return widgets.DefaultButton(th, &clickable, "click me", unit.Dp(100)).Layout(gtx)
					}),
				)
			}),
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if model.Visible() {
					return model.Layout(gtx)
				}
				return layout.Dimensions{}
			}),
		)
	})
	win.Run()
}
