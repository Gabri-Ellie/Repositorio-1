package main

import "log"

func main() {
	x:= 10

	log.Println(x)

	a(x)

	b(&x)

}
func a(x int) {

	x++

	log.Println("na função é ",x)


}
func b(x *int)  {

	*x++

	log.Println("na função é ",*x)

}