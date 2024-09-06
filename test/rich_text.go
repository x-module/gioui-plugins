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
)

func main() {
	var th = theme.NewTheme()
	richText := widgets.NewRichText(th).AddSpan([]richtext.SpanStyle{
		{
			Content:     "Hello ",
			Color:       th.Color.DefaultTextWhiteColor,
			Size:        unit.Sp(24),
			Interactive: true,
		},
	}).OnClick(func(gtx layout.Context, content string) {
		println("clicked")
	}).OnHover(func(gtx layout.Context, content string) {
		println("hovered")
	}).OnUnHover(func(gtx layout.Context, content string) {
		println("unhovered")
	}).OnLongPress(func(gtx layout.Context, content string) {
		println("long pressed")
	})

	card := widgets.NewCard(th)
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return richText.Layout(gtx)
			})
		})
	})
	win.Run()
}
