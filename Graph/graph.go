package main

import (

	//"github.com/goki/gi/giv"
	//"github.com/goki/gi/complete"

	"fmt"
	"log"
	"math"

	"github.com/Knetic/govaluate"
	"github.com/chewxy/math32"
	"github.com/goki/gi"
	"github.com/goki/gi/gimain"
	"github.com/goki/gi/oswin"
	"github.com/goki/gi/svg"
	"github.com/goki/gi/units"
	"github.com/goki/ki"
)

var graph *svg.SVG
var gmin, gmax, gsz, ginc gi.Vec2D
var sz float32

var functions = map[string]govaluate.ExpressionFunction{
	"cos": func(args ...interface{}) (interface{}, error) {
		y := math.Cos(args[0].(float64))
		return y, nil
	},
	"sin": func(args ...interface{}) (interface{}, error) {
		y := math.Sin(args[0].(float64))
		return y, nil
	},
	"tan": func(args ...interface{}) (interface{}, error) {
		y := math.Tan(args[0].(float64))
		return y, nil
	},
	"pow": func(args ...interface{}) (interface{}, error) {
		y := math.Pow(args[0].(float64), args[1].(float64))
		return y, nil
	},
	"abs": func(args ...interface{}) (interface{}, error) {
		y := math.Abs(args[0].(float64))
		return y, nil
	},

	"fact": func(args ...interface{}) (interface{}, error) {

		y := FactorialMemoization(int(args[0].(float64)))

		return y, nil
	},
}

var lineNo = 0
var colors = []string{"black", "red", "blue", "green", "purple", "brown", "orange"}

func main() {

	gimain.Main(func() {
		mainrun()
	})

}

type Marble struct {
	Pos gi.Vec2D
	Vel gi.Vec2D
}

var Marbles []*Marble

var NMarbles = 10

var unitOfTime = 10 // in terms of ms

var StartSpeed = 0.2 // Coordinates per unit of time

var vp *gi.Viewport2D

var Stop = false

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

	oswin.TheApp.SetName("Graphing")
	oswin.TheApp.SetAbout("Graphing is an app that will allow you to enter equations and have them be graphed. There will also be other modes where you can have marbles fall or things like that.")

	win := gi.NewWindow2D("graph", "Graphing App", width, height, true) // true = pixel sizes

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

	title := mfr.AddNewChild(gi.KiT_Label, "title").(*gi.Label)
	title.Text = "<b>Graphing</b>"
	title.SetProp("white-space", gi.WhiteSpaceNormal) // wrap
	title.SetProp("text-align", gi.AlignCenter)       // note: this also sets horizontal-align, which controls the "box" that the text is rendered in..
	title.SetProp("vertical-align", gi.AlignCenter)
	title.SetProp("font-family", "Times New Roman, serif")
	title.SetProp("font-size", "x-large")
	// title.SetProp("letter-spacing", 2)
	// title.SetStretchMaxWidth()

	graphingInput := mfr.AddNewChild(gi.KiT_TextField, "graphingInput").(*gi.TextField)
	graphingInput.Placeholder = "Enter your equation"
	graphingInput.SetProp("min-width", "100ch")
	graphingInput.TextFieldSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.TextFieldDone) {
			updt := vp.UpdateStart()

			//graphResult := mfr.AddNewChild(gi.KiT_Label, "graphResult").(*gi.Label)
			//graphResult.Text = fmt.Sprintf("Your equation is: %v", graphingInput.Text())

			Graph(graphingInput.Text())

			vp.UpdateEnd(updt)
		}
	})

	brow := mfr.AddNewChild(gi.KiT_Layout, "brow").(*gi.Layout)
	brow.Lay = gi.LayoutHoriz
	brow.SetProp("spacing", units.NewValue(2, units.Ex))

	brow.SetProp("horizontal-align", gi.AlignLeft)
	// brow.SetProp("horizontal-align", gi.AlignJustify)
	brow.SetStretchMaxWidth()

	submitGraphingInput := brow.AddNewChild(gi.KiT_Button, "submitGraphingInput").(*gi.Button)
	submitGraphingInput.Text = "Graph equation"
	submitGraphingInput.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			updt := vp.UpdateStart()

			//graphResult := mfr.AddNewChild(gi.KiT_Label, "graphResult").(*gi.Label)
			//graphResult.Text = fmt.Sprintf("Your equation is: %v", graphingInput.Text())

			Graph(graphingInput.Text())

			vp.UpdateEnd(updt)
		}
	})

	resetGraphButton := brow.AddNewChild(gi.KiT_Button, "resetGraphButton").(*gi.Button)
	resetGraphButton.Text = "Reset graph"
	resetGraphButton.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			updt := vp.UpdateStart()

			InitGraph()

			vp.UpdateEnd(updt)
		}
	})

	launchMarbles := brow.AddNewChild(gi.KiT_Button, "launchMarbles").(*gi.Button)
	launchMarbles.Text = "Run!"

	launchMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			go RunMarbles()
		}
	})

	stopMarbles := brow.AddNewChild(gi.KiT_Button, "stopMarbles").(*gi.Button)
	stopMarbles.Text = "Stop"
	stopMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			Stop = true
		}
	})

	frame := mfr.AddNewChild(gi.KiT_Frame, "frame").(*gi.Frame)

	graph = frame.AddNewChild(svg.KiT_SVG, "graph").(*svg.SVG)
	sz = float32(800)
	graph.SetProp("min-width", sz)
	graph.SetProp("min-height", sz)
	graph.SetStretchMaxWidth()
	graph.SetStretchMaxHeight()

	gmin = gi.Vec2D{-10, -10}
	gmax = gi.Vec2D{10, 10}
	gsz = gmax.Sub(gmin)
	ginc = gsz.DivVal(sz)

	graph.ViewBox.Min = gmin
	graph.ViewBox.Size = gsz
	graph.Norm = true
	graph.Fill = true
	graph.SetProp("background-color", "white")
	graph.SetProp("stroke-width", ".2pct")

	InitGraph()

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

}

