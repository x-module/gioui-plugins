/**
 * Created by Goland
 * @file   theme.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/23 16:25
 * @desc   主题定义
 */

package theme

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"image/color"
)

type ElementSize struct {
	TextSize unit.Sp
	Height   unit.Dp
	Inset    layout.Inset
	IconSize unit.Dp
}

type Size struct {
	Tiny   ElementSize
	Small  ElementSize
	Medium ElementSize
	Large  ElementSize

	DefaultElementWidth      unit.Dp
	DefaultTextSize          unit.Sp
	DropdownTextSize         unit.Sp
	DefaultIconSize          unit.Dp
	DefaultElementRadiusSize unit.Dp
	DefaultWidgetRadiusSize  unit.Dp
}
type Color struct {
	DefaultWindowBgGrayColor  color.NRGBA
	DefaultContentBgGrayColor color.NRGBA
	DefaultBgGrayColor        color.NRGBA
	DefaultTextWhiteColor     color.NRGBA
	DefaultBorderGrayColor    color.NRGBA
	DefaultBorderBlueColor    color.NRGBA
	DefaultLineColor          color.NRGBA
	DefaultMaskBgColor        color.NRGBA

	ButtonBorderColor       color.NRGBA
	ButtonDefaultTextColor  color.NRGBA
	ButtonTertiaryBgColor   color.NRGBA
	ButtonTertiaryTextColor color.NRGBA
	ButtonTextBlackColor    color.NRGBA
	ButtonDefaultColor      color.NRGBA
	ButtonTertiaryColor     color.NRGBA

	GreenColor   color.NRGBA
	ErrorColor   color.NRGBA
	WarningColor color.NRGBA
	SuccessColor color.NRGBA
	InfoColor    color.NRGBA
	PrimaryColor color.NRGBA

	IconGrayColor            color.NRGBA
	BorderBlueColor          color.NRGBA
	BorderLightGrayColor     color.NRGBA
	HoveredBorderBlueColor   color.NRGBA
	FocusedBorderBlueColor   color.NRGBA
	ActivatedBorderBlueColor color.NRGBA
	FocusedBgColor           color.NRGBA
	TextSelectionColor       color.NRGBA
	HintTextColor            color.NRGBA

	DropDownBgGrayColor          color.NRGBA
	DropDownItemHoveredGrayColor color.NRGBA

	ActionTipsBgGrayColor color.NRGBA
	ProgressBarColor      color.NRGBA

	MenuHoveredBgColor  color.NRGBA
	MenuSelectedBgColor color.NRGBA
	LogTextWhiteColor   color.NRGBA

	NotificationBgColor        color.NRGBA
	NotificationTextWhiteColor color.NRGBA
	ModalBgGrayColor           color.NRGBA

	DropdownMenuBgColor color.NRGBA
	DropdownTextColor   color.NRGBA

	NoticeInfoColor    color.NRGBA
	NoticeSuccessColor color.NRGBA
	NoticeWaringColor  color.NRGBA
	NoticeErrorColor   color.NRGBA

	JsonStartEndColor color.NRGBA
	JsonKeyColor      color.NRGBA
	JsonStringColor   color.NRGBA
	JsonNumberColor   color.NRGBA
	JsonBoolColor     color.NRGBA
	JsonNullColor     color.NRGBA
}

type Theme struct {
	*material.Theme
	Color Color
	Size  Size
}

func NewTheme(isDark ...bool) *Theme {
	t := &Theme{
		Theme: material.NewTheme(),
	}
	if len(isDark) == 0 || !isDark[0] {
		t = t.dark()
	}
	return t
}

func (t *Theme) Material() *material.Theme {
	return t.Theme
}

