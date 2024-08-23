/**
 * Created by Goland
 * @file   Confirm.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/22 18:56
 * @desc   Confirm.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/gioui-plugins/resource"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/utils"
	"image"
)

type Confirm struct {
	theme            *theme.Theme
	title            string
	message          string
	height           int
	width            int
	visible          bool
	cancelClickable  widget.Clickable
	confirmClickable widget.Clickable
	clickerWidget    *Clickable
	cancelFunc       func()
	confirmFunc      func()
}

func NewConfirm(th *theme.Theme) *Confirm {
	modal := &Confirm{
		theme:         th,
		height:        150,
		width:         300,
		title:         "操作确认",
		clickerWidget: NewClickable(th),
	}
	return modal
}

func (c *Confirm) Confirm(fun func()) {
	c.confirmFunc = fun
}
func (c *Confirm) Cancel(fun func()) {
	c.cancelFunc = fun
}

func (c *Confirm) SetWidth(width int) {
	c.width = width
}
func (c *Confirm) Visible() bool {
	return c.visible
}
func (c *Confirm) SetTitle(title string) {
	c.title = title
}

func (c *Confirm) SetHeight(height int) {
	c.height = height
}

func (c *Confirm) Message(message string) {
	c.message = message
	c.visible = true
}
func (c *Confirm) Close() {
	c.visible = false
}

func (c *Confirm) Layout(gtx layout.Context) layout.Dimensions {
	if !c.visible {
		return layout.Dimensions{}
	}
	if c.visible {
		// 绘制全屏半透明遮罩层
		paint.Fill(gtx.Ops, c.theme.Color.DefaultMaskBgColor)
	}
	for c.cancelClickable.Clicked(gtx) {
		c.visible = false
		if c.cancelFunc != nil {
			c.cancelFunc()
		}
	}
	for c.confirmClickable.Clicked(gtx) {
		c.visible = false
		if c.confirmFunc != nil {
			c.confirmFunc()
		}
	}
	width := gtx.Dp(unit.Dp(c.width))
	height := gtx.Dp(unit.Dp(c.height))
	return c.clickerWidget.SetWidget(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{
			Top: unit.Dp(10),
		}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				// Set the size of the Confirm
				gtx.Constraints = layout.Exact(image.Point{X: width, Y: height})
				rc := clip.RRect{
					Rect: image.Rectangle{Max: image.Point{
						X: gtx.Constraints.Min.X,
						Y: gtx.Constraints.Min.Y,
					}},
					NW: 10, NE: 10, SE: 10, SW: 10,
				}
				paint.FillShape(gtx.Ops, c.theme.Color.DefaultContentBgGrayColor, rc.Op(gtx.Ops))
				// Center the text inside the Confirm
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: 0, Right: 10, Bottom: 10, Top: 10}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										gtx.Constraints.Max.X = gtx.Dp(unit.Dp(20))
										return resource.ActionInfoOutlineIcon.Layout(gtx, c.theme.Color.GreenColor)
									}),
									layout.Rigid(layout.Spacer{Width: unit.Dp(5)}.Layout),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return Label(c.theme, c.title).Layout(gtx)
									}),
								)

							})
						})
					}),
					utils.DrawLineFlex(c.theme.Color.DefaultLineColor, unit.Dp(1), unit.Dp(c.width)),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(50))
						return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: 5, Right: 5, Bottom: 2, Top: 2}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return Label(c.theme, c.message).Layout(gtx)
							})
						})
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.UniformInset(unit.Dp(15)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										but := DefaultButton(c.theme, &c.cancelClickable, "取消", unit.Dp(70), layout.Inset{
											Top: 3, Bottom: 3,
											Left: 5, Right: 5,
										})
										return but.Layout(gtx)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Dimensions{Size: image.Point{X: 20, Y: 0}}
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										but := SuccessButton(c.theme, &c.confirmClickable, "确认", unit.Dp(70), layout.Inset{
											Top: 3, Bottom: 3,
											Left: 5, Right: 5,
										})
										return but.Layout(gtx)
									}),
								)
							})
						})

					}),
				)
			})
		})
	}).Layout(gtx)
}
