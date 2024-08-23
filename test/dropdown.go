package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/widgets"
	"github.com/gioui-plugins/window"
)

func main() {
	th := theme.NewTheme()

	dropDown := widgets.NewDropDown(th, []string{"a", "b", "c", "d"}...)
	dropDown2 := widgets.NewDropDown(th, []string{"aaa", "bbb", "cccc", "dddd"}...)

	dropDown2.SetOnChanged(func(value string) {
		println(value)
	})

	card := widgets.NewCard(th)
	// dropDown.SetOptions(options)
	dropDown.SetWidth(unit.Dp(300))
	win := window.NewInitialize()
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return dropDown.Layout(gtx, th)
					})
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return dropDown2.Layout(gtx, th)
					})
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
			)
		})
	})
	win.Run()
}
