package main

import (
	"fmt"
	"os"
	"strconv"
)

//defer
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func main() {
	// if文
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("other")
	}

	//簡易文を設定したif文
	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	//if文でエラーハンドリング
	var s string = "A"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	//for文
	// int := 0
	// for {
	// 	int++
	// 	if int == 3 {
	// 		break
	// 	}
	// 	fmt.Println(int)
	// }

	// point := 0
	// for point < 10 {
	// 	point++
	// 	fmt.Println(point)
	// }

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	TestDefer()

	/*
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()
	*/

	RunDefer()

	//deferを応用したリソースの解放処理
	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello World"))

	//recover
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	//panic
	panic("runtime error")
	fmt.Println("start")
}
