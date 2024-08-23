/**
 * Created by Goland
 * @file   tips.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/20 23:29
 * @desc   tips.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioui-plugins/theme"
	"image"
	"sync"
)

type Tips struct {
	sync.Mutex
	visible bool
	message string
	theme   *theme.Theme
}

func NewTips(th *theme.Theme) *Tips {
	return &Tips{
		theme: th,
	}
}

func (t *Tips) Message(message string) {
	t.message = message
	t.visible = true
}
func (t *Tips) Close() {
	t.visible = false
}

func (t *Tips) Visible() bool {
	return t.visible
}

func (t *Tips) Layout(gtx layout.Context) layout.Dimensions {
	if !t.visible {
		return layout.Dimensions{}
	}
	if t.visible {
		// 绘制全屏半透明遮罩层
		paint.Fill(gtx.Ops, t.theme.Color.DefaultMaskBgColor)
	}
	// Define the size of the Tips
	width := gtx.Dp(250)
	height := gtx.Dp(50)
	return new(widget.Clickable).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Top: unit.Dp(0),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			// Center the Tips in the parent container
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Set the size of the Tips
				gtx.Constraints = layout.Exact(image.Point{X: width, Y: height})
				// 绘制全屏半透明遮罩层
				rc := clip.UniformRRect(image.Rectangle{Max: image.Point{
					X: gtx.Constraints.Min.X,
					Y: gtx.Constraints.Min.Y,
				}}, gtx.Dp(t.theme.Size.DefaultWidgetRadiusSize))
				paint.FillShape(gtx.Ops, t.theme.Color.ActionTipsBgGrayColor, rc.Op(gtx.Ops))
				// Center the text inside the Tips
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {

					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Max.X = gtx.Dp(17)
							loader := material.Loader(t.theme.Material())
							loader.Color = t.theme.Color.GreenColor
							return loader.Layout(gtx)
						}),
						layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return Label(t.theme, t.message).Layout(gtx)
						}),
					)

					// l := material.Label(t.theme.Material(), unit.Sp(16), "")
					// l.Font.Typeface = "MaterialIcons"
					// l.Color = t.theme.TextColor
					// l.Text = t.message
					// return l.Layout(gtx)
				})
			})
		})
	})
}
