/**
 * Created by Goland
 * @file   line.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/23 14:39
 * @desc   日志
 */

package widgets

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
	"strings"
)

type Log struct {
	th      *theme.Theme
	scroll  *Scroll
	editor  widget.Editor
	win     *app.Window
	logData []string
}

func NewLog(th *theme.Theme, win *app.Window) *Log {
	return &Log{
		th:     th,
		scroll: NewScroll(th),
		win:    win,
	}
}

func (l *Log) SetLogData(data string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("SetLogData panic: ", err)
			l.SetLogData(data)
		}
	}()
	l.logData = append(l.logData, data)
	l.editor.SetText(strings.Join(l.logData, "\n"))
	l.win.Invalidate()
}

func (l *Log) Layout(gtx layout.Context) layout.Dimensions {
	l.scroll.SetElementList([]layout.Widget{
		func(gtx layout.Context) layout.Dimensions {
			med := material.Editor(l.th.Material(), &l.editor, "")
			med.TextSize = unit.Sp(16)
			med.Color = l.th.Color.HintTextColor
			return med.Layout(gtx)
		},
	})
	return l.scroll.Layout(gtx)
}
