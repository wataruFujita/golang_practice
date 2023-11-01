package main

import "fmt"

func main() {
	// 配列の宣言
	var arr1 [3]int
	fmt.Println(arr1)
	// 型を表示
	fmt.Printf("%T\n", arr1)
	
	// 値を指定していない部分はから文字が入る
	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)

	//暗黙的な宣言
	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//与えられた値の数で要素数が決まる
	arr4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr4)

	fmt.Println(arr2[0])

	//len関数で要素数を取得
	fmt.Println(len(arr2))

}
