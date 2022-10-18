package main

import (
	"log"
	"runtime"
	"sync"
)

func main() {

	log.Println("cpus:", runtime.NumCPU())
	log.Println("goroutines", runtime.NumGoroutine())

	c := 0

	total := 10

	var wg  sync.WaitGroup

	wg.Add(total)

	var mu = sync.Mutex{}

	for i := 0; i < total; i++ {

	go func() {

		mu.Lock()

		v := c

		runtime.Gosched()

		v++

		c = v

		mu.Unlock()

		wg.Done()

		}()
	log.Println("goroutines", runtime.NumGoroutine())
	}
	wg.Wait()

	log.Println("goroutines", runtime.NumGoroutine())

	log.Println("valor final", c)
}

