package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Zero values
	var f64 float64
	var i16 int16
	var t bool
	var s string
	var i int

	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(i16, reflect.TypeOf(i16))
	fmt.Println(t, reflect.TypeOf(t))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(i, reflect.TypeOf(i))
}
