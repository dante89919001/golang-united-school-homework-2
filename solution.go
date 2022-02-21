package square

import "math"

type intCustomType int

func CalcSquare(sideLen float64, sidesNum intCustomType) (result float64) {
	if sidesNum == 3 {
		return result = (math.Pow(sideLen, 2) * math.Sqrt(3)) / 4
	} else if sidesNum == 4 {
		return result = math.Pow(sideLen, 2)
	} else if sidesNum == 0 {
		return result = math.Pi * math.Pow(sideLen, 2)
	}
	return 0
}
