package main

import (
	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/svg"
	"github.com/goki/gi/units"
	"github.com/goki/ki"
)

var Vp *gi.Viewport2D
var Graph *svg.SVG
var SvgLines *svg.Group
var SvgMarbles *svg.Group
var SvgCoords *svg.Group

var gmin, gmax, gsz, ginc gi.Vec2D
var GraphSize float32 = 800

func main() {
	gimain.Main(func() {
		mainrun()
	})
}

func mainrun() {
	width := 1024
	height := 1024

	Pars.Defaults()
	Lns.Defaults()

	rec := ki.Node{}          // receiver for events
	rec.InitName(&rec, "rec") // this is essential for root objects not owned by other Ki tree nodes

	oswin.TheApp.SetName("Graphing")
	oswin.TheApp.SetAbout("Graphing is an app that will allow you to enter equations and have them be graphed. There will also be other modes where you can have marbles fall or things like that.")

	win := gi.NewWindow2D("graph", "Graphing App", width, height, true) // true = pixel sizes

	Vp = win.WinViewport2D()
	updt := Vp.UpdateStart()

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
	Vp.CSS = css

	mfr := win.SetMainFrame()

	lns := mfr.AddNewChild(giv.KiT_TableView, "lns").(*giv.TableView)
	lns.Viewport = Vp
	lns.SetSlice(&Lns, nil)

	pstru := mfr.AddNewChild(giv.KiT_StructViewInline, "pstru").(*giv.StructViewInline)
	pstru.SetStruct(&Pars, nil)

	brow := mfr.AddNewChild(gi.KiT_Layout, "brow").(*gi.Layout)
	brow.Lay = gi.LayoutHoriz
	brow.SetProp("spacing", units.NewValue(2, units.Ex))

	brow.SetProp("horizontal-align", gi.AlignLeft)
	// brow.SetProp("horizontal-align", gi.AlignJustify)
	brow.SetStretchMaxWidth()

	doGraph := brow.AddNewChild(gi.KiT_Button, "doGraph").(*gi.Button)
	doGraph.Text = "Graph"
	doGraph.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			updt := Vp.UpdateStart()
			ResetMarbles()
			Lns.Graph()
			Vp.UpdateEnd(updt)
		}
	})

	runMarbles := brow.AddNewChild(gi.KiT_Button, "runMarbles").(*gi.Button)
	runMarbles.Text = "Run!"
	runMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			go RunMarbles()
		}
	})

	stepMarbles := brow.AddNewChild(gi.KiT_Button, "stepMarbles").(*gi.Button)
	stepMarbles.Text = "Step"
	stepMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			go UpdateMarbles()
		}
	})

	stopMarbles := brow.AddNewChild(gi.KiT_Button, "stopMarbles").(*gi.Button)
	stopMarbles.Text = "Stop"
	stopMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			Stop = true
		}
	})

	resetMarbles := brow.AddNewChild(gi.KiT_Button, "resetMarbles").(*gi.Button)
	resetMarbles.Text = "Reset Marbles"
	resetMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			ResetMarbles()
		}
	})

	frame := mfr.AddNewChild(gi.KiT_Frame, "frame").(*gi.Frame)

	Graph = frame.AddNewChild(svg.KiT_SVG, "graph").(*svg.SVG)
	Graph.SetProp("min-width", GraphSize)
	Graph.SetProp("min-height", GraphSize)
	Graph.SetStretchMaxWidth()
	Graph.SetStretchMaxHeight()

	SvgLines = Graph.AddNewChild(svg.KiT_Group, "SvgLines").(*svg.Group)
	SvgMarbles = Graph.AddNewChild(svg.KiT_Group, "SvgMarbles").(*svg.Group)
	SvgCoords = Graph.AddNewChild(svg.KiT_Group, "SvgCoords").(*svg.Group)

	gmin = gi.Vec2D{-10, -10}
	gmax = gi.Vec2D{10, 10}
	gsz = gmax.Sub(gmin)
	ginc = gsz.DivVal(GraphSize)

	Graph.ViewBox.Min = gmin
	Graph.ViewBox.Size = gsz
	Graph.Norm = true
	Graph.Fill = true
	Graph.SetProp("background-color", "white")
	Graph.SetProp("stroke-width", ".2pct")

	InitCoords()
	ResetMarbles()
	Lns.Graph()

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

	Vp.UpdateEndNoSig(updt)

	win.StartEventLoop()
}
