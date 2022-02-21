package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type side int

func CalcSquare(sideLen float64, sidesNum side) float64 {
	const SidesTriangle side = 3
	const SidesSquare side = 4
	const SidesCircle side = 0
	switch {
	case sidesNum == SidesTriangle:
		return math.Pow(sideLen, 2) * math.Sqrt(3) / 4
	case sidesNum == SidesSquare:
		return math.Pow(sideLen, 2)
	case sidesNum == SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0
	}
}
