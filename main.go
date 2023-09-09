package main

import "fmt"

type Vec2 struct {
	x int
	y int
}

const SIZE = 32

type Pixel int

const (
	EMPTY = iota
	APPLE
	SNAKE
	HEAD
)

type Grid struct {
	size   Vec2
	buffer [SIZE][SIZE]Pixel
}

type Game struct {
	score int
	grid  Grid
}

func main() {
	var game = Game{
		score: 0,
		grid: Grid{
			size: Vec2{
				x: SIZE,
				y: SIZE,
			},
			buffer: [SIZE][SIZE]Pixel{},
		},
	}
	game.render()
}

func (game Game) render() {
	game.grid.render()
}

const (
	CURSOR_HOME  = "\033[H"
	TOP_LEFT     = "╭"
	TOP_RIGHT    = "╮"
	BOTTOM_LEFT  = "╰"
	BOTTOM_RIGHT = "╯"
	HORIZONTAL   = "│"
	VERTICAL     = "─"
)

func (grid Grid) render() {
	fmt.Print(CURSOR_HOME)
	for i := 0; i < grid.size.y+2; i++ {
		if i == 0 {
			fmt.Print(TOP_LEFT)
		} else if i == grid.size.y+1 {
			fmt.Print(BOTTOM_LEFT)
		} else {
			fmt.Print(HORIZONTAL)
		}
		for j := 0; j < 2*grid.size.x+2; j++ {
			if i == 0 || i == grid.size.y+1 {
				fmt.Print(VERTICAL)
			} else if j == 0 || j == 2*grid.size.x+1 {
				fmt.Print(" ")
			} else {
				var x = (j - 1) / 2
				var y int = i - 1
				switch grid.buffer[y][x] {
				case EMPTY:
					fmt.Print(" ")
					break
				case APPLE:
					fmt.Print("@")
					break
				case SNAKE:
					fmt.Print("S")
					break
				case HEAD:
					fmt.Print("^")
					break
				}
			}
		}
		if i == 0 {
			fmt.Print(TOP_RIGHT)
		} else if i == grid.size.y+1 {
			fmt.Print(BOTTOM_RIGHT)
		} else {
			fmt.Print(HORIZONTAL)
		}
		fmt.Println()
	}
}
