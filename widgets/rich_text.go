/**
 * Created by Goland
 * @file   rich_text.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/6 12:07
 * @desc   富文本.go
 */

package widgets

import (
	"gioui.org/f32"
	"gioui.org/gesture"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/x/richtext"
	"github.com/x-module/gioui-plugins/theme"
	"log"
)

type actionFun func(gtx layout.Context, content string)
type RichText struct {
	theme *theme.Theme
	state richtext.InteractiveText
	spans []richtext.SpanStyle

	click     actionFun
	hover     actionFun
	unHover   actionFun
	longPress actionFun
}

func NewRichText(th *theme.Theme) *RichText {
	return &RichText{theme: th}
}

func (r *RichText) AddSpan(spans []richtext.SpanStyle) *RichText {
	for key := range spans {
		spans[key].Interactive = true
	}
	r.spans = append(r.spans, spans...)
	return r
}
func (r *RichText) UpdateSpan(spans []richtext.SpanStyle) *RichText {
	for key := range spans {
		spans[key].Interactive = true
	}
	r.spans = spans
	return r
}

func (r *RichText) OnClick(f actionFun) *RichText {
	r.click = f
	return r
}

func (r *RichText) OnHover(f actionFun) *RichText {
	r.hover = f
	return r
}
func (r *RichText) OnUnHover(f actionFun) *RichText {
	r.unHover = f
	return r
}
func (r *RichText) OnLongPress(f actionFun) *RichText {
	r.longPress = f
	return r
}

func (r *RichText) Layout(gtx layout.Context) layout.Dimensions {
	for {
		span, event, ok := r.state.Update(gtx)
		if !ok {
			break
		}
		content, _ := span.Content()
		switch event.Type {
		case richtext.Click:
			log.Println(event.ClickData.Kind)
			if event.ClickData.Kind == gesture.KindClick {
				if r.click != nil {
					r.click(gtx, content)
				}
			}
		case richtext.Hover:
			if r.hover != nil {
				r.hover(gtx, content)
			}
		case richtext.Unhover:
			if r.unHover != nil {
				r.unHover(gtx, content)
			}
		case richtext.LongPress:
			if r.longPress != nil {
				r.longPress(gtx, content)
			}
		}
	}
	// render the rich text into the operation list
	return richtext.Text(&r.state, r.theme.Shaper, r.spans...).Layout(gtx)
}

func (r RichText) UnderLineLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	dims := widget(gtx)
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return widget(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return NewLine(r.theme).Color(r.theme.Color.MarkdownDefaultColor).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(dims.Size.X), 0)).Layout(gtx)
		}),
	)
}
func (r RichText) DeleteLineLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	dims := widget(gtx)
	return layout.Stack{Alignment: layout.W}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return widget(gtx)
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return NewLine(r.theme).Color(r.theme.Color.MarkdownDefaultColor).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(dims.Size.X), 0)).Layout(gtx)
		}),
	)
}
func (r RichText) MarkLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return layout.Background{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		defer clip.Rect{
			Max: gtx.Constraints.Min,
		}.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, r.theme.Color.GreenColor)
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}, widget)
}
