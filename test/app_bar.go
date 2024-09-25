package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/x/component"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"image/color"
)

func main() {
	var clickable widget.Clickable

	var th = theme.NewTheme()
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})

	th.Palette.Fg = th.Color.DefaultContentBgGrayColor
	th.Palette.ContrastBg = th.Color.DefaultContentBgGrayColor
	actions := []component.AppBarAction{
		{
			Layout: func(gtx layout.Context, bg, fg color.NRGBA) layout.Dimensions {
				return widgets.Label(th, "Favorite").Layout(gtx)
			},
		},
		component.SimpleIconAction(&clickable, resource.EditIcon,
			component.OverflowAction{
				Name: "Create",
				Tag:  "&p.plusBtn",
			},
		),
	}
	// bar := component.AppBar{
	// 	NavigationIcon:   resource.ActionFullIcon,
	// 	NavigationButton: component.SimpleIconAction(th.Color.DefaultIconColor, resource.CloseIcon, component.AppBarNavigationClicked{}),
	// }

	bar := component.NewAppBar(component.NewModal())

	bar.NavigationIcon = resource.ActionFullIcon
	bar.Title = "Name"
	bar.SetActions(actions, nil)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		bar.Layout(gtx, th.Material(), "Hello, Gio!", "asdfasdfasdfasd")
	})
	win.Run()
}
func Overflow() []component.OverflowAction {
	return []component.OverflowAction{
		{
			Name: "Example 1",
			Tag:  "&p.exampleOverflowState",
		},
		{
			Name: "Example 2",
			Tag:  "&p.exampleOverflowState",
		},
	}
}
