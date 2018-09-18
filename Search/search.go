// Copyright (c) 2018, The KaiOS Authors. All rights reserved.
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
	//"github.com/goki/gi/giv"
	"github.com/goki/gi/units"
	"github.com/goki/ki/kit"
	//"github.com/goki/gi/units"
	"github.com/goki/ki"
	//"github.com/goki/ki/kit"

	"log"
	//"time"
)



type SearchRec struct {
	Title     string
	Content   string
	
}

func (pr *SearchRec) Special(prompt string) {
	fmt.Printf("this is a special function!  %v", prompt)
}

type SearchTable []*SearchRec

func (pr *SearchTable) SaveAs(filename gi.FileName) error {
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

func (pr *SearchTable) SaveDefault() error {
	return pr.SaveAs("search_recs_table.json")
}

func (pr *SearchTable) Load(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

func (pr *SearchTable) LoadDefault() error {
	return pr.Load("default_plan_table.json")
}

var KiT_SearchTable = kit.Types.AddType(&SearchTable{}, SearchTableProps)

var SearchTableProps = ki.Props{
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

var ThePlan SearchTable

// var PlannerDB *bolt.DB

// func LoadSearchTable() []*SearchRec {

// 	lt := make([]*SearchRec, 0, 100) // 100 is the starting capacity of slice -- increase if you expect more users.

// 	PlannerDB.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte("SearchTable"))

// 		if b != nil {
// 			b.ForEach(func(k, v []byte) error {
// 				if v != nil {
// 					rec := SearchRec{}
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
// 	lt := LoadSearchTable()
// 	for _, lr := range lt {
// 		if lr.Goal == usr && lr.Role == passwd {
// 			return true
// 		}
// 	}
// 	return false
// }

// func SaveNewLogin(rec *SearchRec) {
// 	PlannerDB.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucketIfNotExists([]byte("SearchTable"))
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

	win := gi.NewWindow2D("kplanner", "Search | Kataform", width, height, true) // true = pixel sizes

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

	
	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: never gets here..
	fmt.Printf("ending\n")
}
