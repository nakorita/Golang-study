package main

import "fmt"

// func fibonacci2(x int) int {
// 	if x < 2 {
// 		return 1
// 	}
// 	return fibonacci2((x - 1)) * fibonacci2((x - 2))
// }

// func main() {
// 	fmt.Println(fibonacci2(4))
// }

func fibonacci2(x int) int {
	answer := 0
	f0, f1 := 1

	if x < 2{
		return x
	}
	for i range(1, x){
		answer = f1 + f0
		f0 = f1
		f1 = answer
	}
	return answer
}

func main() {
	fmt.Println(fibonacci2(5))
}

// https://excelsior-cjh.tistory.com/31
// 위 사이트 예제 코드임