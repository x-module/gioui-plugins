package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	th := theme.NewTheme()
	dropDown := widgets.NewSearchDropDown(th)
	dropDown.SetOnChanged(func(value string) {
		println(dropDown.GetSelected())
	})
	dropDown.SetOptions([]*widgets.SearchDropDownOption{
		{
			Value: "1",
			Text:  "1",
		},
		{
			Value: "2",
			Text:  "2",
		},
		{
			Value: "3",
			Text:  "3",
		},
		{
			Value: "4",
			Text:  "4",
		},
	})
	dropDown.SetWidth(unit.Dp(300))
	card := widgets.NewCard(th)
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
						return dropDown.Layout(gtx, th)
					})
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
			)
		})
	})
	win.Run()
}
