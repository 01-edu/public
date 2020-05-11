package main

import "fmt"

const (
	CLOSE = iota
	OPEN
)

type Door struct {
	State int
}

func CloseDoor(ptrDoor *Door) {
	fmt.Println("Door Closing...")
	ptrDoor.State = CLOSE
}

func OpenDoor(ptrdoor *Door) {
	fmt.Println("Door Opening...")
	ptrdoor.State = OPEN
}

func IsDoorOpened(ptrDoor *Door) bool {
	fmt.Println("is the Door opened ?")
	return ptrDoor.State == OPEN
}

func IsDoorClosed(ptrDoor *Door) bool {
	fmt.Println("is the Door closed ?")
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
