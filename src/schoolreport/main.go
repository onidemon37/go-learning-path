package main

import (
	"fmt"
	"strings"
)

func main() {
	// var name string = "Dent, Arthur"
	// var score = 87
	// or 
	// name := "Dent, Arthur"
	// score := 87
    // or short declaration syntax
	//name, score := "Dent, Arthur", 87
	students := []string{"Dent, Arthur",
		"MacMillan, Tricia",
		"Prefect, Ford",
	}
	scores := []int{87, 96, 64}
	
	fmt.Println("Student Scores")
	fmt.Println(strings.Repeat("-", 14)) // https://pkg.go.dev/strings@go1.23.5#Repeat
	fmt.Println(students[0], scores[0])
	fmt.Println(students[1], scores[1])
	fmt.Println(students[2], scores[2])
}
