package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var name string
	// var id int

	// name = "Kim Inha"
	// id = 1000

	// fmt.Println(name, reflect.TypeOf(name))
	// fmt.Println(id, reflect.TypeOf(id))

	// var name string = "Kim Inha"
	// var id int = 1000

	// var name = "Kim Inha"
	// var id = 1000

	name := "Kim Inha"
	id := 1000

	fmt.Println(name, reflect.TypeOf(name))
	fmt.Println(id, reflect.TypeOf(id))
}
