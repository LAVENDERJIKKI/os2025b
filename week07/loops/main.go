package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var month time.Month = now.Month()
	//month :=now.Month()
	var day int = now.Day()
	fmt.Println(month, day)

	univ := "GO$ Inha$"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)
	fmt.Println(changed)
}
