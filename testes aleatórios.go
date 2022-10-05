package main

import (
	"log"
)
import
	"math"

func main() {

	log.Println("4+3=", soma())
	log.Println("7-4=", subtração())
	log.Println("35/7=", divisão())
	log.Println("3*7=", multiplicação())
	log.Println("2⁷", exp())
	log.Println("a circunferencia de um circulo de raio 7 é =", circunferencia())
	log.Println("7!=", fatorial(7))

}

func soma() int{

	return 4+3

}

func subtração() int {

	return 7-4

}

func divisão() int {

	return 35/5

}

func multiplicação() int {

	return 3*7

}

func exp() float64 {

	return math.Exp2(7)

}

func circunferencia() float64 {

	return math.Pi * 2 * 7

}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)

}
	