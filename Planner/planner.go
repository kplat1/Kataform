// Copyright (c) 2018, The KaiOS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	//"reflect"

	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin"
	//"github.com/goki/gi/units"
	"github.com/goki/ki"
	//"github.com/goki/ki/kit"

	bolt "github.com/coreos/bbolt"
	"log"
	//"time"
)

type PlannerRec struct {
	Goal string
	Role string
}

var PlannerDB *bolt.DB

func LoadPlannerTable() []*PlannerRec {

	lt := make([]*PlannerRec, 0, 100) // 100 is the starting capacity of slice -- increase if you expect more users.

	PlannerDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PlannerTable"))

		if b != nil {
			b.ForEach(func(k, v []byte) error {
				if v != nil {
					rec := PlannerRec{}
					json.Unmarshal(v, &rec) // loads v value as json into rec
					lt = append(lt, &rec)   // adds record to login table

				}
				return nil
			})
		}
		return nil
	})

	return lt
}

func CheckLogin(usr, passwd string) bool {
	lt := LoadPlannerTable()
	for _, lr := range lt {
		if lr.Goal == usr && lr.Role == passwd {
			return true
		}
	}
	return false
}

func SaveNewLogin(rec *PlannerRec) {
	PlannerDB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("PlannerTable"))
		jb, err := json.Marshal(rec) // converts rec to json, as bytes jb

		err = b.Put([]byte(rec.Goal), jb)
		return err
	})
}

func main() {
	var err error
	PlannerDB, err = bolt.Open("Planner.db", 0600, nil)

	if err != nil {
		log.Fatal(err)
	}
	defer PlannerDB.Close()

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

	win := gi.NewWindow2D("gogi-widgets-demo", "Planner | Kataform", width, height, true) // true = pixel sizes

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
	vlay.SetProp("font-family", "Georgia, serif")
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
		`<b>Planner</b>`
	title.SetProp("text-align", gi.AlignCenter)
	title.SetProp("align-vert", gi.AlignTop)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	title.SetProp("line-height", 1.5)
	trow.AddNewChild(gi.KiT_Stretch, "str2")

	p1 := trow.AddNewChild(gi.KiT_Label, "p1").(*gi.Label)
	p1.Text = "<b>Planner</b>, based on the <b>7 habits of highly effective people</b>, using <b>weekly planning</b>, and <b>habit 3</b>"
	p1.SetProp("text-align", gi.AlignCenter)

	///
	motivationText := trow.AddNewChild(gi.KiT_Label, "motivationText").(*gi.Label)
	motivationText.Text = "<b>Create Q2 goals for each of your roles! Make sure to do this weekly</b>"
	motivationText.SetProp("text-align", gi.AlignCenter)

	buttonStartResult := trow.AddNewChild(gi.KiT_Label, "buttonStartResult").(*gi.Label)
	buttonStartResult.Text = "<b>Add new goal:</b>"
	userText := trow.AddNewChild(gi.KiT_TextField, "userText").(*gi.TextField)
	userText.SetText("Goal")
	userText.SetProp("width", "20em")
	passwdText := trow.AddNewChild(gi.KiT_TextField, "passwdText").(*gi.TextField)
	passwdText.SetText("Role")
	passwdText.SetProp("width", "20em")

	signUpButton := trow.AddNewChild(gi.KiT_Button, "signUpButton").(*gi.Button)
	signUpButton.Text = "<b>Create!</b>"

	signUpButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
			// vp.Win.Quit()
			//gi.PromptDialog(vp, "Button1 Dialog", "This is a dialog!  Various specific types of dialogs are available.", true, true, nil, nil)
			updt := vp.UpdateStart()
			usr := userText.Text()
			passwd := passwdText.Text()

			newPlannerRec := PlannerRec{Goal: usr, Role: passwd}
			SaveNewLogin(&newPlannerRec)

			vp.UpdateEnd(updt)
		}
	})

	/*lt := LoadPlannerTable()

	gi.StructTableView(lt)
	*/
	////
	trow.AddNewChild(gi.KiT_Space, "spc1")

	viewlogins := trow.AddNewChild(gi.KiT_Button, "viewlogins").(*gi.Button)
	viewlogins.SetText("View PlannerTable")
	viewlogins.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			lt := LoadPlannerTable()

			giv.TableViewDialog(vp, &lt, giv.DlgOpts{Title: "Login Table"}, nil, nil, nil)
		}
	})

	addlogin := trow.AddNewChild(gi.KiT_Button, "addlogin").(*gi.Button)
	addlogin.SetText("Add Login")
	addlogin.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			rec := PlannerRec{}
			giv.StructViewDialog(vp, &rec, giv.DlgOpts{Title: "Enter Login Info"}, recv, func(recv, send ki.Ki, sig int64, data interface{}) {
				if sig == int64(gi.DialogAccepted) {
					SaveNewLogin(&rec)
				}
			})
		}
	})

	quit := trow.AddNewChild(gi.KiT_Button, "quit").(*gi.Button)
	quit.SetText("Quit")
	quit.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			gi.PromptDialog(vp, gi.DlgOpts{Title: "Quit", Prompt: "Quit: Are You Sure?"}, true, true, recv, func(recv, send ki.Ki, sig int64, data interface{}) {
				if sig == int64(gi.DialogAccepted) {
					PlannerDB.Close()
					oswin.TheApp.Quit()
				}
			})
		}
	})

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: never gets here..
	fmt.Printf("ending\n")
}
