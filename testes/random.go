package main

import "log"

func main() {
	x := Soma(1, 2, 3)
	log.Println(x)
}
func Soma(i ...int) int {
	soma := 0

	for _, v := range i {
		soma += v
	}
	return soma
}
