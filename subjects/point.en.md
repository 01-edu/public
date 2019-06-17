## Point

### Instructions

Create a `.go` file and copy the code below into our file
and add the code necessary so the program works.

- The program must be submitted inside a folder with the name `point`.

```go
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	fmt.Printf("x = %d, y = %d",points.x, points.y)
	fmt.Println()
}
```

### Expected output

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
x = 42, y = 21
student@ubuntu:~/piscine/test$
```