func (t *Theme) dark() *Theme {
	t.Color.DefaultWindowBgGrayColor = color.NRGBA{R: 17, G: 15, B: 20, A: 255}
	t.Color.DefaultContentBgGrayColor = color.NRGBA{R: 24, G: 24, B: 28, A: 255}

	t.Color.DefaultBgGrayColor = color.NRGBA{R: 53, G: 54, B: 56, A: 255}
	t.Color.DefaultTextWhiteColor = color.NRGBA{R: 223, G: 223, B: 224, A: 255}
	t.Color.DefaultBorderGrayColor = color.NRGBA{R: 53, G: 54, B: 56, A: 255}
	t.Color.DefaultBorderBlueColor = color.NRGBA{R: 127, G: 231, B: 196, A: 255}

	t.Color.DefaultLineColor = color.NRGBA{R: 44, G: 44, B: 47, A: 255}
	t.Color.DefaultMaskBgColor = color.NRGBA{R: 10, G: 10, B: 12, A: 230}

	t.Color.IconGrayColor = color.NRGBA{R: 136, G: 136, B: 137, A: 255}
	t.Color.BorderBlueColor = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	t.Color.BorderLightGrayColor = color.NRGBA{R: 65, G: 65, B: 68, A: 255}
	t.Color.HoveredBorderBlueColor = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	t.Color.FocusedBorderBlueColor = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	t.Color.ActivatedBorderBlueColor = color.NRGBA{R: 127, G: 231, B: 196, A: 255}
	t.Color.FocusedBgColor = color.NRGBA{R: 33, G: 50, B: 46, A: 255}
	t.Color.TextSelectionColor = color.NRGBA{R: 92, G: 136, B: 177, A: 255}
	t.Color.HintTextColor = color.NRGBA{R: 136, G: 136, B: 137, A: 255}

	t.Color.DropDownBgGrayColor = color.NRGBA{R: 72, G: 72, B: 77, A: 255}
	t.Color.DropDownItemHoveredGrayColor = color.NRGBA{R: 90, G: 90, B: 96, A: 255}

	t.Color.ButtonBorderColor = color.NRGBA{R: 76, G: 76, B: 79, A: 255}
	t.Color.ButtonTertiaryBgColor = color.NRGBA{R: 24, G: 24, B: 28, A: 255}
	t.Color.ButtonTertiaryTextColor = color.NRGBA{R: 149, G: 149, B: 150, A: 255}
	t.Color.ButtonDefaultTextColor = color.NRGBA{R: 216, G: 216, B: 217, A: 255}
	t.Color.ButtonTextBlackColor = color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	t.Color.GreenColor = color.NRGBA{R: 101, G: 231, B: 188, A: 255}
	t.Color.ErrorColor = color.NRGBA{R: 232, G: 127, B: 127, A: 255}
	t.Color.WarningColor = color.NRGBA{R: 242, G: 201, B: 126, A: 255}
	t.Color.SuccessColor = color.NRGBA{R: 99, G: 226, B: 184, A: 255}
	t.Color.InfoColor = color.NRGBA{R: 113, G: 192, B: 231, A: 255}
	t.Color.PrimaryColor = color.NRGBA{R: 99, G: 226, B: 184, A: 255}
	t.Color.ButtonDefaultColor = color.NRGBA{R: 24, G: 24, B: 28, A: 255}
	t.Color.ButtonTertiaryColor = color.NRGBA{R: 24, G: 24, B: 28, A: 255}

	t.Color.ActionTipsBgGrayColor = color.NRGBA{A: 255, R: 48, G: 48, B: 51}
	t.Color.ProgressBarColor = color.NRGBA{R: 127, G: 200, B: 235, A: 255}

	t.Color.MenuHoveredBgColor = color.NRGBA{R: 45, G: 45, B: 48, A: 255}
	t.Color.MenuSelectedBgColor = color.NRGBA{R: 35, G: 54, B: 51, A: 255}
	t.Color.LogTextWhiteColor = color.NRGBA{R: 202, G: 202, B: 203, A: 255}

	t.Color.NotificationBgColor = color.NRGBA{R: 72, G: 72, B: 77, A: 255}
	t.Color.NotificationTextWhiteColor = color.NRGBA{R: 219, G: 219, B: 220, A: 255}
	t.Color.ModalBgGrayColor = color.NRGBA{R: 44, G: 44, B: 50, A: 255}

	t.Color.DropdownMenuBgColor = color.NRGBA{R: 72, G: 72, B: 77, A: 255}
	t.Color.DropdownTextColor = color.NRGBA{R: 212, G: 212, B: 213, A: 255}

	t.Color.NoticeInfoColor = color.NRGBA{R: 108, G: 184, B: 221, A: 255}
	t.Color.NoticeSuccessColor = color.NRGBA{R: 101, G: 231, B: 188, A: 255}
	t.Color.NoticeWaringColor = color.NRGBA{R: 242, G: 201, B: 126, A: 255}
	t.Color.NoticeErrorColor = color.NRGBA{R: 231, G: 127, B: 127, A: 255}

	t.Color.JsonStartEndColor = color.NRGBA{R: 194, G: 196, B: 202, A: 255}
	t.Color.JsonKeyColor = color.NRGBA{R: 159, G: 101, B: 150, A: 255}
	t.Color.JsonStringColor = color.NRGBA{R: 105, G: 168, B: 114, A: 255}
	t.Color.JsonNumberColor = color.NRGBA{R: 41, G: 159, B: 171, A: 255}
	t.Color.JsonBoolColor = color.NRGBA{R: 161, G: 112, B: 88, A: 255}
	t.Color.JsonNullColor = color.NRGBA{R: 170, G: 118, B: 93, A: 255}

	t.Size.Tiny = ElementSize{TextSize: unit.Sp(9), Height: unit.Dp(10), Inset: layout.UniformInset(unit.Dp(4)), IconSize: unit.Dp(14)}
	t.Size.Small = ElementSize{TextSize: unit.Sp(12), Height: unit.Dp(15), Inset: layout.UniformInset(unit.Dp(6)), IconSize: unit.Dp(18)}
	t.Size.Medium = ElementSize{TextSize: unit.Sp(14), Height: unit.Dp(20), Inset: layout.UniformInset(unit.Dp(8)), IconSize: unit.Dp(24)}
	t.Size.Large = ElementSize{TextSize: unit.Sp(20), Height: unit.Dp(25), Inset: layout.UniformInset(unit.Dp(10)), IconSize: unit.Dp(30)}

	t.Size.DefaultElementWidth = unit.Dp(350)
	t.Size.DefaultTextSize = unit.Sp(14)
	t.Size.DropdownTextSize = unit.Sp(13)
	t.Size.DefaultIconSize = unit.Dp(20)
	t.Size.DefaultElementRadiusSize = unit.Dp(4)
	t.Size.DefaultWidgetRadiusSize = unit.Dp(8)

	return t
}
