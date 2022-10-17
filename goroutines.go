package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)
var wg sync.WaitGroup
func main() {

	log.Println(runtime.NumCPU())
	log.Println(runtime.NumGoroutine())

	wg.Add(2)

	go unu()
	go du()

	log.Println(runtime.NumGoroutine())

	wg.Wait()
}
func unu()  {

	for i := 0; i < 20; i++{
		log.Println("unu:", i)
		time.Sleep(20)
	}
	wg.Done()

}
func du() {

	for i := 0; i <20; i++ {
		log.Println("du:", i)
		time.Sleep(20)

	}
	wg.Done()
}
