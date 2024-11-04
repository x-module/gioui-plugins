package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable
	var th = theme.NewTheme()
	card := widgets.NewCard(th)
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	customWindow := widgets.NewWindow(th, win.GetWin())
	menuItemOptions := []*widgets.ListMenuOption{
		{
			Text:  "属性",
			Value: "profile",
		},
		{
			Text:  "通知",
			Value: "notice",
		},
		{
			Text:  "退出",
			Value: "exit",
		},
	}
	menus := widgets.NewListMenu(th)
	menus.SetLabel("土豆").SetOptions(menuItemOptions)
	menus.SetLabelWidth(unit.Dp(50))
	menus.SetMenuWidth(unit.Dp(100))
	menus.Clicked(func(key int, menu string) {
		fmt.Println(key, menu)
	})
	customWindow.SetContent(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Top: unit.Dp(5)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// return widgets.Label(th, "Hello World").Layout(gtx)
			return menus.Layout(gtx)
		})
	})
	customWindow.SetCloseWinHook(func() bool {
		fmt.Println("close window")
		return true
	})
	customWindow.SetMinWinHook(func() bool {
		fmt.Println("minimize window")
		return true
	})
	customWindow.SetFullWinHook(func() bool {
		fmt.Println("full window")
		return true
	})
	customWindow.SetUnMaximizeWinHook(func() bool {
		fmt.Println("unmaximize window")
		return true
	})

	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.NoActionBar().CenterWindow()
	win.Frame(func(gtx layout.Context, ops op.Ops, w *app.Window) {
		if clickable.Clicked(gtx) {
			win.Size(window.ElementStyle{
				Width:  1200,
				Height: 900,
			}).ReCenterWindow()
		}
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return customWindow.Layout(gtx)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return widgets.NewLine(th).Color(th.Color.DefaultLineColor).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
			}),
			// content
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
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
			}),
		)
	})
	win.Run()
}
