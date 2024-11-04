package main

import (
	"fmt"
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
	th := theme.NewTheme()
	card := widgets.NewCard(th)
	menuItemOptions := []*widgets.ListMenuOption{
		{
			Text:  "属性",
			Value: "profile",
		},
		{
			Text:  "通知",
			Value: "notice",
		},
		{
			Text:  "退出",
			Value: "exit",
		},
	}
	menus := widgets.NewListMenu(th)
	menus.SetLabel("土豆").SetOptions(menuItemOptions)
	menus.SetLabelWidth(unit.Dp(50))
	menus.SetMenuWidth(unit.Dp(200))
	menus.Clicked(func(key int, menu string) {
		fmt.Println(key, menu)
	})
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								gtx.Constraints.Min.X = gtx.Dp(unit.Dp(700))
								return layout.Spacer{Width: unit.Dp(10)}.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return menus.Layout(gtx)
							}),
						)
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
	})
	win.Run()
}
