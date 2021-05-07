package main

import "fmt"

func main() {
	count1, count2 := 1, 10.4
	// count3 := count1++

	// --count1
	count1++
	count2--

	fmt.Println("count1++ :", count1)
	fmt.Println("count2-- :", count2)
	// fmt.Prinln(count2++)
	count2++
	fmt.Println(count2)
}

// Go에서는 전위 연산자를 사용할 수 없다.(주석처리 코드)
// 또한, 변수 선언과 동시에 증감 연산자를 사용해서 대입할 수 없다.
// 값이 반환되지 않으므로 Println에서 바로 사용할 수 없다.
