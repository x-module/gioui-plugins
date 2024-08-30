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
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	grid := widgets.NewTable(th)
	data := []map[string]any{
		{"name": "one", "age": 1, "class": 3, "birthDay": "2022-01-01", "address": "beijing"},
		{"name": "two", "age": 2, "class": 3, "birthDay": "2022-01-01", "address": "beijing"},
		{"name": "three", "age": 3, "class": 3, "birthDay": "2022-01-01", "address": "beijing"},
		{"name": "four", "age": 4, "class": 3, "birthDay": "2022-01-01", "address": "beijing"},
		{"name": "five", "age": 5, "class": 3, "birthDay": "2022-01-01", "address": "beijing"},
		{"name": "six", "age": 6, "class": 3, "birthDay": "2022-01-01", "address": "beijing"},
	}
	grid.SetData(data)
	grid.SetHeader([]string{"Name", "Age", "Class", "BirthDay", "Address"})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return grid.Layout(gtx)
		})
	})
	win.Run()
}
