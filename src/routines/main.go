package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	ch := make(chan string)
	doneCh := make(chan bool)

	go genAbc(ch)
	go printer(ch, doneCh)

	println("This prints first")
	<-doneCh
}

func genAbc(ch chan string) {
	for l := byte('a'); l < byte('z'); l++ {
		// go println(string(l))
		ch <- string(l)
	}
	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		println(l)
	}

	doneCh <- true
}
