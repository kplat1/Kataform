package main

import (
	"fmt"
	//"go/token"

	"github.com/goki/gi/gi"
	//"github.com/goki/gi/giv"
	//"github.com/goki/gi/complete"
// 	"math/rand"

	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"

	"github.com/goki/gi/units"
	"github.com/goki/ki"

	"github.com/goki/ki/kit"
	// 	"github.com/goki/gi/svg"
	//"strconv"
	//"math
	"github.com/goki/gi/oswin/key"
	"github.com/goki/gi/svg"
// 	"time"
)

func main() {

	gimain.Main(func() {
		mainrun()
	})
}

type GameFrame struct {
	Row *gi.Layout
	gi.Frame
}

var KiT_GameFrame = kit.Types.AddType(&GameFrame{}, nil)

func (gf *GameFrame) ConnectEvents2D() {
	// 	fmt.Printf("Hi \n")
	gf.ConnectEvent(oswin.KeyChordEvent, gi.HiPri, func(recv, send ki.Ki, sig int64, d interface{}) {
		// fvv := recv.Embed(KiT_DomFrame).(*DomFrame)
		kt := d.(*key.ChordEvent)
		ch := kt.Chord()

		// fmt.Printf("HI2 \n")
		switch ch {
		case "w":
			kt.SetProcessed()
			gf.UpAction()
			case "a":
			kt.SetProcessed()
			gf.LeftAction()
			case "s":
			kt.SetProcessed()
			gf.DownAction()
			case "d":
			kt.SetProcessed()
			gf.RightAction()
		}
	})

}

func (gf *GameFrame) HasFocus2D() bool {
	return true // always.. we're typically a dialog anyway
}
func (gf *GameFrame) UpAction() {

	
	up, _ := gf.Row.ChildByName("upAction", 0)
	up.(*gi.Action).Trigger()
}

func (gf *GameFrame) DownAction() {

	
	down, _ := gf.Row.ChildByName("downAction", 0)
	down.(*gi.Action).Trigger()
}

func (gf *GameFrame) RightAction() {

	
	right, _ := gf.Row.ChildByName("rightAction", 0)
	right.(*gi.Action).Trigger()
}
func (gf *GameFrame) LeftAction() {

	
	left, _ := gf.Row.ChildByName("leftAction", 0)
	left.(*gi.Action).Trigger()
}



var SvgGame *svg.SVG
var SvgPeople *svg.Group
var SvgMap *svg.Group

var gmin, gmax, gsz, ginc gi.Vec2D
var GameSize float32 = 200


var trow *gi.Layout

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

	oswin.TheApp.SetName("Game")
	oswin.TheApp.SetAbout("Game")

	win := gi.NewWindow2D("game", "Game", width, height, true) // true = pixel sizes

	vp := win.WinViewport2D()
	updt := vp.UpdateStart()

	// style sheet
	var css = ki.Props{
		"Action": ki.Props{
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
	// 	dfr := mfr.AddNewChild(KiT_DomFrame, "domframe").(*DomFrame)
	mfr.SetProp("spacing", units.NewValue(1, units.Ex))
	// mfr.SetProp("background-color", "linear-gradient(to top, red, lighter-80)")
	// dfr.SetProp("background-color", "linear-gradient(to right, red, orange, yellow, green, blue, indigo, violet)")
	// dfr.SetProp("background-color", "linear-gradient(to right, rgba(255,0,0,0), rgba(255,0,0,1))")
	// dfr.SetProp("background-color", "radial-gradient(red, lighter-80)")

	// vars in here :

	// end of vars

	trow = mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	trow.Lay = gi.LayoutVert
	trow.SetStretchMaxWidth()

	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>BlueRed</b>"
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
	
	gfr := mfr.AddNewChild(KiT_GameFrame, "gameframe").(*GameFrame)
	gfr.SetProp("background-color", "white")

	gfr.Row = mfr.AddNewChild(gi.KiT_Layout, "brow").(*gi.Layout)
	gfr.Row.Lay = gi.LayoutHoriz
	
	
	start := gfr.Row.AddNewChild(gi.KiT_Action, "start").(*gi.Action)
	start.Text = "Start!"

	start.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
	  updt := SvgGame.UpdateStart()
	  
	go MainLoop()
	
	SvgGame.UpdateEnd(updt)

	})
	
	

	upAction := gfr.Row.AddNewChild(gi.KiT_Action, "upAction").(*gi.Action)
	upAction.Text = "Move Up!"

	upAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
	  updt := SvgGame.UpdateStart()
	  
	player.Pos.Y = player.Pos.Y + 0.5
	
	SvgGame.UpdateEnd(updt)

	})

	downAction := gfr.Row.AddNewChild(gi.KiT_Action, "downAction").(*gi.Action)
	downAction.Text = "Move Down!"

	downAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		
updt := SvgGame.UpdateStart()
	  
	player.Pos.Y = player.Pos.Y - 0.5
	
	SvgGame.UpdateEnd(updt)
	})
	
	rightAction := gfr.Row.AddNewChild(gi.KiT_Action, "rightAction").(*gi.Action)
	rightAction.Text = "Move Right!"
	
	rightAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		
