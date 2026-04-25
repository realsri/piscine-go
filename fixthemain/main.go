package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

type Door struct {
	state int
}

const (
	CLOSE int = 1
	OPEN  int = 0
)

func CloseDoor(Door *Door) bool {
	PrintStr("Door Closing...")
	Door.state = CLOSE
	return true
}

func OpenDoor(Door *Door) bool {
	PrintStr("Door Opening...")
	Door.state = OPEN
	return true
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	if Door.state == OPEN {
		return true
	}
	return false
}

func IsDoorClose(Door Door) bool {
	PrintStr("is the Door closed ?")
	if Door.state == CLOSE {
		return true
	}
	return false
}

func main() {
	var door Door

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}
	if IsDoorOpen(door) {
		CloseDoor(&door)
	}
	if door.state == OPEN {
		CloseDoor(&door)
	}
}
