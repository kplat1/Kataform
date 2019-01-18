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
	"strconv"
	"strings"
)

type player struct {
	Name     string
	Position string
	Number   int
	Hometown string
	Right    bool
	College  string
	Height   float32
	Color    string
	GameFilm string
}

type team struct {
	Name string
}

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
		case "z":
			kt.SetProcessed()
			gf.UpAction()
		}
	})

}

func (gf *GameFrame) HasFocus2D() bool {
	return true // always.. we're typically a dialog anyway
}
func (gf *GameFrame) UpAction() {

	// 	fmt.Printf("Up action!!\n")
	// 	up, _ := gf.Row.ChildByName("upAction", 0)
	// 	up.(*gi.Action).Trigger()
}

var SvgGame *svg.SVG
var SvgPeople *svg.Group
var SvgMap *svg.Group

var gmin, gmax, gsz, ginc gi.Vec2D
var GameSize float32 = 200

var trow *gi.Layout
var Players []player
var Teams []team
var tab0frame *gi.Layout
var TEAM string
var tv *gi.TabView
var vp *gi.Viewport2D

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

	vp = win.WinViewport2D()
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
	Players = append(Players, player{"Isaiah Love", "Quarterback", 4, "Cleveland, Ohio", true, "Ohio State", 6.2, "Light Brown", "33 yard td run @ Michigan State"})

	Players = append(Players, player{"TJ Bowers", "Quarterback", 14, "Fort Lauderdale, Florida", false, "Miami", 6.2, "White", "53 yard td pass made at the right side of the end-zone vs Duke"})
	Players = append(Players, player{"LeShaun Benjamin", "Wide Receiver", 11, "Fresno, California", true, "UCLA", 6.4, "Brown", "30 yard leaping td catch made at the back of the end-zone vs Arizona"})
	Players = append(Players, player{"Kyler Henley", "Wide Receiver", 12, "Jacksonville, Florida", true, "Georgia", 6.1, "Black", " 26 yard one-handed td catch @ Alabama. (Catch made at left side of field at the 4 yard line, and runs it i)"})
	Players = append(Players, player{"Gavin Sczepaniak", "Tight End", 87, "Grantsburg, Wisconsin", true, "Wisconsin", 6.4, "White", "10 yard td catch at the right side of the end-zone @ Northwestern. (Catch made at 1 yard line, and walks in)"})
	Players = append(Players, player{"Henri Fisher", "Tight End", 80, "Detroit, Michigan", true, "Michigan State", 6.5, "White", "19 yard one handed td catch with his right hand made at the left side of the end-zone vs Ohio State"})
	Players = append(Players, player{"Deion Humphrey", "Running Back", 28, "Pittsburgh, Pennsylvania", true, "Michigan", 5.11, "Black", " 41 yard td run vs Nebraska"})
	Players = append(Players, player{"Myles Fuller", "Running Back", 9, "Little Rock, Arkansas", false, "Oklahoma", 6.0, "Black", "99 yard kickoff return @ West Virginia"})
	Players = append(Players, player{"Leo Haynes", "Offensive Lineman (Tackle)", 66, "Hunstville, Alabama", true, "Alabama", 6.6, "White", "No game film for OLs or Punters"})
	Players = append(Players, player{"Connor Owens", "Offensive Lineman (Guard)", 78, "Fort Wayne, Indiana", false, "Notre Dame", 6.6, "White", "No game film for OLs or Punters"})
	Players = append(Players, player{"Peyton Lowe", "Defensive End", 99, "Buffalo, New York", true, "Syracuse", 6.3, "White", "7 yard sack at the 27 (of other team) vs NC State"})
	Players = append(Players, player{"Jamar Rogers", "Defensive End", 97, "Salem, Oregon", true, "Washington", 6.4, "Dark Black", "10 yard sack at the 48 (of his team) vs Colorado"})
	Players = append(Players, player{"Liam Van Dijk", "Defensive Tackle", 91, "Salt Lake City, Utah", true, "Utah", 6.3, "White", "13 yard sack at the 8 (of other team) vs Washington State"})
	Players = append(Players, player{"Demetrius Williamson", "Defensive Tackle", 86, "Akron, Ohio", true, "Penn State", 6.4, "Black", "11 yard sack at the 22 (of his team) vs Michigan"})
	Players = append(Players, player{"Devonta Knox", "Linebacker", 57, "Los Angeles, California", true, "Stanford", 6.2, "Light Black", "Interception at the 46 (of his team) vs California"})
	Players = append(Players, player{"Jaire Hopkins", "Linebacker", 47, "Atlanta, Georgia", true, "Georgia", 6.3, "Black", "ISack made at the 20 (of his team) @ LSU"})
	Players = append(Players, player{"Antonio Alexander", "Defensive Back (Cornerback)", 21, "Harrisberg, Pennsylvania", true, "Boston College", 6.0, "Light Black", "Interception at the 33 (of his team) vs Clemson"})
	Players = append(Players, player{"Mario Hayes", "Defensive Back (Safety)", 34, "Charlotte, North Carolina", true, "Georgia Tech", 5.11, "Dark Black", "Pick Six (Caught at 30 of other team) @ Florida State"})
	Players = append(Players, player{"Kai McFarland", "Placekicker", 95, "Omaha, Nebraska", true, "Nebraska", 6.0, "White", "45 yard field goal good vs Purdue"})
	Players = append(Players, player{"Sam Gonzales", "Placekicker", 91, "Oakland, California", false, "California", 5.11, "White", "41 yard field goal good vs Oregon"})
	Players = append(Players, player{"Grayson Cox", "Punter", 7, "Boise, Idaho", true, "Washington", 6.0, "White", "No game film for OLs and Punters"})
	Players = append(Players, player{"Levi Wheeler", "Punter", 1, "New Orleans, Louisiana", true, "LSU", 6.1, "White", "No game film for OLS and Punters"})
	// end of vars

	Teams = append(Teams, team{"Arizona Cardinals"})
	Teams = append(Teams, team{"Atlanta Falcons"})
	Teams = append(Teams, team{"Baltimore Ravens"})
	Teams = append(Teams, team{"Buffalo Bills"})
	Teams = append(Teams, team{"Carolina Panthers"})
	Teams = append(Teams, team{"Chicago Bears"})
	Teams = append(Teams, team{"Cincinnati Bengals"})
	Teams = append(Teams, team{"Cleveland Browns"})
	Teams = append(Teams, team{"Dallas Cowboys"})
	Teams = append(Teams, team{"Denver Broncos"})
	Teams = append(Teams, team{"Detroit Lions"})
	Teams = append(Teams, team{"Green Bay Packers"})
	Teams = append(Teams, team{"Houston Texans"})
	Teams = append(Teams, team{"Indianapolis Colts"})
	Teams = append(Teams, team{"Jacksonville Jaguars"})
	Teams = append(Teams, team{"Kansas City Chiefs"})
	Teams = append(Teams, team{"Los Angeles Chargers"})
	Teams = append(Teams, team{"Los Angeles Rams"})
	Teams = append(Teams, team{"Miami Dolphins"})
	Teams = append(Teams, team{"Minnesota Vikings"})
	Teams = append(Teams, team{"New England Patriots"})
	Teams = append(Teams, team{"New Orleans Saints"})
	Teams = append(Teams, team{"New York Giants"})
	Teams = append(Teams, team{"New York Jets"})
	Teams = append(Teams, team{"Oakland Raiders"})
	Teams = append(Teams, team{"Philadelphia Eagles"})
	Teams = append(Teams, team{"Pittsburgh Steelers"})
	Teams = append(Teams, team{"San Francisco 49ers"})
	Teams = append(Teams, team{"Seattle Seahawks"})
	Teams = append(Teams, team{"Tampa Bay Buccaneers"})
	Teams = append(Teams, team{"Tennesee Titans"})
	Teams = append(Teams, team{"Washington Redskins"})

	tv = mfr.AddNewChild(gi.KiT_TabView, "tv").(*gi.TabView)

	tab0k, _ := tv.AddNewTab(gi.KiT_Layout, "Team Setup")
	tab0frame = tab0k.(*gi.Layout)
	tab0frame.Lay = gi.LayoutHoriz

	teamSelect()

	tab0frame.SetProp("white-space", "normal")
	heading := tab0frame.AddNewChild(gi.KiT_Label, "heading").(*gi.Label)
	heading.Text = "<b>Pick a team:</b>"
	heading.SetProp("font-size", "x-large")
	for i := 0; i < len(Teams); i++ {
		team := tab0frame.AddNewChild(gi.KiT_Button, fmt.Sprintf("team%v", i)).(*gi.Button)
		team.Text = fmt.Sprintf("<b>%v</b>", Teams[i].Name)
		team.SetProp("font-size", "large")
		team.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			// fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
			if sig == int64(gi.ButtonClicked) {
				sendstuff := send.Name()
				s := strings.Split(sendstuff, "m")
				id := s[1]
				fmt.Printf("id: %v \n", id)
				iid, _ := strconv.Atoi(id)
				TEAM = Teams[iid].Name
				fmt.Printf("Team: %v \n", TEAM)

				gi.StringPromptDialog(vp, "", "Don't enter anything here -- Nothing will be done if you enter anything.",
					gi.DlgOpts{Title: "Front Letter", Prompt: fmt.Sprintf("Congratulations, player on being the new head coach of the %v. After an 0-16 season, the board has highered you to revive this franchise. Click OK to continue and then click on the Draft Tab.", TEAM)},
					rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						//dlg := send.(*gi.Dialog)
						if sig == int64(gi.DialogAccepted) {
							// 	val := gi.StringPromptDialogValue(dlg)
							// 	fmt.Printf("got string value: %v\n", val)
							tv.DeleteTabIndex(0, true)
							playerSetup()
						}
					})
			}
		})

	}

	frontLetter()

	// 	trow = mfr.AddNewChild(gi.KiT_Layout, "trow").(*gi.Layout)
	// 	trow.Lay = gi.LayoutVert
	// 	trow.SetStretchMaxWidth()

	// 	title := trow.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	// 	title.Text = "<b>Game - Play me now!</b>"
	// 	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	// 	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	// 	title.SetProp("vertical-align", gi.AlignCenter)
	// 	title.SetProp("font-family", "Times New Roman, serif")
	// 	title.SetProp("font-size", "x-large")
	// 	// title.SetProp("letter-spacing", 2)
	// 	title.SetProp("line-height", 1.5)
	// 	title.SetStretchMaxWidth()
	// 	title.SetStretchMaxHeight()

	// 	trow.AddNewChild(gi.KiT_Space, "spc1")

	// 	gfr := mfr.AddNewChild(KiT_GameFrame, "gameframe").(*GameFrame)
	// 	gfr.SetProp("background-color", "white")

	// 	gfr.Row = mfr.AddNewChild(gi.KiT_Layout, "brow").(*gi.Layout)
	// 	gfr.Row.Lay = gi.LayoutHoriz

	// 	doSomethingOne := gfr.Row.AddNewChild(gi.KiT_Action, "doSomethingOne").(*gi.Action)
	// 	doSomethingOne.Text = "Do something One"

	// 	doSomethingOne.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {

	// 	})

	// 	doSomethingTwo := gfr.Row.AddNewChild(gi.KiT_Action, "doSomethingTwo").(*gi.Action)
	// 	doSomethingTwo.Text = "Do something Two"

	// 	doSomethingTwo.ActionSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {

	// 	})

	// 	gfr.AddNewChild(gi.KiT_Space, "spc2")

	// 	SvgGame = gfr.AddNewChild(svg.KiT_SVG, "SvgGame").(*svg.SVG)
	// 	SvgGame.SetProp("min-width", GameSize)
	// 	SvgGame.SetProp("min-height", GameSize)
	// 	SvgGame.SetStretchMaxWidth()
	// 	SvgGame.SetStretchMaxHeight()

	// 	SvgPeople = SvgGame.AddNewChild(svg.KiT_Group, "SvgPeople").(*svg.Group)
	// 	SvgMap = SvgGame.AddNewChild(svg.KiT_Group, "SvgMap").(*svg.Group)

	// 	gmin = gi.Vec2D{-10, -10}
	// 	gmax = gi.Vec2D{10, 10}
	// 	gsz = gmax.Sub(gmin)
	// 	ginc = gsz.DivVal(GameSize)

	// 	SvgGame.ViewBox.Min = gmin
	// 	SvgGame.ViewBox.Size = gsz
	// 	SvgGame.Norm = true
	// 	SvgGame.InvertY = true
	// 	SvgGame.Fill = true
	// 	SvgGame.SetProp("background-color", "white")
	// 	SvgGame.SetProp("stroke-width", ".8pct")

	// 	InitMap()
	// 	InitPlayer()

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

	SvgGame.UpdateEnd(updt)

}

