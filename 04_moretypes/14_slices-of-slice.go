package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Printf("board 출력: %v\n", board)

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("len 출력: %v\n", len(board))
		fmt.Printf("cap 출력: %v\n", cap(board))
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
}

// 기초 - 더 많은 타입들 - 슬라이스의 슬라이스
// 슬라이스는 다른 슬라이스를 포함하여
// 모든 타입을 담을 수 있다.
