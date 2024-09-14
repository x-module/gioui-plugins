package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// TreeNode 表示树中的一个节点
type TreeNode struct {
	text       string
	children   []*TreeNode
	isExpanded bool
}

// TreeView 用于渲染树形结构
type TreeView struct {
	theme     *material.Theme
	rootNodes []*TreeNode
}

func (tv *TreeView) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Vertical{}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			// 这里可以添加头部或者其他固定的内容
			return layout.Dimensions{}
		}),
		layout.Flexed(1, func(gtx C) D {
			return tv.renderTree(gtx, tv.rootNodes)
		}),
	)
}

func (tv *TreeView) renderTree(gtx layout.Context, nodes []*TreeNode) layout.Dimensions {
	if len(nodes) == 0 {
		return layout.Dimensions{}
	}

	var dims layout.Dimensions
	for _, node := range nodes {
		dims = layout.Flex{}.Layout(gtx, 1, func(gtx C) D {
			return tv.renderNode(gtx, node)
		})
	}

	return dims
}

func (tv *TreeView) renderNode(gtx layout.Context, node *TreeNode) layout.Dimensions {
	if node == nil {
		return layout.Dimensions{}
	}

	// 创建一个 Clickable 对象来处理点击事件
	clickable := widget.Clickable{}

	// 渲节点标题
	clickable.Do(gtx, func(gtx C) D {
		return layout.Inset{
			Top:    unit.Dp(8),
			Bottom: unit.Dp(8),
			Left:   unit.Dp(16),
			Right:  unit.Dp(16),
		}.Layout(gtx, func(gtx C) D {
			return tv.theme.Text.Hint.Layout(gtx, node.text)
		})
	})

	// 检查是否点击了节点
	if clickable.Clicked() {
		node.isExpanded = !node.isExpanded
	}

	// 绘制展开/折叠图标
	if len(node.children) > 0 {
		if node.isExpanded {
			tv.theme.Icon.ChevronDown.Layout(gtx, clip.Rect{
				Max: gtx.Px(unit.Dp(24)),
			}.Op())
		} else {
			tv.theme.Icon.ChevronRight.Layout(gtx, clip.Rect{
				Max: gtx.Px(unit.Dp(24)),
			}.Op())
		}
	}

	// 递归渲染子节点
	if node.isExpanded {
		var childDims layout.Dimensions
		for _, child := range node.children {
			childDims = layout.Flex{}.Layout(gtx, 1, func(gtx C) D {
				return tv.renderNode(gtx, child)
			})
		}
		dims = childDims
	}

	return dims
}

func main() {
	go func() {
		w := app.NewWindow(app.Title("Tree View Example"))
		tv := &TreeView{
			theme: material.NewTheme(&material.Style{
				Font: widget.Font{Face: material.Monospace},
			}),
			rootNodes: []*TreeNode{
				{
					text: "Root Node 1",
					children: []*TreeNode{
						{
							text: "Child Node 1.1",
							children: []*TreeNode{
								{text: "Grandchild Node 1.1.1"},
								{text: "Grandchild Node 1.1.2"},
							},
						},
						{
							text: "Child Node 1.2",
							children: []*TreeNode{
								{text: "Grandchild Node 1.2.1"},
							},
						},
					},
				},
				{
					text: "Root Node 2",
					children: []*TreeNode{
						{
							text: "Child Node 2.1",
							children: []*TreeNode{
								{text: "Grandchild Node 2.1.1"},
							},
						},
					},
				},
			},
		}
		for {
			gtx := layout.NewContext(&w.Events())
			dims := tv.Layout(gtx)
			op.InvalidateOp{}.Add(gtx.Ops)
			w.MainLoop()
		}
	}()

	app.Main()
}
