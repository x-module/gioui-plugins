package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/utils"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	th := theme.NewTheme()
	menu := widgets.NewMenu(th)
	menuItemOptions := []widgets.MenuItem{
		{Icon: resource.ActionPermIdentityIcon, Text: "Account",
			SubMenu: []widgets.MenuItem{
				{Text: "Login"},
				{Text: "Logout"},
			},
		},
		{Icon: resource.EditorFunctionsIcon, Text: "RPC"},
		{Icon: resource.EditorBorderAllIcon, Text: "Generate",
			SubMenu: []widgets.MenuItem{
				{Text: "config"},
				{Text: "code"},
			},
		},
		{Icon: resource.MapsDirectionsRunIcon, Text: "Match"},
		{Icon: resource.SettingsIcon, Text: "Settings",
			SubMenu: []widgets.MenuItem{
				{Text: "setting"},
				{Text: "Exit"},
			},
		},
	}
	menu.SetMenuItemOptions(menuItemOptions)
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return menu.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return utils.DrawLine(gtx, th.Color.DefaultLineColor, unit.Dp(gtx.Constraints.Max.Y), unit.Dp(1))
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
