package main

import (
	"fmt"
)

func main() {
	//マップ
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4)

	//エラーハンドリングしながら値を取り出す
	s, ok := m4[3]
	fmt.Println(s, ok)
	if !ok {
		fmt.Println("error")
	}

	m4[3] = "CHINA"
	fmt.Println(m4)
	
	//要素の削除
	delete(m4, 3)
	fmt.Println(m4)

	m5 := map[string]int {
		"Apple": 100,
		"Banana": 200,
	}
	//forによる取り出し
	for k, v := range m5 {
		fmt.Println(k, v)
	}
}
