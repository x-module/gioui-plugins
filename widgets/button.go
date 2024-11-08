/**
 * Created by Goland
 * @file   button.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/23 16:25
 * @desc   Button
 */

package widgets

import (
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/utils"
	"image"
	"image/color"
	"math"

	"gioui.org/io/semantic"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type ButtonStyle struct {
	TextColor   color.NRGBA
	TextSize    unit.Sp
	BorderColor color.NRGBA
	BgColor     color.NRGBA
}

type Button struct {
	Text         string
	Icon         *widget.Icon
	IconPosition int
	theme        *theme.Theme
	// Color is the text color.
	Color        color.NRGBA
	TextSize     unit.Sp
	IconSize     unit.Dp
	Background   color.NRGBA
	CornerRadius unit.Dp
	Inset        layout.Inset
	Button       *widget.Clickable
	width        unit.Dp
	bdColor      color.NRGBA
	hovered      func(gtx layout.Context)

	defaultButtonStyle  ButtonStyle
	tertiaryButtonStyle ButtonStyle
	infoButtonStyle     ButtonStyle
	successButtonStyle  ButtonStyle
	errorButtonStyle    ButtonStyle
}

func (b *Button) SetTextSize(size unit.Sp) {
	b.TextSize = size
}

// set iconSize
func (b *Button) SetIconSize(size unit.Dp) {
	b.IconSize = size
}
func (b *Button) Hovered(f func(gtx layout.Context)) {
	b.hovered = f
}
func DefaultButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ButtonDefaultTextColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ButtonDefaultColor,
		TextSize:     th.Size.DefaultTextSize,
		IconSize:     th.Size.DefaultIconSize,
		bdColor:      th.Color.ButtonBorderColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func TertiaryButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ButtonTertiaryTextColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ButtonTertiaryBgColor,
		TextSize:     th.Size.DropdownTextSize,
		bdColor:      th.Color.ButtonBorderColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func PrimaryButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ButtonTextBlackColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.PrimaryColor,
		TextSize:     th.Size.DefaultTextSize,
		IconSize:     th.Size.DefaultIconSize,
		bdColor:      th.Color.PrimaryColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func PrimaryOutlineButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.PrimaryColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ButtonDefaultColor,
		TextSize:     th.Size.DefaultTextSize,
		IconSize:     th.Size.DefaultIconSize,
		bdColor:      th.Color.PrimaryColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func InfoButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ButtonTextBlackColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.InfoColor,
		TextSize:     th.Size.DefaultTextSize,
		IconSize:     th.Size.DefaultIconSize,
		bdColor:      th.Color.InfoColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func InfoOutlineButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.InfoColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ButtonDefaultColor,
		TextSize:     th.Size.DefaultTextSize,
		IconSize:     th.Size.DefaultIconSize,
		bdColor:      th.Color.InfoColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func SuccessButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ButtonTextBlackColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.SuccessColor,
		TextSize:     th.Size.DropdownTextSize,
		bdColor:      th.Color.SuccessColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func BlueButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.DefaultTextWhiteColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.BlueColor,
		TextSize:     th.Size.DropdownTextSize,
		bdColor:      th.Color.BlueColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func SuccessOutlineButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.SuccessColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ButtonDefaultColor,
		TextSize:     th.Size.DropdownTextSize,
		bdColor:      th.Color.SuccessColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func WarningButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ButtonTextBlackColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.WarningColor,
		TextSize:     th.Size.DefaultTextSize,
		IconSize:     th.Size.DefaultIconSize,
		bdColor:      th.Color.WarningColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func WarningOutlineButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.WarningColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ButtonDefaultColor,
		TextSize:     th.Size.DefaultTextSize,
		IconSize:     th.Size.DefaultIconSize,
		bdColor:      th.Color.WarningColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func ErrorButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ButtonTextBlackColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ErrorColor,
		TextSize:     th.Size.DropdownTextSize,
		bdColor:      th.Color.ErrorColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func ErrorOutlineButton(th *theme.Theme, button *widget.Clickable, txt string, width unit.Dp, inset ...layout.Inset) Button {
	b := Button{
		theme:        th,
		Text:         txt,
		Color:        th.Color.ErrorColor,
		CornerRadius: th.Size.DefaultElementRadiusSize,
		Background:   th.Color.ButtonDefaultColor,
		TextSize:     th.Size.DropdownTextSize,
		bdColor:      th.Color.ErrorColor,
		Inset: layout.Inset{
			Top: 8, Bottom: 8,
			Left: 8, Right: 8,
		},
		Button: button,
		width:  width,
	}
	if len(inset) > 0 {
		b.Inset = inset[0]
	}
	return b
}
func (b Button) SetIcon(icon *widget.Icon, iconPosition int) Button {
	b.Icon = icon
	b.IconPosition = iconPosition
	return b
}
func (b Button) Layout(gtx layout.Context) layout.Dimensions {
	minWidth := gtx.Constraints.Min
	if b.width > 0 {
		minWidth.X = gtx.Dp(b.width)
	}
	border := widget.Border{
		Color:        b.bdColor,
		Width:        unit.Dp(1),
		CornerRadius: b.theme.Size.DefaultElementRadiusSize,
	}

	return b.Button.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		semantic.Button.Add(gtx.Ops)
		if b.Button.Hovered() {
			border.Color = b.theme.Color.BorderBlueColor
			if b.hovered != nil {
				b.hovered(gtx)
			}
		}
		return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Background{}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, gtx.Dp(b.theme.Size.DefaultElementRadiusSize)).Push(gtx.Ops).Pop()
					if b.Background != (color.NRGBA{}) {
						background := b.Background
						paint.Fill(gtx.Ops, background)
					}
					for _, c := range b.Button.History() {
						drawInk(gtx, c)
					}
					return layout.Dimensions{Size: gtx.Constraints.Min}
				},
				func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Min = minWidth
					return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						iconDims := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if b.Icon != nil {
								return layout.Inset{Right: unit.Dp(5)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									gtx.Constraints.Min.X = gtx.Dp(b.IconSize)
									gtx.Constraints.Min.Y = gtx.Dp(b.IconSize)
									return b.Icon.Layout(gtx, b.Color)
								})
							}
							return layout.Dimensions{}
						})
						labelDims := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							lb := material.Label(b.theme.Material(), b.TextSize, b.Text)
							lb.Color = b.Color
							lb.Alignment = text.Middle
							return lb.Layout(gtx)
						})

						items := []layout.FlexChild{iconDims, labelDims}
						if b.IconPosition == 1 {
							items = []layout.FlexChild{labelDims, iconDims}
							b.Inset.Right = unit.Dp(5)
						}

						return b.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
								items...,
							)
						})
					})
				},
			)
		})
	})
}
func drawInk(gtx layout.Context, c widget.Press) {
	// duration is the number of seconds for the
	// completed animation: expand while fading in, then
	// out.
	const (
		expandDuration = float32(0.5)
		fadeDuration   = float32(0.9)
	)

	now := gtx.Now

	t := float32(now.Sub(c.Start).Seconds())

	end := c.End
	if end.IsZero() {
		// If the press hasn't ended, don't fade-out.
		end = now
	}

	endt := float32(end.Sub(c.Start).Seconds())

	// Compute the fade-in/out position in [0;1].
	var alphat float32
	{
		var haste float32
		if c.Cancelled {
			// If the press was cancelled before the inkwell
			// was fully faded in, fast forward the animation
			// to match the fade-out.
			if h := 0.5 - endt/fadeDuration; h > 0 {
				haste = h
			}
		}
		// Fade in.
		half1 := t/fadeDuration + haste
		if half1 > 0.5 {
			half1 = 0.5
		}

		// Fade out.
		half2 := float32(now.Sub(end).Seconds())
		half2 /= fadeDuration
		half2 += haste
		if half2 > 0.5 {
			// Too old.
			return
		}

		alphat = half1 + half2
	}

	// Compute the expand position in [0;1].
	sizet := t
	if c.Cancelled {
		// Freeze expansion of cancelled presses.
		sizet = endt
	}
	sizet /= expandDuration

	// Animate only ended presses, and presses that are fading in.
	if !c.End.IsZero() || sizet <= 1.0 {
		gtx.Execute(op.InvalidateCmd{})
	}

	if sizet > 1.0 {
		sizet = 1.0
	}

	if alphat > .5 {
		// Start fadeout after half the animation.
		alphat = 1.0 - alphat
	}
	// Twice the speed to attain fully faded in at 0.5.
	t2 := alphat * 2
	// Beziér ease-in curve.
	alphaBezier := t2 * t2 * (3.0 - 2.0*t2)
	sizeBezier := sizet * sizet * (3.0 - 2.0*sizet)
	size := gtx.Constraints.Min.X
	if h := gtx.Constraints.Min.Y; h > size {
		size = h
	}
	// Cover the entire constraints min rectangle and
	// apply curve values to size and color.
	size = int(float32(size) * 2 * float32(math.Sqrt(2)) * sizeBezier)
	alpha := 0.7 * alphaBezier
	const col = 0.8
	ba, bc := byte(alpha*0xff), byte(col*0xff)
	rgba := utils.MulAlpha(color.NRGBA{A: 0xff, R: bc, G: bc, B: bc}, ba)
	ink := paint.ColorOp{Color: rgba}
	ink.Add(gtx.Ops)
	rr := size / 2
	defer op.Offset(c.Position.Add(image.Point{
		X: -rr,
		Y: -rr,
	})).Push(gtx.Ops).Pop()
	defer clip.UniformRRect(image.Rectangle{Max: image.Pt(size, size)}, rr).Push(gtx.Ops).Pop()
	paint.PaintOp{}.Add(gtx.Ops)
}
