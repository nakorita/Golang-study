package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// 기초 - 패키지와 변수,함수 - 기본 자료형
// Go의 기본 Type
// bool, String,
// int int8 int16 int32 int 64
// uint uint8 uint16 uint32 uint64 uintptr
// byte, rune, float32, float64
// cf) byte는 uint8과 똑같은 자료형.
//		바이트 값을 8비트 부호없는 정수 값과 구별하는 데 사용함.
// cf) rune은 int32와 똑같은 자료형.
// complex64, complex128
// int와 uintptr type은 보통 32-bit 시스템에서는 32bit,
// 64-bit 시스템에서는 64 bit의 길이.
// 정수값이 필요할 때에는 특정한 이유로
// 크기를 정해야하거나 unsigned 정수 type을
// 사용해야하는 게 아니라면 int를 사용해야 한다.
