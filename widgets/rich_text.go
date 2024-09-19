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
	"gioui.org/font"
	"gioui.org/gesture"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/x/styledtext"
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
	"log"
	"time"
)

type actionFun func(gtx layout.Context, content string)
type RichText struct {
	th    *theme.Theme
	state InteractiveText
	spans []SpanStyle

	click     actionFun
	hover     actionFun
	unHover   actionFun
	longPress actionFun
}

func (r *RichText) AddSpan(spans []SpanStyle) *RichText {
	for key := range spans {
		spans[key].Interactive = true
	}
	r.spans = append(r.spans, spans...)
	return r
}
func (r *RichText) UpdateSpan(spans []SpanStyle) *RichText {
	for key := range spans {
		spans[key].Interactive = true
	}
	r.spans = spans
	return r
}
func NewRichText(th *theme.Theme) *RichText {
	return &RichText{th: th}
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
func (r *RichText) update(gtx layout.Context) {
	for {
		span, event, ok := r.state.Update(gtx)
		if !ok {
			break
		}
		content, _ := span.Content()
		switch event.Type {
		case RcClick:
			log.Println(event.ClickData.Kind)
			if event.ClickData.Kind == gesture.KindClick {
				if r.click != nil {
					r.click(gtx, content)
				}
			}
		case RcHover:
			if r.hover != nil {
				r.hover(gtx, content)
			}
		case RcUnHover:
			if r.unHover != nil {
				r.unHover(gtx, content)
			}
		case RcLongPress:
			if r.longPress != nil {
				r.longPress(gtx, content)
			}
		}
	}
}

func (r *RichText) Layout(gtx layout.Context) layout.Dimensions {
	r.update(gtx)
	return r.Text(&r.state, r.th.Shaper, r.spans...).Layout(gtx)
}
func (t RichText) UnderLineLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	dims := widget(gtx)
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return widget(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return NewLine(t.th).Color(t.th.Color.MarkdownDefaultColor).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(dims.Size.X), 0)).Layout(gtx)
		}),
	)
}
func (t RichText) DeleteLineLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	dims := widget(gtx)
	return layout.Stack{Alignment: layout.W}.Layout(gtx,
		layout.Stacked(func(gtx layout.Context) layout.Dimensions {
			return widget(gtx)
		}),
		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
			return NewLine(t.th).Color(t.th.Color.MarkdownDefaultColor).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(dims.Size.X), 0)).Layout(gtx)
		}),
	)
}
func (t RichText) MarkLayout(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return layout.Background{}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		defer clip.Rect{
			Max: gtx.Constraints.Min,
		}.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, t.th.Color.GreenColor)
		return layout.Dimensions{Size: gtx.Constraints.Min}
	}, widget)
}

// ======================================================================================

// LongPressDuration is the default duration of a long press gesture.
// Override this variable to change the detection threshold.
var LongPressDuration time.Duration = 250 * time.Millisecond

// EventType describes a kind of iteraction with rich text.
type EventType uint8

const (
	RcHover EventType = iota
	RcUnHover
	RcLongPress
	RcClick
)

// Event describes an interaction with rich text.
type Event struct {
	Type EventType
	// ClickData is only populated if Type == Clicked
	ClickData gesture.ClickEvent
}

// InteractiveSpan holds the persistent state of rich text that can
// be interacted with by the user. It can report clicks, hovers, and
// long-presses on the text.
type InteractiveSpan struct {
	click        gesture.Click
	pressing     bool
	hovering     bool
	longPressed  bool
	pressStarted time.Time
	contents     string
	metadata     map[string]interface{}
}

func (i *InteractiveSpan) Update(gtx layout.Context) (Event, bool) {
	if i == nil {
		return Event{}, false
	}
	for {
		e, ok := i.click.Update(gtx.Source)
		if !ok {
			break
		}
		switch e.Kind {
		case gesture.KindClick:
			i.pressing = false
			if i.longPressed {
				i.longPressed = false
			} else {
				return Event{Type: RcClick, ClickData: e}, true
			}
		case gesture.KindPress:
			i.pressStarted = gtx.Now
			i.pressing = true
		case gesture.KindCancel:
			i.pressing = false
			i.longPressed = false
		}
	}
	if isHovered := i.click.Hovered(); isHovered != i.hovering {
		i.hovering = isHovered
		if isHovered {
			return Event{Type: RcHover}, true
		} else {
			return Event{Type: RcUnHover}, true
		}
	}

	if !i.longPressed && i.pressing && gtx.Now.Sub(i.pressStarted) > LongPressDuration {
		i.longPressed = true
		return Event{Type: RcLongPress}, true
	}
	return Event{}, false
}

