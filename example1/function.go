package main

import "fmt"

//関数
func Plus(x, y int) int {
	return x + y
}
// 返り値が複数の場合
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値に変数を指定する場合
func Double(price int) (result int) {
	result = price * 2
	//returnの中身は省略できる
	return
}
//返り値のない関数
func NoReturn() {
	fmt.Println("No Return")
	return
}

//関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

//関数を引数に取る関数
func Callfunction(f func()) {
	f()
}

// クロージャー
func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

//クロージャー応用
func multiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

//ジェネレーター
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	//返り値の破棄
	i2, _ := Div(9, 4)
	fmt.Println(i2)

	fmt.Println(Double(100))

	NoReturn()

	//無名関数
	f := func(x, y int) int {
		return x + y
	}
	i3 := f(1, 2)
	fmt.Println(i3)

	//無名関数の即時実行
	i4 := func(x, y int) int {
		return x + y
	}(1, 2)
	fmt.Println(i4)

	f2 := ReturnFunc()
	f2()

	Callfunction(func() {
		fmt.Println("I'm a function")
	})
	f3 := Later()
	fmt.Println(f3("Hello"))
	fmt.Println(f3("My"))
	fmt.Println(f3("Name"))
	fmt.Println(f3("Is"))
	fmt.Println(f3("Golang"))

	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println(double(10))
	fmt.Println(double(5))
	fmt.Println(triple(10))
	fmt.Println(triple(5))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	ints2 := integers()
	fmt.Println(ints2())
}
