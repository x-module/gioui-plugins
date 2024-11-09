/**
 * Created by Goland
 * @file   step.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/8 14:33
 * @desc   step.go
 */

package widgets

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"image/color"
)

var unDoColor = color.NRGBA{R: 53, G: 54, B: 56, A: 255}

var doneColor = color.NRGBA{R: 46, G: 204, B: 113, A: 255}

type StepItem struct {
	Title string
}

type Step struct {
	th          *theme.Theme
	steps       []StepItem
	currentStep int
}

func NewStep(th *theme.Theme) *Step {
	return &Step{
		th: th,
	}
}

func (s *Step) SetCurrentStep(currentStep int) *Step {
	s.currentStep = currentStep
	return s
}
func (s *Step) SetSteps(steps []StepItem) *Step {
	s.steps = steps
	return s
}

func (s *Step) Label(txt string, color color.NRGBA) material.LabelStyle {
	label := material.Label(s.th.Material(), s.th.Size.DefaultTextSize, txt)
	label.Color = color
	// label.Font.Weight = font.Bold
	label.TextSize = s.th.Size.DefaultTextSize
	return label
}

func (s *Step) LayoutStep() ([]layout.FlexChild, []layout.FlexChild) {
	var stepChildren = []layout.FlexChild{
		layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(3)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: unit.Dp(6)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return resource.AVFiberManualRecordIcon.Layout(gtx, doneColor)
				})
			})
		}),
	}
	var tipChildren = []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Dp(unit.Dp(105))
			return layout.Inset{Top: unit.Dp(3), Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return s.Label("Start", doneColor).Layout(gtx)
			})
		}),
	}
	for key, item := range s.steps {
		stepChildren = append(stepChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(13)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if key < s.currentStep {
					return s.Label("──────────", doneColor).Layout(gtx)
				}
				return s.Label("──────────", unDoColor).Layout(gtx)
			})
		}))
		stepChildren = append(stepChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Top: unit.Dp(3)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Inset{Top: unit.Dp(6)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					if key < s.currentStep {
						return resource.AVFiberManualRecordIcon.Layout(gtx, doneColor)
					}
					return resource.AVFiberManualRecordIcon.Layout(gtx, unDoColor)
				})
			})
		}))
		tipChildren = append(tipChildren, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			gtx.Constraints.Min.X = gtx.Dp(unit.Dp(109))
			return layout.Inset{Top: unit.Dp(3)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if key < s.currentStep {
					return s.Label(item.Title, doneColor).Layout(gtx)
				}
				return s.Label(item.Title, unDoColor).Layout(gtx)
			})
		}))
	}
	return stepChildren, tipChildren
}

func (s *Step) Layout(gtx layout.Context) layout.Dimensions {
	stepChildren, tipChildren := s.LayoutStep()
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				stepChildren...,
			)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				tipChildren...,
			)
		}),
	)
}
