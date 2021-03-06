// Copyright (c) 2018, The GoKi Authors. All rights reserved.
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
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/units"
	"github.com/goki/ki/kit"
	//"github.com/goki/gi/units"
	"github.com/goki/ki"
	//"github.com/goki/ki/kit"
	//"oswin"
	"log"
	//"time"
)

// data stuff

type NoteRec struct {
	Title string
	Note  string
}

//func (pr *NoteRec) Special(prompt string) {
//	fmt.Printf("this is a special function!  %v", prompt)
//}

type NoteTable []*NoteRec

func (pr *NoteTable) SaveAs(filename gi.FileName) error {
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

func (pr *NoteTable) SaveDefault() error {
	return pr.SaveAs("notes.json")
}

func (pr *NoteTable) Load(filename gi.FileName) error {
	b, err := ioutil.ReadFile(string(filename))
	if err != nil {
		gi.PromptDialog(nil, gi.DlgOpts{Title: "File Not Found", Prompt: err.Error()}, true, false, nil, nil)
		log.Println(err)
		return err
	}
	return json.Unmarshal(b, pr)
}

func (pr *NoteTable) LoadDefault() error {
	return pr.Load("notes.json")
}

var KiT_NoteTable = kit.Types.AddType(&NoteTable{}, NoteTableProps)

var NoteTableProps = ki.Props{
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

var TheNote NoteTable

// end of data stuff

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

	oswin.TheApp.SetName("widgets")
	oswin.TheApp.SetAbout(`This is a demo of the main widgets and general functionality of the <b>GoGi</b> graphical interface system, within the <b>GoKi</b> tree framework.  See <a href="https://github.com/goki">GoKi on GitHub</a>.
<p>The <a href="https://github.com/goki/gi/blob/master/examples/widgets/README.md">README</a> page for this example app has lots of further info.</p>`)

	win := gi.NewWindow2D("gogi-widgets-demo", "GoGi Widgets Demo", width, height, true) // true = pixel sizes

	//icnm := "widget-wedge-down"

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

	//giedsc := gi.ActiveKeyMap.ChordForFun(gi.KeyFunGoGiEditor)
	//prsc := gi.ActiveKeyMap.ChordForFun(gi.KeyFunPrefs)

	// content here:

	// tabview

	tv := mfr.AddNewChild(gi.KiT_TabView, "tv").(*gi.TabView)
	//	tv.NewTabButton = true

	// first tab

	tab1k, _ := tv.AddNewTab(gi.KiT_Layout, "Welcome screen")
	trow := tab1k.(*gi.Layout)
	trow.SetProp("white-space", gi.WhiteSpaceNormal) // wrap

	//	trow := tab1.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert

	trow.SetStretchMaxWidth()
	trow.SetStretchMaxHeight()

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label) // creates the text element
	title.Text = "Welcome to Kataform Notes"                     // sets the text value

	button := trow.AddNewChild(gi.KiT_Button, "button").(*gi.Button) // creates the text element
	button.Text = "Add note"                                         // sets the text value

	button.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
		//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
		if sig == int64(gi.ButtonClicked) {
			// Result:

			gi.StringPromptDialog(vp, "", "Enter title of new note",
				gi.DlgOpts{Title: "Enter note title", Prompt: "Enter the title of your new note."},
				rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {

					if sig == int64(gi.DialogAccepted) {
						updt := vp.UpdateStart()
						dlg := send.(*gi.Dialog)
						val := gi.StringPromptDialogValue(dlg)
						//fmt.Printf("got string value: %v\n", val)
						fmt.Printf("Got value. Your title is: %v", val)

						newNotek, _ := tv.AddNewTab(gi.KiT_Layout, fmt.Sprintf("%v", val))
						newNoteRow := newNotek.(*gi.Layout)

						newNoteRow.SetProp("white-space", gi.WhiteSpaceNormal) // wrap

						//	trow := tab1.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
						newNoteRow.Lay = gi.LayoutVert

						newNoteRow.SetStretchMaxWidth()
						newNoteRow.SetStretchMaxHeight()

						noteSave := newNoteRow.AddNewChild(gi.KiT_Button, fmt.Sprintf("noteSave%v", val)).(*gi.Button)
						noteSave.Text = "Save your note"


						noteText := newNoteRow.AddNewChild(giv.KiT_TextView, fmt.Sprintf("noteText%v", val)).(*giv.TextView)
						noteText.Placeholder = "Enter note contents here..."
						// edit1.SetText("Edit this text")
						//noteText.SetProp("min-width", "20em")

						txbuf := giv.NewTextBuf()
						txbuf.Hi.Lang = ""
						txbuf.Hi.Style = "emacs"
						txbuf.New(1) // add blank line
						//\txbuf.Open(samplefile)

						noteText.SetBuf(txbuf)
						vp.UpdateEnd(updt)

						noteSave.ButtonSig.Connect(rec.This, func(recv, send ki.Ki, sig int64, data interface{}) {
							//fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
							if sig == int64(gi.ButtonClicked) {

								Title := fmt.Sprintf("%v", val)
								Note := fmt.Sprintf("%v", txbuf.Text)
								rec := &NoteRec{Title, Note}
								

								TheNote = append(TheNote, rec)
								TheNote.SaveDefault()
							}
						})


					}
				})

			// End result
		}
	})

	// example text:

	// END of example text

	// end of content

	win.OSWin.SetCloseCleanFunc(func(w oswin.Window) {
		fmt.Printf("Doing final Close cleanup here..\n")
	})

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}
