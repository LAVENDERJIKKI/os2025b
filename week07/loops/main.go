// pass_fail reports whether a grade is passing or failing.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	score, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Println(score, status)
}
