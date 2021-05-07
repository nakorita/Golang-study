package main

import "fmt"

func main() {
	ny, kim := "first name", "family name"

	p := &ny
	fmt.Printf("p: %v\n", p)
	fmt.Printf("&ny: %p\n", &ny)
	fmt.Printf("&p: %p\n", &p)
	fmt.Println(ny)
	fmt.Println(*p)
	*p = "Nayoung"
	fmt.Println(*p)
	fmt.Printf("p: %v\n", p)
	fmt.Printf("&p: %p\n", &p)
	fmt.Println(ny)

	// p := &kim
	// fmt.Println(p)
	fmt.Println(kim)
}
