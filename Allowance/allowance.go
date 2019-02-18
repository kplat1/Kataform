// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin"

	"encoding/json"
	"fmt"
	"io/ioutil"

	//"github.com/goki/gi/units"
	"github.com/goki/ki/kit"

	//"github.com/goki/gi/units"
	"github.com/goki/ki"
	//"github.com/goki/ki/kit"

	"log"
	//"time"
)

// data stuff starts here

type AllowanceRec struct {
	Person   string
	Spending float64
	Saving   float64
	Charity  float64
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

var TheAllowance AllowanceTable

// event data now

type EventRec struct {
	Person  string
	Event   string
	Money   float64
	Account string
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

// data stuff ends here

func main() {
	gimain.Main(func() {
		mainrun()
	})
}

func mainrun() {
	width := 1024
	height := 768

	rec := ki.Node{}          // receiver for events
	rec.InitName(&rec, "rec") // this is essential for root objects not owned by other Ki tree nodes

	fmt.Printf(fmt.Sprintf("The plan is: %v", TheAllowance))

	win := gi.NewWindow2D("Allowance", "Allowance", width, height, true) // pixel sizes

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()

	tv := mfr.AddNewChild(gi.KiT_TabView, "tv").(*gi.TabView)

	lbl1k, _ := tv.AddNewTab(gi.KiT_SplitView, "Accounts")

	// put first tab stuff in here

	/*lbl1 := lbl1k.(*gi.Label)
	  lbl1.Text = "Hello"
	*/

	split := lbl1k.AddNewChild(gi.KiT_SplitView, "split").(*gi.SplitView)

	grid := split.AddNewChild(gi.KiT_Layout, "grid").(*gi.Layout)
	grid.Lay = gi.LayoutGrid
	//fmt.Printf(fmt.Sprintf("Num:%v", pr.Load("allowance.json")))

	TheAllowance.Load("allowance.json")

	for i := 0; i < len(TheAllowance); i++ {

		fmt.Printf("Hello")

		grid_sub := grid.AddNewChild(gi.KiT_Layout, fmt.Sprintf("grid_sub_%v", i)).(*gi.Layout)
		grid_sub.Lay = gi.LayoutVert

		grid_sub.AddNewChild(gi.KiT_Space, fmt.Sprintf("spc_%v", i))

		grid_text_1 := grid_sub.AddNewChild(gi.KiT_Label, "grid_text_1").(*gi.Label)
		grid_text_1.Text = fmt.Sprintf("<b>%v</b>", TheAllowance[i].Person)
		grid_text_1.SetProp("font-size", "x-large")

		grid_text_2 := grid_sub.AddNewChild(gi.KiT_Label, "grid_text_2").(*gi.Label)
		grid_text_2.Text = fmt.Sprintf("<b>Spending</b>: %v", TheAllowance[i].Spending)

		grid_text_3 := grid_sub.AddNewChild(gi.KiT_Label, "grid_text_3").(*gi.Label)
		grid_text_3.Text = fmt.Sprintf("<b>Saving</b>: %v", TheAllowance[i].Saving)

		grid_text_4 := grid_sub.AddNewChild(gi.KiT_Label, "grid_text_4").(*gi.Label)
		grid_text_4.Text = fmt.Sprintf("<b>Charity</b>: %v", TheAllowance[i].Charity)

	}

	//lbl3 := split.AddNewChild(gi.KiT_Label, "lbl3").(*gi.Label)
	//lbl3.Text = "Text"

	/*	tablev := split.AddNewChild(giv.KiT_TableView, "tablev").(*giv.TableView)
		tablev.Viewport = vp

		sv := split.AddNewChild(giv.KiT_StructView, "sv").(*giv.StructView)
		sv.Viewport = vp

		split.SetSplits(.5, .5)


		tablev.SetSlice(&TheAllowance, nil)

		tablev.WidgetSig.Connect(sv.This, func(recv, send ki.Ki, sig int64, data interface{}) {
			if sig == int64(gi.WidgetSelected) {
				idx := tablev.SelectedIdx
				if idx >= 0 {
					rec := TheAllowance[idx]

					sv.SetStruct(rec, nil)
				}
			}
		})


	*/

	// next will be 2nd tab stuff
	lbl2k, _ := tv.AddNewTab(gi.KiT_TabView, "Transaction / Events")
	lbl2 := lbl2k.(*gi.TabView)
	fmt.Printf(fmt.Sprintf("%v", lbl2k))

	transtvk, _ := lbl2.AddNewTab(gi.KiT_TabView, "transtv")
	transtv := transtvk.(*gi.TabView)
	fmt.Printf(fmt.Sprintf("%v", transtv))

	TheAllowance.Load("allowance.json")
	TheEvent.Load("event.json")
	fmt.Printf("EVENTS: %v", len(TheAllowance))

	for i := 0; i < len(TheAllowance); i++ {
		tab, _ := transtv.AddNewTab(gi.KiT_Layout, fmt.Sprintf("%v", TheAllowance[i].Person))
		addNewTrans := tab.AddNewChild(gi.KiT_Button, fmt.Sprintf("addNewTrans_%v", i)).(*gi.Button)
		addNewTrans.Text = "Click here to add new transaction"
		fmt.Printf(fmt.Sprintf("AA_%v", tab))
		//split2 := tab.AddNewChild(gi.KiT_SplitView, "split2").(*gi.SplitView)
		//split_lay := tab.AddNewChild(gi.KiT_Layout, fmt.Sprintf("split_lay_%v", i)).(*gi.Layout)
		//split_lay.Lay = gi.LayoutVert

		for j := 0; j < len(TheEvent); j++ {
			fmt.Printf("BB_")
			//fmt.Printf(fmt.Sprintf("CONTENT: %v"), TheEvent[j].Person)
			if TheEvent[j].Person == TheAllowance[i].Person {

				addNewTrans.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
					//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
					if sig == int64(gi.ButtonClicked) { // note: 3 diff ButtonSig sig's possible -- important to check
						rec := &EventRec{}
						giv.StructViewDialog(vp, rec, giv.DlgOpts{Title: "Transaction", Prompt: "Enter Transaction Info", Ok: true, Cancel: true}, recv, func(recv, send ki.Ki, sig int64, data interface{}) {
							if sig == int64(gi.DialogAccepted) {
								TheEvent = append(TheEvent, rec)
								TheEvent.SaveDefault()

								// add allowance updater
							}
						})

					}
				})

				split_lay_text := tab.AddNewChild(gi.KiT_Label, fmt.Sprintf("event_tab_lay_%v", j)).(*gi.Label)

				split_lay_text.Text = fmt.Sprintf("<b>Event:</b> %v", TheEvent[j].Event)
				tab.AddNewChild(gi.KiT_Space, fmt.Sprintf("spc_%v", j))

			}

		}

	}

	tv.SelectTabIndex(0)

	// main menu
	appnm := oswin.TheApp.Name()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

	amen := win.MainMenu.KnownChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.KnownChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	win.OSWin.SetCloseCleanFunc(func(w oswin.Window) {
		go oswin.TheApp.Quit() // once main window is closed, quit
	})

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()
}
