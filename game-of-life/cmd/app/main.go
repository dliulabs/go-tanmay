package main

import (
	"game-of-life/gameoflife"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func Run() {
	size := float64(20)
	width := float64(400)
	height := float64(400)
	cfg := pixelgl.WindowConfig{
		Title:  "Conway's Game of Life",
		Bounds: pixel.R(0, 0, width*size, height*size),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	gol := gameoflife.NewGameOfLife(int(width), int(height), int(size))
	gol.Random()
	for !win.Closed() {
		gol.PlayRound()
		win.Canvas().SetPixels(gol.Pixels.Pix)
		win.Update()
	}
}

func main() {
	pixelgl.Run(Run)
}
