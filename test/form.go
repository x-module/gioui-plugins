package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable

	var th = theme.NewTheme()
	th.Size.DefaultTextSize = unit.Sp(14)
	form := widgets.NewForm(th, unit.Dp(30), unit.Dp(120))
	// dropDown := widgets.NewDropDown(th, []string{"a", "b", "c", "d"}...)

	userName := widgets.NewInput(th, "please input username11")
	codeInput := widgets.NewInput(th, "请输入CODE...")
	password := widgets.NewInput(th, "please input password")
	dropDown := widgets.NewDropDown(th, []string{"a", "b", "c", "d"}...)

	// codeInput.SetSize(th.Size.Large)
	// userName.SetSize(th.Size.Medium)
	// password.SetSize(th.Size.Medium)
	// password.SetSize(th.Size.Medium)
	password.SetBefore(func(gtx layout.Context) layout.Dimensions {
		return resource.ActionPermIdentityIcon.Layout(gtx, th.Color.DefaultIconColor)
	})
	password.Password()
	form.Add("username", userName.Layout)
	form.Add("password", password.Layout)
	form.Add("dropDown", dropDown.Layout)
	form.Add("code", func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				codeInput.SetWidth(unit.Dp(200))
				return codeInput.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				but := widgets.DefaultButton(th, &clickable, "submit", unit.Dp(100))
				return but.Layout(gtx)
			}),
		)
	})
	form.AddButton(widgets.BlueButton(th, &clickable, "submit", unit.Dp(100)).Layout)

	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
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
