 # Point

## Instructions

Create a `.go` file and copy the code below into our file

- The main task is to return a working program.

- 

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

## Expected output

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
x = 42, y = 21 
student@ubuntu:~/piscine/test$
```