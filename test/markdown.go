package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/x/markdown"
	"gioui.org/x/richtext"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {

	var th = theme.NewTheme()
	card := widgets.NewCard(th)
	renderer := markdown.NewRenderer()
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	var TextState richtext.InteractiveText
	// cache := []richtext.SpanStyle{
	// 	{
	// 		Content: "Hello, Gio!",
	// 		Size:    unit.Sp(20),
	// 		Color:   th.Color.GreenColor,
	// 	},
	// }

	cache1, _ := renderer.Render([]byte("* aasdfasdfasdf \n" +
		"* asdffffffff \n" +
		"<u>underline</u> \n"))

	for key := range cache1 {
		cache1[key].Color = th.Color.WhiteColor
	}

	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return richtext.Text(&TextState, th.Material().Shaper, cache1...).Layout(gtx)
						})
					})
				}),
			)
		})
	})
	win.Run()
}
