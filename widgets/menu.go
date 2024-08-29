/**
 * Created by Goland
 * @file   menu.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/7 14:20
 * @desc   menu.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/utils"
)

type Menu struct {
	theme           *theme.Theme
	list            *widget.List
	Items           []MenuItem
	menuItemOptions []MenuItemOption
	clicks          [][]*widget.Clickable
	clicked         int
	clickCallback   func(main int, sub int)

	parentIndex int
	subIndex    int
}

func NewMenu(th *theme.Theme) *Menu {
	return &Menu{
		theme:   th,
		clicked: 100,
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
				// ScrollToEnd: true,
			},
		},
	}
}

type MenuItemOption struct {
	Icon        *widget.Icon
	Text        string
	MarginRight int
	MarginLeft  int
	SubMenu     []MenuItemOption
	Target      func()
}

type MenuItem struct {
	Content layout.Widget
}

func (m *Menu) SetClickCallback(callback func(main int, sub int)) {
	m.clickCallback = callback
}

// SetMenuItemOptions 设置menuItemOptions
func (m *Menu) SetMenuItemOptions(options []MenuItemOption) {
	m.menuItemOptions = options
	for key := range m.menuItemOptions {
		d := key
		m.Items = append(m.Items, MenuItem{
			Content: func(gtx layout.Context) layout.Dimensions {
				return m.generateMenu(gtx, d, m.theme)
			},
		})
	}
}

func (m *MenuItem) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return m.Content(gtx)
		}),
	)
}

func (m *Menu) copyMenuItem(item MenuItem) MenuItem {
	return MenuItem{
		Content: item.Content,
	}
}

// Layout
func (m *Menu) action(gtx layout.Context) {
	for i, cs := range m.clicks {
		for s, c := range cs {
			for c.Clicked(gtx) {
				if s == 0 {
					if m.clicked == i {
						m.clicked = 1000
					} else {
						m.clicked = i
					}
				}
				m.parentIndex = i
				m.subIndex = s
				if m.clickCallback != nil {
					m.clickCallback(i, s)
				}
			}
		}
	}
}
func (m *Menu) Layout(gtx layout.Context) layout.Dimensions {
	m.action(gtx)

	return m.list.Layout(gtx, len(m.Items), func(gtx layout.Context, index int) layout.Dimensions {
		return m.Items[index].Layout(gtx)
	})
}
func (m *Menu) makeButton(option MenuItemOption, clickable *widget.Clickable) FlatButton {
	return FlatButton{
		theme:                  m.theme,
		Icon:                   option.Icon,
		Text:                   option.Text,
		TextSize:               unit.Sp(11),
		IconPosition:           FlatButtonIconStart,
		Clickable:              clickable,
		SpaceBetween:           unit.Dp(3),
		BackgroundPadding:      unit.Dp(0),
		CornerRadius:           0,
		MinWidth:               unit.Dp(160),
		BackgroundColor:        m.theme.Color.DefaultContentBgGrayColor,
		HoveredBackgroundColor: m.theme.Color.MenuHoveredBgColor,
		ClickedBackgroundColor: m.theme.Color.MenuSelectedBgColor,
		TextColor:              m.theme.Color.DefaultTextWhiteColor,
		ContentPadding:         unit.Dp(4),
		MarginRight:            option.MarginRight,
		MarginLeft:             option.MarginLeft,
	}
}
func (m *Menu) makeSubButton(option MenuItemOption, clickable *widget.Clickable) FlatButton {
	return FlatButton{
		theme:                  m.theme,
		Icon:                   option.Icon,
		Text:                   option.Text,
		IconPosition:           FlatButtonIconStart,
		Clickable:              clickable,
		SpaceBetween:           unit.Dp(2),
		BackgroundPadding:      unit.Dp(0),
		CornerRadius:           0,
		MinWidth:               unit.Dp(160),
		BackgroundColor:        m.theme.Color.DefaultContentBgGrayColor,
		HoveredBackgroundColor: m.theme.Color.MenuHoveredBgColor,
		ClickedBackgroundColor: m.theme.Color.MenuSelectedBgColor,
		TextColor:              m.theme.Color.DefaultTextWhiteColor,
		ContentPadding:         unit.Dp(3),
		MarginRight:            option.MarginRight,
		MarginLeft:             25,
		IconSize:               unit.Dp(15),
	}
}
func (m *Menu) generateMenu(gtx layout.Context, key int, th *theme.Theme) layout.Dimensions {
	for key := range m.menuItemOptions {
		scli := []*widget.Clickable{
			new(widget.Clickable),
		}
		if m.menuItemOptions[key].SubMenu != nil {
			for range m.menuItemOptions[key].SubMenu {
				scli = append(scli, new(widget.Clickable))
			}
		}
		m.clicks = append(m.clicks, scli)
	}
	return layout.Inset{Left: unit.Dp(0), Right: unit.Dp(0), Top: unit.Dp(0), Bottom: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						button := m.makeButton(m.menuItemOptions[key], m.clicks[key][0])
						if m.subIndex != 0 && m.parentIndex == key {
							button.TextColor = m.theme.Color.GreenColor
						}
						return button.Layout(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return utils.DrawLine(gtx, m.theme.Color.DefaultLineColor, unit.Dp(1), unit.Dp(160))
					}),
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if m.clicked != key {
					return layout.Dimensions{}
				} else {
					if m.menuItemOptions[key].SubMenu != nil {
						var subMenu []layout.FlexChild
						for i := range m.menuItemOptions[key].SubMenu {
							buttons := m.makeSubButton(m.menuItemOptions[key].SubMenu[i], m.clicks[key][i+1])
							subMenu = append(subMenu, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Inset{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											if m.parentIndex == key && m.subIndex == (i+1) {
												buttons.TextColor = m.theme.Color.GreenColor
											}
											return buttons.Layout(gtx)
										})
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return utils.DrawLine(gtx, m.theme.Color.DefaultLineColor, unit.Dp(1), unit.Dp(160))
									}),
								)
							}))
						}
						return layout.N.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								return layout.Flex{Axis: layout.Vertical}.Layout(gtx, subMenu...)
							})
						})

					}
					return layout.Dimensions{}
				}
			}),
			// utils.DrawLineFlex(resource.DefaultLineColor, unit.Dp(1), unit.Dp(160)),
		)
	})
}
