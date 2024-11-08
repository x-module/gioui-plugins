package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"time"
)

func main() {
	var th = theme.NewTheme()
	var ed widget.Editor
	card := widgets.NewCard(th)
	logScroll := widgets.NewScroll(th)
	txt := " Network showdown-nakama_default  Creating"

	logScroll.SetElementList([]layout.Widget{
		func(gtx layout.Context) layout.Dimensions {
			med := material.Editor(th.Material(), &ed, "")
			med.TextSize = unit.Sp(16)
			med.Color = th.Color.HintTextColor
			return med.Layout(gtx)
		},
	})
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)

	go func() {
		for {
			time.Sleep(time.Second * 1) // 每5秒添加一条日志
			txt += "\n" + fmt.Sprintf("New log entry at %s", time.Now())
			ed.SetText(txt)
			win.GetWin().Invalidate()
		}
	}()

	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		// ==============================================
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return logScroll.Layout(gtx)
			})
		})
	})
	win.Run()
}
