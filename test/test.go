package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable
	th := theme.NewTheme()
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			component.Shadow(unit.Dp(2), unit.Dp(100)).Layout(gtx)
			component.Shadow(unit.Dp(2), unit.Dp(100)).Layout(gtx)
			component.Shadow(unit.Dp(2), unit.Dp(100)).Layout(gtx)
			btn := material.Button(th.Material(), &clickable, "Right Aligned Button")
			return btn.Layout(gtx)
		})
	})
	win.Run()
}
