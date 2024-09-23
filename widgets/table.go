/**
 * Created by Goland
 * @file   grid.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/30 17:18
 * @desc   grid.go
 */

package widgets

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	text2 "gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/utils"
	"image"
)

type Table struct {
	theme       *theme.Theme
	height      unit.Dp
	grid        component.GridState
	headerFun   layout.ListElement
	dataFun     outlay.Cell
	headers     []string
	data        [][]any
	dataContent []widget.Bool
	card        *Card
}

func NewTable(th *theme.Theme) *Table {
	table := &Table{
		theme:  th,
		height: unit.Dp(30),
		card:   NewCard(th).SetPadding(0).SetRadius(0),
	}
	return table
}

func (t *Table) SetHeader(header []string) *Table {
	t.headers = header
	return t
}
func (t *Table) SetData(data [][]any) *Table {
	t.data = data
	for range data {
		t.dataContent = append(t.dataContent, widget.Bool{})
	}
	return t
}

func (t *Table) SetHeaderFun(headerFun layout.ListElement) *Table {
	t.headerFun = headerFun
	return t
}
func (t *Table) SetDataFun(dataFun outlay.Cell) *Table {
	t.dataFun = dataFun
	return t
}

func (t *Table) LayoutHoverTable(gtx layout.Context) layout.Dimensions {
	if len(t.data) == 0 {
		return layout.Dimensions{}
	}
	inset := layout.UniformInset(unit.Dp(2))
	orig := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	macro := op.Record(gtx.Ops)
	dims := inset.Layout(gtx, layout.Spacer{Height: t.height}.Layout)
	_ = macro.Stop()
	gtx.Constraints = orig
	if t.headerFun == nil {
		t.headerFun = func(gtx layout.Context, index int) layout.Dimensions {
			utils.DrawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.TableHeaderBgColor)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return Label(t.theme, t.headers[index], true).Layout(gtx)
			})
		}
	}
	if t.dataFun == nil {
		t.dataFun = func(gtx layout.Context, row, col int) layout.Dimensions {
			return t.dataContent[row].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if t.dataContent[row].Hovered() {
					utils.DrawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.DefaultContentBgGrayColor)
				} else {
					utils.DrawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.DefaultWindowBgGrayColor)
				}
				NewLine(t.theme).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
				labelDims := layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return Label(t.theme, fmt.Sprint(t.data[row][col])).Layout(gtx)
				})
				return labelDims
			})
		}
	}
	return component.Table(t.theme.Material(), &t.grid).Layout(gtx, len(t.data), len(t.data[0]),
		func(axis layout.Axis, index, constraint int) int {
			switch axis {
			case layout.Horizontal:
				return constraint / len(t.headers)
			default:
				return dims.Size.Y
			}
		},
		t.headerFun,
		t.dataFun,
	)
}

func (t *Table) LayoutBorderTable(gtx layout.Context) layout.Dimensions {
	return widget.Border{
		Color: t.theme.Color.BorderLightGrayColor,
		Width: unit.Dp(1),
	}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		table := component.Table(t.theme.Material(), &t.grid)
		table.AnchorStrategy = material.Overlay
		return table.Layout(gtx, len(t.data), len(t.headers),
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Horizontal:
					return constraint / len(t.headers)
				default:
					return 100
				}
			},
			func(gtx layout.Context, index int) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return NewLayout().CenterLayout(gtx, func(gtx layout.Context) layout.Dimensions {
									return layout.UniformInset(unit.Dp(1)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
										return t.card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											return NewLayout().CenterLayout(gtx, func(gtx layout.Context) layout.Dimensions {
												return Label(t.theme, t.headers[index], true).Layout(gtx)
											})
										})
									})
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								t.drawHorizontalLine(gtx)
								return layout.Dimensions{}
							}),
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						t.drawVerticalLine(gtx)
						return layout.Dimensions{}
					}),
				)
			},
			func(gtx layout.Context, row, col int) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return NewLayout().CenterLayout(gtx, func(gtx layout.Context) layout.Dimensions {
									label := material.Label(t.theme.Theme, unit.Sp(15), fmt.Sprint(t.data[row][col]))
									label.Alignment = text2.Middle
									label.Color = t.theme.Color.DefaultTextWhiteColor
									label.TextSize = t.theme.Size.DefaultTextSize
									return label.Layout(gtx)
								})
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								t.drawHorizontalLine(gtx)
								return layout.Dimensions{}
							}),
						)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						t.drawVerticalLine(gtx)
						return layout.Dimensions{}
					}),
				)
			},
		)
	})

}

func (t *Table) drawHorizontalLine(gtx layout.Context) {
	// 定义线条宽度和长度
	lineWidth := unit.Dp(1)
	lineLength := gtx.Constraints.Max.X
	// 创建一个矩形作为线条
	line := clip.Rect{
		Max: image.Point{X: lineLength, Y: gtx.Dp(lineWidth)},
	}.Push(gtx.Ops)
	// 设置线条颜色
	paint.ColorOp{Color: t.theme.Color.BorderLightGrayColor}.Add(gtx.Ops)
	// 填充矩形以绘制线条
	paint.PaintOp{}.Add(gtx.Ops)
	// 完成绘制
	line.Pop()
}
func (t *Table) drawVerticalLine(gtx layout.Context) {
	// 定义线条宽度和长度
	lineWidth := gtx.Constraints.Max.Y
	lineLength := unit.Dp(1)
	// 创建一个矩形作为线条
	line := clip.Rect{
		Max: image.Point{X: int(lineLength), Y: lineWidth},
	}.Push(gtx.Ops)
	// 设置线条颜色
	paint.ColorOp{Color: t.theme.Color.BorderLightGrayColor}.Add(gtx.Ops)
	// 填充矩形以绘制线条
	paint.PaintOp{}.Add(gtx.Ops)
	// 完成绘制
	line.Pop()
}
