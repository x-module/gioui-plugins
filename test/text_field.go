package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
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
	card := widgets.NewCard(th)
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	text := component.TextField{
		CharLimit: 10,
		Helper:    "HelperHelperHelper",
	}
	text.Prefix = func(gtx layout.Context) layout.Dimensions {
		th.Palette.Fg = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
		return widgets.Label(th, "$").Layout(gtx)
	}
	text.Suffix = func(gtx layout.Context) layout.Dimensions {
		th.Palette.Fg = color.NRGBA{R: 100, G: 100, B: 100, A: 255}
		return widgets.Label(th, ".00").Layout(gtx)
	}
	th.Palette.ContrastBg = th.Color.WhiteColor
	th.Palette.Fg = th.Color.DefaultTextWhiteColor
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return text.Layout(gtx, th.Material(), "Name")
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
