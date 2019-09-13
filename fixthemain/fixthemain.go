package main

import "fmt"

type Door struct{
	state bool
}
const OPEN = false
const CLOSE =true

func OpenDoor(ptrDoor *Door){
	fmt.Print("Door Opening...\n")
	ptrDoor.state = OPEN
}

func CloseDoor(ptrDoor *Door){
	fmt.Print("Door closing...\n")
	ptrDoor.state = CLOSE
}

func IsDoorOpen(ptrDoor Door) bool {
	fmt.Print("is the Door opened ?\n")
	return ptrDoor.state
}

func IsDoorClose(ptrDoor Door) bool {
	fmt.Print("is the Door closed ?\n")
	return ptrDoor.state
}

func main() {
	door := Door{}

	OpenDoor(&door)
	if IsDoorClose(door) {
		OpenDoor(&door)
	}else if IsDoorOpen(door) {
		CloseDoor(&door)
	} else if door.state == OPEN {
		CloseDoor(&door)
	}
}