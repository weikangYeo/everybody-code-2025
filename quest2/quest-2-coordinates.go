package main

import "fmt"

// nolint: DuplicateDecl
func main() {
	quest1Result := coordinate{0, 0}
	for range 3 {
		quest1Result = runCycle(quest1Result, coordinate{155, 53}, 10)
	}
	fmt.Printf("Quest 1: %#v\n", quest1Result)
	fmt.Printf("Quest 2: %#v\n", getGridTotalEngravedPoint(-4581, 67892, 101, 10))
	fmt.Printf("Quest 3: %#v\n", getGridTotalEngravedPoint(-4581, 67892, 1001, 1))
}

type coordinate struct {
	x, y int64
}

func sum(co1, co2 coordinate) coordinate {
	return coordinate{
		x: co1.x + co2.x,
		y: co1.y + co2.y,
	}
}

func multiply(co1, co2 coordinate) coordinate {
	// X1 * X2 - Y1 * Y2, X1 * Y2 + Y1 * X2
	return coordinate{
		x: co1.x*co2.x - co1.y*co2.y,
		y: co1.x*co2.y + co1.y*co2.x,
	}
}

func divide(co1, co2 coordinate) coordinate {
	return coordinate{
		x: co1.x / co2.x,
		y: co1.y / co2.y,
	}
}

func runCycle(startCoordinate, coordinateToAdd coordinate, dividerCoordinate int64) coordinate {
	result := multiply(startCoordinate, startCoordinate)
	result = divide(result, coordinate{
		x: dividerCoordinate, y: dividerCoordinate,
	})
	result = sum(result, coordinateToAdd)
	return result
}

func getGridTotalEngravedPoint(x, y, gridSize, gridIncremental int64) int64 {
	var count int64
	count = 0
	// x = col, y = row
	baseInput := coordinate{x, y}
	input := baseInput

	// navigate y axis (row++)
	for range gridSize {
		// navigate x axis (col++)
		for range gridSize {
			if isPointEngravable(coordinate{0, 0}, input) {
				//fmt.Print("x")
				count++
			} else {
				//fmt.Print(".")
			}
			input = coordinate{input.x + gridIncremental, input.y}
		}
		//fmt.Println()
		// reset x back to originalInput
		// set y to next row
		input.x = baseInput.x
		input.y = input.y + gridIncremental
	}

	return count

}

func isPointEngravable(startCoordinate, coordinateToAdd coordinate) bool {
	pointResult := startCoordinate
	for range 100 {
		pointResult = runCycle(pointResult, coordinateToAdd, 100000)
	}
	return isWithinRange(pointResult.x) && isWithinRange(pointResult.y)
}

func isWithinRange(num int64) bool {
	return num >= -1000000 && num <= 1000000
}
