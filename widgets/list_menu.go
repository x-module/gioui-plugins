/**
 * Created by Goland
 * @file   list_menu.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/7/27 23:25
 * @desc   list_menu.go
 */

package widgets

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/x/component"
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
)

type ListMenu struct {
	Label           string
	optionsItems    []string
	theme           *theme.Theme
	menuContextArea component.ContextArea
	menuState       component.MenuState
	labelWidth      unit.Dp
	menuWidth       unit.Dp
	clickFun        func(key int, menu string)
	options         []*ListMenuOption
	click           widget.Clickable
	Show            bool
	clickable       widget.Clickable
	content         widget.Bool
}

type ListMenuOption struct {
	Text      string
	Value     string
	clickable widget.Clickable

	Icon      *widget.Icon
	IconColor color.NRGBA
}

func NewListMenu(th *theme.Theme, label string, options []*ListMenuOption) *ListMenu {
	listMenu := &ListMenu{
		theme:      th,
		Label:      label,
		labelWidth: unit.Dp(200),
		options:    options,
		menuContextArea: component.ContextArea{
			Activation:       pointer.ButtonPrimary,
			AbsolutePosition: true,
		},
		// menuState: component.MenuState{
		//	OptionList: layout.List{Axis: layout.Vertical},
		//	Options:    []func(gtx layout.Context) layout.Dimensions{},
		// },
	}
	return listMenu
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
			fmt.Println("clicked,i:", i)
			l.clickFun(i, opt.Value)
		}
	}
	l.updateMenuItems()
	return l.content.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return l.clickable.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						l.Show = false
						if l.content.Hovered() {
							l.Show = true
							gtx.Execute(op.InvalidateCmd{})
						}
						return Label(l.theme, l.Label).Layout(gtx)

					})
				})
			}),
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				if !l.Show {
					return layout.Dimensions{}
				}
				return layout.E.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{
						Top:  unit.Dp(30),
						Left: unit.Dp(2),
					}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Max.X = gtx.Dp(l.menuWidth)
						gtx.Constraints.Min = gtx.Constraints.Max
						menu := component.Menu(l.theme.Material(), &l.menuState)
						menu.SurfaceStyle.Fill = l.theme.Color.DropdownMenuBgColor
						return menu.Layout(gtx)
					})
				})
			}),
		)
	})
}

// updateMenuItems creates or updates menu items based on options and calculates minWidth.
func (l *ListMenu) updateMenuItems() {
	l.menuState.Options = l.menuState.Options[:0]
	for _, opt := range l.options {
		opt := opt
		l.menuState.Options = append(l.menuState.Options, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					itm := component.MenuItem(l.theme.Material(), &opt.clickable, opt.Text)
					if opt.Icon != nil {
						itm.Icon = opt.Icon
						itm.IconColor = opt.IconColor
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
