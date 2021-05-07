package main

import "fmt"

// 방법 1
var num1 int = 1

//잘못된 방법
// num4 := 4

func main() {
	// 방법 2
	var num2 = 2
	// 방법 3
	num3 := 3
	fmt.Println(num1, num2, num3)

	red := 10
	fmt.Println(red)
	blue := &red
	fmt.Println(blue)
	fmt.Println(&red)
	// blue = 2
	fmt.Println(*blue)
	fmt.Println(&blue)
}

// &연산자 는 해당 변수가 메모리상으로
// 저장되어있는 주소 반환
// *연산자는 해당 포인터 변수가 갖고있는
// 원래의 값을 반환한다.
// &연산자는 해당 변수가 메모리상을 저장되어
// 있는 주소를 반환한다.
// *연산자는 해당 포인터 변수가 갖고있는
// 원래 값을 반환한다.
