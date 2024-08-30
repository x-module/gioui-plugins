package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	th := theme.NewTheme()
	fileSelect := widgets.NewFileSelector(th, "请选择文件...")
	fileSelect.SetWidth(unit.Dp(600))
	// fileSelect.SetFilter(".json")
	fileSelect.SetOnSelectFile(func(dir string) {
		fmt.Println("file:", dir)
	})
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(400))
				return widgets.Label(th, "&clickable, nil, 0,  unit.Dp(100)").Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return fileSelect.Layout(gtx)
			}),
		)
	})
	win.Run()
}
