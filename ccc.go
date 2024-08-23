package main

//
// func main() {
// 	var clickable widget.Clickable
// 	var overlay widget.Clickable // This will be our overlay to block clicks.
// 	var th = material.NewTheme()
// 	// w := new(app.Window)
// 	var ops op.Ops
// 	go func() {
// 		w := new(app.Window)
// 		for {
// 			e := w.Event()
// 			switch e := e.(type) {
// 			case app.DestroyEvent:
// 				panic(e.Err)
// 			case app.FrameEvent:
// 				gtx := app.NewContext(&ops, e)
// 				// ==============================================
// 				layout.Flex{Axis: layout.Vertical}.Layout(gtx,
// 					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
// 						if clickable.Clicked(gtx) {
// 							log.Println("Button clicked!")
// 						}
// 						return material.Button(th, &clickable, "Click Me").Layout(gtx)
// 					}),
// 				)
//
// 				// Conditionally layout an overlay that blocks interaction with underlying components.
// 				if true { // Set to true to enable the overlay, false to disable it.
// 					layOverlay(gtx, &overlay)
// 				}
//
// 				e.Frame(gtx.Ops)
// 			}
// 		}
// 	}()
// 	app.Main()
// }
//
// func layOverlay(gtx layout.Context, overlay *widget.Clickable) {
// 	// Fill the whole area with a semi-transparent overlay.
// 	dr := image.Rectangle{Max: gtx.Constraints.Max}
// 	paint.FillShape(gtx.Ops, color.NRGBA{R: 0, G: 0, B: 0, A: 128}, clip.Rect(dr).Op())
//
// 	// Create an invisible clickable widget that covers the entire area.
// 	overlay.Layout(gtx)
// 	if overlay.Clicked(gtx) {
// 		log.Println("Overlay clicked, blocking underlying clicks.")
// 	}
// }
