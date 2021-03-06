// Copyright (c) 2018, The KaiOS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	//"reflect"

	"github.com/goki/gi/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/giv"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/units"
	"github.com/goki/ki/kit"

	//"github.com/goki/gi/units"
	"github.com/goki/ki"
	//"github.com/goki/ki/kit"

	"log"
	//"time"
)

//
// type PlanTime time.Time
//
// func (ft PlanTime) String() string {
// 	return (time.Time)(ft).Format("Mon Jan  2 15 MST 2006")
// }

type DayTime struct {
	Day  string
	Time string
}

type PlannerRec struct {
	Goal      string
	Role      string
	DayTimes  []DayTime
	Completed bool
}

func (pr *PlannerRec) Special(prompt string) {
	fmt.Printf("this is a special function!  %v", prompt)
}

type PlannerTable []*PlannerRec

func (pr *PlannerTable) SaveAs(filename gi.FileName) error {
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

func (pr *PlannerTable) SaveDefault() error {
	return pr.SaveAs("default_plan_table.json")
}

func (pr *PlannerTable) Load(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

func (pr *PlannerTable) LoadDefault() error {
	return pr.Load("default_plan_table.json")
}

func (pr *PlannerTable) GetGoals() []string {
	sa := make([]string, len(*pr)+1)
	sa[0] = SelectGoalStr
	for r := range *pr {
		rec := (*pr)[r]
		sa[r+1] = rec.Goal
	}
	return sa
}
func (pr *PlannerTable) SetGoalTime(goal, day, time string) {
	newdt := DayTime{day, time}

	fmt.Printf("Goal: %v Day: %v Time: %v \n", goal, day, time)
	for r := range *pr {
		rec := (*pr)[r]
		if rec.Goal == goal {
			rec.DayTimes = append(rec.DayTimes, newdt)

		} else {
			for di, dt := range rec.DayTimes { // check if it has been changed
				if dt == newdt {
					rec.DayTimes = append(rec.DayTimes[0:di], rec.DayTimes[di+1:]...) // delete index di

				}
			}
		}
	}
	pr.SaveDefault()
	pr.UpdateCalendar(CalendarGrid)
}

func (pr *PlannerTable) UpdateCalendar(cal *gi.Layout) {
	updt := cal.UpdateStart()
	// fmt.Printf("Goal: %v Day: %v Time: %v \n", goal, day, time)
	pr.ClearCalendar(CalendarGrid)
	for r := range *pr {
		rec := (*pr)[r]
		for _, dt := range rec.DayTimes {
			if dt.Day == "" || dt.Time == "" {
				continue
			}
			row := TimeToRow[dt.Time]
			col := DayToCol[dt.Day]
			index := (8 * row) + col
			frame := cal.KnownChild(index).(*gi.Frame)
			cb := frame.KnownChild(0).(*gi.ComboBox)
			cb.SetCurVal(rec.Goal)

			fmt.Printf("EVERYTHING: Dt: %v Row: %v Column: %v Index: %v \n", dt, row, col, index)

		}
	}
	cal.UpdateEnd(updt)
}

func (pr *PlannerTable) ClearCalendar(cal *gi.Layout) {
	goals := pr.GetGoals()
	for row := 1; row < len(Times)+1; row++ {
		for col := 1; col < len(Days)+1; col++ {
			index := (8 * row) + col
			frame := cal.KnownChild(index).(*gi.Frame)
			cb := frame.KnownChild(0).(*gi.ComboBox)

			cb.ItemsFromStringList(goals, true, 0)
			cb.SetCurVal(SelectGoalStr)

		}
	}
}

var KiT_PlannerTable = kit.Types.AddType(&PlannerTable{}, PlannerTableProps)

var PlannerTableProps = ki.Props{
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
		// {"SaveDefault", ki.Props{
		// 	"label": "Save",
		// 	"icon":  "file-save",
		// }},
		// {"LoadDefault", ki.Props{
		// 	"label": "Load",
		// 	"icon":  "file-save",
		// }},
	},
}

var CalendarGrid *gi.Layout

var ThePlan PlannerTable

var SelectGoalStr = "?"

var Times = map[int]string{
	1:  "All Day",
	2:  "7 AM",
	3:  "8 AM",
	4:  "9 AM",
	5:  "10 AM",
	6:  "11 AM",
	7:  "12 PM",
	8:  "1 PM",
	9:  "2 PM",
	10: "3 PM",
	11: "4 PM",
	12: "5 PM",
	13: "6 PM",
	14: "7 PM",
	15: "8 PM",
}

var TimeToRow map[string]int

func initTimeToRow() {
	TimeToRow = make(map[string]int, len(Times))
	for k, v := range Times {
		TimeToRow[v] = k
	}
}

var Days = map[int]string{
	1: "Sunday",
	2: "Monday",
	3: "Tuesday",
	4: "Wednesday",
	5: "Thursday",
	6: "Friday",
	7: "Saturday",
}

var DayToCol map[string]int

func initDayToCol() {
	DayToCol = make(map[string]int, len(Days))
	for k, v := range Days {
		DayToCol[v] = k
	}
}

func init() {
	initTimeToRow()
	initDayToCol()
}

// var PlannerDB *bolt.DB

// func LoadPlannerTable() []*PlannerRec {

// 	lt := make([]*PlannerRec, 0, 100) // 100 is the starting capacity of slice -- increase if you expect more users.

// 	PlannerDB.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("PlannerTable"))

// 		if b != nil {
// 			b.ForEach(func(k, v []byte) error {
// 				if v != nil {
// 					rec := PlannerRec{}
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
// 	lt := LoadPlannerTable()
// 	for _, lr := range lt {
// 		if lr.Goal == usr && lr.Role == passwd {
// 			return true
// 		}
// 	}
// 	return false
// }

// func SaveNewLogin(rec *PlannerRec) {
// 	PlannerDB.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucketIfNotExists([]byte("PlannerTable"))
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

	win := gi.NewWindow2D("kplanner", "Planner | Kataform", width, height, true) // true = pixel sizes

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
	title.Text = `<b>Planner</b>, based on the <b>7 habits of highly effective people</b>, using <b>weekly planning</b>, and <b>habit 3</b>`
	title.SetProp("text-align", gi.AlignCenter)
	title.SetProp("align-vert", gi.AlignTop)
	title.SetProp("font-size", "x-large")

	tabview := mfr.AddNewChild(gi.KiT_TabView, "tabview").(*gi.TabView)
	tab1, _ := tabview.AddNewTab(gi.KiT_SplitView, "Goals")

	split := tab1.(*gi.SplitView)
	// split.SetProp("height", "20em")

	tab2, _ := tabview.AddNewTab(gi.KiT_Layout, "Calendar")

	CalendarGrid = tab2.(*gi.Layout)
	CalendarGrid.Lay = gi.LayoutGrid
	rows := 16
	cols := 8

	CalendarGrid.SetProp("columns", cols)
	CalendarGrid.SetProp("max-width", -1)
	CalendarGrid.SetProp("max-height", -1)

	ThePlan = make(PlannerTable, 0, 1000)

	ThePlan.LoadDefault()

	// goals := ThePlan.GetGoals()

	for r := 0; r < rows; r++ {

		for c := 0; c < cols; c++ {
			cell := CalendarGrid.AddNewChild(gi.KiT_Frame, fmt.Sprintf("cell_%v_%v", r, c)).(*gi.Frame)
			cell.SetProp("background-color", "white")

			cell.SetProp("border-color", "black")
			cell.SetProp("border-width", "2px")
			cell.SetProp("max-width", -1)
			cell.SetProp("max-height", -1)
			cell.SetProp("min-height", "2em")

			if r == 0 {
				text := cell.AddNewChild(gi.KiT_Label, fmt.Sprintf("cell_%v_%v", r, c)).(*gi.Label)
				cell.SetProp("background-color", "lightgreen")
				if c == 0 {
					cell.SetProp("background-color", "black")
				} else {
					text.Text = Days[c]
				}

			} else if c == 0 {
				text := cell.AddNewChild(gi.KiT_Label, fmt.Sprintf("text_%v_%v", r, c)).(*gi.Label)
				cell.SetProp("background-color", "lightgreen")
				text.Text = Times[r]
			} else {
				combo := cell.AddNewChild(gi.KiT_ComboBox, fmt.Sprintf("combo_%v_%v", r, c)).(*gi.ComboBox)

				combo.ComboSig.Connect(vp.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
					goal := data.(string)
					cb := send.(*gi.ComboBox)
					// fmt.Printf("%v \n")
					var row, col int
					fmt.Sscanf(cb.Nm, "combo_%v_%v", &row, &col)
					ThePlan.SetGoalTime(goal, Days[col], Times[row])
				})

			}

		}

	}

	ThePlan.UpdateCalendar(CalendarGrid)

	tv := split.AddNewChild(giv.KiT_TableView, "tv").(*giv.TableView)
	// tv.SetProp("height", "20em")
	tv.Viewport = vp
	sv := split.AddNewChild(giv.KiT_StructView, "sv").(*giv.StructView)
	sv.Viewport = vp
	// split0.SetSplits(.5, .5)
	split.SetSplits(.5, .5)

	tv.SetSlice(&ThePlan, nil)

	tv.WidgetSig.Connect(sv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.WidgetSelected) {
			idx := tv.SelectedIdx
			if idx >= 0 {
				rec := ThePlan[idx]
				sv.SetStruct(rec, nil)
			}
		}
	})

	tv.ViewSig.Connect(sv.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		fmt.Printf("Saving \n")
		ThePlan.SaveDefault()
		ThePlan.UpdateCalendar(CalendarGrid)
	})

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

	// 		newPlannerRec := PlannerRec{Goal: usr, Role: passwd}
	// 		SaveNewLogin(&newPlannerRec)

	// 		vp.UpdateEnd(updt)
	// 	}
	// })

	/*lt := LoadPlannerTable()
	gi.StructTableView(lt)
	*/
	////
	// trow.AddNewChild(gi.KiT_Space, "spc1")

	// viewlogins := trow.AddNewChild(gi.KiT_Button, "viewlogins").(*gi.Button)
	// viewlogins.SetText("View PlannerTable")
	// viewlogins.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 	if sig == int64(gi.ButtonClicked) {
	// 		lt := LoadPlannerTable()

	// 		giv.TableViewDialog(vp, &lt, giv.DlgOpts{Title: "Login Table"}, nil, nil, nil)
	// 	}
	// })

	// addlogin := trow.AddNewChild(gi.KiT_Button, "addlogin").(*gi.Button)
	// addlogin.SetText("Add Login")
	// addlogin.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
	// 	if sig == int64(gi.ButtonClicked) {
	// 		rec := PlannerRec{}
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

	// note: never gets here..
	fmt.Printf("ending\n")
}
