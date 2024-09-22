/**
 * Created by Goland
 * @file   link.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/9/22 00:22
 * @desc   link.go
 */

package widgets

import (
	"fmt"
	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
	"log"
	"os/exec"
	"runtime"
)

type Link struct {
	th   *theme.Theme
	link *widget.Clickable
	name string
	url  string
}

func NewLink(th *theme.Theme) *Link {
	return &Link{
		th:   th,
		link: &widget.Clickable{},
	}
}

func (l *Link) SetLink(name, url string) *Link {
	l.name = name
	l.url = url
	return l
}

func (l *Link) Layout(gtx layout.Context) layout.Dimensions {
	// 设置链接文本样式
	label := material.Label(l.th.Material(), unit.Sp(16), l.name)
	label.Color = l.th.Color.DefaultLinkColor
	label.TextSize = l.th.Size.DefaultTextSize
	label.Alignment = text.Middle
	// 绘制链接文本
	// 处理点击事件
	for l.link.Clicked(gtx) {
		l.openBrowser(l.url)
	}
	dims := label.Layout(gtx)
	return l.link.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		// 设置鼠标光标为小手
		pointer.CursorPointer.Add(gtx.Ops)
		defer op.Offset(gtx.Constraints.Min).Push(gtx.Ops).Pop()
		if l.link.Hovered() {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return dims
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return NewLine(l.th).Color(l.th.Color.DefaultLinkColor).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(dims.Size.X), 0)).Layout(gtx)
				}),
			)
		}
		return dims
	})
}

// openBrowser 在默认浏览器中打开指定的 URL
func (l *Link) openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
