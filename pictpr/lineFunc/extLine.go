package lineFunc

func UnderCut(cl CLine) GLine {
	gl := MkGLine(CLine.Width)

	for x := 0; x < cl.Width; x++ {
		cv := cl.Px[x][1]
		if cl.Px[x][2] < cv {
			cv = cl.Px[x][2]
		}
		if cl.Px[x][3] < cv {
			cv = cl.Px[x][3]
		}

		gl[x] = cv
	}

	return gl
}

func Mul(a, b Pict) Pict {
	res := MkPict(a.Width, a.Height)

	for x := 0; x < a.Width; x++ {
		for y := 0; y < a.Height; y++ {
			res.Px[x][y][0] = 255

			for i := 1; i < 4; i++ {
					res.Px[x][y][i] = (uint8)(
						255.0 * 
						(float32)(a.Px[x][y][i]) * 
						(float32)(b.Px[x][y][i]) / 
						(255.0 * 255.0))
			}
		}
	}

	return res
}