updt := SvgGame.UpdateStart()
	  
	player.Pos.X = player.Pos.X + 0.5
	
	SvgGame.UpdateEnd(updt)
	})
	
	
	
	leftAction := gfr.Row.AddNewChild(gi.KiT_Action, "leftAction").(*gi.Action)
	leftAction.Text = "Move Left!"
	
	leftAction.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		
updt := SvgGame.UpdateStart()
	  
	player.Pos.X = player.Pos.X - 0.5
	
	SvgGame.UpdateEnd(updt)
	})
	
	

	gfr.AddNewChild(gi.KiT_Space, "spc2")

	SvgGame = gfr.AddNewChild(svg.KiT_SVG, "SvgGame").(*svg.SVG)
	SvgGame.SetProp("min-width", GameSize)
	SvgGame.SetProp("min-height", GameSize)
	SvgGame.SetStretchMaxWidth()
	SvgGame.SetStretchMaxHeight()

	
	SvgPeople = SvgGame.AddNewChild(svg.KiT_Group, "SvgPeople").(*svg.Group)
	SvgMap = SvgGame.AddNewChild(svg.KiT_Group, "SvgMap").(*svg.Group)

	gmin = gi.Vec2D{-10, -10}
	gmax = gi.Vec2D{10, 10}
	gsz = gmax.Sub(gmin)
	ginc = gsz.DivVal(GameSize)

	SvgGame.ViewBox.Min = gmin
	SvgGame.ViewBox.Size = gsz
	SvgGame.Norm = true
	SvgGame.InvertY = true
	SvgGame.Fill = true
	SvgGame.SetProp("background-color", "white")
	SvgGame.SetProp("stroke-width", ".8pct")

	InitMap()
	InitPlayer()

	//////////////////////////////////////////
	//      Main Menu

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

	// note: may eventually get down here on a well-behaved quit, but better
	// to handle cleanup above using QuitCleanFunc, which happens before all
	// windows are closed etc
	fmt.Printf("main loop ended\n")
}

func InitMap() {
		updt := SvgGame.UpdateStart()
		
			bottomLine := SvgMap.AddNewChild(svg.KiT_Line, "bottomLine").(*svg.Line)

	bottomLine.SetProp("stroke", "black")
	bottomLine.Start = gi.Vec2D{-10, -8}
	bottomLine.End = gi.Vec2D{10, -8}
	
	
	
	topLine := SvgMap.AddNewChild(svg.KiT_Line, "topLine").(*svg.Line)

	topLine.SetProp("stroke", "black")
	topLine.Start = gi.Vec2D{-10, 8}
	topLine.End = gi.Vec2D{10, 8}
	
	
	rightLine := SvgMap.AddNewChild(svg.KiT_Line, "rightLine").(*svg.Line)

	rightLine.SetProp("stroke", "black")
	rightLine.Start = gi.Vec2D{10, 8}
	rightLine.End = gi.Vec2D{10, -8}
	
	leftLine := SvgMap.AddNewChild(svg.KiT_Line, "leftLine").(*svg.Line)

	leftLine.SetProp("stroke", "black")
	leftLine.Start = gi.Vec2D{-10, 8}
	leftLine.End = gi.Vec2D{-10, -8}
	


	SvgGame.UpdateEnd(updt)

}

 var player *svg.Rect

func InitPlayer() {
	updt := SvgGame.UpdateStart()
	SvgPeople.DeleteChildren(true)

	player = SvgPeople.AddNewChild(svg.KiT_Rect, "player").(*svg.Rect)

	player.SetProp("fill", "blue")
	player.SetProp("stroke", "green")
	player.Size = gi.Vec2D{2, 2}
	player.Pos = gi.Vec2D{-8, 6}

	SvgGame.UpdateEnd(updt)

}



// func JumpLoop() {
//   fmt.Printf("HIII \n")

//   for y := -9.9; y > -10; y++ {
//     updt := SvgGame.UpdateStart()
//     if y < 10 {

//       if VertSpeed == 1 {
//       player.Pos.Y = float32(y)
//       } else {
//         player.Pos.Y = float32(y) - 2
//       }

//     } else {
//       fmt.Printf("Coming down \n")
//           SvgGame.UpdateEnd(updt)

//       break
//     }
//     SvgGame.UpdateEnd(updt)
//     time.Sleep(1 * time.Millisecond)

//   }
//   JumpLoopDown()

// }

// func JumpLoopDown() {
//   fmt.Printf("Coming down func \n")

//   for y := player.Pos.Y; y >= -10; y-- {
//     updt := SvgGame.UpdateStart()
//     player.Pos.Y = float32(y)
//     SvgGame.UpdateEnd(updt)
//     fmt.Printf("Updated before this! \n")
//     time.Sleep(1 * time.Millisecond)

//   }

// }

var obstacle *svg.Rect

func MainLoop() {
	
	for i := 0; i > 0; i++ {

		updt := SvgGame.UpdateStart()
		
		if (player.Pos.X + player.Size.X > 8) || (player.Pos.X < -8) {
		  if player.Pos.X > 0 {
		    player.Pos.X = 6
		  } else {
		    player.Pos.X = -8
		  }
		}
 
		SvgGame.UpdateEnd(updt)

	}
}