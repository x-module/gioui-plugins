package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
)

type RadioButton struct {
	radioButton *widget.Bool
	key         string
	label       string
	group       *widget.Enum
	theme       *theme.Theme
	iconSize    unit.Dp
	textSize    unit.Sp
}

// NewRadioButton returns a RadioButton with a label. The key specifies
// the value for the Enum.
func NewRadioButton(th *theme.Theme, group *widget.Enum, key, label string) *RadioButton {
	r := &RadioButton{
		radioButton: &widget.Bool{Value: true},
		theme:       th,
		group:       group,
		key:         key,
		label:       label,
		iconSize:    th.Size.DefaultIconSize,
		textSize:    th.Size.DefaultTextSize,
	}
	return r
}

func (r *RadioButton) SetSize(size theme.ElementSize) {
	r.iconSize = size.IconSize
	r.textSize = size.TextSize
}

// Layout updates enum and displays the radio button.
func (r *RadioButton) Layout(gtx layout.Context) layout.Dimensions {
	iconColor := r.theme.Color.BorderLightGrayColor
	if r.group.Value == r.key {
		iconColor = r.theme.Color.ActivatedBorderBlueColor
	}
	return r.radioButton.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		if r.radioButton.Hovered() {
			iconColor = r.theme.Color.HoveredBorderBlueColor
		}
		rb := material.RadioButton(r.theme.Material(), r.group, r.key, r.label)
		rb.IconColor = iconColor
		rb.Color = r.theme.Color.DefaultTextWhiteColor
		rb.Size = r.iconSize
		rb.TextSize = r.textSize
		return rb.Layout(gtx)
	})
}
