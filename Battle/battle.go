// Copyright (c) 2018, The Kataform Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	//"encoding/json"
	"fmt"
	//"reflect"

	"github.com/goki/gi"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/oswin/driver"

	//"github.com/goki/gi/units"
	"github.com/goki/ki"
	//"github.com/goki/ki/kit"
	//bolt "github.com/coreos/bbolt"
	//"log"
	//"time"
)

func main() {
	//var err error

	driver.Main(func(app oswin.App) {
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

	win := gi.NewWindow2D("gogi-widgets-demo", "Write", width, height, true) // true = pixel sizes

	//icnm := "widget-wedge-down"

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()
	vp.Fill = true

	// style sheet
	var css = ki.Props{
		"button": ki.Props{
			"background-color": gi.Color{255, 240, 240, 255},
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

	vlay := vp.AddNewChild(gi.KiT_Frame, "vlay").(*gi.Frame)
	vlay.Lay = gi.LayoutVert
	// vlay.SetProp("background-color", "linear-gradient(to top, red, lighter-80)")
	// vlay.SetProp("background-color", "linear-gradient(to right, red, orange, yellow, green, blue, indigo, violet)")
	// vlay.SetProp("background-color", "linear-gradient(to right, rgba(255,0,0,0), rgba(255,0,0,1))")
	// vlay.SetProp("background-color", "radial-gradient(red, lighter-80)")

	trow := vlay.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()

	trow.AddNewChild(gi.KiT_Stretch, "str1")

	battle_pos_x := 0
	battle_pos_y := 0

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.SetProp("align-horiz", gi.AlignCenter)
	title.SetProp("align-vert", gi.AlignTop)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	title.Text = "Battle App"

	welcome := trow.AddNewChild(gi.KiT_Label, "welcome").(*gi.Label)
	welcome.Text = "Welcome to the Battle App. This is the first battle game that will take place in Kataform, an app that will allow basic fighting. The Battle App remains in the early channels."

	battle_row := trow.AddNewChild(gi.KiT_Layout, "battle_row").(*gi.Layout)
	battle_row.Lay = gi.LayoutVert

	move_right_button := battle_row.AddNewChild(gi.KiT_Button, "move_right_button").(*gi.Button)
	move_right_button.Text = "Move right"

	move_right_button.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
			// vp.Win.Quit()
			//gi.PromptDialog(vp, "Button1 Dialog", "This is a dialog!  Various specific types of dialogs are available.", true, true, nil, nil)
			updt := vp.UpdateStart()

			battle_pos_x = battle_pos_x + 1

			vp.UpdateEnd(updt)
		}
	})

	move_left_button := battle_row.AddNewChild(gi.KiT_Button, "move_left_button").(*gi.Button)
	move_left_button.Text = "Move left"

	move_left_button.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
			// vp.Win.Quit()
			//gi.PromptDialog(vp, "Button1 Dialog", "This is a dialog!  Various specific types of dialogs are available.", true, true, nil, nil)
			updt := vp.UpdateStart()

			battle_pos_x = battle_pos_x - 1

			vp.UpdateEnd(updt)
		}
	})

	// sig stuff

	move_down_button := battle_row.AddNewChild(gi.KiT_Button, "move_down_button").(*gi.Button)
	move_down_button.Text = "Move down"
	// sig stuff

	move_down_button.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
			// vp.Win.Quit()
			//gi.PromptDialog(vp, "Button1 Dialog", "This is a dialog!  Various specific types of dialogs are available.", true, true, nil, nil)
			updt := vp.UpdateStart()

			battle_pos_y = battle_pos_y + 1

			vp.UpdateEnd(updt)
		}
	})

	move_up_button := battle_row.AddNewChild(gi.KiT_Button, "move_up_button").(*gi.Button)
	move_up_button.Text = "Move Up"
	// sig stuff

	move_up_button.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
			// vp.Win.Quit()
			//gi.PromptDialog(vp, "Button1 Dialog", "This is a dialog!  Various specific types of dialogs are available.", true, true, nil, nil)
			updt := vp.UpdateStart()

			battle_pos_y = battle_pos_y - 1

			vp.UpdateEnd(updt)
		}
	})

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: never gets here..
	fmt.Printf("ending\n")
}
