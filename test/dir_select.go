package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/widgets"
	"github.com/gioui-plugins/window"
)

func main() {
	var th = theme.NewTheme()
	var dirSelect = widgets.NewDirSelector(th, "请选择目录...")
	dirSelect.SetWidth(unit.Dp(400))
	dirSelect.SetOnSelectDir(func(dir string) {
		fmt.Println("dir:", dir)
	})
	win := window.NewInitialize()
	win.Title("Hello, Gio!").Size(800, 600)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(400))
					return widgets.Label(th, "&clickable, nil, 0,  unit.Dp(100)").Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return dirSelect.Layout(gtx)
				}),
			)
		})
	})
	win.Run()
}
