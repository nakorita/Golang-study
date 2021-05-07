package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - Nil 인터페이스 값
