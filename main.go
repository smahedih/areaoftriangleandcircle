package main

import "fmt"

const Pi = 3.1416

func main() {

	var base, height float64
	var area1, area2, radius float64

	fmt.Print("Enter base: ")
	fmt.Scan(&base)

	fmt.Print("Enter height: ")
	fmt.Scan(&height)

	fmt.Print("Enter radius: ")
	fmt.Scan(&radius)

	area1 = 0.5 * base * height
	area2 = Pi * radius * radius

	fmt.Printf("The area of triangle: %v\n", area1)
	fmt.Printf("The area of circle: %v\n", area2)
}
