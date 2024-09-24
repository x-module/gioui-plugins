/**
 * Created by Goland
 * @file   Tree.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/14 22:57
 * @desc   Tree.go
 */

package widgets

import (
	"fmt"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"image"
)

type ClickAction func(gtx layout.Context, node *TreeNode)
type Tree struct {
	theme       *theme.Theme
	nodes       []*TreeNode
	width       unit.Dp
	clickedNode *TreeNode
	click       ClickAction
}

func NewTree(th *theme.Theme) *Tree {
	return &Tree{
		theme: th,
		width: unit.Dp(200),
	}
}

func (t *Tree) OnClick(fun ClickAction) *Tree {
	t.click = fun
	return t
}
func (t *Tree) SetWidth(width unit.Dp) *Tree {
	t.width = width
	return t
}

func (t *Tree) SetNodes(nodes []*TreeNode) *Tree {
	for _, node := range nodes {
		t.setClick(node)
	}
	t.setPath(nodes, []int{})
	t.nodes = nodes
	return t
}

func (t *Tree) AddTopNode(newNode *TreeNode) error {
	t.setClick(newNode)
	t.nodes = append(t.nodes, newNode)
	t.setPath(t.nodes, []int{})
	return nil
}

func (t *Tree) AddSonNode(newNode *TreeNode) error {
	if t.clickedNode == nil {
		return fmt.Errorf("no node selected")
	}
	t.setClick(newNode)
	path := t.clickedNode.Path
	parentNode, err := t.getNode(t.nodes, path)
	if err != nil {
		return err
	}
	parentNode.Children = append(parentNode.Children, newNode)
	t.setPath(t.nodes, []int{})
	return nil
}

func (t *Tree) getNode(nodes []*TreeNode, paths []int) (*TreeNode, error) {
	if nodes == nil {
		nodes = t.nodes
	}
	for _, path := range paths {
		if len(nodes) <= path {
			return nil, fmt.Errorf("err path")
		}
		if len(paths) > 1 {
			if nodes[path].Children != nil {
				return t.getNode(nodes[path].Children, paths[1:])
			}
			return nil, fmt.Errorf("err path")
		}
		return nodes[path], nil
	}
	return nil, fmt.Errorf("err path")
}

func (t *Tree) setPath(nodes []*TreeNode, path []int) {
	if nodes == nil {
		nodes = t.nodes
	}
	for key, node := range nodes {
		sign := []int{}
		if len(path) == 0 {
			sign = []int{key}
		} else {
			sign = append(path, key)
		}
		node.Path = sign
		if len(node.Children) > 0 {
			t.setPath(node.Children, sign)
		}
	}
}

func (t *Tree) setClick(nodes *TreeNode) {
	nodes.clickable = &widget.Clickable{}
	if len(nodes.Children) > 0 {
		for _, child := range nodes.Children {
			t.setClick(child)
		}
	}
}

type CallbackFun func(gtx layout.Context)

type TreeNode struct {
	Text          string
	Icon          *widget.Icon
	Children      []*TreeNode
	expanded      bool
	selected      bool
	clickable     *widget.Clickable
	ClickCallback CallbackFun
	Path          []int
}

func (t *Tree) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// 这里可以添加头部或者其他固定的内容
			return layout.Dimensions{}
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, t.renderTree(gtx, t.nodes)...)
		}),
	)
}

func (t *Tree) renderTree(gtx layout.Context, nodes []*TreeNode) []layout.FlexChild {
	if len(nodes) == 0 {
		return []layout.FlexChild{}
	}
	var dims []layout.FlexChild
	for _, node := range nodes {
		dims = append(dims, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.renderNode(gtx, node, 0, true)
		}))
	}
	return dims
}

func (t *Tree) renderNode(gtx layout.Context, node *TreeNode, loop int, isParent bool) layout.Dimensions {
	// 渲节点标题
	bgColor := t.theme.Color.CardBgColor

	if node.clickable.Clicked(gtx) {
		node.expanded = !node.expanded
		t.clickedNode = node
		if node.ClickCallback != nil {
			node.ClickCallback(gtx)
		}
		if t.click != nil {
			t.click(gtx, node)
		}
	}
	if node.clickable.Hovered() {
		bgColor = t.theme.Color.TreeHoveredBgColor
	}
	var sonItems []layout.FlexChild
	// 绘制展开/折叠图标
	if len(node.Children) > 0 {
		if node.expanded {
			sonItems = append(sonItems, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return node.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Top: unit.Dp(5)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Max.X = gtx.Dp(t.theme.Size.DefaultIconSize)
						return resource.ArrowDownIcon.Layout(gtx, t.theme.Color.TreeIconColor)
					})
				})
			}))
		} else {
			sonItems = append(sonItems, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return node.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Top: unit.Dp(5)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Max.X = gtx.Dp(t.theme.Size.DefaultIconSize)
						return resource.ArrowRightIcon.Layout(gtx, t.theme.Color.TreeIconColor)
					})
				})
			}))
		}
	}
	sonItems = append(sonItems, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Min.X = gtx.Dp(t.width)
		return node.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(6)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return Label(t.theme, node.Text).Layout(gtx)
			})
		})
	}))

	items := []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Background{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// defer clip.Rect(image.Rectangle{Max: gtx.Constraints.Max}, gtx.Dp(t.theme.Size.DefaultElementRadiusSize)).Push(gtx.Ops).Pop()
				if t.clickedNode == node {
					bgColor = t.theme.Color.TreeClickedBgColor
				}
				defer clip.Rect{
					Max: image.Point{
						X: gtx.Constraints.Max.X,
						Y: gtx.Constraints.Min.Y,
					},
				}.Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, bgColor)
				return layout.Dimensions{Size: gtx.Constraints.Min}
			}, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(25))
				return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Left: unit.Dp(loop * 20)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, sonItems...)
					})
				})
			})
		}),
	}
	// 递归渲染子节点
	if node.expanded && len(node.Children) > 0 {
		var dims []layout.FlexChild
		for _, child := range node.Children {
			dims = append(dims, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				level := loop + 1
				return t.renderNode(gtx, child, level, false)
			}))
		}
		items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx, dims...)
			// })
		}))
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx, items...)
}

func (t *Tree) GetCurrentNode() *TreeNode {
	return t.clickedNode
}
func (t *Tree) MinTree(gtx layout.Context, nodes []*TreeNode) {
	if nodes == nil {
		nodes = t.nodes
	}
	for _, node := range nodes {
		if node.expanded {
			node.expanded = false
		}
		if len(node.Children) > 0 {
			t.MinTree(gtx, node.Children)
		}
	}
	gtx.Execute(op.InvalidateCmd{})
}
