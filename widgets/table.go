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
	"gioui.org/unit"
	"gioui.org/widget"
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
	headerFun   layout.ListElement
	dataFun     outlay.Cell
	header      []string
	data        []map[string]any
	dataContent []widget.Bool
	keys        []string
}

func NewTable(th *theme.Theme) *Table {
	table := &Table{
		theme:  th,
		height: unit.Dp(30),
	}
	return table
}

func (t *Table) SetHeader(header []string) *Table {
	t.header = header
	return t
}
func (t *Table) SetData(data []map[string]any) *Table {
	t.data = data
	for range data {
		t.dataContent = append(t.dataContent, widget.Bool{})
	}
	t.keys = maps.Keys(t.data[0])
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

func (t *Table) LayoutTable(gtx layout.Context) D {
	if len(t.data) == 0 {
		return layout.Dimensions{}
	}
	inset := layout.UniformInset(unit.Dp(2))
	orig := gtx.Constraints
	gtx.Constraints.Min = image.Point{}
	macro := op.Record(gtx.Ops)
	dims := inset.Layout(gtx, layout.Spacer{Height: unit.Dp(30)}.Layout)
	_ = macro.Stop()
	gtx.Constraints = orig
	if t.headerFun == nil {
		t.headerFun = func(gtx layout.Context, index int) layout.Dimensions {
			t.drawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.TableHeaderBgColor)
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return Label(t.theme, t.header[index], true).Layout(gtx)
			})
		}
	}
	if t.dataFun == nil {
		t.dataFun = func(gtx layout.Context, row, col int) layout.Dimensions {
			return t.dataContent[row].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if t.dataContent[row].Hovered() {
					t.drawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.DefaultContentBgGrayColor)
				} else {
					t.drawBackground(gtx, layout.Spacer{}.Layout(gtx).Size, t.theme.Color.DefaultWindowBgGrayColor)
				}
				dims := layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return Label(t.theme, fmt.Sprint(t.data[row][t.keys[col]])).Layout(gtx)
				})
				NewLine(t.theme).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
				return dims
			})
		}
	}
	return component.Table(t.theme.Material(), &t.grid).Layout(gtx, len(t.data), len(t.data[0]),
		func(axis layout.Axis, index, constraint int) int {
			switch axis {
			case layout.Horizontal:
				return constraint / len(t.header)
			default:
				return dims.Size.Y
			}
		},
		t.headerFun,
		t.dataFun,
	)
}

// drawBackground 在给定的尺寸上绘制一个背景颜色
func (t *Table) drawBackground(gtx layout.Context, size image.Point, col color.NRGBA) {
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, col)
}