var Lines []*govaluate.EvaluableExpression

func Graph(exstr string) {
	path1 := graph.AddNewChild(svg.KiT_Path, "path1").(*svg.Path)
	path1.SetProp("fill", "none")
	clr := colors[lineNo%len(colors)]
	path1.SetProp("stroke", clr)

	expr, err := govaluate.NewEvaluableExpressionWithFunctions(exstr, functions)
	Lines = append(Lines, expr)

	if err != nil {
		log.Println(err)
		return
	}

	params := make(map[string]interface{}, 8)
	params["x"] = float64(0)
	ps := ""
	start := true
	for x := gmin.X; x < gmax.X; x += ginc.X {
		params["x"] = x
		yi, _ := expr.Evaluate(params)
		y := float32(-yi.(float64))
		if start {
			ps += fmt.Sprintf("M %v %v ", x, y)

			start = false
		} else {
			ps += fmt.Sprintf("L %v %v ", x, y)
		}
	}
	path1.SetData(ps)
	lineNo++
}

var SvgMarbles *svg.Group
var MarbleRadius = .1

func InitGraph() {

	graph.DeleteChildren(true)

	Stop = false

	Lines = make([]*govaluate.EvaluableExpression, 0)

	xAxis := graph.AddNewChild(svg.KiT_Line, "xAxis").(*svg.Line)
	xAxis.Start = gi.Vec2D{-10, 0}
	xAxis.End = gi.Vec2D{10, 0}
	xAxis.SetProp("stroke", "#888")

	yAxis := graph.AddNewChild(svg.KiT_Line, "yAxis").(*svg.Line)
	yAxis.Start = gi.Vec2D{0, -10}
	yAxis.End = gi.Vec2D{0, 10}
	yAxis.SetProp("stroke", "#888")

	SvgMarbles = graph.AddNewChild(svg.KiT_Group, "SvgMarbles").(*svg.Group)

	InitMarbles()

	for _, m := range Marbles {

		circle := SvgMarbles.AddNewChild(svg.KiT_Circle, "circle").(*svg.Circle)
		circle.SetProp("stroke", "black")
		circle.SetProp("fill", "purple")
		circle.Radius = float32(MarbleRadius)
		circle.Pos = m.Pos
		circle.Pos.Y = -circle.Pos.Y

	}

}
func UpdateMarbles() {
	updt := vp.UpdateStart()
	params := make(map[string]interface{}, 8)
	params["x"] = float64(0)
	for i, m := range Marbles {

		m.Pos = m.Pos.Add(m.Vel)

		for _, ln := range Lines {

			params["x"] = m.Pos.X
			yi, _ := ln.Evaluate(params)
			y := float32(yi.(float64))

			if m.Pos.Y < y {
				params["x"] = m.Pos.X - .01
				yi, _ = ln.Evaluate(params)
				yl := float32(yi.(float64))

				params["x"] = m.Pos.X + .01
				yi, _ = ln.Evaluate(params)
				yr := float32(yi.(float64))

				//slp := (yr - yl) / .02
				ang := math32.Atan2(yl-yr, 0.02)
				ang *= 2

				nvx := m.Vel.X*math32.Cos(ang) - m.Vel.Y*math32.Sin(ang)
				nvy := m.Vel.Y*math32.Sin(ang) + m.Vel.X*math32.Cos(ang)

				m.Vel = gi.Vec2D{nvx, nvy}

				m.Pos.Y = y
				//m.Vel.Y = -m.Vel.Y
			}

		}

		circle := SvgMarbles.KnownChild(i).(*svg.Circle)
		circle.Pos = m.Pos
		circle.Pos.Y = -circle.Pos.Y

	}
	vp.UpdateEnd(updt)
}

func InitMarbles() {
	Marbles = make([]*Marble, 0)
	for n := 0; n < NMarbles; n++ {

		diff := float32(n) / 2

		m := Marble{Pos: gi.Vec2D{0, 10 - diff}, Vel: gi.Vec2D{0, float32(-StartSpeed)}}

		Marbles = append(Marbles, &m)

	}

}

func RunMarbles() {

	for i := 0; i < 1000; i++ {
		//fmt.Printf("Update: %v \n", i)
		UpdateMarbles()
		//time.Sleep(time.Duration(unitOfTime) * time.Millisecond)

		if Stop {
			break
		}

	}
}

const LIM = 100

var facts [LIM]float64

func FactorialMemoization(n int) (res float64) {

	if n < 0 {
		return 1
	}

	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = float64(n) * FactorialMemoization(n-1)
		return res
	}

	return 1
}