// Layout adds the pointer input op for this interactive span and updates its
// state. It uses the most recent pointer.AreaOp as its input area.
func (i *InteractiveSpan) Layout(gtx layout.Context) layout.Dimensions {
	for {
		_, ok := i.Update(gtx)
		if !ok {
			break
		}
	}
	if i.pressing && !i.longPressed {
		gtx.Execute(op.InvalidateCmd{})
	}
	defer clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops).Pop()

	pointer.CursorPointer.Add(gtx.Ops)
	i.click.Add(gtx.Ops)
	return layout.Dimensions{}
}

// Content returns the text content of the interactive span as well as the
// metadata associated with it.
func (i *InteractiveSpan) Content() (string, map[string]interface{}) {
	return i.contents, i.metadata
}

// Get looks up a metadata property on the interactive span.
func (i *InteractiveSpan) Get(key string) interface{} {
	return i.metadata[key]
}

// InteractiveText holds persistent state for a block of text containing
// spans that may be interactive.
type InteractiveText struct {
	Spans       []InteractiveSpan
	lastUpdate  time.Time
	updateIndex int
}

// resize makes sure that there are exactly n interactive spans.
func (i *InteractiveText) resize(n int) {
	if n == 0 && i == nil {
		return
	}

	if cap(i.Spans) >= n {
		i.Spans = i.Spans[:n]
	} else {
		i.Spans = make([]InteractiveSpan, n)
	}
}

// Update returns the first span with unprocessed events and the events that
// need processing for it.
func (i *InteractiveText) Update(gtx layout.Context) (*InteractiveSpan, Event, bool) {
	if i == nil {
		return nil, Event{}, false
	}
	if i.lastUpdate != gtx.Now {
		i.lastUpdate = gtx.Now
		i.updateIndex = 0
	}
	for k := i.updateIndex; k < len(i.Spans); k++ {
		i.updateIndex = k
		span := &i.Spans[k]
		for {
			ev, ok := span.Update(gtx)
			if !ok {
				break
			}
			return span, ev, true
		}
	}
	return nil, Event{}, false
}

// SpanStyle describes the appearance of a span of styled text.
type SpanStyle struct {
	Font           font.Font
	Size           unit.Sp
	Color          color.NRGBA
	Content        string
	Interactive    bool
	metadata       map[string]interface{}
	interactiveIdx int
}

// Set configures a metadata key-value pair on the span that can be
// retrieved if the span is interacted with. If the provided value
// is empty, the key will be deleted from the metadata.
func (ss *SpanStyle) Set(key string, value interface{}) {
	if value == "" {
		if ss.metadata != nil {
			delete(ss.metadata, key)
			if len(ss.metadata) == 0 {
				ss.metadata = nil
			}
		}
		return
	}
	if ss.metadata == nil {
		ss.metadata = make(map[string]interface{})
	}
	ss.metadata[key] = value
}

// DeepCopy returns an identical SpanStyle with its own copy of its metadata.
func (ss SpanStyle) DeepCopy() SpanStyle {
	out := ss
	if len(ss.metadata) > 0 {
		md := make(map[string]interface{})
		for k, v := range ss.metadata {
			md[k] = v
		}
		out.metadata = md
	}
	return out
}

// TextStyle presents rich text.
type TextStyle struct {
	th         *theme.Theme
	State      *InteractiveText
	Styles     []SpanStyle
	Alignment  text.Alignment
	WrapPolicy styledtext.WrapPolicy
	*text.Shaper
}

// Text constructs a TextStyle.
func (r *RichText) Text(state *InteractiveText, shaper *text.Shaper, styles ...SpanStyle) TextStyle {
	return TextStyle{
		th:     r.th,
		State:  state,
		Styles: styles,
		Shaper: shaper,
	}
}
func (t TextStyle) Layout(gtx layout.Context) layout.Dimensions {
	for {
		_, _, ok := t.State.Update(gtx)
		if !ok {
			break
		}
	}
	// OPT(dh): it'd be nice to avoid this allocation
	styles := make([]styledtext.SpanStyle, len(t.Styles))
	numInteractive := 0
	for i := range t.Styles {
		st := &t.Styles[i]
		if st.Interactive {
			st.interactiveIdx = numInteractive
			numInteractive++
		}
		styles[i] = styledtext.SpanStyle{
			Font:    st.Font,
			Size:    st.Size,
			Color:   st.Color,
			Content: st.Content,
		}
	}
	t.State.resize(numInteractive)
	text := styledtext.Text(t.Shaper, styles...)
	text.WrapPolicy = t.WrapPolicy
	text.Alignment = t.Alignment
	return text.Layout(gtx, func(gtx layout.Context, i int, _ layout.Dimensions) {
		span := &t.Styles[i]
		if !span.Interactive {
			return
		}
		st := &t.State.Spans[span.interactiveIdx]
		st.contents = span.Content
		st.metadata = span.metadata
		st.Layout(gtx)
	})
}
