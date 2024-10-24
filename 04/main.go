package main

import (
	"fmt"
	"golang_course_one/04/model"
)

func main() {
	// personOne := model.Person{
	// 	name:   "Mobin",
	// 	family: "Shaterian",
	// }
	// fmt.Println(personOne)

	p2 := model.New("Ali", "Alavi")
	fmt.Println(p2)

	fmt.Println(p2.GetFamily())
	p2.SetName("Reza")
	fmt.Println(p2.GetName())

}
