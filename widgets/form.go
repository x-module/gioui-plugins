/**
 * Created by Goland
 * @file   form.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/27 15:26
 * @desc   form.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/theme"
)

type Form struct {
	theme       *theme.Theme
	label       []string
	element     []layout.Widget
	spaceHeight unit.Dp
	labelWidth  unit.Dp

	marginTop    unit.Dp
	marginBottom unit.Dp
	marginRight  unit.Dp
	marginLeft   unit.Dp
}

func NewForm(th *theme.Theme, spaceHeight unit.Dp, labelWidth unit.Dp) *Form {
	return &Form{
		theme:        th,
		spaceHeight:  spaceHeight,
		labelWidth:   labelWidth,
		marginTop:    unit.Dp(0),
		marginBottom: unit.Dp(0),
		marginRight:  unit.Dp(0),
		marginLeft:   unit.Dp(0),
	}
}

// SetMargin 设置margin
func (f *Form) SetMargin(top, bottom, right, left unit.Dp) *Form {
	f.marginTop = top
	f.marginBottom = bottom
	f.marginRight = right
	f.marginLeft = left

	return f
}
func (f *Form) Add(label string, widget layout.Widget) *Form {
	f.label = append(f.label, label)
	f.element = append(f.element, widget)
	return f
}
func (f *Form) AddButton(widget layout.Widget) *Form {
	f.label = append(f.label, "")
	f.element = append(f.element, widget)
	return f
}

func (f *Form) Layout(gtx layout.Context) layout.Dimensions {
	var elements []layout.FlexChild
	for i := range f.label {
		elements = append(elements, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					gtx.Constraints.Max.X = gtx.Dp(f.labelWidth)
					return layout.Inset{Top: f.spaceHeight}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
							layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
								return layout.Spacer{}.Layout(gtx)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return layout.Inset{Top: unit.Dp(4), Right: unit.Dp(4)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
									labelText := f.label[i]
									if labelText != "" {
										labelText = labelText + "："
									}
									return Label(f.theme, labelText).Layout(gtx)
								})
							}),
						)
					})

				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.Inset{Top: f.spaceHeight}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return f.element[i](gtx)
					})
				}),
			)
		}))
	}
	return layout.Inset{Top: f.marginTop, Bottom: f.marginBottom, Right: f.marginRight, Left: f.marginLeft}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx, elements...)
	})
}
