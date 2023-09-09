package main

import (
	"fmt"
	"math/rand"
)

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

type Direction int

const (
	LEFT = iota
	DOWN
	UP
	RIGHT
)

type Snake struct {
	body   []Vec2
	facing Direction
	head   Vec2
}

type Game struct {
	score int
	grid  Grid
	apple Vec2
	snake Snake
}

func (game Game) render() {
	for i := 0; i < len(game.snake.body); i++ {
		var body_cell = game.snake.body[i]
		game.grid.buffer[body_cell.y][body_cell.x] = SNAKE
	}
	game.grid.buffer[game.apple.y][game.apple.x] = APPLE
	game.grid.buffer[game.snake.head.y][game.snake.head.x] = HEAD
	game.grid.render()
}

func (game *Game) update() {
	var head = game.snake.body[len(game.snake.body)-1]
	var new_head Vec2 = head
	switch game.snake.facing {
	case LEFT:
		new_head.x -= 1
		break
	case RIGHT:
		new_head.x += 1
		break
	case UP:
		new_head.y -= 1
		break
	case DOWN:
		new_head.y += 1
		break
	}
	if head != game.apple {
		game.snake.body = game.snake.body[1:]
	} else {
		game.apple.x = rand.Intn(SIZE)
		game.apple.y = rand.Intn(SIZE)
	}
	game.snake.body = append(game.snake.body, new_head)
	game.snake.head = new_head
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
				var y = i - 1
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

func print_help() {
	fmt.Print(`
Symbols:
	^ -> Head
	@ -> Apple
	S -> Snake Body

Directions:
	0 -> Left
	1 -> Down
	2 -> Up
	3 -> Right
Input:	`)
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
		snake: Snake{
			body: []Vec2{
				{
					x: SIZE / 2,
					y: SIZE / 2,
				},
			},
			facing: LEFT,
		},
		apple: Vec2{
			x: SIZE / 4,
			y: SIZE / 4,
		},
	}
	for {
		game.update()
		game.render()
		print_help()
		var input int
		fmt.Scan(&input)
		game.snake.facing = Direction(input)
	}
}
