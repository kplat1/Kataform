// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	//"go/token"

	"github.com/goki/gi/gi"
	//"github.com/goki/gi/giv"
	//"github.com/goki/gi/complete"
	"math/rand"

	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/key"
	"github.com/goki/gi/units"
	"github.com/goki/ki"
	"github.com/goki/ki/kit"
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

type DomFrame struct {
	gi.Frame
	ButtonRow *gi.Layout
}

var KiT_DomFrame = kit.Types.AddType(&DomFrame{}, nil)

func (df *DomFrame) ConnectEvents2D() {
	df.ConnectEvent(oswin.KeyChordEvent, gi.HiPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		// fvv := recv.Embed(KiT_DomFrame).(*DomFrame)
		kt := d.(*key.ChordEvent)
		ch := kt.Chord()
		switch ch {
		case "s":
			kt.SetProcessed()
			df.DownAction()
		case "w":
			kt.SetProcessed()
			df.UpAction()
		case "a":
			kt.SetProcessed()
			df.LeftAction()
		case "d":
			kt.SetProcessed()
			df.RightAction()
		case "1":
			kt.SetProcessed()
			df.ModeAttackAction()
		case "2":
			kt.SetProcessed()
			df.ModeDefenseAction()
		case "3":
			kt.SetProcessed()
			df.ModeSurrenderAction()

		}
	})
}

func (df *DomFrame) HasFocus2D() bool {
	return true // always.. we're typically a dialog anyway
}

func (df *DomFrame) DownAction() {
	fmt.Printf("Down action!!\n")
	down, _ := df.ButtonRow.ChildByName("downAction", 0)
	down.(*gi.Action).Trigger()
}
func (df *DomFrame) UpAction() {
	fmt.Printf("Up action!!\n")
	up, _ := df.ButtonRow.ChildByName("upAction", 0)
	up.(*gi.Action).Trigger()
}
func (df *DomFrame) LeftAction() {
	fmt.Printf("Left action!!\n")
	left, _ := df.ButtonRow.ChildByName("leftAction", 0)
	left.(*gi.Action).Trigger()
}
func (df *DomFrame) RightAction() {
	fmt.Printf("Right action!!\n")
	right, _ := df.ButtonRow.ChildByName("rightAction", 0)
	right.(*gi.Action).Trigger()
}

func (df *DomFrame) ModeAttackAction() {
	fmt.Printf("Mode attack action!!\n")
	modeAttack, _ := df.ButtonRow.ChildByName("modeAttackAction", 0)
	modeAttack.(*gi.Action).Trigger()
}

func (df *DomFrame) ModeDefenseAction() {
	fmt.Printf("Mode defense action!!\n")
	modeDefense, _ := df.ButtonRow.ChildByName("modeDefenseAction", 0)
	modeDefense.(*gi.Action).Trigger()
}
func (df *DomFrame) ModeSurrenderAction() {
	fmt.Printf("Mode surrender action!!\n")
	modeSurrender, _ := df.ButtonRow.ChildByName("modeSurrenderAction", 0)
	modeSurrender.(*gi.Action).Trigger()
}

