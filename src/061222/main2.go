package main

import "fmt"

func main() {
	a, b := 5, 7
	c, d := b, a
	if b > 0 && a > 0 {
		for a != b {
			if a > b {
				a, d = a-b, c+d
			} else {
				b, c = b-a, d+c
			}
		}
		fmt.Printf("Ausgabe: (%d + %d)/2 = %d, (%d + %d)/2 = %d\n", a, b, (a+b)/2, c, d, (c+d)/2)
	} else {
		fmt.Printf("Negative Zahlen können nicht berechnet werden.")
	}
}
