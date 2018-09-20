// Copyright (c) 2018, The Kataform Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"time"
	//"reflect"

	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/units"
	"github.com/goki/ki/kit"

	//"github.com/goki/gi/units"
	"github.com/goki/ki"
	//"github.com/goki/ki/kit"

	"log"
	//"time"
)

type AllowanceRec struct {
	Person  string
	Balance float64
}

// ** Work on allowance app and action **

func (pr *AllowanceRec) Special(prompt string) {
	fmt.Printf("this is a special function!  %v", prompt)
}

type AllowanceTable []*AllowanceRec




func (pr *AllowanceTable) SaveAs(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return nil
}

func (pr *AllowanceTable) SaveDefault() error {
	return pr.SaveAs("allowance.json")
}

func (pr *AllowanceTable) Load(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

func (pr *AllowanceTable) LoadDefault() error {
	return pr.Load("allowance.json")
}

var KiT_AllowanceTable = kit.Types.AddType(&AllowanceTable{}, AllowanceTableProps)

var AllowanceTableProps = ki.Props{
	"MainMenu": ki.PropSlice{
		{"AppMenu", ki.BlankProp{}},
		{"File", ki.PropSlice{
			{"LoadDefault", ki.Props{
				"shortcut": "Command+O",
			}},
			{"SaveDefault", ki.Props{
				"shortcut": "Command+S",
			}},
			{"sep-close", ki.BlankProp{}},
			{"Close Window", ki.BlankProp{}},
		}},
		{"Edit", "Copy Cut Paste"},
		{"Window", "Windows"},
	},
	"ToolBar": ki.PropSlice{
		{"SaveDefault", ki.Props{
			"label": "Save",
			"icon":  "file-save",
		}},
		{"LoadDefault", ki.Props{
			"label": "Load",
			"icon":  "file-save",
		}},
	},
}

var ThePlan AllowanceTable



// event data now


type EventRec struct {
	Person  string
	Event string
}

// ** Work on Event app and action **

func (pr *EventRec) Special(prompt string) {
	fmt.Printf("this is a special function!  %v", prompt)
}

type EventTable []*EventRec




func (pr *EventTable) SaveAs(filename gi.FileName) error {
	b, err := json.MarshalIndent(pr, "", "  ")
	if err != nil {
		log.Println(err) // unlikely
		return err
	}
	err = ioutil.WriteFile(string(filename), b, 0644)
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "Could not Save to File", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
	}
	return nil
}

func (pr *EventTable) SaveDefault() error {
	return pr.SaveAs("event.json")
}

