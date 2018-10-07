// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	//"go/token"

	"github.com/goki/gi"
	//"github.com/goki/gi/complete"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/units"
	"github.com/goki/ki"
	"strconv"
	//"math"
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

	oswin.TheApp.SetName("GameOne")
	oswin.TheApp.SetAbout("This is the first game on Kataform, a basic game. More info later.")

	win := gi.NewWindow2D("game-one", "Game One", width, height, true) // true = pixel sizes

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

	currentNum := 0

	trow := mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>GameOne - a basic game.</b>"
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
	welcomeText.Text = "Welcome to game one, a basic Kataform game. GameOne is currently in its early stages, so what the game will look like in the end is not know yet."
	welcomeText.SetProp("text-align", gi.AlignCenter)

	start_button := trow.AddNewChild(gi.KiT_Button, "start_button").(*gi.Button)
	start_button.Text = "Click here to start."

	start_button.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
			// vp.Win.Quit()

			updt := vp.UpdateStart()
			gi.StringPromptDialog(vp, "", "Enter first number here",
				gi.DlgOpts{Title: "Enter first number", Prompt: "Enter the first number you want to have for the game."},
				rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
					updt := vp.UpdateStart()
					dlg := send.(*gi.Dialog)
					if sig == int64(gi.DialogAccepted) {
						val := gi.StringPromptDialogValue(dlg)

						currentNum, _ = strconv.Atoi(fmt.Sprintf("%v", val))

						fmt.Printf("First number: %v", currentNum)

						// next prompt

						gi.StringPromptDialog(vp, "", "Enter second number here:",
							gi.DlgOpts{Title: "Enter second number", Prompt: "Enter the second number you want to have for the game."},
							rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
								dlg := send.(*gi.Dialog)
								if sig == int64(gi.DialogAccepted) {

									updt := vp.UpdateStart()
									 
									
									val := gi.StringPromptDialogValue(dlg)
									
									testNum, _ := strconv.Atoi(fmt.Sprintf("%v", val))
									
									fmt.Printf("BIG: %v", testNum)
									
									//currentNum, _ = strconv.Atoi(val)
									
									
									var newCurrentNum, _ = strconv.Atoi(fmt.Sprintf("%v", cacNewNum(currentNum, testNum)))
									
									fmt.Printf("    BIG2: %v", currentNum)

									// next info

									gi.StringPromptDialog(vp, "", "Do not need to type here",
										gi.DlgOpts{Title: "Starting Number result", Prompt: fmt.Sprintf( "You have completed the first section. Your current number is %v. Click ok to continue. ", newCurrentNum)},
										rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
										//	dlg := send.(*gi.Dialog)
											if sig == int64(gi.DialogAccepted) {

												updt := vp.UpdateStart()
												//val := gi.StringPromptDialogValue(dlg)
												
												
												
												
												
 var currentNumPasser = float64(newCurrentNum)

gimain.Main(func() {
		toolWindow(currentNumPasser, float64(currentNum))
	})




												vp.UpdateEndNoSig(updt)

											}
										})

									vp.UpdateEndNoSig(updt)

								}
							})
						vp.UpdateEndNoSig(updt)
					}
				})
			vp.UpdateEndNoSig(updt)
		}

	})

	appnm := oswin.TheApp.Name()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "File", "Edit", "Window"})

	amen := win.MainMenu.KnownChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	// note: Command in shortcuts is automatically translated into Control for
	// Linux, Windows or Meta for MacOS
	fmen := win.MainMenu.KnownChildByName("File", 0).(*gi.Action)
	fmen.Menu = make(gi.Menu, 0, 10)
	fmen.Menu.AddAction(gi.ActOpts{Label: "New", Shortcut: "Command+N"},
		rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
			fmt.Printf("File:New menu action triggered\n")
		})
	fmen.Menu.AddAction(gi.ActOpts{Label: "Open", Shortcut: "Command+O"},
		rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
			fmt.Printf("File:Open menu action triggered\n")
		})
	fmen.Menu.AddAction(gi.ActOpts{Label: "Save", Shortcut: "Command+S"},
		rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
			fmt.Printf("File:Save menu action triggered\n")
		})
	fmen.Menu.AddAction(gi.ActOpts{Label: "Save As..", Shortcut: "Shift+Command+S"},
		rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
			fmt.Printf("File:SaveAs menu action triggered\n")
		})
	fmen.Menu.AddSeparator("csep")
	fmen.Menu.AddAction(gi.ActOpts{Label: "Close Window", Shortcut: "Command+W"},
		win.This, func(recv, send ki.Ki, sig int64, data interface{}) {
			win.OSWin.CloseReq()
		})

	emen := win.MainMenu.KnownChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	inQuitPrompt := false
	oswin.TheApp.SetQuitReqFunc(func() {
		if !inQuitPrompt {
			inQuitPrompt = true
			gi.PromptDialog(vp, gi.DlgOpts{Title: "Really Quit?",
				Prompt: "Are you <i>sure</i> you want to quit?"}, true, true,
				win.This, func(recv, send ki.Ki, sig int64, data interface{}) {
					if sig == int64(gi.DialogAccepted) {
						oswin.TheApp.Quit()
					} else {
						inQuitPrompt = false
					}
				})
		}
	})

	oswin.TheApp.SetQuitCleanFunc(func() {
		fmt.Printf("Doing final Quit cleanup here..\n")
	})

	inClosePrompt := false
	win.OSWin.SetCloseReqFunc(func(w oswin.Window) {
		if !inClosePrompt {
			inClosePrompt = true
			gi.PromptDialog(vp, gi.DlgOpts{Title: "Really Close Window?",
				Prompt: "Are you <i>sure</i> you want to close the window?  This will Quit the App as well."}, true, true,
				win.This, func(recv, send ki.Ki, sig int64, data interface{}) {
					if sig == int64(gi.DialogAccepted) {
						oswin.TheApp.Quit()
					} else {
						inClosePrompt = false
					}
				})
		}
	})

	win.OSWin.SetCloseCleanFunc(func(w oswin.Window) {
		fmt.Printf("Doing final Close cleanup here..\n")
	})

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}

