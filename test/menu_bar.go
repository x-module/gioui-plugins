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
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	menus := []widgets.MenuBarItem{
		{
			Title: "File",
			Items: []widgets.MenuBarItemElement{
				{
					Name: "New",
					Action: func() {
						println("New")
					},
				},
				{
					Name: "Open",
					Action: func() {
						println("Open")
					},
				},
				{
					Name: "Save",
					Action: func() {
						println("Save")
					},
				},
				{
					Name: "Exit",
					Action: func() {
						println("Exit")
					},
				},
				{
					Name: "Close",
					Action: func() {
						println("Close")
					},
				},
			},
		},
		{
			Title: "Edit",
			Items: []widgets.MenuBarItemElement{
				{
					Name: "Copy",
					Action: func() {
						println("Copy")
					},
				},
			},
		},
		{
			Title: "Help",
			Items: []widgets.MenuBarItemElement{
				{
					Name: "About",
					Action: func() {
						println("About")
					},
				},
			},
		},
	}

	menuBar := widgets.NewMenuBar(th)
	menuBar.AddMenuBarItem(menus)

	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return menuBar.Layout(gtx)
		})
	})
	win.Run()
}
