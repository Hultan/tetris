package tetris

import (
	"image/color"
)

type tetromino struct {
	id     int
	color  color.Color
	blocks [tetrominoHeight][tetrominoWidth]bool
}

var tetrominos = [tetrominoCount]tetromino{
	{
		id:    1,
		color: color.RGBA{R: 0, G: 255, B: 255, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]bool{
			{false, false, true, false},
			{false, false, true, false},
			{false, false, true, false},
			{false, false, true, false},
		},
	},
	{
		id:    2,
		color: color.RGBA{R: 0, G: 0, B: 255, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]bool{
			{false, false, false, false},
			{false, true, false, false},
			{false, true, true, true},
			{false, false, false, false},
		},
	},
	{
		id:    3,
		color: color.RGBA{R: 255, G: 128, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]bool{
			{false, false, false, false},
			{false, false, false, true},
			{false, true, true, true},
			{false, false, false, false},
		},
	},
	{
		id:    4,
		color: color.RGBA{R: 255, G: 255, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]bool{
			{false, false, false, false},
			{false, true, true, false},
			{false, true, true, false},
			{false, false, false, false},
		},
	},
	{
		id:    5,
		color: color.RGBA{R: 0, G: 255, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]bool{
			{false, false, false, false},
			{false, false, true, true},
			{false, true, true, false},
			{false, false, false, false},
		},
	},
	{
		id:    6,
		color: color.RGBA{R: 200, G: 100, B: 200, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]bool{
			{false, false, false, false},
			{false, false, true, false},
			{false, true, true, true},
			{false, false, false, false},
		},
	},
	{
		id:    7,
		color: color.RGBA{R: 255, G: 0, B: 0, A: 255},
		blocks: [tetrominoHeight][tetrominoWidth]bool{
			{false, false, false, false},
			{false, true, true, false},
			{false, false, true, true},
			{false, false, false, false},
		},
	},
}
