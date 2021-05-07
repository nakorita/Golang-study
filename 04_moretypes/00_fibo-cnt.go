package main

import "fmt"

func outterFibo() func() int {
	cnt := 0

	// 밑에 return 부분에 fibo를 사용하려면
	// 미리 fibo라는 변수를 선언해줘야 사용가능.
	// 왜냐하면 fibo := 이렇게 하면 아직 선언하는 중이라고
	// 받아들이기 때문에.
	var fibo func(k int) int
	fibo = func(k int) int { //k를 넣었을 때 k번째 피보나치 항 구하는 함수
		if k < 2 {
			return k
		}
		return fibo(k-1) + fibo(k-2)
	}

	return func() int {
		// cnt를 하나 증가시킴
		cnt += 1
		return fibo(cnt) //cnt번째 피보나치 수열 항
	}
}

func main() {
	fibonacci := outterFibo()
	fmt.Println(fibonacci())
	fmt.Println(fibonacci())
	fmt.Println(fibonacci())
	fmt.Println(fibonacci())
	fmt.Println(fibonacci())

}
