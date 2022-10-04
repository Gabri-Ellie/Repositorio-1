package main

import "log"

func main() {
	log.Println(fatorial(6))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial (x-1)
}