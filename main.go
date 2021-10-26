package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

//go:embed helloworld.txt
var helloworld string

func clear() {
	fmt.Print("\x1B[2J\x1B[H")
}

func main() {
	frames := strings.Split(helloworld, "\n\n")
	clear()
	for i, frame := range frames {
		fmt.Println(frame)
		time.Sleep(30 * time.Millisecond)
		if i < len(frames)-1 {
			clear()
		}
	}
}