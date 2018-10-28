package main

import (

	//"github.com/goki/gi/giv"
	//"github.com/goki/gi/complete"

	"fmt"

	"github.com/chewxy/math32"
	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/svg"
	"github.com/goki/ki"
)

func main() {

	gimain.Main(func() {
		mainrun()
	})
}
func mainrun() {
	width := 1024
	height := 768

	// turn these on to see a traces of various stages of processing..
	// gi.Update2DTrace = true
	// gi.Render2DTrace = true
	// gi.Layout2DTrace = true
	// ki.SignalTrace = true

	rec := ki.Node{}          // receiver for events
	rec.InitName(&rec, "rec") // this is essential for root objects not owned by other Ki tree nodes

	oswin.TheApp.SetName("Graphing")
	oswin.TheApp.SetAbout("This is a graphing app.")

	win := gi.NewWindow2D("graph", "Graphing App", width, height, true) // true = pixel sizes

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	// style sheet
	var css = ki.Props{
		"Action": ki.Props{
			"background-color": gi.Prefs.Colors.Control, // gi.Color{255, 240, 240, 255},
		},
		"#combo": ki.Props{
			"background-color": gi.Color{240, 255, 240, 255},
		},
		".hslides": ki.Props{
			"background-color": gi.Color{240, 225, 255, 255},
		},
		"kbd": ki.Props{
			"color": "blue",
		},
	}
	vp.CSS = css

	mfr := win.SetMainFrame()

	title := mfr.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>Graphing</b>"
	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	title.SetProp("vertical-align", gi.AlignCenter)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	title.SetStretchMaxWidth()
	title.SetStretchMaxHeight()

	aboutText := mfr.AddNewChild(gi.KiT_Label, "aboutText").(*gi.Label)
	aboutText.Text = "Graphing is an app that will allow you to enter equations and have them be graphed. There will also be other modes where you can have marbles fall or things like that."
	aboutText.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	aboutText.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..

	aboutText.SetStretchMaxWidth()
	aboutText.SetStretchMaxHeight()

	graphingInput := mfr.AddNewChild(gi.KiT_TextField, "graphingInput").(*gi.TextField)
	graphingInput.Placeholder = "Enter your equation"

	submitGraphingInput := mfr.AddNewChild(gi.KiT_Button, "submitGraphingInput").(*gi.Button)
	submitGraphingInput.Text = "Graph equation"
	submitGraphingInput.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			updt := vp.UpdateStart()

			graphResult := mfr.AddNewChild(gi.KiT_Label, "graphResult").(*gi.Label)
			graphResult.Text = fmt.Sprintf("Your equation is: %v", graphingInput.Text())

			vp.UpdateEnd(updt)
		}
	})

	graph := mfr.AddNewChild(svg.KiT_SVG, "graph").(*svg.SVG)
	sz := float32(800)
	graph.SetProp("width", sz)
	graph.SetProp("height", sz)

	gmin := gi.Vec2D{-10, -10}
	gmax := gi.Vec2D{10, 10}
	gsz := gmax.Sub(gmin)
	ginc := gsz.DivVal(100)

	graph.ViewBox.Min = gmin
	graph.ViewBox.Size = gsz
	graph.Norm = true
	graph.Fill = true
	graph.SetProp("background-color", "white")
	graph.SetProp("stroke-width", "1pct")

	path1 := graph.AddNewChild(svg.KiT_Path, "path1").(*svg.Path)
	path1.SetProp("fill", "none")
	ps := ""
	start := true
	for x := gmin.X; x < gmax.X; x += ginc.X {
		y := 5 * math32.Cos(x)
		if start {
			ps += fmt.Sprintf("M %v %v ", x-gmin.X, y-gmin.Y)
			start = false
		} else {
			ps += fmt.Sprintf("L %v %v ", x-gmin.X, y-gmin.Y)
		}
	}
	path1.SetData(ps)

	//////////////////////////////////////////
	//      Main Menu

	appnm := oswin.TheApp.Name()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

	amen := win.MainMenu.KnownChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.KnownChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

}