var currentMode string = "none"

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
	dfr := mfr.AddNewChild(KiT_DomFrame, "domframe").(*DomFrame)
	dfr.SetProp("spacing", units.NewValue(1, units.Ex))
	// mfr.SetProp("background-color", "linear-gradient(to top, red, lighter-80)")
	// dfr.SetProp("background-color", "linear-gradient(to right, red, orange, yellow, green, blue, indigo, violet)")
	// dfr.SetProp("background-color", "linear-gradient(to right, rgba(255,0,0,0), rgba(255,0,0,1))")
	// dfr.SetProp("background-color", "radial-gradient(red, lighter-80)")

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

	trow := dfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
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

	trow.AddNewChild(gi.KiT_Space, "spc1")

	dfr.ButtonRow = trow.AddNewChild(gi.KiT_Layout, "brow").(*gi.Layout)
	dfr.ButtonRow.Lay = gi.LayoutHoriz
	dfr.ButtonRow.SetProp("spacing", units.NewValue(2, units.Ex))

	dfr.ButtonRow.SetProp("horizontal-align", gi.AlignLeft)
	// dfr.ButtonRow.SetProp("horizontal-align", gi.AlignJustify)
	dfr.ButtonRow.SetStretchMaxWidth()

	dfr.SetProp("background-color", "white")

	upAction := dfr.ButtonRow.AddNewChild(gi.KiT_Action, "upAction").(*gi.Action)
	upAction.Text = "Move up"
	upAction.Shortcut = "W"

	downAction := dfr.ButtonRow.AddNewChild(gi.KiT_Action, "downAction").(*gi.Action)
	downAction.Text = "Move down"
	downAction.Shortcut = "S"

	rightAction := dfr.ButtonRow.AddNewChild(gi.KiT_Action, "rightAction").(*gi.Action)
	rightAction.Text = "Move right"
	rightAction.Shortcut = "D"

	leftAction := dfr.ButtonRow.AddNewChild(gi.KiT_Action, "leftAction").(*gi.Action)
	leftAction.Text = "Move Left"
	leftAction.Shortcut = "A"

	modeAttackAction := dfr.ButtonRow.AddNewChild(gi.KiT_Action, "modeAttackAction").(*gi.Action)
	modeAttackAction.Text = "Set mode to attack"
	modeAttackAction.Shortcut = "1"

	modeDefenseAction := dfr.ButtonRow.AddNewChild(gi.KiT_Action, "modeDefenseAction").(*gi.Action)
	modeDefenseAction.Text = "Set mode to defend"
	modeDefenseAction.Shortcut = "2"

	modeSurrenderAction := dfr.ButtonRow.AddNewChild(gi.KiT_Action, "modeSurrenderAction").(*gi.Action)
	modeSurrenderAction.Text = "Set mode to surrender"
	modeSurrenderAction.Shortcut = "3"

	playingGrid := trow.AddNewChild(gi.KiT_Layout, "playingGrid").(*gi.Layout)
	playingGrid.Lay = gi.LayoutGrid

	playingGrid.SetProp("columns", 4)

	upAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received Action signal: %v from Action: %v\n", gi.ActionSignals(sig), send.Name())
		if Players[0].curGrid-4 < 0 {

		} else {
			Players[0].curGrid -= 4
			redrawPlayingGrid(playingGrid, Players[0].curGrid+4, "up")
		}
	})

	downAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received Action signal: %v from Action: %v\n", gi.ActionSignals(sig), send.Name())
		fmt.Printf("DOWN")
		if Players[0].curGrid+4 > 15 {
			fmt.Printf("\n CAN'T MOVE DOWN \n")
		} else {
			Players[0].curGrid += 4

			fmt.Printf("\n %v \n", Players[0].curGrid)

			redrawPlayingGrid(playingGrid, Players[0].curGrid-4, "down")
		}

	})

	rightAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		//	fmt.Printf("Received Action signal: %v from Action: %v\n", gi.ActionSignals(sig), send.Name())

		if Players[0].curGrid+1 > 15 {

		} else {
			Players[0].curGrid += 1
			redrawPlayingGrid(playingGrid, Players[0].curGrid-1, "right")

		}
	})

	leftAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received Action signal: %v from Action: %v\n", gi.ActionSignals(sig), send.Name())

		if Players[0].curGrid-1 < 0 {

		} else {
			Players[0].curGrid -= 1
			redrawPlayingGrid(playingGrid, Players[0].curGrid+1, "left")

		}
	})

	modeAttackAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received Action signal: %v from Action: %v\n", gi.ActionSignals(sig), send.Name())

		fmt.Printf("Change mode to attack")

		currentMode = "attack"
		fmt.Printf("Current mode is %v \n", currentMode)

	})

	modeDefenseAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received Action signal: %v from Action: %v\n", gi.ActionSignals(sig), send.Name())

		fmt.Printf("Change mode to defense")

		currentMode = "defense"
		fmt.Printf("Current mode is %v \n", currentMode)

	})
	modeSurrenderAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received Action signal: %v from Action: %v\n", gi.ActionSignals(sig), send.Name())

		fmt.Printf("Change mode to surrender!")

		currentMode = "surrender"
		fmt.Printf("Current mode is %v \n", currentMode)

	})
	win.AddShortcut("Alt+S", downAction)
	win.AddShortcut("Alt+W", upAction)
	win.AddShortcut("Alt+D", rightAction)
	win.AddShortcut("Alt+A", leftAction)

	drawPlayingGrid(playingGrid)

	//win.AddShortcut("s", redrawPlayingGrid(playingGrid, Players[0].curGrid, "down"))

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

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}

