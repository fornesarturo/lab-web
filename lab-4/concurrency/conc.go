package main

import (
	"fmt"
	// "time"
	"sync"
)

func f1(w *sync.WaitGroup) {
	fmt.Println("F1")
	w.Done()
}
func f2(ch chan int) {
	fmt.Println("F2")
	i := <-ch
	fmt.Println("I: ", i)
}

func f3() {
	fmt.Println("F3")
}

func f4(w *sync.WaitGroup, ch chan int) {
	go f1(w)
	go f2(ch)
	go f3()
}

func main() {
	var w sync.WaitGroup
	ch := make(chan int)

	w.Add(2)
	f4(&w, ch)
	ch <- 1
	w.Done()
	w.Wait()
	fmt.Println("Hey")
	// time.Sleep(2 * time.Second)
}
