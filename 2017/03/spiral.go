package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(value int) int {
	if value < 0 {
		value = -value
	}

	return value
}

func findMiddle(size int) (x int, y int) {
	x = (size - 1) / 2
	y = size / 2

	return
}

func findPosition(size, value int) (x int, y int) {
	square := size * size
	diff := square - value

	if size%2 == 1 {
		x = size - 1
		y = size - 1

		if diff >= size {
			x = x - (size - 1)
		}

		y = y - diff%size
	} else {
		x = 0
		y = 0

		if diff >= size {
			x = x + (size - 1)
		}

		y = y + diff%size
	}

	return
}

func main() {
	input := 289326

	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		check(err)
		input = i
	}

	size := int(math.Ceil(math.Sqrt(float64(input))))

	x1, y1 := findMiddle(size)
	x2, y2 := findPosition(size, input)

	distance := abs(x2-x1) + abs(y2-y1)

	fmt.Println(distance)
}
