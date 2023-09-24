package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func dist(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
