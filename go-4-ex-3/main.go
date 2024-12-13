package main
import "fmt"
import "math"
// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula() (float64,float64,float64){
	var a float64 = 2
	var b float64 = 7
	var c float64 = 9

	d := math.Pow(b, 2) - 4*a*c
    switch {
    case d > 0:
        amountSolution := 2
		fmt.Print("Amont of solutions: " , amountSolution)
		var positiveValue float64 = ((b-(2*b)) + math.Sqrt((math.Pow(b,2)- 4*a*c)))/(2*a)
		var negativeValue float64 = ((b-(2*b)) - math.Sqrt((math.Pow(b,2)- 4*a*c)))/(2*a)
		return positiveValue, negativeValue
    case d == 0:
        amountSolution := 1
		fmt.Print("Amont of solutions: " , amountSolution)
        x:= (b-(2*b))/2*a
        return x
    case d < 0:
        amountSolution := 0
		fmt.Print("Amont of solutions: " , amountSolution)
        return 0
    }

	return 0
}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Print(computeQuadraticFormula())
}
