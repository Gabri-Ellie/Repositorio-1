package main

import "log"

func main() {
	x := []int{42,43,44,45,46,47,48,49,50,51}
	x = append(x, 52)
	x = append(x, 53,54,55)
	log.Println(x)
	y := []int{56,57,58,59,60}
	x = append(x, y...)
	log.Println(x)
}
// cap.9 ex.4
// Nos próximos exercicios, ir variando os tipos de inteiros 
