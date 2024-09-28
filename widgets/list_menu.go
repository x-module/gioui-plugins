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
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
)

type ListMenu struct {
	theme *theme.Theme

	label            string
	optionsItems     []string
	menuContextArea  component.ContextArea
	menuState        component.MenuState
	labelWidth       unit.Dp
	menuWidth        unit.Dp
	clickFun         func(key int, menu string)
	options          []*ListMenuOption
	click            widget.Clickable
	clickable        widget.Clickable
	icon             *IconButton
	contentMarginTop unit.Dp
}

type ListMenuOption struct {
	Text      string
	Value     string
	clickable widget.Clickable

	Icon      *widget.Icon
	IconColor color.NRGBA
}

func NewListMenu(th *theme.Theme) *ListMenu {
	listMenu := &ListMenu{
		theme:      th,
		labelWidth: unit.Dp(200),
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonPrimary,
			AbsolutePosition: false,
			PositionHint:     layout.E,
		},
	}
	return listMenu
}

func (l *ListMenu) ContentMarginTop(top unit.Dp) *ListMenu {
	l.contentMarginTop = top
	return l
}

func (l *ListMenu) SetLabel(label string) *ListMenu {
	l.label = label
	return l
}

func (l *ListMenu) SetIcon(icon *widget.Icon) *ListMenu {
	l.icon = NewIconButton(l.theme, icon)
	return l
}

func (l *ListMenu) SetOptions(options []*ListMenuOption) *ListMenu {
	l.options = options
	return l
}

func (l *ListMenu) SetMenuWidth(width unit.Dp) {
	l.menuWidth = width
}
func (l *ListMenu) SetLabelWidth(width unit.Dp) {
	l.labelWidth = width
}
func (l *ListMenu) Clicked(fun func(key int, menu string)) {
	l.clickFun = fun
}

func (l *ListMenu) Layout(gtx layout.Context) layout.Dimensions {
	for i, opt := range l.options {
		for opt.clickable.Clicked(gtx) {
			if l.clickFun != nil {
				l.clickFun(i, opt.Value)
			}
		}
	}
	if l.contentMarginTop == 0 {
		l.contentMarginTop = unit.Dp(30)
	}
	l.updateMenuItems()
	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return l.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if l.icon != nil {
								return l.icon.Layout(gtx)
							}
							return layout.Dimensions{}
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if l.label != "" {
								return Label(l.theme, l.label).Layout(gtx)
							}
							return layout.Dimensions{}
						}),
					)
				})
			})
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return l.menuContextArea.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:  l.contentMarginTop,
						Left: unit.Dp(2),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Max.X = gtx.Dp(l.menuWidth)
						gtx.Constraints.Min = gtx.Constraints.Max
						menu := component.Menu(l.theme.Material(), &l.menuState)
						menu.SurfaceStyle.Fill = l.theme.Color.DropdownMenuBgColor
						return menu.Layout(gtx)
					})
				})
			})
		}),
	)
}

// updateMenuItems creates or updates menu items based on options and calculates minWidth.
func (l *ListMenu) updateMenuItems() {
	l.menuState.Options = l.menuState.Options[:0]
	for _, opt := range l.options {
		currentOpt := opt
		l.menuState.Options = append(l.menuState.Options, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					itm := component.MenuItem(l.theme.Material(), &currentOpt.clickable, currentOpt.Text)
					if currentOpt.Icon != nil {
						itm.Icon = currentOpt.Icon
						itm.IconColor = currentOpt.IconColor
					}
					itm.Label.TextSize = l.theme.Size.DefaultTextSize
					itm.Label.Color = l.theme.Color.DropdownTextColor
					return itm.Layout(gtx)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return NewLine(l.theme).Color(l.theme.Color.DefaultLineColor).Width(1).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X), 0)).Layout(gtx)
				}),
			)
		})
	}
}
