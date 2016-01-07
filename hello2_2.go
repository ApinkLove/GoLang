package main

import (
	"bufio"
	"fmt"
	_ "math"
	"os"
	"unsafe"
)

func main() {
	var num uint16 = 42
	var num1 uint16 = 24
	fmt.Println(num % num1)  // 나머지
	fmt.Println(num1 << num) // 비트연산
	fmt.Println(^num1)       // 비트 반전연산 0->1 , 1->0
	var num2 = 3
	var num3 = 2.2
	//	fmt.Println(num2 + num3) 자료형이 다른 두개는 계산불능
	fmt.Println(float64(num2) + num3)
	fmt.Println(num2 + int(num3))
	//	fmt.Println(math.MaxInt64 + 1) : overflow상태임 허용x
	//	var num4 uint8 = 0 - 1
	//	fmt.Println(num4) : underflow 상태임 부호가 없는 수에 음수를 냈으니..
	fmt.Println(unsafe.Sizeof(num))
	fmt.Println(unsafe.Sizeof(num1))
	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num3))
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
