/**
 * Created by Goland
 * @file   center.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/2 22:05
 * @desc   center.go
 */

package layout

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

func HorizontalCenter(gtx layout.Context, width unit.Dp, widget layout.Widget) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(10)}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Dp(width)
			return widget(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{Width: unit.Dp(10)}.Layout(gtx)
		}),
	)
}
