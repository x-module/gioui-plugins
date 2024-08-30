package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	th := theme.NewTheme()

	accountType := &widget.Enum{
		Value: "old",
	}

	newAccount := widgets.NewRadioButton(th, accountType, "new", "新用户")
	oldAccount := widgets.NewRadioButton(th, accountType, "old", "老用户")

	oldAccount.SetSize(th.Size.Large)
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.Stack{Alignment: layout.Center}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return newAccount.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: unit.Dp(20)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return oldAccount.Layout(gtx)
					}),
				)
			}),
		)
	})
	win.Run()
}