func cacNewNum(num int, num2 int) string {
	fmt.Printf("cacNewNum working ")

	//returnValueOne := fmt.Sprintf("FIRST NUM: %v, SECOND NUM: %v", num, num2)
	//fmt.Printf("RETURN: %v", returnValue)

	numberOne := num + num2

	numberTwo := (num + num2) * (num - num2)

	numberThree := numberTwo / numberOne

	//	numberFour := math.Round(numberThree)

	returnValue := strconv.Itoa(numberThree)

	return returnValue

}

func toolWindow (startNum float64, num1 float64) {
  
  
  addAmount := startNum
  
  diff1 := 2.5
  diff2 := 3.5
  
  subtractAmount := diff1 * startNum
  divideAmount := diff2 * startNum
  
  
  currentNum := startNum
  
  
  width := 1024
	height := 768

	// turn these on to see a traces of various stages of processing..
	// gi.Update2DTrace = true
	// gi.Render2DTrace = true
	// gi.Layout2DTrace = true
	// ki.SignalTrace = true

	rec := ki.Node{}          // receiver for events
	rec.InitName(&rec, "rec") // this is essential for root objects not owned by other Ki tree nodes

	oswin.TheApp.SetName("GameOne ToolBar")
	oswin.TheApp.SetAbout("This is the first game on Kataform, a basic game. This is the toolbar.")

	win := gi.NewWindow2D("game-one-toolbar", "Game One Toolbar", width, height, true) // true = pixel sizes

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

	//currentNum := 0

	trow := mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>Game One toolbar</b>"
	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	title.SetProp("vertical-align", gi.AlignCenter)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	title.SetStretchMaxWidth()
	title.SetStretchMaxHeight()


goalNum := startNum * (-2.5 * num1)


current_num_text := trow.AddNewChild(gi.KiT_Label, "current_num_text").(*gi.Label)
current_num_text.Text = fmt.Sprintf("Your current number is %v. You are trying to get to %v.", startNum, goalNum)


addButton := trow.AddNewChild(gi.KiT_Button, "addButton").(*gi.Button)
addButton.Text = fmt.Sprintf("Click here to add %v to your number.", addAmount)


addButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
		  
		  	updt := vp.UpdateStart()
		  
			currentNum += addAmount
			current_num_text.Text = fmt.Sprintf("Your current number is %v. You are trying to get to %v.", currentNum, goalNum)
			
			
			fmt.Printf("Your current number is %v. You are trying to get to %v.", currentNum, goalNum)
			
			vp.UpdateEnd(updt)
			
		}
	})


subtractButton := trow.AddNewChild(gi.KiT_Button, "subtractButton").(*gi.Button)
subtractButton.Text = fmt.Sprintf("Click here to subtract %v from your number.", subtractAmount)



subtractButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
		  
		  	updt := vp.UpdateStart()
		  
			currentNum -= subtractAmount
			current_num_text.Text = fmt.Sprintf("Your current number is %v. You are trying to get to %v.", currentNum, goalNum)
			
			
			fmt.Printf("Your current number is %v. You are trying to get to %v.", currentNum, goalNum)
			
			vp.UpdateEnd(updt)
			
		}
	})



divideButton := trow.AddNewChild(gi.KiT_Button, "divideButton").(*gi.Button)
divideButton.Text = fmt.Sprintf("Click here to divide your number by %v.", divideAmount)


divideButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
		  
		  	updt := vp.UpdateStart()
		  
			currentNum = currentNum / divideAmount
			current_num_text.Text = fmt.Sprintf("Your current number is %v. You are trying to get to %v.", currentNum, goalNum)
			
			
			fmt.Printf("Your current number is %v. You are trying to get to %v.", currentNum, goalNum)
			
			vp.UpdateEnd(updt)
			
		}
	})
	
	
	
	checkResult := trow.AddNewChild(gi.KiT_Label, "checkResult").(*gi.Label)
	
	
	checkButton := trow.AddNewChild(gi.KiT_Button, "checkButton").(*gi.Button)
	checkButton.Text = "Click here to check your answer."
	
	
	
	checkButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
		  
		  	updt := vp.UpdateStart()
		  
			if currentNum == goalNum {
			  checkResult.Text = "VICTORY! Now run the app again and try to get other numbers to work."
			} else {
			  checkResult.Text = "Try again. Your number is incorrect."
			}
			
			vp.UpdateEnd(updt)
			
		}
	})


vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

}