// var player *svg.Rect

func InitPlayer() {
	// 	updt := SvgGame.UpdateStart()
	// 	SvgPeople.DeleteChildren(true)

	// 	player = SvgPeople.AddNewChild(svg.KiT_Rect, "player").(*svg.Rect)

	// 	player.SetProp("fill", "red")
	// 	player.SetProp("stroke", "darkred")
	// 	player.Size = gi.Vec2D{2, 2}
	// 	player.Pos = gi.Vec2D{-5, -10}

	// 	SvgGame.UpdateEnd(updt)

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

	for i := 0; i > -1; i++ {

		updt := SvgGame.UpdateStart()

		SvgGame.UpdateEnd(updt)

	}
}

func frontLetter() {

}

func teamSelect() {

}

var draftRound = 1
var Player1 string

func playerSetup() {

	tab1k, _ := tv.AddNewTab(gi.KiT_Layout, "Draft")

	updt := tab1k.UpdateStart()
	defer tab1k.UpdateEnd(updt)
	rec := ki.Node{}          // receiver for events
	rec.InitName(&rec, "rec") // this is essential for root objects not owned by other Ki tree nodes

	bigFrame := tab1k.(*gi.Layout)
	bigFrame.Lay = gi.LayoutVert

	draftHeader := bigFrame.AddNewChild(gi.KiT_Label, "draftHeader").(*gi.Label)
	draftHeader.Text = "<b>Welcome to the first round of the draft! Pick wisely.                      </b>"
	draftHeader.Redrawable = true
	draftHeader.SetProp("font-size", "x-large")

	draftFrame := bigFrame.AddNewChild(gi.KiT_Layout, "draftFrame").(*gi.Layout)
	draftFrame.Lay = gi.LayoutHoriz

	for i := 0; i < len(Players); i++ {
		playerFrame := draftFrame.AddNewChild(gi.KiT_Layout, fmt.Sprintf("playerFrame%v", i)).(*gi.Layout)
		playerFrame.Lay = gi.LayoutVert

		playerName := playerFrame.AddNewChild(gi.KiT_Button, fmt.Sprintf("playerName%v", i)).(*gi.Button)
		playerName.Text = fmt.Sprintf("<b>%v</b>", Players[i].Name)
		// playerName.SetProp("label", ki.Props{"font-size": "x-large"})

		property1 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property1%v", i)).(*gi.Label)
		property2 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property2%v", i)).(*gi.Label)
		property3 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property3%v", i)).(*gi.Label)
		property4 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property4%v", i)).(*gi.Label)
		property5 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property5%v", i)).(*gi.Label)
		property6 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property6%v", i)).(*gi.Label)
		property7 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property7%v", i)).(*gi.Label)
		property8 := playerFrame.AddNewChild(gi.KiT_Label, fmt.Sprintf("property8%v", i)).(*gi.Label)

		property1.Text = fmt.Sprintf("<b>Position</b>: %v", Players[i].Position)
		property2.Text = fmt.Sprintf("<b>Player Number</b>: %v", Players[i].Number)
		property3.Text = fmt.Sprintf("<b>Hometown</b>: %v", Players[i].Hometown)
		if Players[i].Right {
			property4.Text = fmt.Sprintf("<b>Right Handed?</b> Yes")
		} else {
			property4.Text = fmt.Sprintf("<b>Right Handed?</b> No")

		}
		property5.Text = fmt.Sprintf("<b>College</b>: %v", Players[i].College)
		property6.Text = fmt.Sprintf("<b>Height</b>: %v", Players[i].Height)
		property7.Text = fmt.Sprintf("<b>Skin Color</b>: %v", Players[i].Color)
		property8.Text = fmt.Sprintf("<b>Game Film</b>: %v", Players[i].GameFilm)
		property8.SetProp("width", "50ch")
		property8.SetProp("white-space", "normal")
		// prop5text := playerFrame.KnownChild(5).(*gi.Label)
		// fmt.Printf("Property 5: %v \n", )

		playerName.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
			fmt.Printf("Received button signal: %v from button: %v\n", gi.ButtonSignals(sig), send.Name())
			if sig == int64(gi.ButtonClicked) {

				fmt.Printf("Player Name: %v \n", playerName.Text)
				splits := strings.Split(playerName.Text, "<b>")
				player1 := splits[1]

				player2 := strings.Split(player1, "</b>")
				player3 := player2[0]
				fmt.Printf("Player3: %v \n", player3)

				Player1 = player3

				gi.StringPromptDialog(vp, "", "Don't enter anything here -- Nothing will be done if you enter anything.",
					gi.DlgOpts{Title: "Confirmation", Prompt: fmt.Sprintf("Are you sure you want to select %v and move on to the second round of the draft?", Player1)},
					rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
						//dlg := send.(*gi.Dialog)
						if sig == int64(gi.DialogAccepted) {
							draftRound = 2
							draftHeader.SetText("<b>Welcome to the second round of the draft! Pick wisely.</b>")
						}
					})

			}
		})

	}
}
