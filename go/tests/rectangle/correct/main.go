package main

import "fmt"

type point struct {
	x int
	y int
}

type rectangle struct {
	upLeft    point
	downRight point
}

func defineRectangle(vPoint []point, n int) *rectangle {
	xmin := vPoint[0].x
	xmax := vPoint[0].x
	ymin := vPoint[0].y
	ymax := vPoint[0].y

	ptr := &rectangle{}

	for i := 0; i < n; i++ {
		if vPoint[i].x < xmin {
			xmin = vPoint[i].x
		}
		if vPoint[i].x > xmax {
			xmax = vPoint[i].x
		}
		if vPoint[i].y < ymin {
			ymin = vPoint[i].y
		}
		if vPoint[i].y > ymax {
			ymax = vPoint[i].y
		}
	}
	ptr.upLeft.x = xmin
	ptr.upLeft.y = ymax
	ptr.downRight.x = xmax
	ptr.downRight.y = ymin

	return ptr
}

func calcArea(ptr *rectangle) int {
	return (ptr.upLeft.x - ptr.downRight.x) * (ptr.downRight.y - ptr.upLeft.y)
}

func main() {
	vPoint := []point{}
	rectangle := &rectangle{}
	n := 7

	for i := 0; i < n; i++ {
		val := point{
			x: i%2 + 1,
			y: i + 2,
		}
		vPoint = append(vPoint, val)
	}

	rectangle = defineRectangle(vPoint, n)
	fmt.Println("area of the rectangle:", calcArea(rectangle))
}
