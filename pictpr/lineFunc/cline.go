package lineFunc

import (
	"fmt"
	"os"
	"image"
	"image/png"
	"image/color"
)

type CLine struct{
	Width int
	Px [][]uint8
}

func (this *CLine) ShowSize(){
    fmt.Printf("%d\n", this.Width)
}

func (this *CLine) Extract(p Pict, y int){
	this.Width = p.Width

	this.Px = make([][]uint8, this.Width)
	for x := 0; x < this.Width; x++ {
		this.Px[x] = make([]uint8, 4)
	}

	for x := 0; x < this.Width; x++ {
		for i := 0; i < 4; i++ {
			this.Px[x][i] = p.Px[x][y][i]
		}
	}
}

func MkCLine(w int) Pict{
	var cl CLine
	cl.Width = w

	cl.Px = make([][]uint8, w)
	for x := 0; x < w; x++ {
		cl.Px[x] = make([]uint8, 4)
	}	

	for x := 0; x < w; x++ {
		for i := 0; i < 4; i++ {
			cl.Px[x][i] = 0
		}
	}

	return cl
}
