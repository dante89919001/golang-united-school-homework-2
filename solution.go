package square
import "math"
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type intCustomType float64

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var S float64 = 0;
	if(sidesNum == 3) {
		return S = ((sideLen * sideLen) * math.sqrt(3) )/ 4;
	}
	else if(sidesNum == 4){
		return S = sideLen*sideLen;
	}
	else if(sidesNum == 0){
		return S = math.Pi * (sideLen*sideLen);
	}
	else{
		return 0;
	}
}
