/**
 * Created by Goland
 * @file   list_menu.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/27 23:25
 * @desc   list_menu.go
 */

package widgets

import (
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
	"gioui.org/x/outlay"
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
)

type RightMenu struct {
	theme *theme.Theme

	optionsItems     []string
	menuContextArea  component.ContextArea
	menuState        component.MenuState
	menuWidth        unit.Dp
	clickFun         func(key int, menu string)
	options          []*RightMenuOption
	click            widget.Clickable
	clickable        widget.Clickable
	contentMarginTop unit.Dp
}

type RightMenuOption struct {
	Text      string
	Value     string
	clickable widget.Clickable

	Line      int
	Icon      *widget.Icon
	IconColor color.NRGBA
}

func NewRightMenu(th *theme.Theme) *RightMenu {
	return &RightMenu{
		theme:     th,
		menuWidth: unit.Dp(150),
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonSecondary,
			AbsolutePosition: false,
			PositionHint:     layout.E,
		},
	}
}

func (l *RightMenu) ContentMarginTop(top unit.Dp) *RightMenu {
	l.contentMarginTop = top
	return l
}
func (l *RightMenu) SetOptions(options []*RightMenuOption) *RightMenu {
	l.options = options
	return l
}

func (l *RightMenu) SetMenuWidth(width unit.Dp) {
	l.menuWidth = width
}
func (l *RightMenu) Clicked(fun func(key int, menu string)) {
	l.clickFun = fun
}

func (l *RightMenu) Layout(gtx layout.Context, widgets layout.Widget) layout.Dimensions {
	for i, opt := range l.options {
		for opt.clickable.Clicked(gtx) {
			if l.clickFun != nil {
				l.clickFun(i, opt.Value)
			}
		}
	}
	if l.contentMarginTop == 0 {
		l.contentMarginTop = unit.Dp(20)
	}
	l.updateMenuItems()
	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return l.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return widgets(gtx)
				})
			})
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return l.menuContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: unit.Dp(15), Bottom: unit.Dp(5), Left: unit.Dp(5), Right: unit.Dp(5)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					// return NewCard(l.theme).SetRadius(0).SetBgColor(l.theme.Color.RightMenuBorderColor).SetPadding(unit.Dp(0)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					// 	menu := component.Menu(l.theme.Material(), &l.menuState)
					// 	menu.SurfaceStyle.Fill = l.theme.Color.RightMenuBgColor
					// 	return menu.Layout(gtx)
					// })
					return widget.Border{
						Color:        l.theme.Color.RightMenuBorderColor,
						Width:        unit.Dp(1),
						CornerRadius: l.theme.Size.DefaultElementRadiusSize,
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Max.X = gtx.Dp(l.menuWidth)
						gtx.Constraints.Max.Y = gtx.Dp(unit.Dp(165))
						// gtx.Constraints.Min = gtx.Constraints.Max
						menu := component.Menu(l.theme.Material(), &l.menuState)
						menu.SurfaceStyle.Fill = l.theme.Color.RightMenuBgColor
						return menu.Layout(gtx)
					})
				})
			})
		}),
	)
}

// updateMenuItems creates or updates menu items based on options and calculates minWidth.
func (l *RightMenu) updateMenuItems() {
	l.menuState.Options = l.menuState.Options[:0]
	for _, opt := range l.options {
		currentOpt := opt
		l.menuState.Options = append(l.menuState.Options, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					if opt.Line > 0 {
						return layout.Inset{Top: unit.Dp(4), Bottom: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return NewLine(l.theme).Color(l.theme.Color.RightMenuLineColor).Width(1).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
						})
					}
					itm := component.MenuItem(l.theme.Material(), &currentOpt.clickable, currentOpt.Text)
					if currentOpt.Icon != nil {
						itm.Icon = currentOpt.Icon
						itm.IconColor = currentOpt.IconColor
					}

					itm.HoverColor = l.theme.Color.RightMenuHoveredBgColor

					itm.IconSize = l.theme.Size.DefaultIconSize
					itm.Label.TextSize = l.theme.Size.DefaultTextSize
					itm.Label.Color = l.theme.Color.DropdownTextColor
					itm.IconInset = outlay.Inset{
						Top:    unit.Dp(5),
						Bottom: unit.Dp(5),
						Start:  unit.Dp(5),
						// End:    unit.Dp(5),
					}
					itm.LabelInset = outlay.Inset{
						Top:    unit.Dp(5),
						Bottom: unit.Dp(5),
						Start:  unit.Dp(10),
						// End:    unit.Dp(5),
					}
					return itm.Layout(gtx)
				}),
			)
		})
	}
}
