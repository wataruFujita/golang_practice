package main

import "fmt"

func Double(i int){
	i = i * 2
}

func Double2(i *int){
	*i = *i * 2
}

func Doublev3(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 100
	fmt.Println(n)

	fmt.Println(&n)

	Double(n)

	fmt.Println(n)

	//nのアドレスを渡す
	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	/*
	*p = 300
	fmt.Println(n)
	n = 200
	fmt.Println(*p)
	*/

	Double2(&n)
	//本体nの価が変わる
	fmt.Println(n)

	//スライスは参照型のため参照渡しを行わなくても値が変わる
	var sl []int = []int{1, 2, 3}
	Doublev3(sl)
	fmt.Println(sl)
}
