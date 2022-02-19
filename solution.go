package lab1

import (
	"fmt"
	"math"
)
// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type myInt int64

const (
	SidesTriangle = 3
	SidesSquare = 4
	SidesCircle = 0
)

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	var S float64 = 0
	if sidesNum == SidesTriangle {
		S := math.Pow(sideLen, 2) * math.Sqrt(3) / 4
		return S
	} else if sidesNum == SidesSquare {
		S := math.Pow(sideLen, 2)
		return S
	} else if sidesNum == SidesCircle {
		S := math.Pow(sideLen,2)*math.Pi
		return S
	} else {
		fmt.Println("Unknown shape")
		return S
	}
}

func main(){
	CalcSquare(10.0, SidesTriangle)
	CalcSquare(10.0, SidesSquare)
	CalcSquare(10.0, SidesCircle)
	CalcSquare(10.0, 5)
}