package widgets

import (
	"gioui.org/gesture"
	"gioui.org/io/event"
	"gioui.org/io/input"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"image"
	"image/color"
)

type state uint8

const (
	inactive state = iota
	hovered
	activated
	focused
)

type ActionFun func(gtx layout.Context)
type Input struct {
	theme *theme.Theme

	editor    widget.Editor
	height    unit.Dp
	before    layout.Widget
	after     layout.Widget
	icon      *widget.Icon
	iconClick widget.Clickable

	click       gesture.Click
	state       state
	borderColor color.NRGBA
	bgColor     color.NRGBA
	hint        string
	radius      unit.Dp
	size        theme.ElementStyle
	width       unit.Dp

	showPassword bool

	lineHeight unit.Sp
	hasBorder  bool

	onIconClick ActionFun
	onFocus     ActionFun
	onLostFocus ActionFun
	onChange    ActionFun
}

func NewInput(th *theme.Theme, hint string, text ...string) *Input {
	t := &Input{
		theme:     th,
		editor:    widget.Editor{},
		hasBorder: true,
		// width:  th.Size.DefaultElementWidth,
	}
	t.size = th.Size.Medium
	t.hint = hint
	t.radius = th.Size.DefaultElementRadiusSize
	if len(text) > 0 {
		t.editor.SetText(text[0])
	}
	t.editor.SingleLine = true
	return t
}

func NewTextArea(th *theme.Theme, hint string, text ...string) *Input {
	t := &Input{
		theme:     th,
		editor:    widget.Editor{},
		height:    unit.Dp(100),
		hasBorder: true,
		// width:  th.Size.DefaultElementWidth,
	}
	t.size = th.Size.Medium
	t.hint = hint
	t.radius = th.Size.DefaultElementRadiusSize
	if len(text) > 0 {
		t.editor.SetText(text[0])
	}
	t.editor.SingleLine = false
	t.editor.LineHeightScale = 1.5
	return t
}

// set hasBorder
func (i *Input) SetHasBorder(hasBorder bool) *Input {
	i.hasBorder = hasBorder
	return i
}

// lineHeight
func (i *Input) SetLineHeight(height unit.Sp) *Input {
	i.lineHeight = height
	return i
}

func (i *Input) SetOnFocus(f ActionFun) *Input {
	i.onFocus = f
	return i
}
func (i *Input) SetOnLostFocus(f ActionFun) *Input {
	i.onLostFocus = f
	return i
}

func (i *Input) SetHeight(height unit.Dp) *Input {
	i.height = height
	return i
}

func (i *Input) SetWidth(width unit.Dp) *Input {
	i.width = width
	return i
}

func (i *Input) SetOnIconClick(f ActionFun) *Input {
	i.onIconClick = f
	return i
}
func (i *Input) SetonChanged(f ActionFun) *Input {
	i.onChange = f
	return i
}

func (i *Input) Password() *Input {
	i.editor.Mask = '*'
	i.icon, _ = widget.NewIcon(icons.ActionVisibilityOff)
	// t.IconPosition = IconPositionEnd
	i.showPassword = false
	return i
}
func (i *Input) SetIcon(icon *widget.Icon) *Input {
	i.icon = icon
	return i
}

// SetRadius 设置radius
func (i *Input) SetRadius(radius unit.Dp) *Input {
	i.radius = radius
	return i
}
func (i *Input) ReadOnly() *Input {
	i.editor.ReadOnly = true
	return i
}
func (i *Input) SetBefore(before layout.Widget) *Input {
	i.before = before
	return i
}
func (i *Input) SetAfter(after layout.Widget) *Input {
	i.after = after
	return i
}

func (i *Input) SetSize(size theme.ElementStyle) *Input {
	i.size = size
	return i
}

func (i *Input) SetText(text string) *Input {
	i.editor.SetText(text)
	return i
}
func (i *Input) GetText() string {
	return i.editor.Text()
}
func (i *Input) update(gtx layout.Context, th *theme.Theme) {
	if gtx.Focused(&i.editor) {
		if i.onFocus != nil {
			i.onFocus(gtx)
		}
	} else {
		if i.onLostFocus != nil {
			i.onLostFocus(gtx)
		}
	}
	disabled := gtx.Source == (input.Source{})
	for {
		ev, ok := i.click.Update(gtx.Source)
		if !ok {
			break
		}
		switch ev.Kind {
		case gesture.KindPress:
			gtx.Execute(key.FocusCmd{Tag: &i.editor})
		case gesture.KindClick:

		default:

		}
	}
	i.state = inactive
	if i.click.Hovered() && !disabled {
		i.state = hovered
	}
	// if t.editor.Len() > 0 {
	// 	t.state = activated
	// }
	if gtx.Source.Focused(&i.editor) && !disabled {
		i.state = focused
	}

	i.bgColor = i.theme.Color.DefaultBgGrayColor

	if i.editor.ReadOnly {
		return
	}

	switch i.state {
	case inactive:
		i.borderColor = i.theme.Color.InputInactiveBorderColor
	case hovered:
		i.borderColor = i.theme.Color.InputHoveredBorderColor
	case focused:
		i.bgColor = i.theme.Color.InputFocusedBgColor
		i.borderColor = i.theme.Color.InputFocusedBorderColor
	case activated:
		i.borderColor = i.theme.Color.InputActivatedBorderColor
	}
	for {
		e, ok := i.editor.Update(gtx)
		if !ok {
			break
		}
		if _, ok := e.(widget.ChangeEvent); ok {
			if i.onChange != nil {
				i.onChange(gtx)
			}
		}
	}
}

