package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)

	for {
		v, ok := <-fanin
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}

}

// SEND TO CHANNEL

func send(even, odd chan<- int) {

	for i := 0; i <= 10; i++ {

		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}
	close(odd)
	close(even)
}

// RECEIVE FROM CHANNELS AND FANIN

func receive(odd, even <-chan int, fanin chan<- int) {

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			v, ok := <-even
			if ok {
				fanin <- v
			} else {
				break
			}

		}
		wg.Done()
	}()

	go func() {
		for {
			v, ok := <-odd
			if ok {
				fanin <- v
			} else {
				break
			}

		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
