## Fix the Main

### Instructions

Write and fix the following program.

### Expected functions

```go
package piscine

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door closing...")
	state = CLOSE
	return true
}

func IsDoorOpen(Door Door) {
	PrintStr("Door is open ?")
	return Door.state = OPEN
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("Door is close ?")
}

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