func (pr *EventTable) Load(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

func (pr *EventTable) LoadDefault() error {
	return pr.Load("event.json")
}

var KiT_EventTable = kit.Types.AddType(&EventTable{}, EventTableProps)

var EventTableProps = ki.Props{
	"MainMenu": ki.PropSlice{
		{"AppMenu", ki.BlankProp{}},
		{"File", ki.PropSlice{
			{"LoadDefault", ki.Props{
				"shortcut": "Command+O",
			}},
			{"SaveDefault", ki.Props{
				"shortcut": "Command+S",
			}},
			{"sep-close", ki.BlankProp{}},
			{"Close Window", ki.BlankProp{}},
		}},
		{"Edit", "Copy Cut Paste"},
		{"Window", "Windows"},
	},
	"ToolBar": ki.PropSlice{
		{"SaveDefault", ki.Props{
			"label": "Save",
			"icon":  "file-save",
		}},
		{"LoadDefault", ki.Props{
			"label": "Load",
			"icon":  "file-save",
		}},
	},
}

var TheEvent EventTable




// var PlannerDB *bolt.DB

// func LoadAllowanceTable() []*AllowanceRec {

// 	lt := make([]*AllowanceRec, 0, 100) // 100 is the starting capacity of slice -- increase if you expect more users.

// 	PlannerDB.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("AllowanceTable"))

// 		if b != nil {
// 			b.ForEach(func(k, v []byte) error {
// 				if v != nil {
// 					rec := AllowanceRec{}
// 					json.Unmarshal(v, &rec) // loads v value as json into rec
// 					lt = append(lt, &rec)   // adds record to login table

// 				}
// 				return nil
// 			})
// 		}
// 		return nil
// 	})

// 	return lt
// }

// func CheckLogin(usr, passwd string) bool {
// 	lt := LoadAllowanceTable()
// 	for _, lr := range lt {
// 		if lr.Goal == usr && lr.Role == passwd {
// 			return true
// 		}
// 	}
// 	return false
// }

// func SaveNewLogin(rec *AllowanceRec) {
// 	PlannerDB.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucketIfNotExists([]byte("AllowanceTable"))
// 		jb, err := json.Marshal(rec) // converts rec to json, as bytes jb

// 		err = b.Put([]byte(rec.Goal), jb)
// 		return err
// 	})
// }

func main() {
	// var err error
	// PlannerDB, err = bolt.Open("Planner.db", 0600, nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer PlannerDB.Close()

	gimain.Main(func() {
		mainrun()
	})
}




func mainrun() {
	width := 1200
	height := 900

	win := gi.NewWindow2D("kplanner", "Allowance | Kataform", width, height, true) // true = pixel sizes

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()
	vp.Fill = true

	// // style sheet
	// var css = ki.Props{
	// 	"button": ki.Props{
	// 		"background-color": gi.Color{255, 240, 240, 255},
	// 	},
	// 	"#combo": ki.Props{
	// 		"background-color": gi.Color{240, 255, 240, 255},
	// 	},
	// 	".hslides": ki.Props{
	// 		"background-color": gi.Color{240, 225, 255, 255},
	// 	},
	// 	"kbd": ki.Props{
	// 		"color": "blue",
	// 	},
	// }
	// vp.CSS = css

	mfr := win.SetMainFrame()
	mfr.SetProp("spacing", units.NewValue(1, units.Ex))
	mfr.SetProp("font-family", "Georgia, serif")

	trow := mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()








	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = `<b>Allowance</b>, made easy. This is a product of Kataform to be used to track allowance.`
	title.SetProp("text-align", gi.AlignCenter)
	title.SetProp("align-vert", gi.AlignTop)
	title.SetProp("font-size", "x-large")


	split := mfr.AddNewChild(gi.KiT_SplitView, "split").(*gi.SplitView)

	tv := split.AddNewChild(giv.KiT_TableView, "tv").(*giv.TableView)
	tv.Viewport = vp
	sv := split.AddNewChild(giv.KiT_StructView, "sv").(*giv.StructView)
	sv.Viewport = vp

	split.SetSplits(.5, .5)
	tv.SetSlice(&ThePlan, nil)

	tv.WidgetSig.Connect(sv.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.WidgetSelected) {
			idx := tv.SelectedIdx
			if idx >= 0 {
				rec := ThePlan[idx]
				sv.SetStruct(rec, nil)
			}
		}
	})
	
	
	
	
	
	// events below here
	
	split2 := mfr.AddNewChild(gi.KiT_SplitView, "split2").(*gi.SplitView)

	tv2 := split2.AddNewChild(giv.KiT_TableView, "tv2").(*giv.TableView)
	tv2.Viewport = vp
	sv2 := split.AddNewChild(giv.KiT_StructView, "sv2").(*giv.StructView)
	sv2.Viewport = vp

	split2.SetSplits(.5, .5)
	tv2.SetSlice(&TheEvent, nil)

	tv2.WidgetSig.Connect(sv.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.WidgetSelected) {
			idx := tv.SelectedIdx
			if idx >= 0 {
				rec := TheEvent[idx]
				sv2.SetStruct(rec, nil)
			}
		}
	})
	/*
	//next tab here
	win2 := gi.NewWindow2D("kplanner", "Allowance2 | Kataform", width, height, true) // true = pixel sizes
	
	vp2 := win2.WinViewport2D()
	updt2 := vp.UpdateStart()
	//vp2.Fill = true

	// // style sheet
	// var css = ki.Props{
	// 	"button": ki.Props{
	// 		"background-color": gi.Color{255, 240, 240, 255},
	// 	},
	// 	"#combo": ki.Props{
	// 		"background-color": gi.Color{240, 255, 240, 255},
	// 	},
	// 	".hslides": ki.Props{
	// 		"background-color": gi.Color{240, 225, 255, 255},
	// 	},
	// 	"kbd": ki.Props{
	// 		"color": "blue",
	// 	},
	// }
	// vp.CSS = css

	mfr2 := win2.SetMainFrame()
	mfr2.SetProp("spacing", units.NewValue(1, units.Ex))
	mfr2.SetProp("font-family", "Georgia, serif")

	trow2 := mfr2.AddNewChild(gi.KiT_Layout, "trow2").(*gi.Layout)
	trow2.Lay = gi.LayoutVert
	trow2.SetStretchMaxWidth()




	title2 := trow2.AddNewChild(gi.KiT_Label, "title2").(*gi.Label)
	title2.Text = `<b>Allowance</b>, made easy. This is a product of Kataform to be used to track allowance.`
	title2.SetProp("text-align", gi.AlignCenter)
	title2.SetProp("align-vert", gi.AlignTop)
	title2.SetProp("font-size", "x-large")

*/
	
	// motivationText := trow.AddNewChild(gi.KiT_Label, "motivationText").(*gi.Label)
	// motivationText.Text = "<b>Create Q2 goals for each of your roles! Make sure to do this weekly</b>"
	// motivationText.SetProp("text-align", gi.AlignCenter)

	// buttonStartResult := trow.AddNewChild(gi.KiT_Label, "buttonStartResult").(*gi.Label)
	// buttonStartResult.Text = "<b>Add new goal:</b>"
	// userText := trow.AddNewChild(gi.KiT_TextField, "userText").(*gi.TextField)
	// userText.SetText("Goal")
	// userText.SetProp("width", "20em")
	// passwdText := trow.AddNewChild(gi.KiT_TextField, "passwdText").(*gi.TextField)
	// passwdText.SetText("Role")
	// passwdText.SetProp("width", "20em")

	// signUpButton := trow.AddNewChild(gi.KiT_Button, "signUpButton").(*gi.Button)
	// signUpButton.Text = "<b>Create!</b>"

	// signUpButton.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 	//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
	// 	if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
	// 		// vp.Win.Quit()
	// 		//gi.PromptDialog(vp, "Button1 Dialog", "This is a dialog!  Various specific types of dialogs are available.", true, true, nil, nil)
	// 		updt := vp.UpdateStart()
	// 		usr := userText.Text()
	// 		passwd := passwdText.Text()

	// 		newAllowanceRec := AllowanceRec{Goal: usr, Role: passwd}
	// 		SaveNewLogin(&newAllowanceRec)

	// 		vp.UpdateEnd(updt)
	// 	}
	// })

	/*lt := LoadAllowanceTable()

	gi.StructTableView(lt)
	*/
	////
	// trow.AddNewChild(gi.KiT_Space, "spc1")

	// viewlogins := trow.AddNewChild(gi.KiT_Button, "viewlogins").(*gi.Button)
	// viewlogins.SetText("View AllowanceTable")
	// viewlogins.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 	if sig == int64(gi.ButtonClicked) {
	// 		lt := LoadAllowanceTable()

	// 		giv.TableViewDialog(vp, &lt, giv.DlgOpts{Title: "Login Table"}, nil, nil, nil)
	// 	}
	// })

	// addlogin := trow.AddNewChild(gi.KiT_Button, "addlogin").(*gi.Button)
	// addlogin.SetText("Add Login")
	// addlogin.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 	if sig == int64(gi.ButtonClicked) {
	// 		rec := AllowanceRec{}
	// 		giv.StructViewDialog(vp, &rec, giv.DlgOpts{Title: "Enter Login Info"}, recv, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 			if sig == int64(gi.DialogAccepted) {
	// 				SaveNewLogin(&rec)
	// 			}
	// 		})
	// 	}
	// })

	// quit := trow.AddNewChild(gi.KiT_Button, "quit").(*gi.Button)
	// quit.SetText("Quit")
	// quit.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 	if sig == int64(gi.ButtonClicked) {
	// 		gi.PromptDialog(vp, gi.DlgOpts{Title: "Quit", Prompt: "Quit: Are You Sure?"}, true, true, recv, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 			if sig == int64(gi.DialogAccepted) {
	// 				PlannerDB.Close()
	// 				oswin.TheApp.Quit()
	// 			}
	// 		})
	// 	}
	// })

	
vp.UpdateEndNoSig(updt)


	win.StartEventLoop()
	

	// note: never gets here..
	fmt.Printf("ending\n")
}
