package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
)

func main() {
	var username *widgets.Input
	var password *widgets.Input
	var password2 *widgets.Input
	var age *widgets.Input
	var profile *widgets.Input
	// var clickable widget.Clickable
	var th = theme.NewTheme()

	// w := new(app.Window)
	var ops op.Ops
	username = widgets.NewInput(th, "请输入名称...")
	age = widgets.NewInput(th, "请输入年龄...", "á3452345234523452345")
	password = widgets.NewInput(th, "请输入密码...")
	password2 = widgets.NewInput(th, "请输入确认密码===...")
	profile = widgets.NewTextArea(th, "请输入属性...")

	username.SetSize(th.Size.Tiny)
	password.SetSize(th.Size.Small)
	password.SetRadius(unit.Dp(8))
	password2.SetSize(th.Size.Medium)
	age.SetSize(th.Size.Large)

	age.ReadOnly()
	// password2.SetAfter(func(gtx layout.Context) layout.Dimensions {
	// 	return widgets2.NavigationSubdirectoryArrowRightIcon.Layout(gtx, resource.IconColor)
	// })
	password2.SetBefore(func(gtx layout.Context) layout.Dimensions {
		return resource.ActionPermIdentityIcon.Layout(gtx, th.Color.IconGrayColor)
	})

	password2.Password()
	go func() {
		w := new(app.Window)
		for {
			e := w.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				panic(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				rect := clip.Rect{
					Max: gtx.Constraints.Max,
				}
				paint.FillShape(gtx.Ops, th.Color.DefaultWindowBgGrayColor, rect.Op())
				// =============================================
				// ==============================================
				layout.Stack{Alignment: layout.Center}.Layout(gtx,
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
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
