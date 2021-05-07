package main

import "fmt"

var star = true

func main() {
	vest := 4
	fmt.Printf("vest is of type %T\n", vest)
	var flower uint = uint(vest)
	fmt.Printf("flower is type of %T\n", flower)
	fmt.Println(flower)
	fmt.Println(vest)
	var ice int
	jerry := ice
	fmt.Printf("ice is of type %T\n", ice)
	fmt.Printf("jerry is of type %T\n", jerry)
	fmt.Printf("star is of type %T\n", star)
}

// 기초 - 패키지와 변수,함수 - Inference Type
// := 혹은 var = 표현을 이용해 명시적인 type을
// 정의하지 않고 변수를 선언할 때,
// 그 변수 type은 오른 편에 있는 값으로부터 유추.
