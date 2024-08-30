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
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"
	"github.com/x-module/gioui-plugins/theme"
	"golang.org/x/exp/maps"
	"image"
	"image/color"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Table struct {
	theme       *theme.Theme
	height      unit.Dp
	grid        component.GridState
	sizeFun     outlay.Dimensioner
	headerFun   layout.ListElement
	dataFun     outlay.Cell
	header      []string
	data        []map[string]any
	dataContent []widget.Bool
	dataBgColor color.NRGBA
}

func NewTable(th *theme.Theme) *Table {
	table := &Table{
		theme:  th,
		height: unit.Dp(30),
	}
	return table
}

func (g *Table) SetHeader(header []string) *Table {
	g.header = header
	return g
}
func (g *Table) SetData(data []map[string]any) *Table {
	g.data = data
	for range data {
		g.dataContent = append(g.dataContent, widget.Bool{})
	}
	return g
}

func (g *Table) SetSizeFun(sizeFun outlay.Dimensioner) *Table {
	g.sizeFun = sizeFun
	return g
}
func (g *Table) SetHeaderFun(headerFun layout.ListElement) *Table {
	g.headerFun = headerFun
	return g
}
func (g *Table) SetDataFun(dataFun outlay.Cell) *Table {
	g.dataFun = dataFun
	return g
}

func (g *Table) LayoutTable(gtx layout.Context) D {
	// Configure width based on available space and a minimum size.
	// minSize := gtx.Dp(unit.Dp(200))

	inset := layout.UniformInset(unit.Dp(2))

	// Configure a label styled to be a heading.
	headingLabel := material.Body1(g.theme.Material(), "")
	headingLabel.Font.Weight = font.Bold
	headingLabel.Alignment = text.Middle
	headingLabel.MaxLines = 1

	// Configure a label styled to be a data element.
	dataLabel := material.Body1(g.theme.Material(), "")
	dataLabel.Font.Typeface = "Go Mono"
	dataLabel.MaxLines = 1
	dataLabel.Alignment = text.End

	// Measure the height of a heading row.
	orig := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	macro := op.Record(gtx.Ops)
	dims := inset.Layout(gtx, headingLabel.Layout)
	_ = macro.Stop()
	gtx.Constraints = orig
	keys := maps.Keys(g.data[0])

	return component.Table(g.theme.Material(), &g.grid).Layout(gtx, len(g.data), len(g.data[0]),
		func(axis layout.Axis, index, constraint int) int {
			switch axis {
			case layout.Horizontal:
				return constraint / len(g.header)
			default:
				return dims.Size.Y
			}
		},
		func(gtx layout.Context, index int) layout.Dimensions {
			drawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, g.theme.Color.TableHeaderBgColor)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return inset.Layout(gtx, func(gtx C) D {
					return Label(g.theme, g.header[index], true).Layout(gtx)
				})
			})
		},
		func(gtx C, row, col int) D {
			return inset.Layout(gtx, func(gtx C) D {
				dims1 := layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return Label(g.theme, fmt.Sprint(g.data[row][keys[col]])).Layout(gtx)
				})
				NewLine(g.theme).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
				return dims1
			})
		},
	)
}

func (g *Table) Layout(gtx layout.Context) layout.Dimensions {
	g.check(gtx)
	if len(g.data) == 0 {
		return layout.Dimensions{}
	}
	table := component.Table(g.theme.Material(), &g.grid)
	return table.Layout(gtx, len(g.data), len(g.data[0]),
		g.sizeFun,
		g.headerFun,
		g.dataFun,
	)
}

func (g *Table) check(gtx layout.Context) {
	keys := maps.Keys(g.data[0])
	if g.header == nil {
		g.header = keys
	}

	if g.sizeFun == nil {
		g.sizeFun = func(axis layout.Axis, index, constraint int) int {
			switch axis {
			case layout.Horizontal:
				return constraint / len(g.header)
			default:
				return 100
			}
		}
	}
	if g.headerFun == nil {
		g.headerFun = func(gtx layout.Context, index int) layout.Dimensions {
			drawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, g.theme.Color.TableHeaderBgColor)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return Label(g.theme, g.header[index], true).Layout(gtx)
			})
		}
	}
	if g.dataFun == nil {
		g.dataFun = func(gtx layout.Context, row, col int) layout.Dimensions {
			return g.dataContent[row].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if g.dataContent[row].Hovered() {
					drawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, g.theme.Color.DefaultContentBgGrayColor)
					g.dataBgColor = g.theme.Color.DefaultContentBgGrayColor
				} else {
					drawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, g.theme.Color.DefaultWindowBgGrayColor)
					g.dataBgColor = g.theme.Color.DefaultWindowBgGrayColor
				}
				g.dataBgColor = g.theme.Color.DefaultWindowBgGrayColor
				dims := layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return Label(g.theme, fmt.Sprint(g.data[row][keys[col]])).Layout(gtx)
				})
				NewLine(g.theme).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
				return dims
			})
		}
	}
}

// drawBackground 在给定的尺寸上绘制一个背景颜色
func drawBackground(gtx layout.Context, size image.Point, col color.NRGBA) {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, col)
}
