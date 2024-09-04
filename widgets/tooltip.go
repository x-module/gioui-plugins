/**
 * Created by Goland
 * @file   tooltip.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/4 11:37
 * @desc   tooltip.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/x/component"
	"github.com/x-module/gioui-plugins/theme"
)

type Tooltip struct {
	theme   *theme.Theme
	tipArea component.TipArea
}

func NewTooltip(th *theme.Theme) *Tooltip {
	return &Tooltip{theme: th}
}

func (t *Tooltip) Layout(gtx layout.Context, content layout.Widget, tips string) layout.Dimensions {
	tip := component.DesktopTooltip(t.theme.Material(), tips)
	return t.tipArea.Layout(gtx, tip, func(gtx layout.Context) layout.Dimensions {
		return content(gtx)
	})
}
