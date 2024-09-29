/**
 * Created by Goland
 * @file   card.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/22 11:54
 * @desc   内容展示的卡片
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
	"image"
	"image/color"
)

type Card struct {
	theme   *theme.Theme
	radius  int
	padding unit.Dp
	bgColor color.NRGBA
}

func NewCard(theme *theme.Theme) *Card {
	return &Card{
		theme:   theme,
		radius:  15,
		padding: unit.Dp(20),
		bgColor: theme.Color.CardBgColor,
	}
}

func (c *Card) SetRadius(radius int) *Card {
	c.radius = radius
	return c
}

func (c *Card) SetBgColor(bgColor color.NRGBA) *Card {
	c.bgColor = bgColor
	return c
}
func (c *Card) SetPadding(padding unit.Dp) *Card {
	c.padding = padding
	return c
}

func fill(gtx layout.Context, color color.NRGBA) layout.Dimensions {
	cs := gtx.Constraints
	d := image.Point{X: cs.Max.X, Y: cs.Min.Y}
	track := image.Rectangle{
		Max: d,
	}
	defer clip.Rect(track).Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color)
	return layout.Dimensions{Size: d}
}

func (c *Card) Layout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			rect := clip.UniformRRect(image.Rectangle{Max: image.Point{
				X: gtx.Constraints.Max.X,
				Y: gtx.Constraints.Min.Y,
			}}, c.radius)
			defer rect.Push(gtx.Ops).Pop()
			return fill(gtx, c.bgColor)
		}),
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(c.padding).Layout(gtx, widget)
		}),
	)
}
