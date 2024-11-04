package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"time"
)

func main() {
	var clickable widget.Clickable
	var clickable1 widget.Clickable

	th := theme.NewTheme()
	var card = widgets.NewCard(th)
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		if clickable.Clicked(gtx) {
			widgets.SendSystemNotice("登录成功")
		}
		if clickable1.Clicked(gtx) {
			ts := time.Now().Unix()
			widgets.SendAppInfoNotice(th, "登录成功 ts:"+fmt.Sprint(ts))
			// widgets.SendAppSuccessNotice(th, "登录成功 ts:"+fmt.Sprint(ts))
			// widgets.SendAppErrorNotice(th, "登录成功 ts:"+fmt.Sprint(ts))
			// widgets.SendAppWaringNotice(th, "登录成功 ts:"+fmt.Sprint(ts))
		}
		layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return widgets.DefaultButton(th, &clickable, "system notice", unit.Dp(100)).Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return widgets.DefaultButton(th, &clickable1, "application notice", unit.Dp(100)).Layout(gtx)
						}),
					)
				})
			}),
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				return widgets.NotificationController.Layout(gtx, th)
			}),
		)
	})
	win.Run()
}
