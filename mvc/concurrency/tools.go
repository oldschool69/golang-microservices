package main

import (
	"fmt"
)

func main() {
	testBufferChannel()
}

func testBufferChannel() {
	c := make(chan string, 4)
	go func(input chan string) {
		input <- "Hello 1"
		input <- "Hello 2"
		input <- "Hello 3"
		input <- "Hello 4"
		close(input)
	}(c)

	for greeting := range c {
		fmt.Println(greeting)
	}
}
