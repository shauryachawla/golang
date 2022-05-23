package main

import (
	"fmt"

	"github.com/golang/basics/imports/greetings" // <group-id>/<artifact>/<path_to_package> <--- group id, artifact id is in go.mod at root --!>

	"rsc.io/quote"
)

func main() {
	s := quote.Glass()
	fmt.Println(s)
	greetings.Hello("Shaurya")
}
