package main

import "fmt"

func main() {
	//interface型の変数を宣言
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)

	//計算はできなそう
	//エラー　./interface.go:14:14: invalid operation: x + 3.14 (mismatched types interface{} and float64)
	x = 3.14
	fmt.Println(x + 3.14)
}
