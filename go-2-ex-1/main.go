package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   byte
	Month byte
	Year  int
}
type Profile struct {
	// TODO: embed full name and birth date information
	FullName         FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName:         FullName{FirstName: "Artur", LastName: "Ferreira"},
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       '♍', // TODO: adjust
		BirthDate:        BirthDate{Day: 12, Month: 9, Year: 2007},
	}

	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
