package main

import "log"

func main() {
	log.Println(inteiro())
	log.Println(intestring())
}
func inteiro() int {
	return 10
}
func intestring() (int, string) {
	return 20, "vinte"
}