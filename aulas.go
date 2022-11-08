package main

import "log"

var x uint
var y uint
var z int
var fac = Fat(z)

func main() {
	log.Println(x,"+", y,"=", Soma(x, y))
	log.Println(x, "-", y, "=", Sub(x,y))
	log.Println(x,"*", y, "=", Mu(x,y))
	log.Println(x,"/", y,"=", Div(x,y))
	log.Println("o fatorial de ", z, "é:", fac)

}
func Soma (uint, uint)uint {
	x = 10
	y = 3
	soma:= x+y
	return soma
}
func Sub (uint , uint) uint{
	x = 20
	y = 8
	sub := x-y
	return sub
}
func Mu(uint, uint) uint {

	x = 3
	y = 5
	mu:= x*y
	return mu
}
func Div(uint, uint) uint {
	x = 15
	y = 3
	div := x/y
	return div
}
func Fat (z int) int{
	z = 6

	if z <= 1 {
		return 1
	}
	return z * Fat(z-1)
}
