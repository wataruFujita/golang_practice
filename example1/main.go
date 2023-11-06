package main

import "fmt"
import "math"
import  "github.com/wataruFujita/golang_practice/foo"


func main() {
	fmt.Println("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(foo.Max)
	// fmt.Println(foo.min)
	//関数経由ならminにアクセスできる
	fmt.Println(foo.ReturnMin())
}

