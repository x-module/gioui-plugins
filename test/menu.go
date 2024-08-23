package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/gioui-plugins/resource"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/utils"
	"github.com/gioui-plugins/widgets"
	"github.com/gioui-plugins/window"
)

func main() {
	th := theme.NewTheme()
	menu := widgets.NewMenu(th)
	menuItemOptions := []widgets.MenuItemOption{
		{Icon: resource.ActionPermIdentityIcon, Text: "ACCOUNT  ", MarginRight: 0,
			SubMenu: []widgets.MenuItemOption{
				{Icon: resource.NavigationSubdirectoryArrowRightIcon, Text: "LOGIN   "},
				{Icon: resource.NavigationSubdirectoryArrowRightIcon, Text: "LOGOUT"},
			},
		},
		{Icon: resource.EditorFunctionsIcon, Text: "RPC          ", MarginRight: 0},
		{Icon: resource.EditorBorderAllIcon, Text: "GENERATE", MarginRight: 0,
			SubMenu: []widgets.MenuItemOption{
				{Icon: resource.NavigationSubdirectoryArrowRightIcon, Text: "CONFIG"},
				{Icon: resource.NavigationSubdirectoryArrowRightIcon, Text: "CODE   "},
			},
		},
		{Icon: resource.MapsDirectionsRunIcon, Text: "MATCH     ", MarginRight: 0},
	}

	menu.SetClickCallback(func(main int, sub int) {
		fmt.Println("===============clicked================ main:", main, " sub:", sub)
	})
	menu.SetMenuItemOptions(menuItemOptions)

	win := window.NewInitialize()
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return menu.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return utils.DrawLine(gtx, th.Palette.Fg, unit.Dp(gtx.Constraints.Max.Y), unit.Dp(1))
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				// utils.ColorBackground(gtx, gtx.Constraints.Max, th.Palette.Bg)
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return widgets.Label(th, "Hello World").Layout(gtx)
				})
			}),
		)
	})
	win.Run()
}
