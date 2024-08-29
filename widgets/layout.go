/**
 * Created by Goland
 * @file   layout.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/27 13:49
 * @desc   layout.go
 */

package widgets

import (
	"gioui.org/layout"
)

type Layout struct {
}

func NewLayout() *Layout {
	return &Layout{}
}

func (l *Layout) CenterLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		// 在按钮之前和之后添加均等的 Flexed 空间以保持水平居中
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			// 这里的按钮将在垂直方向上居中
			return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
				// 在按钮之前和之后添加均等的 Flexed 空间以保持水平居中
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{}.Layout(gtx)
				}),
				layout.Rigid(widget),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.Spacer{}.Layout(gtx)
				}),
			)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
	)
}
func (l *Layout) HorizontalLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
		// 在按钮之前和之后添加均等的 Flexed 空间以保持水平居中
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
		layout.Rigid(widget),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
	)
}
func (l *Layout) VerticalLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical, Alignment: layout.Middle}.Layout(gtx,
		// 在按钮之前和之后添加均等的 Flexed 空间以保持水平居中
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
		layout.Rigid(widget),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Spacer{}.Layout(gtx)
		}),
	)
}
