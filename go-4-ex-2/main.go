package main
import "fmt"
import "math"
// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse() float64{
	var a float64 = 7
	var b float64 = 4
	c := math.Sqrt(math.Pow(a,2)+math.Pow(b,2))
	return c
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Print(computeHypotenuse())
}
