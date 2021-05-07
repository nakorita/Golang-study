package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// 기초 - 더 많은 타입들 - Arrays
// [n]T 타입은 타입이 T인 n값들의 배열
// 쉽게 말해 T타입인 n개의 값들로 이루어진 배열
// ex) var a [10]int
//		변수 a를 10개의 정수들의 배열로 선언.

// 배열의 길이는 그 타입의 일부이다.
// 따라서 배열의 크기를 조정할 수 없다.
