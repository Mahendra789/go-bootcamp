package main

import (
	"fmt"
	"strings"
)

func main() {

	words := strings.Fields("Hi my name is mahendra")

	for r := range words {
		fmt.Printf("#: %q\n", r)

	}

}
