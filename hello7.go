package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	i := 419
	if i < 4 { // 무조건 if나 else if , else 옆에 { 가 있어야 한다.
		fmt.Println("거짓말")
	} else if i < 300 {
		fmt.Println("거짓말임")
	} else {
		fmt.Println("진실")
	}
	if i += 5; i == 424 { // 이런식으로 함수 실행후 조건 할수도 있다.
		fmt.Println("착하네")
	}
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
