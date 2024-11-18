/**
 * Created by Goland
 * @file   line.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/23 14:39
 * @desc   日志
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/op"
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
	logData []string
}

func NewLog(th *theme.Theme) *Log {
	return &Log{
		th:     th,
		scroll: NewScroll(th),
	}
}

func (l *Log) InitData(data string) *Log {
	data = strings.TrimSpace(data)
	l.logData = append(l.logData, data)
	l.editor.SetText(strings.Join(l.logData, "\n"))
	return l
}
func (l *Log) SetData(gtx layout.Context, data string) *Log {
	data = strings.TrimSpace(data)
	l.logData = append(l.logData, data)
	l.editor.SetText(strings.Join(l.logData, "\n"))
	gtx.Execute(op.InvalidateCmd{})
	return l
}

func (l *Log) Reset() *Log {
	l.logData = []string{}
	return l
}

func (l *Log) Layout(gtx layout.Context) layout.Dimensions {
	l.scroll.SetElementList([]layout.Widget{
		func(gtx layout.Context) layout.Dimensions {
			med := material.Editor(l.th.Material(), &l.editor, "")
			med.TextSize = l.th.Size.DefaultTextSize
			med.Color = l.th.Color.HintTextColor
			return med.Layout(gtx)
		},
	})
	border := widget.Border{
		Color:        l.th.Color.DefaultBgGrayColor,
		Width:        unit.Dp(1),
		CornerRadius: unit.Dp(4),
	}
	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return l.scroll.Layout(gtx)
		})
	})
}
