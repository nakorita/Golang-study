package main

import "fmt"

func getSum(x int) func() int {
	//매개변수 x도 getSum함수의 지역변수임
	ret := 0

	return func() int { //클로저 함수(익명함수)
		for i := 1; i <= x; i++ {
			ret += i
		}
		return ret
	}
}

// func useSum() func(x int) int { // 1부터 사용자가 입력한 수까지 더하는 함수
// 	// useSum()의 지역변수 sum 선언
// 	sum := 0

// 	return func(x int) int { //클로저 함수(익명함수)
// 		//매개변수 x도 클로저 함수의 지역변수임
// 		fmt.Println("입력: ")
// 		fmt.Scanln(&x)
// 		for i := 1; i <= x; i++ {
// 			sum += i
// 		}
// 		return sum
// 	}
// }

func main() {
	f := getSum(10)
	fmt.Println(f())

	f2 := getSum(20)
	fmt.Println(f2())

	// f3 := useSum()
	// fmt.Println(f3(var))
}
