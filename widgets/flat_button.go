package widgets

import (
	"gioui.org/io/input"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"github.com/gioui-plugins/theme"
	"github.com/gioui-plugins/utils"
	"image"
	"image/color"

	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
)

const (
	FlatButtonIconStart = 0
	FlatButtonIconEnd   = 1
	FlatButtonIconTop   = 2
	FlatButtonIconDown  = 3
)

type FlatButton struct {
	Icon         *widget.Icon
	IconPosition int
	SpaceBetween unit.Dp

	Clickable              *widget.Clickable
	theme                  *theme.Theme
	MinWidth               unit.Dp
	BackgroundColor        color.NRGBA
	HoveredBackgroundColor color.NRGBA
	ClickedBackgroundColor color.NRGBA
	TextColor              color.NRGBA
	Text                   string

	CornerRadius      int
	BackgroundPadding unit.Dp
	ContentPadding    unit.Dp

	MarginRight int
	MarginLeft  int
	IconSize    unit.Dp
}

func (f *FlatButton) SetIcon(icon *widget.Icon, position int, spaceBetween unit.Dp) {
	f.Icon = icon
	f.IconPosition = position
	f.SpaceBetween = spaceBetween
}

func (f *FlatButton) Layout(gtx layout.Context) layout.Dimensions {
	if f.BackgroundColor == (color.NRGBA{}) {
		f.BackgroundColor = f.theme.Color.DefaultContentBgGrayColor
	}

	if f.IconSize == 0 {
		f.IconSize = unit.Dp(20)
	}

	if f.TextColor == (color.NRGBA{}) {
		f.TextColor = f.theme.Color.DefaultTextWhiteColor
	}

	axis := layout.Horizontal
	labelLayout := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return layout.Inset{Right: unit.Dp(f.MarginRight)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			l := material.Label(f.theme.Material(), unit.Sp(11), f.Text)
			l.Color = f.TextColor
			l.TextSize = unit.Sp(13)
			// return Label(f.theme, f.Text).Layout(gtx)
			return l.Layout(gtx)
		})
	})

	widgets := []layout.FlexChild{labelLayout}

	if f.Icon != nil {
		iconLayout := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// return layout.UniformInset(f.SpaceBetween).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Left: unit.Dp(f.MarginLeft), Right: f.SpaceBetween, Top: f.SpaceBetween, Bottom: f.SpaceBetween}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = gtx.Dp(f.IconSize)
				return f.Icon.Layout(gtx, f.TextColor)
			})
		})

		if f.IconPosition == FlatButtonIconTop || f.IconPosition == FlatButtonIconDown {
			axis = layout.Vertical
		}

		switch f.IconPosition {
		case FlatButtonIconStart, FlatButtonIconTop:
			widgets = []layout.FlexChild{iconLayout, labelLayout}
		case FlatButtonIconEnd, FlatButtonIconDown:
			widgets = []layout.FlexChild{labelLayout, iconLayout}
		}
	}

	return f.Clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(f.BackgroundPadding).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			semantic.Button.Add(gtx.Ops)
			// gtx.Constraints.Min.Y = gtx.Dp(20)
			return layout.Background{}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min.X = gtx.Dp(f.MinWidth)
					defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, f.CornerRadius).Push(gtx.Ops).Pop()
					background := f.BackgroundColor
					if gtx.Source == (input.Source{}) {
						background = utils.Disabled(f.BackgroundColor)
					} else if f.Clickable.Hovered() {
						background = f.HoveredBackgroundColor
					} else if gtx.Focused(f.Clickable) {
						background = f.ClickedBackgroundColor
					}
					paint.Fill(gtx.Ops, background)
					return layout.Dimensions{Size: gtx.Constraints.Min}
				},
				func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(f.ContentPadding).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: axis, Alignment: layout.Middle, Spacing: layout.SpaceStart}.Layout(gtx, widgets...)
					})
				},
			)
		})
	})
}
