package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"

	"fyne.io/fyne/v2"
	fapp "fyne.io/fyne/v2/app"
	fwidget "fyne.io/fyne/v2/widget"
)

var myApp = fapp.New()
var myWindow = myApp.NewWindow("Menu Example")

func menu() {
	myWindow.Resize(fyne.NewSize(0, 0))

	// 创建菜单项
	newItem := fyne.NewMenuItem("New", func() {
		// 这里填写点击"New"时希望执行的动作
		fwidget.ShowPopUp(fwidget.NewLabel("New item clicked"), myWindow.Canvas())
	})
	// 更多菜单项...
	fileMenu := fyne.NewMenu("File", newItem) // "File"是顶层菜单标题

	// 创建顶部菜单栏并添加菜单项
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		// 可以添加更多的菜单...
	)

	// 将菜单栏设置到窗口
	myWindow.SetMainMenu(mainMenu)

	// 设置窗口内容
	// myWindow.SetContent(container.NewVBox(
	// 	fwidget.NewLabel("Hello Fyne!"),
	// 	// 添加更多的widget...
	// ))

	myWindow.Show()
}
func main() {
	go menu()
	var clickable widget.Clickable

	var th = theme.NewTheme()
	card := widgets.NewCard(th)
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
	})
	go myApp.Run()
	go win.Run()
	select {}
}
