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

	type mapRow struct {
		id     string
		typeOf string
		x      float64
		y      float64
		z      float64
	}
	//fmt.Printf("%v", mapRow{"block1", "block", 7, 8})

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "Game"

	//mapRow{"idName", "typeOf", 7, 8}

	var Map []mapRow

	// BLOCK TYPES

	/* TYPES:

	woodBlock DONE
	metalBlock DONE
	stoneBlock DONE
	grassLayer DONE
	dirtLayer  DONE
	roundRock NEED
	metalDoor
	woodStore  DONE
	stoneStore NEED
	metalStore
	waterLayer NEED


	*/

	// END BLOCK TYPES

	// THE MAP
Map = append(Map, mapRow{"dirtLayer1", "dirtLayer", 0, 0, 0})
Map = append(Map, mapRow{"grassLayer1", "grassLayer", 0, 0, 1})
Map = append(Map, mapRow{"woodBlock1", "woodBlock", 1, 1, 2})
Map = append(Map, mapRow{"woodStore1", "woodStore", 2, -2, 2})
Map = append(Map, mapRow{"waterLayer1", "waterLayer",0 ,0 ,5 })
Map = append(Map, mapRow{"woodBlock2", "woodBlock",0 ,0 , 6})
Map = append(Map, mapRow{"stoneStore1", "stoneStore",-2 , 2, 2})
Map = append(Map, mapRow{"stoneStore2", "stoneStore",3 ,2 ,2 })
Map = append(Map, mapRow{"stoneStore3", "stoneStore",6 , 5, 2})
Map = append(Map, mapRow{"roundRock1", "roundRock",-5 ,4 , 2})
Map = append(Map, mapRow{"stoneStore4", "stoneStore",2 , 5, 2})
Map = append(Map, mapRow{"roundRock2", "roundRock",4 , 5, 2})
Map = append(Map, mapRow{"cobblestone1", "cobblestone", -2, 3, 2})

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

	//fmt.Printf("%v", img)

var imgSize float32 = 400

	
		
		for i := 0; i < len(Map); i++ {
		  if Map[i].typeOf == "dirtLayer" {
			updt := vp.UpdateStart()

			dirtLayerText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("dirtLayerText%v", i)).(*gi.Label)
			dirtLayerText.Text = fmt.Sprintf("Row %v is a Dirt Layer at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgDirt := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgDirt%v", i)).(*gi.Bitmap)
			imgDirt.OpenImage("dirt.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		} else if Map[i].typeOf == "cobblestone" { // ELSE IF
		
		updt := vp.UpdateStart()

			cobblestoneText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("cobblestoneText%v", i)).(*gi.Label)
			cobblestoneText.Text = fmt.Sprintf("Row %v is a Cobblestone at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgCobblestone := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgCobblestone%v", i)).(*gi.Bitmap)
			imgCobblestone.OpenImage("cobblestone.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)
		
		
		
		
		} else if Map[i].typeOf == "grassLayer" {
			updt := vp.UpdateStart()

			grassLayerText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("grassLayerText%v", i)).(*gi.Label)
			grassLayerText.Text = fmt.Sprintf("Row %v is a Grass Layer at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgGrass := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgGrass%v", i)).(*gi.Bitmap)
			imgGrass.OpenImage("grass.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		} else if Map[i].typeOf == "waterLayer" {
			updt := vp.UpdateStart()

			waterLayerText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("waterLayerText%v", i)).(*gi.Label)
			waterLayerText.Text = fmt.Sprintf("Row %v is a Water Layer at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgWater := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgWater%v", i)).(*gi.Bitmap)
			imgWater.OpenImage("water.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		} else if Map[i].typeOf == "stoneStore" {
			updt := vp.UpdateStart()

			stoneStoreText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("stoneStoreText%v", i)).(*gi.Label)
			stoneStoreText.Text = fmt.Sprintf("Row %v is a Stone Store at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgStoneStore := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgStoneStore%v", i)).(*gi.Bitmap)
			imgStoneStore.OpenImage("stoneStore.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		}	else if Map[i].typeOf == "roundRock" {
			updt := vp.UpdateStart()

			roundRockText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("roundRockText%v", i)).(*gi.Label)
			roundRockText.Text = fmt.Sprintf("Row %v is a Round Rock at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgRock := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgRock%v", i)).(*gi.Bitmap)
			imgRock.OpenImage("rock.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		} else if Map[i].typeOf == "woodBlock" {
			updt := vp.UpdateStart()

			woodBlockText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("woodBlockText%v", i)).(*gi.Label)
			woodBlockText.Text = fmt.Sprintf("Row %v is a Wood Block at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgWood := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgWood%v", i)).(*gi.Bitmap)
			imgWood.OpenImage("wood.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		}  else if Map[i].typeOf == "metalBlock" {
			updt := vp.UpdateStart()

			metalBlockText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("metalBlockText%v", i)).(*gi.Label)
			metalBlockText.Text = fmt.Sprintf("Row %v is a Metal Block at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgMetal := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgMetal%v", i)).(*gi.Bitmap)
			imgMetal.OpenImage("metal.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		} else if Map[i].typeOf == "stoneBlock" {
			updt := vp.UpdateStart()

			stoneBlockText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("stoneBlockText%v", i)).(*gi.Label)
			stoneBlockText.Text = fmt.Sprintf("Row %v is a Stone Block at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgStone := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgStone%v", i)).(*gi.Bitmap)
			imgStone.OpenImage("stone.jpg", imgSize, imgSize)

			vp.UpdateEnd(updt)

		}	else if Map[i].typeOf == "woodStore" {
			updt := vp.UpdateStart()

			woodStoreText := trow.AddNewChild(gi.KiT_Label, fmt.Sprintf("woodStoreText%v", i)).(*gi.Label)
			woodStoreText.Text = fmt.Sprintf("Row %v is a Wood Store at x position %v and y position %v. It is at height %v.", i, Map[i].x, Map[i].y, Map[i].z)

			imgWoodStore := trow.AddNewChild(gi.KiT_Bitmap, fmt.Sprintf("imgWoodStore%v", i)).(*gi.Bitmap)
			imgWoodStore.OpenImage("woodStore.jpg", imgSize, imgSize)
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
