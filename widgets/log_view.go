package widgets

import (
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/golang-module/carbon"
	config "github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
)

type LogType int
type LogData struct {
	Log  string  `json:"log"`
	Type LogType `json:"type"`
}

const (
	Debug LogType = 1 + iota
	Info
	Warn
	Error
	Fatal
)

type LogViewer struct {
	theme       *theme.Theme
	lines       []LogData
	selectables []*widget.Selectable
	list        *widget.List
	colorMap    map[LogType]color.NRGBA
	signMap     map[LogType]string
	showNumber  bool
	shaper      *text.Shaper
}

func NewLogViewer(th *theme.Theme, scrollToEnd ...bool) *LogViewer {
	logView := &LogViewer{
		theme: th,
		list: &widget.List{
			List: layout.List{
				Axis:        layout.Vertical,
				ScrollToEnd: true,
				Alignment:   layout.Middle,
			},
		},
		colorMap: map[LogType]color.NRGBA{
			Debug: th.Color.LogDebugColor,
			Info:  th.Color.LogInfoColor,
			Warn:  th.Color.LogWarnColor,
			Error: th.Color.LogErrorColor,
			Fatal: th.Color.LogFataCorColor,
		},
		signMap: map[LogType]string{
			Debug: "DEBU",
			Info:  "INFO",
			Warn:  "WARN",
			Error: "ERRO",
			Fatal: "FATA",
		},
	}

	logView.list.List.Position.BeforeEnd = false

	data, err := config.Asset("data/fonts/Monaco.ttf")
	if err != nil {
		panic(err)
	}
	face, err := opentype.ParseCollection(data)
	if err != nil {
		panic(err)
	}
	logView.shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(face))
	if len(scrollToEnd) > 0 {
		logView.list.ScrollToEnd = scrollToEnd[0]
	}
	return logView
}

func (j *LogViewer) ShowNumber() {
	j.showNumber = true
}

func (j *LogViewer) Clean() {
	j.lines = []LogData{}
	j.selectables = []*widget.Selectable{}
}

// Log format log
func (j *LogViewer) formatLog(log LogData) LogData {
	logMsg := fmt.Sprintf("%s [%s] %s", j.signMap[log.Type], carbon.Now().ToDateTimeString(), log.Log)
	log.Log = logMsg
	return log
}

func (j *LogViewer) Debug(log string) {
	j.lines = append(j.lines, j.formatLog(LogData{Log: log, Type: Debug}))
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}

func (j *LogViewer) Info(log string) {
	j.lines = append(j.lines, j.formatLog(LogData{Log: log, Type: Info}))
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}

func (j *LogViewer) Warn(log string) {
	j.lines = append(j.lines, j.formatLog(LogData{Log: log, Type: Warn}))
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}
func (j *LogViewer) Error(log string) {
	j.lines = append(j.lines, j.formatLog(LogData{Log: log, Type: Error}))
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}

func (j *LogViewer) Fatal(log string) {
	j.lines = append(j.lines, j.formatLog(LogData{Log: log, Type: Fatal}))
	j.selectables = make([]*widget.Selectable, len(j.lines))
	for i := range j.selectables {
		j.selectables[i] = &widget.Selectable{}
	}
}

func (j *LogViewer) Layout(gtx layout.Context) layout.Dimensions {
	border := widget.Border{
		Color:        j.theme.Color.DefaultBorderGrayColor,
		Width:        unit.Dp(1),
		CornerRadius: unit.Dp(4),
	}

	return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(3).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			list := material.List(j.theme.Material(), j.list)
			return list.Layout(gtx, len(j.lines), func(gtx layout.Context, i int) layout.Dimensions {
				logColor := j.colorMap[j.lines[i].Type]
				return layout.Inset{Top: unit.Dp(2), Bottom: unit.Dp(0)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							if !j.showNumber {
								return layout.Dimensions{}
							}
							return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								l := material.Label(j.theme.Material(), j.theme.TextSize, fmt.Sprintf("%d", i+1))
								l.Font.Weight = font.Medium
								l.Color = logColor
								// l.Font.Style = font.Regular
								l.SelectionColor = j.theme.Color.TextSelectionColor
								l.Alignment = text.End
								l.TextSize = j.theme.Size.DefaultLogTextSize
								return l.Layout(gtx)
							})
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
								l := material.Label(j.theme.Material(), j.theme.TextSize, j.lines[i].Log)
								l.State = j.selectables[i]
								l.Color = logColor
								l.Font.Weight = font.Black
								l.Shaper = j.shaper
								l.SelectionColor = j.theme.Color.TextSelectionColor
								l.TextSize = j.theme.Size.DefaultLogTextSize
								return l.Layout(gtx)
							})
						}),
					)
				})
			})
		})
	})
}
