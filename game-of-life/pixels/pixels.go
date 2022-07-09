package pixels

import "image/color"

type Pixels struct {
	Pix   []uint8
	Width int
}

func NewPixels(width, height int) *Pixels {
	return &Pixels{
		Pix:   make([]uint8, width*height*4), // each cell can have 4 different values
		Width: width,
	}
}

func (p *Pixels) SetColor(x, y int, rgba color.RGBA) {
	r, g, b, a := rgba.RGBA()
	index := (y*p.Width + x) * 4
	p.Pix[index] = uint8(r)
	p.Pix[index+1] = uint8(g)
	p.Pix[index+2] = uint8(b)
	p.Pix[index+3] = uint8(a)
	// pos := 4 * (x + y*p.Width)
	// copy(p.Pix[pos:pos+4], []uint8{uint8(r), uint8(g), uint8(b), uint8(a)})
}

func (p *Pixels) DrawRectangle(x, y, width, height int, rgba color.RGBA) {
	for idx := 0; idx < width; idx++ {
		for idy := 0; idy < height; idy++ {
			p.SetColor(x+idx, y+idy, rgba)
		}
	}
}
