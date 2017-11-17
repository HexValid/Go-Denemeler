package main

import "fmt"

type Kopek struct {
	yasi, kilo int
	cinsi      string
}

func (k Kopek) havla() {
	fmt.Println(k.cinsi, "Hav hav hav!")
}

func main() {

	alfa := Kopek{
		kilo:  1,
		cinsi: "pitbull",
		yasi:  12,
	}
	fmt.Println("Yaşı", alfa.yasi)
	alfa.havla()

}
