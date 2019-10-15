## Rectangle

### Instructions

Consider that a point is defined by its coordinates and that a rectangle
is defined by the points of the upper left and lower right corners.

-   Define two structures named, `point` and `rectangle`.

-   The struct `point` has to have two variables, `x` and `y`, type `int`.

-   The struct `rectangle` has to have two variables, `upLeft` and `downRight` type `point`.

-   The goal is to make a program that:
    -   Given a slice of points of size `n` returns the smallest rectangle that contains all the points in the vector of points. The name of that function is `defineRectangle`.
    -   And which calculates and prints the area of that rectangle you define.

### Expected main and function for the program

```go
func defineRectangle(ptr *point, n int) *rectangle {

}

func calArea(ptr *rectangle) int {

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
	fmt.Println("area of the rectangle:", calArea(rectangle))
}
```

### Expected output

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
area of the rectangle: 6
student@ubuntu:~/piscine-go/test$
```
