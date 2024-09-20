package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
	"image/color"
	"log"
)

var th = theme.NewTheme()

func init() {
	th.Color.CardBgColor = color.NRGBA{R: 44, G: 44, B: 44, A: 255}
	th.Color.DefaultWindowBgGrayColor = color.NRGBA{R: 30, G: 30, B: 30, A: 255}
}

func main() {
	card := widgets.NewCard(th)
	editor := widget.Editor{
		Alignment:  text.Start,
		SingleLine: true,
	}
	editor1 := widget.Editor{
		Alignment:  text.Start,
		SingleLine: true,
	}
	editor.SetText("### title")
	editor1.SetText("### 22222")
	card.SetRadius(0)
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						ed := material.Editor(th.Material(), &editor, "editor")
						// processKey(gtx, &editor)
						if gtx.Focused(&editor) {
							log.Println("editor focused")
						} else {
							log.Println("editor lost focus")
						}
						return ed.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						ed := material.Editor(th.Material(), &editor1, "editor")
						return ed.Layout(gtx)
					}),
				)
			})
		})
	})
	win.Run()
}
