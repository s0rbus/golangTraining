package chch

import (
	"fmt"
	"strings"
	"sync"
)


func Runit() {
	var w sync.WaitGroup
	w.Add(2)
	c1 := make(chan string)
	c2 := make(chan string)
	cc := make(chan chan string)

	go func(cs chan chan string) {
		ch := <-cs
		for i := 65; i< 76; i++ {
			 ch <- string(i)
		}
		w.Done()
	}(cc)

	go func(cs chan chan string) {
		ch := <-cs
		for i := 65; i< 76; i++ {
			 ch <- strings.ToLower(string(i))
		}
		w.Done()
	}(cc)

	go func(ca chan string, cb chan string) {
		for {
			select {
				case m1 := <-ca:
					fmt.Println("First channel: " + m1)
				case m2 := <-cb:
					fmt.Println("Second channel: " + m2)
			}
		}
	}(c1,c2)

	cc <- c1
	cc <- c2
	w.Wait()
}
