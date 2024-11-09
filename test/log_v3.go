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
	"time"
)

func main() {
	var th = theme.NewTheme()
	card := widgets.NewCard(th)

	win := window.NewApplication(new(app.Window))
	log := widgets.NewLog(th, win.GetWin())
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)

	go func() {
		for {
			time.Sleep(time.Second * 1) // 每5秒添加一条日志
			txt := fmt.Sprintf("New log entry at %s", time.Now())
			log.SetLogData(txt)
		}
	}()
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return log.Layout(gtx)
			})
		})
	})
	win.Run()
}
