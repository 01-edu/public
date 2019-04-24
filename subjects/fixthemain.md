## Fix the Main

### Instructions

Write and fix the folloing functions.

### Expected functions

```go
func PutStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PutStr("Door closing...")
	state = CLOSE
	return true
}

func IsDoorOpen(Door Door) {
	PutStr("Door is open ?")
	return Door.state = OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PutStr("Door is close ?")
}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}
```

And its output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
Door Opening...
Door is close ?
Door is open ?
Door closing...
student@ubuntu:~/piscine/test$
```
