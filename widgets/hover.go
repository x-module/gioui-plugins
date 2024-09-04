/**
 * Created by Goland
 * @file   hover.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/4 11:57
 * @desc   hover.go
 */

package widgets

import (
	"gioui.org/x/component"
	"github.com/x-module/gioui-plugins/utils"
	"image"
	"image/color"
	"time"

	"gioui.org/io/event"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Hover implements a material design tool tip as defined at:
// https://material.io/components/tooltips#specs
type Hover struct {
	// Inset defines the interior padding of the Hover.
	layout.Inset
	// CornerRadius defines the corner radius of the RRect background.
	// of the tooltip.
	CornerRadius unit.Dp
	// Text defines the content of the tooltip.
	layout.Widget
	// Bg defines the color of the RRect background.
	Bg     color.NRGBA
	widget layout.Widget
}

func NewHover(th *material.Theme, w layout.Widget) Hover {
	return Hover{
		Inset: layout.Inset{
			Top:    unit.Dp(6),
			Bottom: unit.Dp(6),
			Left:   unit.Dp(8),
			Right:  unit.Dp(8),
		},
		Bg:           utils.WithAlpha(th.Fg, 220),
		CornerRadius: unit.Dp(4),
		widget:       w,
	}
}

// Layout renders the tooltip.
func (t Hover) Layout(gtx C) D {
	return layout.Stack{}.Layout(gtx,
		layout.Expanded(func(gtx C) D {
			radius := gtx.Dp(t.CornerRadius)
			paint.FillShape(gtx.Ops, t.Bg, clip.RRect{
				Rect: image.Rectangle{
					Max: gtx.Constraints.Min,
				},
				NW: radius,
				NE: radius,
				SW: radius,
				SE: radius,
			}.Op(gtx.Ops))
			return D{}
		}),
		layout.Stacked(func(gtx C) D {
			return t.Inset.Layout(gtx, t.widget)
		}),
	)
}

// InvalidateDeadline helps to ensure that a frame is generated at a specific
// point in time in the future. It does this by always requesting a future
// invalidation at its target time until it reaches its target time. This
// makes animating delays much cleaner.
type InvalidateDeadline struct {
	// The time at which a frame needs to be drawn.
	Target time.Time
	// Whether the deadline is active.
	Active bool
}

// SetTarget configures a specific time in the future at which a frame should
// be rendered.
func (i *InvalidateDeadline) SetTarget(t time.Time) {
	i.Active = true
	i.Target = t
}

// Process checks the current frame time and either requests a future invalidation
// or does nothing. It returns whether the current frame is the frame requested
// by the last call to SetTarget.
func (i *InvalidateDeadline) Process(gtx C) bool {
	if !i.Active {
		return false
	}
	if gtx.Now.Before(i.Target) {
		gtx.Execute(op.InvalidateCmd{At: i.Target})
		return false
	}
	i.Active = false
	return true
}

// ClearTarget cancels a request to invalidate in the future.
func (i *InvalidateDeadline) ClearTarget() {
	i.Active = false
}

// HoverArea holds the state information for displaying a tooltip. The zero
// value will choose sensible defaults for all fields.
type HoverArea struct {
	component.VisibilityAnimation
	Hover     InvalidateDeadline
	Press     InvalidateDeadline
	LongPress InvalidateDeadline
	init      bool
	// HoverDelay is the delay between the cursor entering the tip area
	// and the tooltip appearing.
	HoverDelay time.Duration
	// LongPressDelay is the required duration of a press in the area for
	// it to count as a long press.
	LongPressDelay time.Duration
	// LongPressDuration is the amount of time the tooltip should be displayed
	// after being triggered by a long press.
	LongPressDuration time.Duration
	// FadeDuration is the amount of time it takes the tooltip to fade in
	// and out.
	FadeDuration time.Duration
}

const (
	HoverAreaHoverDelay        = time.Millisecond * 500
	HoverAreaLongPressDuration = time.Millisecond * 1500
	HoverAreaFadeDuration      = time.Millisecond * 250
	longPressTheshold          = time.Millisecond * 500
)

// Layout renders the provided widget with the provided tooltip. The tooltip
// will be summoned if the widget is hovered or long-pressed.
func (t *HoverArea) Layout(gtx C, tip Hover, w layout.Widget) D {
	if !t.init {
		t.init = true
		t.VisibilityAnimation.State = component.Invisible
		if t.HoverDelay == time.Duration(0) {
			t.HoverDelay = HoverAreaHoverDelay
		}
		if t.LongPressDelay == time.Duration(0) {
			t.LongPressDelay = longPressTheshold
		}
		if t.LongPressDuration == time.Duration(0) {
			t.LongPressDuration = HoverAreaLongPressDuration
		}
		if t.FadeDuration == time.Duration(0) {
			t.FadeDuration = HoverAreaFadeDuration
		}
		t.VisibilityAnimation.Duration = t.FadeDuration
	}
	for {
		ev, ok := gtx.Event(pointer.Filter{
			Target: t,
			Kinds:  pointer.Press | pointer.Release | pointer.Enter | pointer.Leave,
		})
		if !ok {
			break
		}
		e, ok := ev.(pointer.Event)
		if !ok {
			continue
		}
		switch e.Kind {
		case pointer.Enter:
			t.Hover.SetTarget(gtx.Now.Add(t.HoverDelay))
		case pointer.Leave:
			t.VisibilityAnimation.Disappear(gtx.Now)
			t.Hover.ClearTarget()
		case pointer.Press:
			t.Press.SetTarget(gtx.Now.Add(t.LongPressDelay))
		case pointer.Release:
			t.Press.ClearTarget()
		case pointer.Cancel:
			t.Hover.ClearTarget()
			t.Press.ClearTarget()
		default:

		}
	}
	if t.Hover.Process(gtx) {
		t.VisibilityAnimation.Appear(gtx.Now)
	}
	if t.Press.Process(gtx) {
		t.VisibilityAnimation.Appear(gtx.Now)
		t.LongPress.SetTarget(gtx.Now.Add(t.LongPressDuration))
	}
	if t.LongPress.Process(gtx) {
		t.VisibilityAnimation.Disappear(gtx.Now)
	}
	return layout.Stack{}.Layout(gtx,
		layout.Stacked(w),
		layout.Expanded(func(gtx C) D {
			defer pointer.PassOp{}.Push(gtx.Ops).Pop()
			defer clip.Rect(image.Rectangle{Max: gtx.Constraints.Min}).Push(gtx.Ops).Pop()
			event.Op(gtx.Ops, t)
			originalMin := gtx.Constraints.Min
			gtx.Constraints.Min = image.Point{}

			if t.Visible() {
				macro := op.Record(gtx.Ops)
				tip.Bg = component.Interpolate(color.NRGBA{}, tip.Bg, t.VisibilityAnimation.Revealed(gtx))
				dims := tip.Layout(gtx)
				call := macro.Stop()
				xOffset := (originalMin.X / 2) - (dims.Size.X / 2)
				yOffset := originalMin.Y
				macro = op.Record(gtx.Ops)
				op.Offset(image.Pt(xOffset, yOffset)).Add(gtx.Ops)
				call.Add(gtx.Ops)
				call = macro.Stop()
				op.Defer(gtx.Ops, call)
			}
			return D{}
		}),
	)
}
