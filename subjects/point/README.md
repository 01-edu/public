## point

### Instructions

Create a new directory called `point`.

- The code below has to be copied in a file called `main.go` inside the `point` directory.

- The necessary changes have to be applied so that the program works.

### Code to be copied

```go
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d\n",points.x, points.y)
}
```

### Usage

```console
$ go run .
x = 42, y = 21
$
```
