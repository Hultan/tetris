package tetris

func col(c uint32) float64 {
	return float64(c) / 65535
}

// Convert playground coords to screen coords
func coordsToScreenCoords(x, y int) (float64, float64) {
	return float64(leftBorder + x*blockWidth), float64(topBorder + (playgroundVisibleHeight-1-y)*blockHeight)
}
