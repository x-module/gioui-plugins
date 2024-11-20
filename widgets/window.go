/**
 * Created by Goland
 * @file   window.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/3 18:56
 * @desc   window.go
 */

package widgets

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	layout2 "github.com/x-module/gioui-plugins/layout"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
)

type ActionHook func() bool

type Window struct {
	theme *theme.Theme
	win   *app.Window

	content layout.Widget
	action  widget.Bool

	closeClick widget.Clickable
	minClick   widget.Clickable
	fullClick  widget.Clickable

	isFullWindow bool

	closeWinHook      ActionHook
	minWinHook        ActionHook
	fullWinHook       ActionHook
	unMaximizeWinHook ActionHook

	title             layout.Widget
	titleContentWidth unit.Dp
}

func NewWindow(th *theme.Theme, win *app.Window) *Window {
	return &Window{
		theme:             th,
		win:               win,
		titleContentWidth: unit.Dp(300),
	}
}
func (w *Window) SetTitle(title layout.Widget) *Window {
	w.title = title
	return w
}

func (w *Window) SetTitleContentWidth(width unit.Dp) *Window {
	w.titleContentWidth = width
	return w
}

func (w *Window) SetCloseWinHook(hook ActionHook) *Window {
	w.closeWinHook = hook
	return w
}

func (w *Window) SetMinWinHook(hook ActionHook) *Window {
	w.minWinHook = hook
	return w
}

func (w *Window) SetFullWinHook(hook ActionHook) *Window {
	w.fullWinHook = hook
	return w
}
func (w *Window) SetUnMaximizeWinHook(hook ActionHook) *Window {
	w.unMaximizeWinHook = hook
	return w
}

func (w *Window) SetContent(content layout.Widget) *Window {
	w.content = content
	return w
}

func (w *Window) Layout(gtx layout.Context) layout.Dimensions {
	if w.closeClick.Clicked(gtx) {
		if w.closeWinHook != nil {
			if w.closeWinHook() {
				w.win.Perform(system.ActionClose)
			}
		} else {
			w.win.Perform(system.ActionClose)
		}
	}
	if w.minClick.Clicked(gtx) {
		if w.minWinHook != nil {
			if w.minWinHook() {
				w.win.Perform(system.ActionMinimize)
			}
		} else {
			w.win.Perform(system.ActionMinimize)
		}
	}
	if w.fullClick.Clicked(gtx) {
		if w.isFullWindow {
			if w.unMaximizeWinHook != nil {
				if w.unMaximizeWinHook() {
					w.win.Perform(system.ActionUnmaximize)
				}
			} else {
				w.win.Perform(system.ActionUnmaximize)
			}
		} else {
			if w.fullWinHook != nil {
				if w.fullWinHook() {
					w.win.Perform(system.ActionFullscreen)
				}
			} else {
				w.win.Perform(system.ActionFullscreen)
			}
		}
		w.isFullWindow = !w.isFullWindow
	}

	return layout.Stack{}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return w.action.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							gtx.Constraints.Min.Y = gtx.Dp(unit.Dp(25))
							var child []layout.FlexChild
							if w.action.Hovered() {
								child = []layout.FlexChild{
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Inset{Left: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											gtx.Constraints.Max.X = gtx.Dp(unit.Dp(22))
											resource.ActionPointIcon.Layout(gtx, w.theme.Color.CloseIconColor)
											return layout.Inset{Top: unit.Dp(4), Left: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return w.closeClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
													gtx.Constraints.Max.X = gtx.Dp(unit.Dp(14))
													return resource.ActionCloseIcon.Layout(gtx, w.theme.Color.DefaultContentBgGrayColor)
												})
											})
										})
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Inset{Left: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											gtx.Constraints.Max.X = gtx.Dp(unit.Dp(22))
											resource.ActionPointIcon.Layout(gtx, w.theme.Color.MinIconColor)
											return layout.Inset{Top: unit.Dp(4), Left: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return w.minClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
													gtx.Constraints.Max.X = gtx.Dp(unit.Dp(14))
													return resource.ActionMinIcon.Layout(gtx, w.theme.Color.DefaultContentBgGrayColor)
												})
											})
										})
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										return layout.Inset{Left: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
											gtx.Constraints.Max.X = gtx.Dp(unit.Dp(22))
											resource.ActionPointIcon.Layout(gtx, w.theme.Color.FullIconColor)
											return layout.Inset{Top: unit.Dp(5), Left: unit.Dp(5)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
												return w.fullClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
													gtx.Constraints.Max.X = gtx.Dp(unit.Dp(12))
													return resource.ActionFullIcon.Layout(gtx, w.theme.Color.DefaultContentBgGrayColor)
												})
											})
										})
									}),
								}
							} else {
								child = []layout.FlexChild{
									// action
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										gtx.Constraints.Max.X = gtx.Dp(unit.Dp(22))
										return resource.ActionPointIcon.Layout(gtx, w.theme.Color.CloseIconColor)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										gtx.Constraints.Max.X = gtx.Dp(unit.Dp(22))
										return resource.ActionPointIcon.Layout(gtx, w.theme.Color.MinIconColor)
									}),
									layout.Rigid(func(gtx layout.Context) layout.Dimensions {
										gtx.Constraints.Max.X = gtx.Dp(unit.Dp(22))
										return resource.ActionPointIcon.Layout(gtx, w.theme.Color.FullIconColor)
									}),
								}
							}

							return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
								layout.Rigid(func(gtx layout.Context) layout.Dimensions {
									gtx.Constraints.Min.X = gtx.Dp(unit.Dp(70))
									return layout.Flex{Axis: layout.Horizontal}.Layout(gtx, child...)
								}),
							)
						})
					})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					if w.content != nil {
						return w.content(gtx)
					}
					return layout.Dimensions{}
				}),
			)
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			if w.title != nil {
				return layout2.HorizontalCenter(gtx, w.titleContentWidth, w.title)
			} else {
				return layout.Dimensions{}
			}
		}),
	)

}
