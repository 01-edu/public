## point

### Instructions

Create a `.go` file.

-   The code below has to be copied in that file.

-   The necessary changes have to be applied so that the program works.

-   The program must be submitted inside a folder with the name `point`.

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

### Expected output

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
x = 42, y = 21
student@ubuntu:~/piscine-go/test$
```
