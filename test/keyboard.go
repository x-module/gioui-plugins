package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var th = theme.NewTheme()
	card := widgets.NewCard(th)

	var username *widgets.Input
	username = widgets.NewInput(th, "请输入名称...")

	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)

	keyBoard := widgets.NewKeyboard(th, win.GetWin())
	keyBoard.SetFilters([]event.Filter{
		key.Filter{
			Required: key.ModShortcut,
			Name:     "V",
		},
		key.Filter{
			Name: key.NameBack,
		},
		key.Filter{
			Name: key.NameEscape,
		}, key.Filter{
			Name: key.NameTab,
		}, key.Filter{
			Name: key.NameReturn,
		},
	})

	keyBoard.Event(func(e event.Event) {
		switch e := e.(type) {
		case key.Event:
			switch e.Name {
			case key.NameEscape:
				fmt.Println("esc")
			case key.NameBack:
				fmt.Println("back")
			case key.NameReturn:
				fmt.Println("NameReturn")
			case "V":
				fmt.Println("cmd+v")
			}
		}
	})
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		keyBoard.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Stack{Alignment: layout.Center}.Layout(gtx,
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									return username.Layout(gtx)
								}),
							)
						}),
					)
				})
			})
		})

	})
	win.Run()
}
