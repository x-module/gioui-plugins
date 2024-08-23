// SPDX-License-Identifier: Unlicense OR MIT

package widgets

import (
	"gioui.org/font"
	"gioui.org/widget/material"
	"github.com/gioui-plugins/theme"
)

func H1(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.H1(th.Material(), txt))
}

func H2(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.H2(th.Material(), txt))
}

func H3(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.H3(th.Material(), txt))
}

func H4(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.H4(th.Material(), txt))
}

func H5(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.H5(th.Material(), txt))
}

func H6(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.H6(th.Material(), txt))
}

func Body1(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.Body1(th.Material(), txt))
}

func Body2(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.Body2(th.Material(), txt))
}

func Caption(th *theme.Theme, txt string) material.LabelStyle {
	return labelWithDefaultColor(th, material.Caption(th.Material(), txt))
}

func ErrorLabel(th *theme.Theme, txt string) material.LabelStyle {
	label := Caption(th, txt)
	label.Color = th.Color.ErrorColor
	return label
}
func Label(th *theme.Theme, txt string, bold ...bool) material.LabelStyle {
	label := material.Label(th.Material(), th.Size.DefaultTextSize, txt)
	label.Color = th.Color.DefaultTextWhiteColor
	if len(bold) > 0 && bold[0] == true {
		label.Font.Weight = font.Bold
	}
	return label
}
func labelWithDefaultColor(th *theme.Theme, entry material.LabelStyle) material.LabelStyle {
	entry.Color = th.Color.DefaultTextWhiteColor
	return entry
}
