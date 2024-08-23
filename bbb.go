package main

import (
	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"image/color"
	"log"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Title("Dashed Line Example"),
			app.Size(unit.Dp(800), unit.Dp(600)),
		)
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func draw(w *app.Window) error {
	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&e.Ops, e)
			drawDashedLine(gtx, color.NRGBA{R: 0, G: 0, B: 0, A: 255}, 10, 5)
			e.Frame(gtx.Ops)
		}
	}
}

func drawDashedLine(gtx layout.Context, col color.NRGBA, dashLength, gapLength float32) {
	ops := gtx.Ops
	var path clip.Path
	path.Begin(ops)

	// Starting point of the line
	start := f32.Pt(100, 100)
	// Ending point of the line
	end := f32.Pt(700, 100)

	// Calculate the total length of the dashed line
	dx := end.X - start.X
	dy := end.Y - start.Y
	length := float32((dx*dx + dy*dy).Sqrt())

	// Draw dashes
	for i := float32(0); i < length; i += dashLength + gapLength {
		dashEnd := i + dashLength
		if dashEnd > length {
			dashEnd = length
		}

		// Calculate start and end points of the current dash
		t0 := i / length
		t1 := dashEnd / length
		p0 := lerpPoint(start, end, t0)
		p1 := lerpPoint(start, end, t1)

		path.MoveTo(p0)
		path.LineTo(p1)
	}

	// Create a stroke style and fill the path with the specified color
	stroke := clip.Stroke{
		Path:  path.End(),
		Width: 2,
	}.Op()
	paint.FillShape(ops, col, stroke)
}

// lerpPoint linearly interpolates between two points based on a t parameter
func lerpPoint(a, b f32.Point, t float32) f32.Point {
	return f32.Pt(
		a.X+(b.X-a.X)*t,
		a.Y+(b.Y-a.Y)*t,
	)
}
