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
			col := color.RGBA
			col = worldCurrent[x][y].DrawColor()
			img.Set(x, y, col)
		}

	}
}
