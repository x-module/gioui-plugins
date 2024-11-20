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
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
	"strings"
	"sync"
)

type Log struct {
	sync.Mutex
	th          *theme.Theme
	scroll      *Scroll
	editor      *widget.Editor
	logData     []string
	editElement material.EditorStyle
	border      widget.Border
	logs        chan string
}

func NewLog(th *theme.Theme) *Log {
	log := &Log{
		th:     th,
		scroll: NewScroll(th),
		editor: &widget.Editor{SingleLine: false},
		border: widget.Border{
			Color:        th.Color.DefaultBgGrayColor,
			Width:        unit.Dp(1),
			CornerRadius: unit.Dp(4),
		},
		logs: make(chan string, 10000),
	}
	log.editElement = material.Editor(th.Material(), log.editor, "")
	log.editElement.TextSize = th.Size.DefaultTextSize
	log.editElement.Color = th.Color.HintTextColor
	log.scroll.SetElementList([]layout.Widget{
		log.editElement.Layout,
	})
	return log
}

func (l *Log) SetTextSize(size unit.Sp) *Log {
	l.editElement.TextSize = size
	return l
}
func (l *Log) SetTextColor(color color.NRGBA) *Log {
	l.editElement.Color = color
	return l
}

func (l *Log) SetData(data string) *Log {
	l.logs <- data
	return l
}

func (l *Log) Reset() *Log {
	l.logData = []string{}
	// 关闭旧通道
	close(l.logs)
	// 创建新通道
	l.logs = make(chan string, 10000) // 使用相同的缓冲区大小
	return l
}

func (l *Log) Layout(gtx layout.Context) layout.Dimensions {
	// 读取l.logs 中的数据并且添加到 l.logData 中,不等待新数据
	count := len(l.logs)
	for i := 0; i < count; i++ {
		if data, ok := <-l.logs; ok {
			l.logData = append(l.logData, strings.TrimSpace(data))
		}
	}
	l.editor.SetText(strings.Join(l.logData, "\n"))
	return l.border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return l.scroll.Layout(gtx)
		})
	})
}
