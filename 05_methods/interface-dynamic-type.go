package main

import "fmt"

// 매개변수로 인터페이스 사용함
func printVal(i interface{}) {
	fmt.Println(i)
}

func main() {
	// 인터페이스 x를 선언했지만 내용 명시X
	var x interface{}

	x = 1
	printVal(x)

	// 위에서는 int형으로 인터페이스 x에 담았지만
	// 아래에서는 string 타입으로 다시 선언.
	// dynamic type
	x = "test"
	printVal(x)

	x = true
	printVal(x)
}

// 함수와 구조체는 형으로 쓸 수 있다.
// 구조체도 형을 사용자 정의로 사용하는 것이고,
// 함수도 형으로 사용할 수 있다.
// 인터페이스도 마찬가지로 형으로 사용할 수 있다.
// 인터페이스의 특징
// 1) 인터페이스는 내용을 따로 선언하지 않아도 형으로 사용 가능.
// 2) 인터페이스는 하나의 형이기 때문에 매개변수로 사용 가능.
// 3) 인터페이스는 어떠한 타입도 담을 수 있는 컨테이너
// 		즉, dynamic type 이다.
