package main

import "fmt"

type vertex struct {
	Lat, Long float64
}

var m map[string]vertex

func main() {
	m = make(map[string]vertex)
	fmt.Println(m)
	// 맵의 zero value는 nil. nil 맵은 키도 없고 키 추가X
	_, exist := m[""]
	fmt.Println(exist)
	m["Bell Labs"] = vertex{
		40.68433, -74.39967,
	}
	m["Jong Labs"] = vertex{
		23.33333, -19.22233,
	}
	// 콘솔 출력 함수에 "맵이름[key]"를 바로 입력하면
	// key값에 해당하는 'value값'만 출력된다.
	fmt.Println(m["Bell Labs"])
	// value값과 true/false 값을 반환받기 위한 변수 두 개 선언
	val, exist := m["Bell Labs"]
	// val : 해당 key의 value값, exist : 해당 키에 값이 존재하느냐
	fmt.Println(val, exist)

	// value값만 반환받겠다. val2에 value값 초기화.
	// 꼭 두 개의 값을 반환하는 게 아니다.
	val2 := m["Bell Labs"]
	fmt.Println(val2)
	// 맵의 value값은 반환받지 않고, 존재 여부만 반환
	_, exist = m["Bell Labs"]
	fmt.Println(exist)

	// 맵도 똑같이 len() 함수 사용 가능.
	// 하지만 cap() 함수는 사용 불가.
	fmt.Println(len(m))
	// fmt.Println(cap(m))

	fmt.Println(m)
}

// 기초 - 더 많은 타입들 - Maps
// 맵은 키를 값에 매핑한다.
// 맵의 zero value는 nil이다.
// nil맵은 키도 없고, 키를 추가할 수도 없다.
// make 함수는 주어진 타입의 초기화되고
// 사용 준비가 된 맵을 반환한다.
