/**
 * Created by Goland
 * @file   utils.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/7 14:05
 * @desc   utils.go
 */

package utils

import (
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image"
	"image/color"
)

func ColorBackground(gtx layout.Context, point image.Point, color color.NRGBA) {
	rect := clip.Rect{
		Max: point,
	}
	paint.FillShape(gtx.Ops, color, rect.Op())
}

func WithAlpha(c color.NRGBA, a uint8) color.NRGBA {
	return color.NRGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: a,
	}
}

// Divider

func DrawLineFlex(background color.NRGBA, height, width unit.Dp) layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return DrawLine(gtx, background, height, width)
	})
}

func DrawLine(gtx layout.Context, background color.NRGBA, height, width unit.Dp) layout.Dimensions {
	w, h := gtx.Dp(width), gtx.Dp(height)
	tabRect := image.Rect(0, 0, w, h)
	paint.FillShape(gtx.Ops, background, clip.Rect(tabRect).Op())
	return layout.Dimensions{Size: image.Pt(w, h)}
}
