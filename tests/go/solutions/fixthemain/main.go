package main

import "github.com/01-edu/z01"

const CLOSE = 0
const OPEN = 1

type Door struct {
	State int
}

func PrintStr(str string) {
	arrayRune := []rune(str)
	for _, s := range arrayRune {
		z01.PrintRune(s)
	}
	z01.PrintRune('\n')
}

func CloseDoor(ptrDoor *Door) {
	PrintStr("Door Closing...")
	ptrDoor.State = CLOSE
}

func OpenDoor(ptrdoor *Door) {
	PrintStr("Door Opening...")
	ptrdoor.State = OPEN
}

func IsDoorOpened(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	if ptrDoor.State == OPEN {
		return true
	}
	return false
}

func IsDoorClosed(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	if ptrDoor.State == CLOSE {
		return true
	}
	return false
}

func main() {
	var door Door

	OpenDoor(&door)
	if IsDoorClosed(&door) {
		OpenDoor(&door)
	}
	if IsDoorOpened(&door) {
		CloseDoor(&door)
	}
	if door.State == OPEN {
		CloseDoor(&door)
	}
}
