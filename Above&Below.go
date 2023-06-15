package main

import (
	"image"
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

const windowWidth = 800
const windowHeight = 600
const pixelSize = 10
const bufferWidth = windowWidth / pixelSize
const bufferHeight = windowHeight / pixelSize

// Game world
var worldCurrent [bufferWidth][bufferHeight]int
var worldNext [bufferWidth][bufferHeight]int

func main() {
	pixelgl.Run(run)
}
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Above&Below",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	img := image.NewRGBA(image.Rect(0, 0, bufferWidth, bufferHeight))
	for !win.Closed() {

		pic := pixel.PictureDataFromImage(img)
		sprite := pixel.NewSprite(pic, pic.Bounds())
		mat := pixel.IM.Scaled(pixel.ZV, float64(pixelSize)).Moved(win.Bounds().Center())
		sprite.Draw(win, mat)
		win.Update()
	}
}

func updateWorld() {


}

func drawWorld(img *image.RGBA, worldCurrent [bufferWidth][bufferHeight]int) {
	for x := 0; x < bufferWidth; x++ {
		for y := 0; y < bufferHeight; y++ {
			col := worldCurrent[x][y].Draw()
			img.Set(x, y, col)
		}

	}
}

// Update Logic
type Cell struct {
	X,Y int
}

func getActiveCells(worldCurrent [bufferWidth][bufferHeight]Element) []Cell {
	var activeCells []Cell
	for x := 0; x < bufferWidth; x++ {
		for y := 0; y < bufferHeight; y++ {
			if worldCurrent[x][y] != nil {
				activeCells = append(activeCells, Cell{X: x, Y: y})
			}
		}
	}
	return activeCells
}

type Element interface{
	Update()
	Draw() color.Color
}

type Solid struct {
	Index int
	Weight float64
	Gravity float64
	Color color.RGBA
	flammable bool
}

func (s *Solid) Draw() color.Color{
	return color.RGBA{128, 128, 128, 255}
}

type Wall struct {
	Index int
	Weight float64
	Gravity float64
	Color color.RGBA
	flammable bool
}

func (w *Wall) Draw() color.Color{
	return color.RGBA{255, 255, 255, 255}
}

type Liquid struct {
	Index int
	Weight float64
	Gravity float64
	Color color.RGBA
	flammable bool
}

func (l *Liquid) Draw() color.Color{
	return color.RGBA{0, 0, 255, 255}
}

func (l *Liquid) Update(x, y int, worldCurrent, worldNext *[bufferWidth][bufferHeight]Element) {
	if y > 0 && worldCurrent[x][y-1]!= nil {
		worldNext[x][y-1] = *worldCurrent[x][y]
		worldNext[x][y] = nil
	}else{
		worldNext[x][y] = worldCurrent[x][y]
	}
}

type Gas struct {
	Index int
	Weight float64
	Gravity float64
	Color color.RGBA
	flammable bool
}
func (g *Gas) Draw() color.Color{
	return color.RGBA{255, 255, 255, 255}
}


//Elements
var Titanium = &Wall{
	Index: 22,
	Weight: 1.0,
    Gravity: 0.0,
    Color: color.RGBA{255, 255, 255, 255},
    flammable: true,
}
	