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
	th := theme.NewTheme()
	var clickable1 widget.Clickable
	card := widgets.NewCard(th)
	iconClick := widgets.NewIconButton(th, resource.NavigationRefreshIcon)
	iconClick.SetOnClick(func(gtx layout.Context) {
		fmt.Println("icon clicked")
	})
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return widgets.DefaultButton(th, &clickable1, "default", unit.Dp(100)).SetIcon(resource.DeleteIcon, 0).Layout(gtx)
					})
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(20))
					return iconClick.Layout(gtx)
				}),
			)
		})
	})
	win.Run()
}
