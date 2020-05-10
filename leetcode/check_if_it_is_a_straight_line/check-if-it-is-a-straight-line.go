package line

func checkStraightLine(coordinates [][]int) bool {
	// return checkStraightLineArea(coordinates)
	return checkStraightLineTheta(coordinates)
}

// Time complexity : O(N)
// Space complexity : O(1)
func checkStraightLineArea(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	for i := 2; i < len(coordinates); i++ {
		if ok := onSameLine(coordinates[0], coordinates[1], coordinates[i]); !ok {
			return false
		}
	}

	return true
}

func onSameLine(a, b, c []int) bool {
	return (a[0]*(b[1]-c[1]) + b[0]*(c[1]-a[1]) + c[0]*(a[1]-b[1])) == 0
}

// Theta
// (x-x1)/(x2-x1) = (y-y2)/(y2-y1)
// Time complexity : O(N)
// Space complexity : O(1)
func checkStraightLineTheta(coordinates [][]int) bool {
	x1 := coordinates[0][0]
	y1 := coordinates[0][1]
	x2 := coordinates[1][0]
	y2 := coordinates[1][1]

	thetaX := x2 - x1
	thetaY := y2 - y1

	for i := 2; i < len(coordinates); i++ {
		x := coordinates[i][0]
		y := coordinates[i][1]

		if (x-x1)*thetaY != (y-y1)*thetaX {
			return false
		}
	}

	return true
}
