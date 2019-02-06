package main

import (
	"fmt"
)

func f1(ch chan int) {
	fmt.Println("F1")
	ch <- 1
}
func f2(ch chan int) {
	fmt.Println("F2")
}

func f3(ch chan int) {
	fmt.Println("F3")
}

func f4(ch chan int) {
	go f1(ch)
	go f2(ch)
	go f3(ch)
}

func main() {
	ch := make(chan int)
	f4(ch)
	i := <-ch
	fmt.Println("Hey", i)
	// time.Sleep(2 * time.Second)
}
