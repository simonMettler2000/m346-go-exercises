package main

import "fmt"

type FullName struct {
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
	FullName         FullName
	BirthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName:         FullName{FirstName: "Simon", LastName: "Mettler"},
		NumberOfSiblings: 2,   // TODO: adjust
		ZodiacSign:       '🗿', // TODO: adjust
		BirthDate:        BirthDate{Day: 11, Month: 2, Year: 2000},
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
