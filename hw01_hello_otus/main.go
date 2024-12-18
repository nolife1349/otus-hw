package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	helloStr := "Hello, OTUS!"

	fmt.Println(reverse.String(helloStr))
}
