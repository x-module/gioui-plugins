package widgets

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/theme"
)

type LogViewer struct {
	theme       *theme.Theme
	lines       []string
	selectables []*widget.Selectable
	list        *widget.List
}

func NewLogViewer(th *theme.Theme, scrollToEnd ...bool) *LogViewer {
	logView := &LogViewer{
		theme: th,
		list: &widget.List{
			List: layout.List{
				Axis:        layout.Vertical,
				ScrollToEnd: true,
				Alignment:   layout.Baseline,
			},
		},
	}
	if len(scrollToEnd) > 0 {
		logView.list.ScrollToEnd = scrollToEnd[0]
	}
	return logView
}

func (j *LogViewer) Clean() {
	j.lines = nil
	j.selectables = nil
}

func (j *LogViewer) SetData(log string) {
	j.lines = append(j.lines, log)
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}

func (j *LogViewer) Layout(gtx layout.Context, theme *theme.Theme) layout.Dimensions {
	border := widget.Border{
		Color:        j.theme.Color.DefaultBorderGrayColor,
		Width:        unit.Dp(1),
		CornerRadius: unit.Dp(4),
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(3).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.List(theme.Material(), j.list).Layout(gtx, len(j.lines), func(gtx layout.Context, i int) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							l := material.Label(theme.Material(), theme.TextSize, fmt.Sprintf("%d", i+1))
							l.Font.Weight = font.Medium
							l.Color = j.theme.Color.LogTextWhiteColor
							l.SelectionColor = j.theme.Color.TextSelectionColor
							l.Alignment = text.End
							return l.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							l := material.Label(theme.Material(), theme.TextSize, j.lines[i])
							l.State = j.selectables[i]
							l.Color = j.theme.Color.LogTextWhiteColor
							l.SelectionColor = j.theme.Color.TextSelectionColor
							l.TextSize = j.theme.Size.DefaultTextSize
							return l.Layout(gtx)
						})
					}),
				)
			})
		})
	})
}
