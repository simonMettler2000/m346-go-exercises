package main
import "fmt"
// TODO: implement the function computeGrade

func computeGrade(gotPoints float64, maxPoints float64) float64 {
	var mark float64 = ((gotPoints/maxPoints)*5)+1
	return mark
}

func main() {
	// Example usage
	var gotPoints float64
	var maxPoints float64
	fmt.Print("Enter your achieved exam points: ")
	fmt.Scanln(&gotPoints) // Reads until the newline
	fmt.Print("Enter max achieveable points: ")
	fmt.Scanln(&maxPoints) // Reads until the newline
	fmt.Print(computeGrade(gotPoints, maxPoints))
	
}