package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	wordMap := make(map[string]int)

	// string.Fields는 white space를 기준으로
	// 문자열을 잘라서 배열 형태로 리턴한다.
	for _, word := range strings.Fields(s) {
		// map에 있는 단어의 개수를 증가시킨다.
		fmt.Println("값 확인: ", wordMap)
		wordMap[word]++

	}
	return wordMap
}

func main() {
	count := wordCount("Word Count Count")
	fmt.Println(count)
	fmt.Println(wordCount("hello hi hello"))
}

// 기초 - 더 많은 타입들 - 맵 연습하기
// wordCount라는 함수를 구현한다.
// 이 함수는 s라는 문자열 내에서
// 각각의 "단어" 등장 횟수를 나타내는 맵을 반환한다.
