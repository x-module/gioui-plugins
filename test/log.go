package main

import (
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"github.com/x-module/gioui-plugins/resource"
	"github.com/x-module/gioui-plugins/theme"
	"github.com/x-module/gioui-plugins/widgets"
	"github.com/x-module/gioui-plugins/window"
)

func main() {
	var clickable widget.Clickable
	th := theme.NewTheme()

	// fontPath := "/System/Library/Fonts/SFNS.ttf" // SF Pro
	// fontPath := "/System/Library/Fonts/Monaco.ttf" // SF Pro
	// fontPath := "/System/Library/Fonts/SFNSMono.ttf" // SF Pro
	// fontData, err := os.ReadFile(fontPath)
	// if err != nil {
	// 	panic(err)
	// }
	// face, err := opentype.ParseCollection(fontData)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(face))
	// th.Shaper = text.NewShaper(text.NoSystemFonts(), text.WithCollection(gofont.Regular()))
	log := widgets.NewLogViewer(th, true)
	card := widgets.NewCard(th)
	win := window.NewApplication(new(app.Window))
	win.Title("Hello, Gio!").Size(window.ElementStyle{
		Height: 600,
		Width:  800,
	})
	win.BackgroundColor(th.Color.DefaultWindowBgGrayColor)
	win.Frame(func(gtx layout.Context, ops op.Ops, win *app.Window) {
		if clickable.Clicked(gtx) {
			log.Debug("this is debug log")
			log.Info("this is info log")
			log.Warn("this is warn log")
			log.Error("this is error log")
			log.Fatal("this is fatal log")
		}
		// ==============================================
		layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return widgets.DefaultButton(th, &clickable, "default", unit.Dp(100)).SetIcon(resource.DeleteIcon, 0).Layout(gtx)
					})
				}),
				layout.Rigid(layout.Spacer{Height: unit.Dp(10)}.Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {

					return card.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						gtx.Constraints.Min = gtx.Constraints.Max
						return log.Layout(gtx)
					})
				}),
			)
		})
	})
	win.Run()
}
