package main

import "fmt"

type Vec2 struct {
	x int
	y int
}

type Game struct {
	score     int
	grid_size Vec2
}

const CURSOR_HOME = "\033[H"
const TOP_LEFT = "╭"
const TOP_RIGHT = "╮"
const BOTTOM_LEFT = "╰"
const BOTTOM_RIGHT = "╯"
const HORIZONTAL = "│"
const VERTICAL = "─"

func main() {
	var game = Game{
		score: 0,
		grid_size: Vec2{
			x: 32,
			y: 32,
		},
	}
	game.render()
}

func (game Game) render() {
	fmt.Print(CURSOR_HOME)
	for i := 0; i < game.grid_size.y+2; i++ {
		if i == 0 {
			fmt.Print(TOP_LEFT)
		} else if i == game.grid_size.y+1 {
			fmt.Print(BOTTOM_LEFT)
		} else {
			fmt.Print(HORIZONTAL)
		}
		for j := 0; j < 2*game.grid_size.x+2; j++ {
			if i == 0 || i == game.grid_size.y+1 {
				fmt.Print(VERTICAL)
			} else if j == 0 || j == 2*game.grid_size.x+1 {
				fmt.Print(" ")
			} else {
				fmt.Print("@")
			}
		}
		if i == 0 {
			fmt.Print(TOP_RIGHT)
		} else if i == game.grid_size.y+1 {
			fmt.Print(BOTTOM_RIGHT)
		} else {
			fmt.Print(HORIZONTAL)
		}
		fmt.Println()
	}
}
