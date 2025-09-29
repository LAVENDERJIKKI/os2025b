package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(2.31))
	fmt.Println(reflect.TypeOf("go developer~"))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(91))
}
