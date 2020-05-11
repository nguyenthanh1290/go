package flood

// Time Complexity: O(N), where N is the number of pixels in the image. We might process every pixel.
// Space Complexity: O(N), the size of the implicit call stack when calling dfs.
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		fill(image, sr, sc, image[sr][sc], newColor)
	}

	return image
}

func fill(image [][]int, sr, sc int, oldColor, newColor int) {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[sr]) || image[sr][sc] != oldColor {
		return
	}

	image[sr][sc] = newColor

	fill(image, sr, sc+1, oldColor, newColor)
	fill(image, sr, sc-1, oldColor, newColor)
	fill(image, sr+1, sc, oldColor, newColor)
	fill(image, sr-1, sc, oldColor, newColor)
}
