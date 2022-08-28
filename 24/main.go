package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func NewPoint(xPos int, yPos int) *Point {
	return &Point{
		x: xPos,
		y: yPos,
	}
}

func main() {
	pos1, pos2 := NewPoint(0, 0), NewPoint(5, 5)
	fmt.Println("Distance:", math.Sqrt(math.Pow(float64(pos1.x - pos2.x), 2) + (math.Pow(float64(pos1.y - pos2.y), 2))))

}
