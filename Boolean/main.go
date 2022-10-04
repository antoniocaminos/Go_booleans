package main

import (
	"fmt"
)

//func main() {
/* apples := 18
oranges := 9

//booleans expressions
fmt.Println(apples == oranges)
fmt.Println(apples != oranges)

fmt.Printf("%d > %d: %t", apples, oranges, apples > oranges)
fmt.Println()
fmt.Printf("%d < %d: %t", apples, oranges, apples < oranges)
fmt.Println()
fmt.Printf("%d >= %d: %t", apples, oranges, apples >= oranges)
fmt.Println()
fmt.Printf("%d <= %d: %t", apples, oranges, apples <= oranges)
fmt.Println() */

type Employee struct {
	Name     string
	Age      int
	Salary   int
	FullTime bool
}

func main() {
	jack := Employee{
		Name:     "Jack Smith",
		Age:      27,
		Salary:   40000,
		FullTime: false,
	}
	jill := Employee{
		Name:     "Jill Jones",
		Age:      33,
		Salary:   60000,
		FullTime: true,
	}
	var employees []Employee
	employees = append(employees, jack)
	employees = append(employees, jill)

	for _, x := range employees {
		if x.Age > 30 {
			fmt.Println(x.Name, "tiene 30 o mas")
		} else {
			fmt.Println(x.Name, "tiene menos de 30")

			if x.Age > 30 && x.Salary > 50000 {
				fmt.Println(x.Name, "makes more than 50000 and is over 30")
			} else {
				fmt.Println("either", x.Name, "makes less 50000 or is under 30")
			}
			if x.Age > 30 || x.Salary > 50000 {
				fmt.Println(x.Name, "makes more than 50000 and is over 30")
			} else {
				fmt.Println(x.Name, "makes less 50000 or is under 30")
			}
		}
	}
}
