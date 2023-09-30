package main

import (
	"fmt"
)

type polynomial struct {
	coeff int
	pow   int
	next  *polynomial
}

func generatePoly(term int) *polynomial {
	var poly *polynomial
	temp := new(polynomial)
	poly = temp
	for i := 0; i < term; i++ {
		var coeff, pow int
		fmt.Scanf("%d %d", &coeff, &pow)
		temp.coeff = coeff
		temp.pow = pow

		temp.next = new(polynomial)
		temp = temp.next
		temp.next = nil
	}
	return poly
}

func addPoly(poly1, poly2 *polynomial) *polynomial {
	var result *polynomial
	temp := new(polynomial)
	temp.next = nil
	result = temp
	for poly1.next != nil && poly2.next != nil {
		switch compare(poly1.pow, poly2.pow) {
		case 1:
			temp.pow = poly1.pow
			temp.coeff = poly1.coeff
			poly1 = poly1.next
		case -1:
			temp.pow = poly2.pow
			temp.coeff = poly2.coeff
			poly2 = poly2.next
		case 0:
			temp.pow = poly1.pow
			temp.coeff = poly1.coeff + poly2.coeff
			poly1 = poly1.next
			poly2 = poly2.next
		}
		if poly1 != nil && poly2 != nil {
			temp.next = new(polynomial)
			temp = temp.next
			temp.next = nil
		}
	}

	for poly1.next != nil || poly2.next != nil {
		temp.next = new(polynomial)
		temp = temp.next
		temp.next = nil

		if poly1.next != nil {
			temp.pow = poly1.pow
			temp.coeff = poly1.coeff
			poly1 = poly1.next
		}
		if poly2.next != nil {
			temp.pow = poly2.pow
			temp.coeff = poly2.coeff
			poly2 = poly2.next
		}
	}
	return result
}

func compare(poly1Pow, poly2Pow int) int {
	if poly1Pow > poly2Pow {
		return 1
	} else if poly1Pow < poly2Pow {
		return -1
	} else {
		return 0
	}
}

func printPoly(result *polynomial) {
	for result != nil {
		if result.coeff != 0 {
			fmt.Printf("%d %d ", result.coeff, result.pow)
		}
		result = result.next
	}
}

func main() {
	var term1, term2 int
	fmt.Scanf("%d", &term1)
	poly1 := generatePoly(term1)
	fmt.Scanf("%d", &term2)
	poly2 := generatePoly(term2)

	answer := addPoly(poly1, poly2)

	printPoly(answer)
}
