package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/widgets"
	"github.com/gioui-plugins/window"
)

func main() {
	var th = theme.NewTheme()
	win := window.NewInitialize()
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(layout.Spacer{Height: 30}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widgets.NewLine(th).Color(th.Color.GreenColor).Width(3).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
			}),
			layout.Rigid(layout.Spacer{Height: 30}.Layout),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widgets.NewLine(th).Width(2).DashedLine(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
			}),
		)
	})
	win.Run()
}
