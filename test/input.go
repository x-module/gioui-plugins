package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var th = theme.NewTheme()
	card := widgets.NewCard(th)

	var username *widgets.Input
	var password *widgets.Input
	var password2 *widgets.Input
	var age *widgets.Input
	var profile *widgets.Input
	username = widgets.NewInput(th, "请输入名称...")
	age = widgets.NewInput(th, "请输入年龄...", "á3452345234523452345")
	password = widgets.NewInput(th, "请输入密码...")
	password2 = widgets.NewInput(th, "请输入确认密码===...")
	profile = widgets.NewTextArea(th, "请输入属性2222...")

	username.SetSize(th.Size.Tiny)
	password.SetSize(th.Size.Small)
	password.SetRadius(unit.Dp(8))
	password2.SetSize(th.Size.Medium)
	age.SetSize(th.Size.Large)
	age.ReadOnly()
	password2.SetBefore(func(gtx layout.Context) layout.Dimensions {
		return resource.ActionPermIdentityIcon.Layout(gtx, th.Color.DefaultIconColor)
	})
	password2.Password()

	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Stack{Alignment: layout.Center}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return username.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return password.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return password2.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return age.Layout(gtx)
							}),
							layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return profile.Layout(gtx)
							}),
						)
					}),
				)
			})
		})
	})
	win.Run()
}