func drawPlayingGrid(playingGrid *gi.Layout) {

	for rows := 0; rows < 4; rows++ {

		for cols := 0; cols < 4; cols++ {

			gridPos := ((cols + 1) + rows*4) - 1
			//
			//		fmt.Printf("\n  GRID POS : %v  \n", gridPos)

			//	fmt.Printf("    ROWS :     %v       COLS :    %v    POS :  %v     \n ", rows, cols, gridPos)

			//fmt.Printf("    GRID LENGTH :  %v     ", len(Grid))

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

func redrawPlayingGrid(playingGrid *gi.Layout, prevCell int, dir string) {

	/*if dir == "up" && !(Players[0].curGrid - 4 < 0) {
	  Players[0].curGrid -= 4
	} else if dir == "down" && !(Players[0].curGrid + 4 > 15) {
	  Players[0].curGrid += 4
	} else if dir == "left" && !(Players[0].curGrid - 1 < 0) {
	  Players[0].curGrid -= 1
	} else if dir == "right" && !(Players[0].curGrid + 1 > 15) {
	  Players[0].curGrid += 1
	} else if dir == "none" {
	  } else {
	  return
	}
	*/

	updt := playingGrid.UpdateStart()

	enemyOldPos := Players[1].curGrid

	enemyRandomNumber := rand.Intn(4)
	fmt.Printf("\n RANDOM: %v \n", enemyRandomNumber)

	var enemyDirection string

	if enemyRandomNumber == 0 {
		enemyDirection = "up"
		if enemyOldPos-4 < 0 {
			redrawPlayingGrid(playingGrid, prevCell, "none")
		} else {
			enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
			enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))

			Players[1].curGrid = enemyOldPos - 4
			enemyNewCell := playingGrid.KnownChild(enemyOldPos - 4)
			enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
			Grid[enemyOldPos-4].color = fmt.Sprintf("dark%v", Players[1].color)
			Grid[enemyOldPos].color = Players[1].color
		}

	} else if enemyRandomNumber == 1 {
		enemyDirection = "down"
		if enemyOldPos+4 > 15 {
			redrawPlayingGrid(playingGrid, prevCell, "none")
		} else {
			enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
			enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))
			Players[1].curGrid = enemyOldPos + 4
			enemyNewCell := playingGrid.KnownChild(enemyOldPos + 4)
			enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
			Grid[enemyOldPos+4].color = fmt.Sprintf("dark%v", Players[1].color)
			Grid[enemyOldPos].color = Players[1].color
		}
	} else if enemyRandomNumber == 2 {
		enemyDirection = "left"

		if enemyOldPos-1 < 0 {
			redrawPlayingGrid(playingGrid, prevCell, "none")
		} else {
			enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
			enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))
			Players[1].curGrid = enemyOldPos - 1
			enemyNewCell := playingGrid.KnownChild(enemyOldPos - 1)
			enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
			Grid[enemyOldPos-1].color = fmt.Sprintf("dark%v", Players[1].color)
			Grid[enemyOldPos].color = Players[1].color
		}
	} else if enemyRandomNumber == 3 {
		enemyDirection = "right"
		if enemyOldPos+1 > 15 {
			redrawPlayingGrid(playingGrid, prevCell, "none")
		} else {
			enemyOldCell := playingGrid.KnownChild(enemyOldPos).(*gi.Frame)
			enemyOldCell.SetProp("background-color", fmt.Sprintf("%v", Players[1].color))
			Players[1].curGrid = enemyOldPos + 1
			enemyNewCell := playingGrid.KnownChild(enemyOldPos + 1)
			enemyNewCell.SetProp("background-color", fmt.Sprintf("dark%v", Players[1].color))
			Grid[enemyOldPos+1].color = fmt.Sprintf("dark%v", Players[1].color)
			Grid[enemyOldPos].color = Players[1].color
		}
	}
	fmt.Printf("%v", enemyDirection)

	for rows := 0; rows < 4; rows++ {

		for cols := 0; cols < 4; cols++ {

			//fmt.Printf("\n  GRID POS : %v  \n", gridPos)

			//fmt.Printf("    ROWS :     %v       COLS :    %v    POS :  %v     \n ", rows, cols, gridPos)

			//fmt.Printf("    GRID LENGTH :  %v     ", len(Grid))

			fmt.Printf("\n Redrawing \n")

			gridPos := ((cols + 1) + rows*4) - 1
			if gridPos == Players[0].curGrid {

				fmt.Printf("\n CURRENT MODE: %v \n", currentMode)

				if currentMode == "attack" {
					fmt.Printf("\n ATTACK! \n")
					Players[0].curGrid = gridPos
					newCell := playingGrid.KnownChild(Players[0].curGrid).(*gi.Frame)

					newCell.SetProp("background-color", fmt.Sprintf("light%v", Players[0].color))

					oldCell := playingGrid.KnownChild(prevCell).(*gi.Frame)
					oldCell.SetProp("background-color", fmt.Sprintf("%v", Players[0].color))
					Grid[gridPos].color = fmt.Sprintf("dark%v", Players[0].color)
					Grid[prevCell].color = Players[0].color
				} else if currentMode == "defense" {
					fmt.Printf("\n DEFENSE! \n")
					if Grid[gridPos].color == "green" {
						Players[0].curGrid = gridPos
						newCell := playingGrid.KnownChild(Players[0].curGrid).(*gi.Frame)

						newCell.SetProp("background-color", fmt.Sprintf("light%v", Players[0].color))

						oldCell := playingGrid.KnownChild(prevCell).(*gi.Frame)
						oldCell.SetProp("background-color", fmt.Sprintf("%v", Players[0].color))
						Grid[gridPos].color = fmt.Sprintf("dark%v", Players[0].color)
						Grid[prevCell].color = Players[0].color
					} else {
						gridPos = prevCell
						Players[0].curGrid = prevCell
					}

				} else if currentMode == "surrender" {
					fmt.Printf("\n SURERENDER!  Current new position's color is: %v \n ", Grid[gridPos].color)
					if Grid[gridPos].color != "green" {
						fmt.Printf("Color is not green \n")
						oldCell := playingGrid.KnownChild(prevCell).(*gi.Frame)
						oldCell.SetProp("background-color", Players[1].color)
						Grid[prevCell].color = Players[1].color

						for i := 0; i < len(Grid); i++ {

							if Grid[i].color == "green" {
								newCell := playingGrid.KnownChild(i).(*gi.Frame)

								newCell.SetProp("background-color", fmt.Sprintf("light%v", Players[0].color))
								Grid[i].color = fmt.Sprintf("light%v", Players[0].color)
								Players[0].curGrid = i
								fmt.Printf("\n Setting color to light \n")

								break

							}

						}
					} else {
						fmt.Printf("Color is green! \n")
						Players[0].curGrid = gridPos
						newCell := playingGrid.KnownChild(gridPos).(*gi.Frame)

						newCell.SetProp("background-color", fmt.Sprintf("light%v", Players[0].color))

						oldCell := playingGrid.KnownChild(prevCell).(*gi.Frame)
						oldCell.SetProp("background-color", fmt.Sprintf("%v", Players[0].color))
						Grid[gridPos].color = fmt.Sprintf("dark%v", Players[0].color)
						Grid[prevCell].color = Players[0].color
					}

				} else {
					fmt.Printf("\n NOTHING \n ")
					gridPos = prevCell
					Players[0].curGrid = prevCell

				}
			}

		}
	}

	playingGrid.SetFullReRender()
	playingGrid.UpdateEnd(updt)

}
