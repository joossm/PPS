package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Eingabe a: ")
	fmt.Scan(&a)
	fmt.Print("Eingabe b: ")
	fmt.Scan(&b)
	p := 1
	c := a
	println("a", "b", "c", "p")
	// Invariante: c^b * p = A^B
	for b != 0 {
		println(a, b, c, p)
		if b%2 == 1 {
			c, p = c*c, p*c
		} else {
			c, p = c*c, p
		}
		//Z
		b = b / 2
	}
	println(a, b, c, p)
	fmt.Printf("Ausgabe: p = %d\n", p) // p = A^B
}
