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
var NumOfEquations = 1

func main() {

	gimain.Main(func() {
		mainrun()
	})

}

var irow *gi.Layout

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

	irow = mfr.AddNewChild(gi.KiT_Layout, "irow").(*gi.Layout)
	irow.Lay = gi.LayoutVert

	graphingInput := irow.AddNewChild(gi.KiT_TextField, "newEquation0").(*gi.TextField)
	graphingInput.Placeholder = "Enter your equation"
	graphingInput.SetProp("min-width", "100ch")
	// graphingInput.TextFieldSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
	// 	if sig == int64(gi.TextFieldDone) {
	// 		Graph(graphingInput.Text())
	// 	}
	// })

	brow := mfr.AddNewChild(gi.KiT_Layout, "brow").(*gi.Layout)
	brow.Lay = gi.LayoutHoriz
	brow.SetProp("spacing", units.NewValue(2, units.Ex))

	brow.SetProp("horizontal-align", gi.AlignLeft)
	// brow.SetProp("horizontal-align", gi.AlignJustify)
	brow.SetStretchMaxWidth()

	addNewEquation := brow.AddNewChild(gi.KiT_Button, "addNewEquation").(*gi.Button)
	addNewEquation.Text = "Add new equation"
	addNewEquation.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			updt := vp.UpdateStart()

			newEquation := irow.AddNewChild(gi.KiT_TextField, fmt.Sprintf("newEquation%v", NumOfEquations)).(*gi.TextField)
			newEquation.Placeholder = "Enter your equation"
			newEquation.SetProp("min-width", "100ch")

			NumOfEquations++

			vp.UpdateEnd(updt)

		}
	})

	submitGraphingInput := brow.AddNewChild(gi.KiT_Button, "submitGraphingInput").(*gi.Button)
	submitGraphingInput.Text = "Graph equation"
	submitGraphingInput.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {

			GraphLoop()
			// Graph(graphingInput.Text())
		}
	})

	resetGraphButton := brow.AddNewChild(gi.KiT_Button, "resetGraphButton").(*gi.Button)
	resetGraphButton.Text = "Reset graph"
	resetGraphButton.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			InitGraph()
		}
	})

	runMarbles := brow.AddNewChild(gi.KiT_Button, "runMarbles").(*gi.Button)
	runMarbles.Text = "Run!"
	runMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			go RunMarbles()
		}
	})

	stepMarbles := brow.AddNewChild(gi.KiT_Button, "stepMarbles").(*gi.Button)
	stepMarbles.Text = "Step"
	stepMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			go UpdateMarbles()
		}
	})

	stopMarbles := brow.AddNewChild(gi.KiT_Button, "stopMarbles").(*gi.Button)
	stopMarbles.Text = "Stop"
	stopMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			Stop = true
			// fmt.Printf("Stop is: %v \n", Stop)
		}
	})

	resetMarbles := brow.AddNewChild(gi.KiT_Button, "resetMarbles").(*gi.Button)
	resetMarbles.Text = "Reset Marbles"
	resetMarbles.ButtonSig.Connect(rec.This(), func(recv, send ki.Ki, sig int64, data interface{}) {
		if sig == int64(gi.ButtonClicked) {
			ResetMarbles()
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

type Marble struct {
	Pos    gi.Vec2D
	Vel    gi.Vec2D
	PrvPos gi.Vec2D
}

func (mb *Marble) Init(diff float32) {
	mb.Pos = gi.Vec2D{0, 10 - diff}
	mb.Vel = gi.Vec2D{0, float32(-StartSpeed)}
	mb.PrvPos = mb.Pos
}

var Marbles []*Marble

// Kai: put all these in a struct, and add a StructInlineView to edit them.
// see your other code for how to do it..

var NMarbles = 10

var NSteps = 1000 // number of steps to take when running -- Stop works now..

var StartSpeed = float32(0) // Coordinates per unit of time

var UpdtRate = float32(.2) // how fast to move along velocity vector -- lower = smoother, more slow-mo

var vp *gi.Viewport2D

var Stop = false

var Lines []*govaluate.EvaluableExpression
var SvgLines *svg.Group

var GraphedLine = false

func GraphLoop() {

	updt := graph.UpdateStart()
	// fmt.Printf("Svg Lines:%vend\n", SvgLines)

	// fmt.Printf("\n Hi :%v:\n", GraphedLine)

	if GraphedLine {

		fmt.Printf("Deleting children \n")

		SvgLines.DeleteChildren(true)

		Lines = nil

	}
	SvgLines = graph.AddNewChild(svg.KiT_Group, "SvgLines").(*svg.Group)

	for i := 0; i < NumOfEquations; i++ {
		equation := irow.KnownChild(i).(*gi.TextField)
		text := equation.Text()
		Graph(text)
		GraphedLine = true
	}
	graph.UpdateEnd(updt)
}

func Graph(exstr string) {
	updt := graph.UpdateStart()

	path1 := SvgLines.AddNewChild(svg.KiT_Path, "path1").(*svg.Path)
	path1.SetProp("fill", "none")
	clr := colors[lineNo%len(colors)]
	path1.SetProp("stroke", clr)

	expr, err := govaluate.NewEvaluableExpressionWithFunctions(exstr, functions)

	if err != nil {
		log.Println(err)
		return
	}
	Lines = append(Lines, expr)

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
	graph.UpdateEnd(updt)
}

var SvgMarbles *svg.Group

var MarbleRadius = .1

func InitGraph() {
	updt := graph.UpdateStart()

	lineNo = 0
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
	graph.UpdateEnd(updt)
}

func RadToDeg(rad float32) float32 {
	return rad * 180 / math.Pi
}

var Friction = float32(0.75)
var Gravity = float32(0.01)

func UpdateMarbles() {
	updt := graph.UpdateStart()
	params := make(map[string]interface{}, 8)
	params["x"] = float64(0)
	for i, m := range Marbles {

		m.Vel.Y -= Gravity
		npos := m.Pos.Add(m.Vel.MulVal(UpdtRate))
		ppos := m.PrvPos

		for _, ln := range Lines {
			if ln == nil {
				continue
			}

			params["x"] = m.Pos.X
			yi, _ := ln.Evaluate(params)
			y := float32(yi.(float64))

			params["x"] = m.PrvPos.X
			yi, _ = ln.Evaluate(params)
			yp := float32(yi.(float64))

			// fmt.Printf("y: %v npos: %v pos: %v\n", y, npos.Y, m.Pos.Y)

			if (npos.Y < y && m.PrvPos.Y >= yp) || (npos.Y > y && m.PrvPos.Y <= yp) {
				//				fmt.Printf("crossed!!\n")

				params["x"] = m.Pos.X - .01
				yi, _ = ln.Evaluate(params)
				yl := float32(yi.(float64))

				params["x"] = m.Pos.X + .01
				yi, _ = ln.Evaluate(params)
				yr := float32(yi.(float64))

				//slp := (yr - yl) / .02
				angLn := math32.Atan2(yr-yl, 0.02)
				angN := angLn + math.Pi/2 // + 90 deg

				angI := math32.Atan2(m.Vel.Y, m.Vel.X)
				angII := angI + math.Pi

				angNII := angN - angII
				angR := math.Pi + 2*angNII

				// fmt.Printf("angLn: %v  angN: %v  angI: %v  angII: %v  angNII: %v  angR: %v\n",
				// 	RadToDeg(angLn), RadToDeg(angN), RadToDeg(angI), RadToDeg(angII), RadToDeg(angNII), RadToDeg(angR))

				nvx := Friction * (m.Vel.X*math32.Cos(angR) - m.Vel.Y*math32.Sin(angR))
				nvy := Friction * (m.Vel.X*math32.Sin(angR) + m.Vel.Y*math32.Cos(angR))

				m.Vel = gi.Vec2D{nvx, nvy}

				m.Pos.Y = y
				//m.Vel.Y = -m.Vel.Y
				break
			}
		}

		m.PrvPos = ppos
		m.Pos = m.Pos.Add(m.Vel.MulVal(UpdtRate))

		circle := SvgMarbles.KnownChild(i).(*svg.Circle)
		circle.Pos = m.Pos
		circle.Pos.Y = -circle.Pos.Y

	}
	graph.UpdateEnd(updt)
}

func ResetMarbles() {
	updt := graph.UpdateStart()

	for i, m := range Marbles {
		diff := float32(i) / 2
		m.Init(diff)
		circle := SvgMarbles.KnownChild(i).(*svg.Circle)
		circle.Pos = m.Pos
		circle.Pos.Y = -circle.Pos.Y
	}
	graph.UpdateEnd(updt)
}

func InitMarbles() {
	Marbles = make([]*Marble, 0)
	for n := 0; n < NMarbles; n++ {

		diff := float32(n) / 2

		m := Marble{}
		m.Init(diff)

		Marbles = append(Marbles, &m)
	}
}

func RunMarbles() {
	Stop = false
	for i := 0; i < NSteps; i++ {
		//fmt.Printf("Update: %v \n", i)
		UpdateMarbles()
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
