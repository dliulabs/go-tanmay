package gameoflife

import (
	"game-of-life/pixels"
	"image/color"
	"math/rand"
)

var (
	White = color.RGBA{0, 0, 0, 255}
	Black = color.RGBA{255, 255, 255, 255}
)

type GameOfLife struct {
	Gameboard [][]int
	Pixels    *pixels.Pixels
	Size      int // size of each pixel
}

func NewGameOfLife(width, height, size int) *GameOfLife {
	gameboard := make([][]int, height)
	for row := range gameboard {
		gameboard[row] = make([]int, width)
	}
	pixels := pixels.NewPixels(width*size, height*size)
	return &GameOfLife{
		Gameboard: gameboard,
		Pixels:    pixels,
		Size:      size,
	}
}

func (gol *GameOfLife) Random() {
	for idy := range gol.Gameboard {
		for idx := range gol.Gameboard[idy] {
			gol.Gameboard[idy][idx] = rand.Intn(2) // either 0 or 1
		}

	}
}

// count number of neighbors that are still alive
/*
	a := [2][]uint8{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
	}
	fmt.Println(len(a)) // len 2
*/
func (gol *GameOfLife) CountNeighbors() [][]int {
	matrix := gol.Gameboard
	neighbors := make([][]int, len(matrix))
	for idx, value := range matrix {
		neighbors[idx] = make([]int, len(value))
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			for rowMod := -1; rowMod < 2; rowMod++ {
				newRow := row + rowMod
				if newRow < 0 || newRow >= len(matrix) {
					continue
				}
				for colMod := -1; colMod < 2; colMod++ {
					newCol := col + colMod
					if newCol < 0 || newCol >= len(matrix[row]) {
						continue
					}
					neighbors[row][col] += matrix[newRow][newCol]
				}
			}
		}
	}
	return neighbors
}

func (gol *GameOfLife) PlayRound() {
	neighbors := gol.CountNeighbors()
	for idy := range gol.Gameboard {
		for idx, value := range gol.Gameboard[idy] {
			n := neighbors[idy][idx]
			if value == 1 && (n == 2 || n == 3) {
				continue
			} else if n == 3 {
				gol.Gameboard[idy][idx] = 1
				gol.Pixels.DrawRectangle(idx*gol.Size, idy*gol.Size, gol.Size, gol.Size, Black)
			} else {
				gol.Gameboard[idy][idx] = 0
				gol.Pixels.DrawRectangle(idx*gol.Size, idy*gol.Size, gol.Size, gol.Size, White)
			}
		}

	}
}
