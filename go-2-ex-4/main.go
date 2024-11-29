package main

import "fmt"

type Student struct {
    FirstName string
    LastName  string
}

type Class struct {
    Name     string
    Students []Student
}

func main() {
    INA23aL := Class{
        Name: "INA23aL",
        Students: []Student{
            {"Simon", "Mettler"},
            {"Gian-Marco", "Wismer"},
            {"Artur", "Ferreira"},
        },
    }

    INA23bL := Class{
        Name: "INA23bL",
        Students: []Student{
            {"Marco", "Wismer"},
            {"Nicolas", "Rapedius"},
            {"Yves", "Troxler"},
        },
    }

    modules := map[int][]Class{
        114: {INA23aL, INA23bL},
        254: {INA23aL},
        320: {INA23bL},
    }

    fmt.Println("Classes:")
    fmt.Println("INF20:", INA23aL)
    fmt.Println("INF21:", INA23bL)
    fmt.Println("\nModules:")
    fmt.Println("Module 346:", modules[346])
    fmt.Println("Module 254:", modules[254])
    fmt.Println("Module 320:", modules[320])
}
