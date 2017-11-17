package main

import "fmt"

func main() {

	/*
	var alfa int = 12
	var beta string = "Selam"
	var gama bool = true
	*/

	/*	alfa := 12
		beta := "Selam"
		gama := true*/

	//	var s1, s2, s3, s4, s5 int

	// fmt.Println(alfa, beta, gama)

	cevap, sabit, isim := topla(10, 4)

	fmt.Println(cevap, sabit, isim)
}

func topla(a, b int) (int, int, string) {
	return (a + b), 10, "erkan"
}

//tavuk5555
