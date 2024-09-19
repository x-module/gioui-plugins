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
)

func main() {
	rootNodes := []*widgets.TreeNode{
		{
			Text: "Root 1",
			Children: []*widgets.TreeNode{
				{
					Text: "Child 1.1",
					Children: []*widgets.TreeNode{
						{Text: "Grandchild 1.1.1"},
						{Text: "Grandchild 1.1.2"},
					},
					ClickCallback: func(gtx layout.Context) {
						println("clicked")
					},
				},
				{
					Text: "Child 1.2",
					Children: []*widgets.TreeNode{
						{Text: "Grandchild 1.2.1"},
					},
				},
			},
		},
		{
			Text: "Root 2",
			Children: []*widgets.TreeNode{
				{
					Text: "Child 2.1",
					Children: []*widgets.TreeNode{
						{Text: "Grandchild 2.1.1"},
					},
				},
			},
		},
	}

	var th = theme.NewTheme()
	tree := widgets.NewTree(th)
	tree.SetNodes(rootNodes)
	tree.OnClick(func(gtx layout.Context, node *widgets.TreeNode) {
		fmt.Println("node:", node.Text, " clicked")
	})
	card := widgets.NewCard(th)
	card.SetPadding(0).SetRadius(0)
	win := window.NewApplication(new(app.Window)).CenterWindow()
	win.Title("Hello, Gio!").Size(window.ElementSize{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Max.X = gtx.Dp(200)
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return tree.Layout(gtx)
					})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Dimensions{}
				}),
			)
		})
	})
	win.Run()
}
