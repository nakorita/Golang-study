package main

import "fmt"

type vertex4 struct {
	Lat, Long float64
}

var march = map[string]vertex4{
	"Bell Labs": {40.34355, -74.55554},
	"Google":    {37.44332, -23.44543},
}

func main() {
	fmt.Println(march)
}

// 기초 - 더 많은 타입들 - 계속해서, 맵 리터럴
// 최상위 타입이 타입 이름일 경우,
// 리터럴의 요소에서 생략할 수 있다.
