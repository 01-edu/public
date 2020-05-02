package main

import "github.com/01-edu/z01"

const CLOSE = 0
const OPEN = 1

type Door struct {
	State int
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
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
	return ptrDoor.State == OPEN
}

func IsDoorClosed(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.State == CLOSE
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
