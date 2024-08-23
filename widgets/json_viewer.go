package widgets

import (
	"encoding/json"
	"fmt"
	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/gioui-plugins/theme"
	"strings"
)

type JsonViewer struct {
	theme       *theme.Theme
	data        string
	lines       []string
	selectables []*widget.Selectable
	list        *widget.List
}

func NewJsonViewer(th *theme.Theme) *JsonViewer {
	return &JsonViewer{
		theme: th,
		list: &widget.List{
			List: layout.List{
				Axis: layout.Vertical,
			},
		},
	}
}

func (j *JsonViewer) SetData(data any) {
	// 使用MarshalIndent序列化map，生成格式化的JSON字符串
	// 第二个参数是每一行输出的前缀（通常为空）
	// 第三个参数是每一级缩进的字符串，这里使用4个空格作为缩进
	formattedJSON, _ := json.MarshalIndent(data, "", "    ")
	jsonData := string(formattedJSON)
	j.data = jsonData
	j.lines = strings.Split(jsonData, "\n")
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}

func (j *JsonViewer) Layout(gtx layout.Context) layout.Dimensions {
	border := widget.Border{
		Color:        j.theme.Color.DefaultBgGrayColor,
		Width:        unit.Dp(1),
		CornerRadius: unit.Dp(4),
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(3).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return material.List(j.theme.Material(), j.list).Layout(gtx, len(j.lines), func(gtx layout.Context, i int) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							l := material.Label(j.theme.Material(), j.theme.Size.DefaultTextSize, fmt.Sprintf("%d", i+1))
							l.Font.Weight = font.Medium
							l.Color = j.theme.Color.DefaultTextWhiteColor
							l.SelectionColor = j.theme.Color.TextSelectionColor
							l.Alignment = text.End
							return l.Layout(gtx)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							l := material.Label(j.theme.Material(), j.theme.Size.DefaultTextSize, j.lines[i])
							l.State = j.selectables[i]
							l.SelectionColor = j.theme.Color.TextSelectionColor
							l.TextSize = j.theme.Size.DefaultTextSize
							l.Font.Weight = font.Medium
							l.Color = j.theme.Color.DefaultTextWhiteColor
							l.Alignment = text.End
							return l.Layout(gtx)
						})
					}),
				)
			})
		})
	})
}
