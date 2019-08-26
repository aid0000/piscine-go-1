package piscine
import (
	"math"
)
func Sqrt(nb float64) float64 {
flnb:=float64(nb)
	if nb<0{
		return 0
	}
	return math.Sqrt(flnb)

}
