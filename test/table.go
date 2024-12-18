package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {

	var th = theme.NewTheme()
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})

	table := widgets.NewTable(th)
	headers := []string{"Name", "Age", "Class", "BirthDay", "Address"}
	data := [][]any{
		{"Tom", 18, "Class 1", "2000-01-01", "Beijing"},
		{"Jerry", 19, "Class 2", "2001-01-01", "Shanghai"},
		{"Lucy", 20, "Class 3", "2002-01-01", "Guangzhou"},
		{"Jack", 21, "Class 4", "2003-01-01", "Shenzhen"},
	}
	table.SetData(data)
	table.SetHeader(headers)

	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// return table.LayoutHoverTable(gtx)
			return table.LayoutHoverTable(gtx)
		})
	})
	win.Run()
}
