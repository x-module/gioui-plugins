package main

import (
	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	_ "github.com/x-module/helper"
)

func main() {
	var th = theme.NewTheme()

	richText := widgets.NewRichText(th).AddSpan([]widgets.SpanStyle{
		{
			Content:     "我们都是好人",
			Color:       th.Color.GreenColor,
			Size:        unit.Sp(14),
			Interactive: true,
			Font: font.Font{
				Typeface: "go",
				Weight:   font.Bold,
				Style:    font.Regular,
				// Style:    font.Italic,
			},
		},
		{
			Content:     "斜体",
			Color:       th.Color.GreenColor,
			Size:        unit.Sp(14),
			Interactive: true,
			Font: font.Font{
				Typeface: "go",
				Weight:   font.Bold,
				Style:    font.Regular,
			},
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
				return richText.MarkLayout(gtx)
			})
		})
	})
	win.Run()
}
