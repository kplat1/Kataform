// Copyright (c) 2018, The Kataform Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	//"go/token"
	//"reflect"

	"github.com/goki/gi"
	//"github.com/goki/gi/complete"
	"github.com/goki/gi/gimain"
	//"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/units"
	"github.com/goki/ki"
//	"github.com/goki/ki/kit"
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

	oswin.TheApp.SetName("game")
	oswin.TheApp.SetAbout("This is a Kataform game.")

	win := gi.NewWindow2D("game", "Game", width, height, true) // true = pixel sizes

	//icnm := "widget-wedge-down"

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	// style sheet
	var css = ki.Props{
		"button": ki.Props{
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
	mfr.SetProp("spacing", units.NewValue(1, units.Ex))
	// mfr.SetProp("background-color", "linear-gradient(to top, red, lighter-80)")
	// mfr.SetProp("background-color", "linear-gradient(to right, red, orange, yellow, green, blue, indigo, violet)")
	// mfr.SetProp("background-color", "linear-gradient(to right, rgba(255,0,0,0), rgba(255,0,0,1))")
	// mfr.SetProp("background-color", "radial-gradient(red, lighter-80)")

	trow := mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()




type mapRow struct{
  id string
  typeOf string
  x float64
  y float64
  z float64
}
	//fmt.Printf("%v", mapRow{"block1", "block", 7, 8})
	

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "Game"
	
	//mapRow{"idName", "typeOf", 7, 8}
	
	 var Map []mapRow
	
	
	
	
	
	// BLOCK TYPES
	
	
	/* TYPES:
	
	woodBlock
	metalBlock
	stoneBlock
	grassLayer
	dirtLayer
	roundRock
	metalDoor
	woodStore
	stoneStore
	metalStore
	waterLayer
	
	
	*/
	
	
	// END BLOCK TYPES
	
	
	
	
	
	// THE MAP
	
	
	
	Map = append(Map, mapRow{"dirtLayer1", "dirtLayer", 0, 0, 0}) // creates a new thing to the map
		Map = append(Map, mapRow{"grassLayer1", "grassLayer", 0, 0, 1})
			Map = append(Map, mapRow{"woodStore1", "woodStore", 0, 0, 2})
	
	// END OF THE MAP
	
	

	
	

	
	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	title.SetProp("vertical-align", gi.AlignCenter)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	title.SetStretchMaxWidth()
	title.SetStretchMaxHeight()

// game goes right here


for i := 0; i < len(Map); i++ {
  
   if Map[i].typeOf == "dirtLayer" {
    updt := vp.UpdateStart()
    
    dirtLayerText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("dirtLayerText%v", i)).(*gi.Label)
    dirtLayerText.Text = fmt.Sprintf("Row %v is a Dirt Layer at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)
    
   	vp.UpdateEnd(updt)
 
  } else if Map[i].typeOf == "grassLayer" {
    updt := vp.UpdateStart()
    
    grassLayerText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("grassLayerText%v", i)).(*gi.Label)
    grassLayerText.Text = fmt.Sprintf("Row %v is a Grass Layer at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)
    
   	vp.UpdateEnd(updt)
 
  } else if Map[i].typeOf == "woodStore" {
    updt := vp.UpdateStart()
    
    woodStoreText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("woodStoreText%v", i)).(*gi.Label)
    woodStoreText.Text = fmt.Sprintf("Row %v is a Wood Store at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)
    
   	vp.UpdateEnd(updt)
 
  } else if Map[i].typeOf == "woodBlock" {
    updt := vp.UpdateStart()
    
    woodBlockText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("woodBlockText%v", i)).(*gi.Label)
    woodBlockText.Text = fmt.Sprintf("Row %v is a Wood Block at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)
    
   	vp.UpdateEnd(updt)
 
  } else if Map[i].typeOf == "pathBlock" {
    updt := vp.UpdateStart()
    
    pathBlockText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("pathBlockText%v", i)).(*gi.Label)
    pathBlockText.Text = fmt.Sprintf("Row %v is a Path Block at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)
    
   	vp.UpdateEnd(updt)
  } else {
   updt := vp.UpdateStart()
    
    errorBlockText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("errorBlockText%v", i)).(*gi.Label)
    errorBlockText.Text = fmt.Sprintf("Error. Type %v does not exist.", Map[i].typeOf)
    
   	vp.UpdateEnd(updt)
  }
}
	
	
	
	
	
	
/*	inputConfirm := trow.AddNewChild(gi.KiT_Button, "inputConfirm").(*gi.Button)
	inputConfirm.Text = "Done"
	// button2.SetProp("background-color", "#EDF")
	inputConfirm.Tooltip = "This button will open the GoGi GUI editor where you can edit this very GUI and see it update dynamically as you change things"
	inputConfirm.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
		  textResult := input1.Text()
			fmt.Printf("%v", mapRow{textResult, "block", 7, 8})
		}
	})
	*/
	
	


	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}
