package main

import "fmt"

const CURSOR_HOME = "\033[H"

func main() {
	cursor_to_home()
	clear_grid(32, 32)
}

func cursor_to_home() {
	fmt.Print(CURSOR_HOME)
}

func clear_grid(x int, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < 2*x; j++ {
			fmt.Print("@")
		}
		fmt.Println()
	}
}
