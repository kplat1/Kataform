package main

import (
	"fmt"
	//"go/token"

	"github.com/goki/gi/gi"
	//"github.com/goki/gi/giv"
	//"github.com/goki/gi/complete"
	// 	"math/rand"

	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"

	"github.com/goki/gi/units"
	"github.com/goki/ki"
	// 	"github.com/goki/gi/svg"
	//"strconv"
	//"math
	// 	"time"
)

func main() {

	gimain.Main(func() {
		mainrun()
	})
}

var trow *gi.Layout

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

	oswin.TheApp.SetName("Base")
	oswin.TheApp.SetAbout("Base")

	win := gi.NewWindow2D("base", "Base", width, height, true) // true = pixel sizes

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
	// 	dfr := mfr.AddNewChild(KiT_DomFrame, "domframe").(*DomFrame)
	mfr.SetProp("spacing", units.NewValue(1, units.Ex))
	// mfr.SetProp("background-color", "linear-gradient(to top, red, lighter-80)")
	// dfr.SetProp("background-color", "linear-gradient(to right, red, orange, yellow, green, blue, indigo, violet)")
	// dfr.SetProp("background-color", "linear-gradient(to right, rgba(255,0,0,0), rgba(255,0,0,1))")
	// dfr.SetProp("background-color", "radial-gradient(red, lighter-80)")

	// vars in here :

	// end of vars

	trow = mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>BASE</b>"
	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	title.SetProp("vertical-align", gi.AlignCenter)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	title.SetStretchMaxWidth()
	title.SetStretchMaxHeight()

	// trow.AddNewChild(gi.KiT_Space, "spc1")

	tv := mfr.AddNewChild(gi.KiT_TabView, "tv").(*gi.TabView)
	tv.NewTabButton = true

	homeBaseTabT, _ := tv.AddNewTab(gi.KiT_Layout, "Home Base")

	homeBaseTab := homeBaseTabT.(*gi.Layout)
	homeBaseTab.Lay = gi.LayoutVert
	homeBaseTab.SetStretchMaxWidth()
	homeBaseTab.SetStretchMaxHeight()

	homeBaseTitle := homeBaseTab.AddNewChild(gi.KiT_Label, "homeBaseTitle").(*gi.Label)
	homeBaseTitle.Text = "<b>Welcome to your Base!</b>"
	homeBaseTitle.SetProp("text-align", gi.AlignCenter)

	homeBaseTitle.SetProp("font-family", "Times New Roman, serif")
	homeBaseTitle.SetProp("font-size", "x-large")
	homeBaseTitle.SetProp("align-vert", gi.AlignTop)

	welcomeText := homeBaseTab.AddNewChild(gi.KiT_Label, "welcomeText").(*gi.Label)

	welcomeText.Text = "<b>Let's get going. View the other planning tabs to start getting stuff done!</b>"
	welcomeText.SetProp("font-size", "large")
	welcomeText.SetProp("text-align", gi.AlignCenter)

	planningProjectTabT, _ := tv.AddNewTab(gi.KiT_Layout, "Planning Project")
	planningProjectTab := planningProjectTabT.(*gi.Layout)
	planningProjectTab.Lay = gi.LayoutVert
	planningProjectTab.SetStretchMaxWidth()
	planningProjectTab.SetStretchMaxHeight()

	planningProjectTitle := planningProjectTab.AddNewChild(gi.KiT_Label, "planningProjectTitle").(*gi.Label)
	planningProjectTitle.Text = "<b>Do the projects in your sheet below!</b>"
	planningProjectTitle.SetProp("text-align", gi.AlignCenter)

	planningProjectTitle.SetProp("font-family", "Times New Roman, serif")
	planningProjectTitle.SetProp("font-size", "x-large")
	planningProjectTitle.SetProp("align-vert", gi.AlignTop)

	errorText := planningProjectTab.AddNewChild(gi.KiT_Label, "errorText").(*gi.Label)

	errorText.Text = "Sheet functions are currently disabled. Please view the spreadsheeet instead."
	errorText.SetProp("font-size", "large")
	errorText.SetProp("text-align", gi.AlignCenter)

	//////////////////////////////////////////
	//      Main Menu

	appnm := oswin.TheApp.Name()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

	amen := win.MainMenu.ChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.ChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}
