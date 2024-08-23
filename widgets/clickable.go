package widgets

import (
	"gioui.org/io/input"
	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/utils"
	"image"
	"image/color"
)

type Clickable struct {
	theme        *theme.Theme
	bgColor      color.NRGBA
	bgColorHover color.NRGBA
	clickable    widget.Clickable
	onClick      func()
	widget       layout.Widget
}

func NewClickable(th *theme.Theme) *Clickable {
	return &Clickable{
		theme: th,
	}
}

func (c *Clickable) SetBgColor(bgColor color.NRGBA) {
	c.bgColor = bgColor
}
func (c *Clickable) SetBgColorHover(bgColorHover color.NRGBA) {
	c.bgColorHover = bgColorHover
}
func (c *Clickable) SetWidget(widget layout.Widget) *Clickable {
	c.widget = widget
	return c
}

func (c *Clickable) SetOnClick(onClick func()) {
	c.onClick = onClick
}

func (c *Clickable) Layout(gtx layout.Context) layout.Dimensions {
	if c.bgColor == (color.NRGBA{}) {
		c.bgColor = c.theme.Color.DefaultBgGrayColor
	}
	if c.bgColorHover == (color.NRGBA{}) {
		c.bgColorHover = c.theme.Color.HoveredBorderBlueColor
	}
	for c.clickable.Clicked(gtx) {
		if c.onClick != nil {
			c.onClick()
		}
	}
	return c.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rect := clip.UniformRRect(image.Rectangle{Max: image.Point{
					X: gtx.Constraints.Min.X,
					Y: gtx.Constraints.Min.Y,
				}}, gtx.Dp(c.theme.Size.DefaultElementRadiusSize))
				defer rect.Push(gtx.Ops).Pop()
				if gtx.Source == (input.Source{}) {
					paint.Fill(gtx.Ops, utils.Disabled(c.bgColorHover))
				} else if c.clickable.Hovered() {
					// paint.Fill(gtx.Ops, c.bgColorHover)
				}
				if gtx.Focused(c.clickable) {
					paint.Fill(gtx.Ops, c.bgColorHover)
				}
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			c.widget,
		)
	})
}