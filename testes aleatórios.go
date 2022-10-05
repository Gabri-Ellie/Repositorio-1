package main

import (
	"log"
	"math"
)

func main() {

	som := soma(3,4)
	sub := subtracao(7,4)
	div := divisao(35,5)
	mul := multiplicacao(3,4)
	expo := exp(7)
	cir := circunferencia(7)
	fat := fatorial(7)

	log.Println(3,"+",4,"=", som)
	log.Println(7,"-",4,"=", sub)
	log.Println(35,"/",5,"=", div)
	log.Println(3,"*",4,"=", mul)
	log.Println(2,"^",7,"=", expo)
	log.Println("pi", "*" ,2,"*",7,"=", cir)
	log.Println(7,"! =", fat)

}

func soma(n1 uint, n2 uint) uint{

	return n1+n2

}

func subtracao(n1 uint,  n2 uint ) uint {

	return n1-n2

}

func divisao(n1 uint, n2 uint) uint {

	return n1/n2

}

func multiplicacao(n1 uint, n2 uint) uint {

	return n1*n2

}

func exp(n1 float64) float64 {

	return math.Exp2(n1)

}

func circunferencia(n1 float64) float64 {

	return math.Pi * n1 * 2

}

func fatorial(n1 int) int {
	if n1 == 1 {
		return n1
	}
	return n1 * fatorial(n1-1)

}
	