package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// 인덱스만을 원해서, 두 번째 변수 생략
	for i := range pow {
		// == 2**i
		pow[i] = 1 << uint(i)
		fmt.Printf("[%d]현재 pow의 값: %d\n", i, pow[i])
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

// 기초 - 더 많은 타입들 - Range continued
// _을 할당하여 인덱스 또는 값을 건너뛸 수 있다.
// ex) for _, value = range pow
//		for i, _ = range pow
// **만약 인덱스만을 원하면, 두번째 변수 생략 가능
// ex) for i := range pow
