package square

import "math"

type myType int

var SidesTriangle myType = 3
var SidesSquare myType = 4
var SidesCircle myType = 0

func CalcSquare(sideLen float64, sidesNum myType) (result float64) {
	if sidesNum == SidesTriangle {
		result = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	} else if sidesNum == SidesSquare {
		result = math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		result = math.Pi * math.Pow(sideLen, 2)
	}
	return
}
