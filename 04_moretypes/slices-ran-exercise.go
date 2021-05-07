package main

import "fmt"

func pic3(dx, dy int) [][]uint8 {
	// dy만큼의 길이를 가지는 slic이라는 슬라이스 생성
	slic := make([][]uint8, dy)

	for i := range slic {
		slic[i] = make([]uint8, dx)
		for j := range slic[i] {
			slic[i][j] = uint8(i ^ j)
		}
	}
	return slic
}

func main() {
	var dx, dy int
	fmt.Println("숫자 입력: ")
	fmt.Scanf("%d %d\n", &dx, &dy)
	fmt.Println(pic3(dx, dy))
}
