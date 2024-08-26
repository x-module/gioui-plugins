/**
 * Created by Goland
 * @file   scroll.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/7 12:59
 * @desc   scroll.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
)

type Scroll struct {
	theme       *theme.Theme
	list        *widget.List
	elementList []layout.Widget
	bgColor     color.NRGBA
}

func NewScroll(th *theme.Theme) *Scroll {
	p := &Scroll{
		bgColor: th.Palette.Bg,
		theme:   th,
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
				// ScrollToEnd: true,
			},
		},
	}
	return p
}
func (n *Scroll) SetElementList(elementList []layout.Widget) {
	n.elementList = elementList
}

func (n *Scroll) SetBgColor(color color.NRGBA) {
	n.bgColor = color
}

// SetScrollToEnd 设置ScrollToEnd
func (n *Scroll) SetScrollToEnd(scrollToEnd bool) {
	n.list.ScrollToEnd = scrollToEnd
}

func (n *Scroll) Layout(gtx layout.Context) layout.Dimensions {
	// gtx.Constraints.Min.X = gtx.Dp(unit.Dp(150))
	// utils.ColorBackground(gtx, gtx.Constraints.Max, n.bgColor)
	gtx.Constraints.Min.X = gtx.Constraints.Max.X
	return material.List(n.theme.Material(), n.list).Layout(gtx, len(n.elementList), func(gtx layout.Context, index int) layout.Dimensions {
		return n.elementList[index](gtx)
	})
}