func (i *Input) Layout(gtx layout.Context) layout.Dimensions {
	if i.width > 0 {
		gtx.Constraints.Max.X = gtx.Dp(i.width)
	}
	i.update(gtx, i.theme)
	// gtx.Constraints.Min.X = gtx.Constraints.Max.X
	// gtx.Constraints.Min.Y = 0
	macro := op.Record(gtx.Ops)
	dims := i.layout(gtx, i.theme)
	call := macro.Stop()
	defer pointer.PassOp{}.Push(gtx.Ops).Pop()
	defer clip.Rect(image.Rectangle{Max: dims.Size}).Push(gtx.Ops).Pop()
	i.click.Add(gtx.Ops)
	event.Op(gtx.Ops, &i.editor)
	call.Add(gtx.Ops)
	return dims
}

func (i *Input) layout(gtx layout.Context, th *theme.Theme) layout.Dimensions {
	borderWidth := unit.Dp(1)
	if !i.hasBorder {
		borderWidth = unit.Dp(0)
	}
	border := widget.Border{
		Color:        i.borderColor,
		Width:        borderWidth,
		CornerRadius: i.radius,
	}
	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Background{}.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				rr := gtx.Dp(i.radius)
				defer clip.UniformRRect(image.Rectangle{Max: gtx.Constraints.Min}, rr).Push(gtx.Ops).Pop()
				paint.Fill(gtx.Ops, i.bgColor)
				return layout.Dimensions{Size: gtx.Constraints.Min}
			},
			func(gtx layout.Context) layout.Dimensions {
				return i.size.Inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					inputLayout := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						if i.width > 0 {
							gtx.Constraints.Max.X = gtx.Dp(i.width)
						}
						editor := material.Editor(th.Material(), &i.editor, i.hint)
						editor.HintColor = i.theme.Color.HintTextColor
						editor.SelectionColor = i.theme.Color.TextSelectionColor

						gtx.Constraints.Min.Y = gtx.Dp(i.size.Height) // 设置最小高度为 100dp
						gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
						editor.TextSize = i.size.TextSize
						editor.Color = i.size.TextColor
						editor.LineHeight = i.lineHeight
						editor.LineHeightScale = 1

						if i.height > 0 {
							gtx.Constraints.Min.Y = gtx.Dp(i.height)      // 设置最小高度为 100dp
							gtx.Constraints.Max.Y = gtx.Constraints.Min.Y // 限制最大高度与最小高度相同
						}
						if i.editor.ReadOnly {
							editor.Color = i.theme.Color.HintTextColor
						} else {
							editor.Color = i.theme.Color.DefaultTextWhiteColor
						}
						return editor.Layout(gtx)
					})

					var widgets []layout.FlexChild
					if i.before != nil {
						widgets = append(widgets, layout.Rigid(i.before))
					}
					widgets = append(widgets, inputLayout)
					if i.icon != nil {
						iconLayout := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if i.iconClick.Clicked(gtx) {
								if i.onIconClick != nil {
									i.onIconClick(gtx)
								}
								if !i.showPassword {
									i.editor.Mask = 0
									i.icon = resource.ActionVisibilityIcon
									i.showPassword = true
								} else {
									i.editor.Mask = '*'
									i.icon = resource.ActionVisibilityOffIcon
									i.showPassword = false
								}
							}
							return i.iconClick.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								gtx.Constraints.Max.X = gtx.Dp(i.size.IconSize)
								return i.icon.Layout(gtx, i.theme.Color.DefaultIconColor)
							})
						})
						widgets = append(widgets, iconLayout)
					} else {
						if i.after != nil {
							widgets = append(widgets, layout.Rigid(i.after))
						}
					}
					spacing := layout.SpaceBetween
					return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle, Spacing: spacing}.Layout(gtx, widgets...)
				})
			},
		)
	})
}
