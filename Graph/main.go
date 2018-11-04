package main

import (
	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/svg"
	"github.com/goki/ki"
)

var Vp *gi.Viewport2D
var SvgGraph *svg.SVG
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

	Gr.Defaults()

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

	// the StructView will also show the Graph Toolbar which is main actions..
	gstru := mfr.AddNewChild(giv.KiT_StructView, "gstru").(*giv.StructView)
	gstru.Viewport = Vp // needs vp early for toolbar
	gstru.SetProp("height", "4em")
	gstru.SetStruct(&Gr, nil)

	lns := mfr.AddNewChild(giv.KiT_TableView, "lns").(*giv.TableView)
	lns.Viewport = Vp
	lns.SetSlice(&Gr.Lines, nil)

	frame := mfr.AddNewChild(gi.KiT_Frame, "frame").(*gi.Frame)

	SvgGraph = frame.AddNewChild(svg.KiT_SVG, "graph").(*svg.SVG)
	SvgGraph.SetProp("min-width", GraphSize)
	SvgGraph.SetProp("min-height", GraphSize)
	SvgGraph.SetStretchMaxWidth()
	SvgGraph.SetStretchMaxHeight()

	SvgLines = SvgGraph.AddNewChild(svg.KiT_Group, "SvgLines").(*svg.Group)
	SvgMarbles = SvgGraph.AddNewChild(svg.KiT_Group, "SvgMarbles").(*svg.Group)
	SvgCoords = SvgGraph.AddNewChild(svg.KiT_Group, "SvgCoords").(*svg.Group)

	gmin = gi.Vec2D{-10, -10}
	gmax = gi.Vec2D{10, 10}
	gsz = gmax.Sub(gmin)
	ginc = gsz.DivVal(GraphSize)

	SvgGraph.ViewBox.Min = gmin
	SvgGraph.ViewBox.Size = gsz
	SvgGraph.Norm = true
	SvgGraph.InvertY = true
	SvgGraph.Fill = true
	SvgGraph.SetProp("background-color", "white")
	SvgGraph.SetProp("stroke-width", ".2pct")

	InitCoords()
	ResetMarbles()
	Gr.Lines.Graph()

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
