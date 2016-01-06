package main

import (
	"bufio"
	"fmt"
	"os"
	_ "time" // import시에도 만족한다.
)

func main() {
	var (
		irene, seulgi = "배어머니", "내사랑"
		wendy, joy    = "940229?", "성재야"
	)
	var duckho int
	_ = duckho // Golang은 안쓰는 변수는 이렇게 처리해두어야 한다.
	fmt.Println("Hello world!")
	fmt.Println(wendy)
	fmt.Println(seulgi)
	fmt.Println(irene)
	fmt.Println(joy)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
