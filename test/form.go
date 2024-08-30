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

	var th = theme.NewTheme()
	form := widgets.NewForm(th, unit.Dp(40), unit.Dp(80))
	// dropDown := widgets.NewDropDown(th, []string{"a", "b", "c", "d"}...)

	userName := widgets.NewInput(th, "please input username")
	password := widgets.NewInput(th, "please input password")
	email := widgets.NewInput(th, "please input email")

	form.Add("username", userName.Layout)
	form.Add("password", password.Layout)
	form.Add("email", email.Layout)

	form.Add("", widgets.BlueButton(th, &clickable, "submit", unit.Dp(100)).Layout)

	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		// form.Add("server", func(gtx layout.Context) layout.Dimensions {
		// 	return dropDown.Layout(gtx, th)
		// })
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return form.Layout(gtx)
		})
	})
	win.Run()
}
