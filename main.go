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
	for i := 0; i < game.grid_size.y; i++ {
		for j := 0; j < 2*game.grid_size.x; j++ {
			fmt.Print("@")
		}
		fmt.Println()
	}
}
