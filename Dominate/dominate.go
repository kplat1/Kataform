// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	//"go/token"

	"github.com/goki/gi"
	//"github.com/goki/gi/giv"
	//"github.com/goki/gi/complete"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/units"
	"github.com/goki/ki"
	"math/rand"
	//"strconv"
	//"math
)

func main() {

	gimain.Main(func() {
		mainrun()
	})
}

type gridBox struct {
	gridNum int
	owner   string
	color   string
}

var Grid []gridBox

type player struct {
	name    string
	color   string
	curGrid int
}

var Players []player

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

	oswin.TheApp.SetName("Dominate")
	oswin.TheApp.SetAbout("This is a simple domination / winner take all game.")

	win := gi.NewWindow2D("game-dominate", "Dominate", width, height, true) // true = pixel sizes

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

	// vars in here :

	for i := 0; i < 16; i++ {

		if i <= 7 {
			Grid = append(Grid, gridBox{i, "you", "green"})
		} else {
			Grid = append(Grid, gridBox{i, "computer1", "red"})
		}

	}

	Players = append(Players, player{"you", "green", 0})
	Players = append(Players, player{"computer1", "red", 15})

	// end of vars

	trow := mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>Dominate - winner take all</b>"
	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	title.SetProp("vertical-align", gi.AlignCenter)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	title.SetStretchMaxWidth()
	title.SetStretchMaxHeight()

	trow.AddNewChild(gi.KiT_Space, "spc1")

	welcomeText := trow.AddNewChild(gi.KiT_Label, "welcomeText").(*gi.Label)
	welcomeText.Text = "Welcome to Dominate. Fight for control over a 4 by 4 grid."
	welcomeText.SetProp("text-align", gi.AlignCenter)

	upButton := trow.AddNewChild(gi.KiT_Button, "upButton").(*gi.Button)
	upButton.Text = "Move up"

	downButton := trow.AddNewChild(gi.KiT_Button, "downButton").(*gi.Button)
	downButton.Text = "Move down"
	downButton.Shortcut = "s"
	
	
		rightButton := trow.AddNewChild(gi.KiT_Button, "rightButton").(*gi.Button)
	rightButton.Text = "Move right"

	leftButton := trow.AddNewChild(gi.KiT_Button, "leftButton").(*gi.Button)
	leftButton.Text = "Move Left"
	
	

	playingGrid := trow.AddNewChild(gi.KiT_Layout, "playingGrid").(*gi.Layout)
	playingGrid.Lay = gi.LayoutGrid

	playingGrid.SetProp("columns", 4)

	upButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
			if Players[0].curGrid-4 < 0 {

			} else {
				Players[0].curGrid -= 4
				redrawPlayingGrid(playingGrid, Players[0].curGrid+4)
			}
		}
	})

	downButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
			if Players[0].curGrid+4 > 15 {
				fmt.Printf("\n CAN'T MOVE DOWN \n")
			} else {
				Players[0].curGrid += 4

				fmt.Printf("\n %v \n", Players[0].curGrid)

				redrawPlayingGrid(playingGrid, Players[0].curGrid-4)
			}
		}
	})
	
	
		rightButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
			if Players[0].curGrid + 1 > 15 {

			} else {
				Players[0].curGrid += 1
				redrawPlayingGrid(playingGrid, Players[0].curGrid - 1)
			}
		}
	})
	
	leftButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
			if Players[0].curGrid - 1 < 0 {

			} else {
				Players[0].curGrid -= 1
				redrawPlayingGrid(playingGrid, Players[0].curGrid + 1)
			}
		}
	})

	drawPlayingGrid(playingGrid)
	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}

func drawPlayingGrid(playingGrid *gi.Layout) {

	for rows := 0; rows < 4; rows++ {

		for cols := 0; cols < 4; cols++ {

			gridPos := ((cols + 1) + rows*4) - 1

			fmt.Printf("\n  GRID POS : %v  \n", gridPos)

			fmt.Printf("    ROWS :     %v       COLS :    %v    POS :  %v     \n ", rows, cols, gridPos)

			fmt.Printf("    GRID LENGTH :  %v     ", len(Grid))

			cell := playingGrid.AddNewChild(gi.KiT_Frame, fmt.Sprintf("cell_%v_%v", rows, cols)).(*gi.Frame)
			cell.SetProp("background-color", Grid[gridPos].color)

			cell.SetProp("border-color", "black")
			cell.SetProp("border-width", "4px")
			cell.SetProp("width", "5em")
			cell.SetProp("height", "5em")

			for i := 0; i < len(Players); i++ {
				if gridPos == Players[i].curGrid {


if Players[i].name == "you" {
  cell.SetProp("background-color", fmt.Sprintf("light%v", Players[i].color))
} else {
  cell.SetProp("background-color", fmt.Sprintf("dark%v", Players[i].color))
}

					
				}
			}

		}
	}

}

func redrawPlayingGrid(playingGrid *gi.Layout, prevCell int) {
	updt := playingGrid.UpdateStart()
	for rows := 0; rows < 4; rows++ {

		for cols := 0; cols < 4; cols++ {

			gridPos := ((cols + 1) + rows*4) - 1

			fmt.Printf("\n  GRID POS : %v  \n", gridPos)

			fmt.Printf("    ROWS :     %v       COLS :    %v    POS :  %v     \n ", rows, cols, gridPos)

			fmt.Printf("    GRID LENGTH :  %v     ", len(Grid))

		

				if gridPos == Players[0].curGrid {
					newCell := playingGrid.KnownChild(Players[0].curGrid).(*gi.Frame)

					newCell.SetProp("background-color", fmt.Sprintf("light%v", Players[0].color))

					oldCell := playingGrid.KnownChild(prevCell).(*gi.Frame)
					oldCell.SetProp("background-color", fmt.Sprintf("%v", Players[0].color))

				}
			

		}
	}
	
	enemyOldPos := Players[1].curGrid
	
	
	
	enemyRandomNumber := rand.Intn(4)
	fmt.Printf("\n RANDOM: %v \n", enemyRandomNumber)
	
	var enemyDirection string
	
	if enemyRandomNumber == 0 {
	  enemyDirection = "up"
	  if enemyOldPos - 4 < 0 {
	    
	  } else {
	    enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
	enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))
	    Players[1].curGrid = enemyOldPos - 4
	    enemyNewCell := playingGrid.KnownChild(enemyOldPos - 4)
	  enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
	  }
	  
	} else if enemyRandomNumber == 1 {
	  enemyDirection = "down"
	  if enemyOldPos + 4 > 15 {
	    
	  } else {
	    enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
	enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))
	    Players[1].curGrid = enemyOldPos + 4
	    enemyNewCell := playingGrid.KnownChild(enemyOldPos + 4)
	  enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
	  }
	} else if enemyRandomNumber == 2 {
	  enemyDirection = "left"
	  
	  if enemyOldPos - 1 < 0 {
	    
	  } else {
	    enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
	enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))
	    Players[1].curGrid = enemyOldPos - 1
	    enemyNewCell := playingGrid.KnownChild(enemyOldPos - 1)
	  enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
	  }
	} else if enemyRandomNumber == 3 {
	  enemyDirection = "right"
	  if enemyOldPos + 1 > 15 {
	    
	  } else {
	    enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
	enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))
	    Players[1].curGrid = enemyOldPos + 1
	    enemyNewCell := playingGrid.KnownChild(enemyOldPos + 1)
	  enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
	  }
	}
	fmt.Printf("%v", enemyDirection)
	
	
	playingGrid.SetFullReRender()
	playingGrid.UpdateEnd(updt)


}
