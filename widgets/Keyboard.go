/**
 * Created by Goland
 * @file   Keyboard.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/12 15:46
 * @desc   Keyboard.go
 */

package widgets

import (
	"gioui.org/app"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"github.com/x-module/gioui-plugins/theme"
)

// eventFunc 定义了一个处理事件的函数类型。
type eventFunc func(e event.Event)

// Keyboard 定义了一个键盘事件处理结构体。
type Keyboard struct {
	theme    *theme.Theme   // theme 指向当前使用的主题。
	win      *app.Window    // win 是与键盘事件关联的窗口。
	filters  []event.Filter // filters 是需要过滤的事件列表。
	keyEvent eventFunc      // keyEvent 是键盘事件的处理函数。
}

// NewKeyboard 创建并返回一个新的 Keyboard 实例。
func NewKeyboard(theme *theme.Theme, win *app.Window) *Keyboard {
	return &Keyboard{
		theme: theme,
		win:   win,
		filters: []event.Filter{
			key.Filter{
				Required: key.ModShortcut,
				Name:     "V",
			},
			key.Filter{
				Name: key.NameReturn,
			},
		},
	}
}

// SetFilters 允许设置额外的事件过滤器。
func (k *Keyboard) SetFilters(filters []event.Filter) *Keyboard {
	k.filters = filters
	return k
}

// Event 为键盘事件设置处理函数。
func (k *Keyboard) Event(keyEvent eventFunc) *Keyboard {
	k.keyEvent = keyEvent
	return k
}

// Layout 定义了键盘组件的布局逻辑。
func (k *Keyboard) Layout(gtx layout.Context, content layout.Widget) layout.Dimensions {
	area := clip.Rect{Max: gtx.Constraints.Max}.Push(gtx.Ops)
	event.Op(gtx.Ops, k.win)
	for {
		keyEvent, ok := gtx.Event(k.filters...)
		if !ok {
			break
		}
		if k.keyEvent != nil {
			k.keyEvent(keyEvent)
		}
	}
	dism := content(gtx)
	area.Pop()
	return dism
}
