package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func show(str string) {
	wg.Add(1)
	for i := 0; i < 5; i++ {
		//time.Sleep(1 * time.Second)
		fmt.Println(str)
		wg.Done()
	}
	wg.Wait()
}

func cal(ch chan int) {
	fmt.Println(200 + <-ch)

}

func main() {
	go show("Welocome")

	show("Dhiraj")

	fmt.Println("Start a main method")
	ch := make(chan int)
	go cal(ch)
	ch <- 25

}
