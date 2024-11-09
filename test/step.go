/**
 * Created by Goland
 * @file   split_window.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/8 13:53
 * @desc   split_window.go
 */

package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

var steps = []widgets.StepItem{
	{
		Title: "Step 1",
	},
	{
		Title: "Step 2",
	},
	{
		Title: "Step 3",
	},
	{
		Title: "Step 4",
	},
	{
		Title: "Step 5",
	},
}

func main() {
	var th = theme.NewTheme()
	card := widgets.NewCard(th)
	step := widgets.NewStep(th)
	step.SetSteps(steps)
	step.SetCurrentStep(2)
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return step.Layout(gtx)
		})
	})
	win.Run()
}
