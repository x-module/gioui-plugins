package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"time"
)

func main() {
	var clickable widget.Clickable
	var th = theme.NewTheme()
	change := false
	change1 := false
	card := widgets.NewCard(th)
	image := widgets.NewImage(th, "test/welcome.jpg")
	win1 := new(app.Window)
	win := window.NewInitialize(win1)
	win.Title("Hello, Gio!").Size(600, 400)
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.NoActionBar().CenterWindow()
	go func() {
		for {
			select {
			case <-time.After(time.Second * 1):
				change = !change
				win.Size(800, 500).ReCenterWindow()
				win1.Invalidate()
				return
			}
		}
	}()
	win.Frame(func(gtx layout.Context, ops op.Ops, w *app.Window) {
		if clickable.Clicked(gtx) {
			fmt.Println("------click-----------")
			if !change1 {
				win.Size(1200, 900).HaveActionBar().ReCenterWindow()
			} else {
				win.Size(1100, 800).HaveActionBar().ReCenterWindow()
			}
			change1 = !change1
		}
		if change {
			layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).SetIcon(resource.DeleteIcon, 0).Layout(gtx)
						})
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).SetIcon(resource.DeleteIcon, 0).Layout(gtx)
						})
					}),
				)
			})
		} else {
			image.Layout(gtx)
		}
	})
	win.Run()
}
