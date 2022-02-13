package tetris

import (
	"image/color"
)

type tetromino struct {
	id     int
	color  color.Color
	blocks [tetrominoHeight][tetrominoWidth]int
}

var tetrominos = [tetrominoCount]tetromino{
	{
		id:    1,
		color: color.RGBA{R: 0, G: 255, B: 255, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]int{
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
			{0, 0, 1, 0},
		},
	},
	{
		id:    2,
		color: color.RGBA{R: 0, G: 0, B: 255, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]int{
			{0, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		},
	},
	{
		id:    3,
		color: color.RGBA{R: 255, G: 128, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]int{
			{0, 0, 0, 0},
			{0, 0, 0, 1},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		},
	},
	{
		id:    4,
		color: color.RGBA{R: 255, G: 255, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
	},
	{
		id:    5,
		color: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]int{
			{0, 0, 0, 0},
			{0, 0, 1, 1},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
	},
	{
		id:    6,
		color: color.RGBA{R: 200, G: 100, B: 200, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]int{
			{0, 0, 0, 0},
			{0, 0, 1, 0},
			{0, 1, 1, 1},
			{0, 0, 0, 0},
		},
	},
	{
		id:    7,
		color: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 1, 1},
			{0, 0, 0, 0},
		},
	},
}
