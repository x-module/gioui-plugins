/**
 * Created by Goland
 * @file   rich_text.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/6 12:07
 * @desc   富文本.go
 */

package widgets

import (
	"gioui.org/gesture"
	"gioui.org/layout"
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

func (r *RichText) AddSpan(span richtext.SpanStyle) *RichText {
	r.spans = append(r.spans, span)
	return r
}

func NewRichText(th *theme.Theme) *RichText {
	return &RichText{theme: th}
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
