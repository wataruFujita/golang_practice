package main

import "fmt"

//スライス
// 可変長引数
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// スライス
	var sl []int 
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1,2,3,4,5}
	fmt.Println(sl5[0])

	//部分取り出し
	fmt.Println(sl5[2:4])
	fmt.Println(sl5[:2])
	fmt.Println(sl5[2:])
	fmt.Println(sl5[:])
	fmt.Println(sl5[len(sl5)-1])

	sl6 := []int{100,200}
	fmt.Println(sl6)

	sl6 = append(sl6, 300)
	fmt.Println(sl6)

	sl7 := make([]int, 5)
	fmt.Println(sl7)
	fmt.Println(len(sl7))

	sl8 := make([]int, 5, 10)
	fmt.Println(len(sl8))
	fmt.Println(cap(sl8))

	sl9 := []int{100, 200}
	sl10 := sl9
	sl10[0] = 1000
	// sl9とsl10は同じメモリを参照しているので、sl10の値を変更するとsl9の値も変更される
	fmt.Println(sl9)

	//コピー関数
	sl11 := []int{1, 2, 3, 4, 5}
	sl12 := make([]int, 5, 10)
	fmt.Println(sl12)
	copy(sl12, sl11)
	fmt.Println(sl12)

	//スライス for文
	sl13 := []string{"A", "B", "C"}
	fmt.Println(sl13)

	for _, v := range sl13 {
		fmt.Println(v)
	}

	fmt.Println(Sum(1, 2, 3, 4, 5))
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//スライスを引数に渡すこともできる
	sl14 := []int{1, 2, 3}
	fmt.Println(Sum(sl14...))
}
