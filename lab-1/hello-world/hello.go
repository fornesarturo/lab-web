package main

import (
	"fmt"
)

func bigger(x int, y int) (int, int) {
	if x > y {
		return -1, x
	} else {
		if x == y {
			return 0, x
		} else {
			return 1, y
		}
	}
}

func main() {
	var lab string = "Laboratory #1"
	lab2 := "Laboratory infered"
	fmt.Println(lab)
	fmt.Println(lab2)

	position, number := bigger(2, 3)
	fmt.Println(position, number)
}
