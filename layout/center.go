/**
 * Created by Goland
 * @file   center.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/1 19:49
 * @desc   center.go
 */

package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

func HorizontalCenter(gtx layout.Context, width unit.Dp, widget layout.Widget) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: layout.SpaceBetween}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(1)}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.X = gtx.Dp(width)
			return widget(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(1)}.Layout(gtx)
		}),
	)
}
func VerticalCenter(gtx layout.Context, height unit.Dp, widget layout.Widget) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle, Spacing: layout.SpaceBetween}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(1)}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Max.Y = gtx.Dp(height)
			return widget(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(1)}.Layout(gtx)
		}),
	)
}
