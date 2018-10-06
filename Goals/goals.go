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

	win := gi.NewWindow2D("gogi-widgets-demo", "Acheive it", width, height, true) // true = pixel sizes

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
	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text =
		`<b>Achieve it</b>`
	title.SetProp("align-horiz", gi.AlignCenter)
	title.SetProp("align-vert", gi.AlignTop)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)

	welcome := trow.AddNewChild(gi.KiT_Label, "welcome").(*gi.Label)
	welcome.Text = "<b>Achive</b> your goals, <b>dream</b> high using <b>7 habits</b> based principals."
	welcome.SetProp("align-horiz", gi.AlignCenter)
	trow.AddNewChild(gi.KiT_Space, "spc1")
	
	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: never gets here..
	fmt.Printf("ending\n")
}

