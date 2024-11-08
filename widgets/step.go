/**
 * Created by Goland
 * @file   step.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/11/8 14:33
 * @desc   step.go
 */

package widgets

import (
	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"golang.org/x/exp/shiny/materialdesign/colornames"
	"image/color"
)

type Step struct {
	th *theme.Theme
}

func NewStep(th *theme.Theme) *Step {
	return &Step{
		th: th,
	}
}

func (s *Step) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Stack{Alignment: layout.Center}.Layout(gtx,
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					NewLine(s.th).Color(s.th.Color.GreenColor).Width(3).Line(gtx, f32.Pt(0, 0), f32.Pt(float32(gtx.Constraints.Max.X-100), 0))
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceBetween}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.Spacer{Width: unit.Dp(110)}.Layout(gtx)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							// utils.DrawBackground(gtx, gtx.Constraints.Max, s.th.Color.BlueColor)
							gtx.Constraints.Min.X = gtx.Dp(unit.Dp(40))
							return resource.AVFiberManualRecordIcon.Layout(gtx, s.th.Color.GreenColor)
						}),
					)
				}),
			)
		}),
	)
}

func drawStepProgressBar(gtx layout.Context, currentStep, totalSteps int) layout.Dimensions {
	// 进度条背景
	// 计算每一步的间距
	stepWidth := gtx.Constraints.Max.X / totalSteps

	// 绘制每一步的圆点
	for i := 0; i < totalSteps; i++ {
		clr := colornames.Grey400
		if i < currentStep {
			clr = colornames.LightGreenA700
		}

		drawCircle(gtx, color.NRGBA(clr), stepWidth*(i+1)-stepWidth/2, 10, 10)
	}

	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func drawCircle(gtx layout.Context, clr color.NRGBA, x, y, radius int) {

}
