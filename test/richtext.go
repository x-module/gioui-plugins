package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/x/richtext"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"strings"
)

func main() {

	var th = theme.NewTheme()
	card := widgets.NewCard(th)
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	var TextState richtext.InteractiveText

	// Replace tabs with spaces
	content := "win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {\n\t\tlayout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {\n\t\t\treturn layout.Flex{Axis: layout.Vertical}.Layout(gtx,\n\t\t\t\tlayout.Rigid(func(gtx layout.Context) layout.Dimensions {\n\t\t\t\t\treturn card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {\n\t\t\t\t\t\treturn layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {\n\t\t\t\t\t\t\treturn richtext.Text(&TextState, th.Material().Shaper, cache...).Layout(gtx)\n\t\t\t\t\t\t})\n\t\t\t\t\t})\n\t\t\t\t}),\n\t\t\t)\n\t\t})\n\t})"
	content = strings.ReplaceAll(content, "\t", "    ")

	cache := []richtext.SpanStyle{
		{
			Content: content,
			Size:    unit.Sp(20),
			Color:   th.Color.GreenColor,
		},
		{
			Content: "gigi!",
			Size:    unit.Sp(20),
			Color:   th.Color.WhiteColor,
		},
	}

	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return richtext.Text(&TextState, th.Material().Shaper, cache...).Layout(gtx)
						})
					})
				}),
			)
		})
	})
	win.Run()
}
