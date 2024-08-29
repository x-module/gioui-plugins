package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable
	th := theme.NewTheme()
	win := window.NewInitialize(new(app.Window))
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			btn := material.Button(th.Material(), &clickable, "Right Aligned Button")
			return btn.Layout(gtx)
		})
	})
	win.Run()
}
