// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/goki/gi/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	"github.com/goki/ki/ki"
)

func main() {
	gimain.Main(func() {
		mainrun()
	})
}

func mainrun() {
	width := 1024
	height := 768

	win := gi.NewWindow2D("efight-main", "EFight Home Screen", width, height, true) // pixel sizes

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	mfr := win.SetMainFrame()
	rec := ki.Node{}
	rec.InitName(&rec, "rec")

	tv := mfr.AddNewChild(gi.KiT_TabView, "tv").(*gi.TabView)
	tv.NewTabButton = false

	homeTabk, _ := tv.AddNewTab(gi.KiT_Frame, "Home Tab")
	homeTab := homeTabk.(*gi.Frame)
	homeTab.Lay = gi.LayoutVert
	homeTab.SetStretchMaxWidth()
	homeTab.SetStretchMaxHeight()

	mainTitle := homeTab.AddNewChild(gi.KiT_Label, "mainTitle").(*gi.Label)
	mainTitle.SetProp("font-size", "x-large")
	mainTitle.SetProp("font-family", "Times New Roman, serif")
	mainTitle.SetProp("text-align", "center")
	mainTitle.Text = "<b>Welcome to EFight, an Energy Based 3D battle game!</b>"
	playButton := homeTab.AddNewChild(gi.KiT_Button, "playButton").(*gi.Button)
	playButton.Text = "<b>Play!</b>"
	playButton.SetProp("horizontal-align", gi.AlignCenter)
	homeTab.SetProp("background-color", "lightgreen")
	tv.SelectTabIndex(0)

	// main menu
	appnm := oswin.TheApp.Name()
	mmen := win.MainMenu
	mmen.ConfigMenus([]string{appnm, "Edit", "Window"})

	amen := win.MainMenu.ChildByName(appnm, 0).(*gi.Action)
	amen.Menu = make(gi.Menu, 0, 10)
	amen.Menu.AddAppMenu(win)

	emen := win.MainMenu.ChildByName("Edit", 1).(*gi.Action)
	emen.Menu = make(gi.Menu, 0, 10)
	emen.Menu.AddCopyCutPaste(win)

	win.OSWin.SetCloseCleanFunc(func(w oswin.Window) {
		go oswin.TheApp.Quit() // once main window is closed, quit
	})

	win.MainMenuUpdated()

	vp.UpdateEndNoSig(updt)

	win.StartEventLoop()
}
