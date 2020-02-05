package pictFunc

//import "fmt"

func Ave (inp Pict) []uint8 {
	res := make([]uint8, 4)
	cache := make([]uint64, 4)

	for c:= 0; c < 4; c++ {
		cache[c] = 0;
	}

	for x := 0; x < inp.Width; x++ {
		for y := 0; y < inp.Height; y++ {
			for c:= 1; c < 4; c++ {
				cache[c] += uint64(inp.Px[x][y][c])
			}
		}
	}

	pnum := inp.Width * inp.Height

	for c:= 0; c < 4; c++ {
		res[c] = uint8(cache[c] / uint64(pnum))
	}

	return res
}

func Sdif (bef, aft []uint8) []uint8 {
	res := make([]uint8, 4)

	for c := 0; c < 4; c++ {
		if bef[c] < aft[c] {
			res[c] = uint8(aft[c] - bef[c])
		} else {
			res[c] = uint8(bef[c] - aft[c])
		}
	}

	return res
}

func Dif (bef, aft Pict) Pict {
	res := MkPict(bef.Width, bef.Height)

	for x := 0; x < bef.Width; x++ {
		for y := 0; y < bef.Height; y++ {
			for c := 1; c < 4; c++ {
				bc := int(bef.Px[x][y][c])
				ac := int(aft.Px[x][y][c])

				if bc < ac {
					res.Px[x][y][c] = uint8(ac - bc)
				} else {
					res.Px[x][y][c] = uint8(bc - ac)
				}
			}

			res.Px[x][y][0] = uint8(255)
		}
	}

	return res
}

func Difv (bef, aft Pict) uint64 {
	res := uint64(0)

	for x := 0; x < bef.Width; x++ {
		for y := 0; y < bef.Height; y++ {
			for c := 1; c < 4; c++ {
				bc := uint64(bef.Px[x][y][c])
				ac := uint64(aft.Px[x][y][c])

				if bc < ac {
					res += uint64(ac - bc)
				} else {
					res += uint64(bc - ac)
				}
			}
		}
	}

	return res
}

func Difc (bef, aft Pict, hol int) uint64 {
	res := uint64(0)

	for x := 0; x < bef.Width; x++ {
		for y := 0; y < bef.Height; y++ {
			pv := 0

			for c := 1; c < 4; c++ {
				bc := uint64(bef.Px[x][y][c])
				ac := uint64(aft.Px[x][y][c])

				if bc < ac {
					pv += int(ac - bc)
				} else {
					pv += int(bc - ac)
				}
			}

			if hol < pv {
				res ++
			}
		}
	}

	return res
}

func MozDiv (inp Pict, div int) Pict {
	res := MkPict(inp.Width / div , inp.Height / div)
	divsq := div * div

	for x := 0; x < inp.Width / div; x++ {
		for y := 0; y < inp.Height / div; y++ {

			box := make([]int, 4)
			for c := 0; c < 4; c++ {
				box[c] = 0
			}

			for w := 0; w < div; w++ {
				for h := 0; h < div; h++ {
					for c := 1; c < 4; c++ {
						box[c] += int(inp.Px[x * div + w][y * div + h][c])
					}
				}
			}

			res.Px[x][y][0] = uint8(255)
			res.Px[x][y][1] = uint8(box[1] / divsq)
			res.Px[x][y][2] = uint8(box[2] / divsq)
			res.Px[x][y][3] = uint8(box[3] / divsq)
		}
	}

	return res
}

func MozDivScaled (inp Pict, div int) Pict {
	res := MkPict(inp.Width, inp.Height)
	divsq := div * div

	for x := 0; x < inp.Width / div; x++ {
		for y := 0; y < inp.Height / div; y++ {

			box := make([]int, 4)
			for c := 0; c < 4; c++ {
				box[c] = 0
			}

			for w := 0; w < div; w++ {
				for h := 0; h < div; h++ {
					for c := 1; c < 4; c++ {
						box[c] += int(inp.Px[x * div + w][y * div + h][c])
					}
				}
			}

			for w := 0; w < div; w++ {
				for h := 0; h < div; h++ {
					res.Px[x * div + w][y * div + h][0] = uint8(255)
					res.Px[x * div + w][y * div + h][1] = uint8(box[1] / divsq)
					res.Px[x * div + w][y * div + h][2] = uint8(box[2] / divsq)
					res.Px[x * div + w][y * div + h][3] = uint8(box[3] / divsq)
				}
			}
		}
	}

	return res
}
















