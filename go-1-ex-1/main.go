package main

import "fmt"

func main() {
	var firstName string = "Patrick"
	lastName := "Bucher"

	var dayOfBirth byte = 24
	monthOfBirth := 6
	yearOfBirth := 1923

	numberOfSiblings := 1
	var heightInMeters float32 = 1.88

	zodiacSign := '\u264b'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
