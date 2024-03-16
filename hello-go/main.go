package main

import (
	"fmt"

	"masasdani.com/hello/quotes"
)

func main() {
	message := quotes.Quote()
	fmt.Println(message)
}
