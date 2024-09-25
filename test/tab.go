package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/utils"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"image"
)

func main() {
	var clickable widget.Clickable
	th := theme.NewTheme()
	card := widgets.NewCard(th)
	// w := new(app.Window)
	tabs := widgets.NewTabs(th)
	tabs.SetTabs([]*widgets.Tab{
		{
			Title: "Email",
			Content: func(gtx layout.Context) layout.Dimensions {
				return widgets.Label(th, "Email").Layout(gtx)
			},
		},
		{
			Title: "Custom",
			Content: func(gtx layout.Context) layout.Dimensions {
				return widgets.Label(th, "Custom").Layout(gtx)
			},
		},
		{
			Title: "Token",
			Content: func(gtx layout.Context) layout.Dimensions {
				return widgets.Label(th, "Token").Layout(gtx)
			},
		},
	})

	tabs.SetWidth(unit.Dp(400))
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		// fmt.Println("selected:", tabs.SelectedTab().Title)
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return tabs.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								// gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(300))
								utils.DrawBackground(gtx, image.Point{
									X: gtx.Constraints.Max.X,
									Y: gtx.Constraints.Min.Y,
								}, th.Color.DefaultWindowBgGrayColor)
								return layout.Inset{Top: unit.Dp(150), Left: unit.Dp(200), Bottom: unit.Dp(150)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									return tabs.CurrentTab(gtx)
								})
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
