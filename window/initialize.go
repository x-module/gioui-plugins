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
	"gioui.org/io/system"
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

type ElementSize struct {
	Width  int `yaml:"width"`
	Height int `yaml:"height"`
}

type Application struct {
	destroy      DestroyFun
	frame        FrameFun
	win          *app.Window
	options      []app.Option
	bgColor      color.NRGBA
	init         bool
	centerWindow bool
}

func NewApplication(win *app.Window) *Application {
	return &Application{
		win: win,
		destroy: func(err error) {
			os.Exit(1)
		},
	}
}

func (i *Application) GetWin() *app.Window {
	return i.win
}

func (i *Application) BackgroundColor(color color.NRGBA) {
	i.bgColor = color
}
func (i *Application) NoActionBar() *Application {
	i.win.Option(app.Decorated(false))
	return i
}
func (i *Application) HaveActionBar() *Application {
	i.win.Option(app.Decorated(true))
	return i
}

func (i *Application) ReCenterWindow() *Application {
	i.win.Option(i.options...)
	i.win.Perform(system.ActionCenter)
	return i
}
func (i *Application) CenterWindow() *Application {
	i.centerWindow = true
	return i
}

func (i *Application) Title(t string) *Application {
	i.options = append(i.options, app.Title(t))
	return i
}

func (i *Application) Size(size ElementSize) *Application {
	i.options = append(i.options, app.Size(unit.Dp(size.Width), unit.Dp(size.Height)))
	// i.options = append(i.options, app.MaxSize(unit.Dp(size.Width), unit.Dp(size.Height)))
	// i.options = append(i.options, app.MinSize(unit.Dp(size.Width), unit.Dp(size.Height)))
	return i
}

func (i *Application) Destroy(f func(err error)) {
	i.destroy = f
}
func (i *Application) Frame(f FrameFun) {
	i.frame = f
}

func (i *Application) Run() {
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
				layout.Stack{}.Layout(gtx,
					layout.Stacked(func(gtx layout.Context) layout.Dimensions {
						if !i.init && i.centerWindow {
							i.win.Option(i.options...)
							i.win.Perform(system.ActionCenter)
							i.init = true
						}
						i.frame(gtx, ops, i.win)
						return layout.Dimensions{}
					}),
				)
				e.Frame(gtx.Ops)
			}
		}
	}()
	app.Main()
}
