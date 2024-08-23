package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/gioui-plugins/theme"
)

type LabeledInput struct {
	Label        string
	theme        *theme.Theme
	SpaceBetween int
	inputWidth   unit.Dp
	labelWidth   unit.Dp
	input        *Input
	Hint         string
}

func NewLabeledInput(th *theme.Theme, label string, hint string) *LabeledInput {
	labeledInput := &LabeledInput{
		theme:        th,
		Label:        label,
		input:        NewInput(th, hint),
		Hint:         hint,
		SpaceBetween: 3,
		labelWidth:   unit.Dp(50),
		inputWidth:   th.Size.DefaultElementWidth,
	}
	labeledInput.input.SetWidth(labeledInput.inputWidth)
	return labeledInput
}

func (l *LabeledInput) SetText(text string) *LabeledInput {
	l.input.SetText(text)
	return l
}
func (l *LabeledInput) ReadOnly() *LabeledInput {
	l.input.ReadOnly()
	return l
}

func (l *LabeledInput) SetSpaceBetween(space int) *LabeledInput {
	l.SpaceBetween = space
	return l
}
func (l *LabeledInput) SetInputWidth(width unit.Dp) *LabeledInput {
	l.inputWidth = width
	return l
}
func (l *LabeledInput) SetLabelWidth(width unit.Dp) *LabeledInput {
	l.labelWidth = width
	return l
}

func (l *LabeledInput) GetText() string {
	return l.input.GetText()
}

func (l *LabeledInput) Layout(gtx layout.Context) layout.Dimensions {
	l.input.SetWidth(l.inputWidth)
	return layout.Flex{
		Axis:      layout.Horizontal,
		Alignment: layout.Middle,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Right: unit.Dp(l.SpaceBetween)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Min.X = gtx.Dp(l.labelWidth)
				return Label(l.theme, l.Label).Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return l.input.Layout(gtx)
			// if l.editorWidth > 0 {
			//	gtx.Constraints.Min.X = gtx.Dp(l.editorWidth)
			// } else {
			//	gtx.Constraints.Min.X = gtx.Constraints.Max.X
			// }
			// return widget.Border{
			//	Color:        borderColor,
			//	Width:        unit.Dp(1),
			//	CornerRadius: unit.Dp(4),
			// }.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//	return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			//		editor := material.Editor(l.theme.Material(), l.Editor, l.Hint)
			//		editor.Color = l.theme.TextColor
			//		editor.HintColor = theme.LightBlue
			//		editor.SelectionColor = l.theme.TextSelectionColor
			//		return editor.Layout(gtx)
			//	})
			// })
		}),
	)
}
