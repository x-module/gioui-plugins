/**
 * Created by Goland
 * @file   initialize.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2024/8/23 11:50
 * @desc   initialize.go
 */

package window

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image/color"
	"os"
)

type (
	DestroyFun func(err error)
	FrameFun   func(gtx layout.Context, ops op.Ops, win *app.Window)
)

type Initialize struct {
	destroy DestroyFun
	frame   FrameFun
	win     *app.Window
	options []app.Option
	bgColor color.NRGBA
}

func NewInitialize() *Initialize {
	return &Initialize{
		win: new(app.Window),
		destroy: func(err error) {
			os.Exit(1)
		},
	}
}

func (i *Initialize) BackgroundColor(color color.NRGBA) {
	i.bgColor = color
}

func (i *Initialize) Title(t string) *Initialize {
	i.options = append(i.options, app.Title(t))
	return i
}

func (i *Initialize) Size(width int, height int) *Initialize {
	i.options = append(i.options, app.Size(unit.Dp(width), unit.Dp(height)))
	i.options = append(i.options, app.MaxSize(unit.Dp(width), unit.Dp(height)))
	i.options = append(i.options, app.MinSize(unit.Dp(width), unit.Dp(height)))
	return i
}

func (i *Initialize) Destroy(f func(err error)) {
	i.destroy = f
}
func (i *Initialize) Frame(f FrameFun) {
	i.frame = f
}

func (i *Initialize) Run() {
	i.win.Option(i.options...)
	var ops op.Ops
	go func() {
		for {
			e := i.win.Event()
			switch e := e.(type) {
			case app.DestroyEvent:
				i.destroy(e.Err)
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				if i.bgColor != (color.NRGBA{}) {
					rect := clip.Rect{
						Max: gtx.Constraints.Max,
					}
					paint.FillShape(gtx.Ops, i.bgColor, rect.Op())
				}
				i.frame(gtx, ops, i.win)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
