package lineFunc

import (
	"fmt"
	"os"
	"image"
	"image/png"
	"image/color"
)

type GLine struct{
	Width int
	Px []uint8
}

func (this *GLine) ShowSize(){
    fmt.Printf("%d\n", this.Width)
}

func MkCLine(w int) Pict{
	var gl GLine
	gl.Width = w

	gl.Px = make([]uint8, w)

	for x := 0; x < w; x++ {
		gl.Px[x] = 0
	}

	return gl
}
