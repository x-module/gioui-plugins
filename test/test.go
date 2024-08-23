package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable
	th := theme.NewTheme()
	win := window.NewInitialize()
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(layout.Spacer{Height: 100}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X = 200
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return material.Button(th.Material(), &clickable, "Hello, Gio!").Layout(gtx)
				})
			}),
		)
	})
	win.Run()
}
