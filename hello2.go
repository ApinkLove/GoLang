package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var a float64 = 10.12
	var num int = 0xa // 16진수에서 a는 10!
	for i := 0; i < 10; i++ {
		a = a - 0.012
	}
	fmt.Println(a)
	fmt.Println(num)
	const epsilon = 1e-14 // 엡실론
	fmt.Println(math.Abs(a-10) <= epsilon)
	/*
		엡실론보다 차이가 작거나 같으면 그 둘의 값은 같다.
		차이라서 음수가 나올수도 있으니까 Abs 즉 절대값을 씌운다.
	*/
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
