package tetris

// Convert 0-65535 color to 0-1 color
func col(c uint32) float64 {
	return float64(c) / 65535
}

// Convert field coords to screen coords
func screenCoords(x, y int) (float64, float64) {
	return float64(leftBorder + x*blockWidth), float64(topBorder + (playfieldVisibleHeight-1-y)*blockHeight)
}
