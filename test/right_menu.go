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
)

var options = []*widgets.RightMenuOption{
	{
		Text:  "菜单笔记",
		Value: "菜单笔记",
	},
	{
		Text:  "在新标签页打开",
		Value: "在新标签页打开",
	},
	{
		Text:  "在新窗口页打开",
		Value: "在新窗口页打开",
	},
	{
		Text:  "保存为自定义模板",
		Value: "保存为自定义模板",
	},
	{
		Line: 1,
	},
	{
		Text:  "添加到快捷方式",
		Value: "添加到快捷方式",
	},
	{
		Line: 1,
	},
	{
		Text:  "演示",
		Value: "演示",
	},
	{
		Text:  "导出笔记",
		Value: "导出笔记",
	},
}

func main() {
	var th = theme.NewTheme()
	var clickable widget.Clickable
	card := widgets.NewCard(th)
	rightMenu := widgets.NewRightMenu(th)
	rightMenu.SetOptions(options)
	rightMenu.Clicked(func(key int, menu string) {
		fmt.Println(key, menu)
	})
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return rightMenu.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return widgets.Label(th, "right menu").Layout(gtx)
						})
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
	})
	win.Run()
}
