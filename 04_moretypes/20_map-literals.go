package main

import "fmt"

type vertex3 struct {
	Lat, Long float64
}

// var map이름 = map[key 자료형]value 자료형
var m2 = map[string]vertex3{
	// kel:value
	// 위에서 vertex3라는 구조체 선언
	"Bell Labs": vertex3{
		40.68433, -74.39967,
	},
	"Google": vertex3{
		37.08987, -122.33004,
	},
}

func main() {
	fmt.Println(m2)
}

// 기초 - 더 많은 타입들 - 맵 리터럴
// 맵 리터럴은 구조체 리터럴과 같지만, 키가 필요하다.
