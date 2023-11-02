package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "B"
	user.Age = 2000
}

func main() {
	var user1 User
	//この場合、フィールドには値をセットしていないので初期値が自動で入る
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	//フィールドを明示的に指定しない場合は宣言時のフィールドの順番通りに指定する必要がある
	user4 := User{"user4", 40}
	fmt.Println(user4)

	// user5  := User{30, "user5"}
	// fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	//newで定義すると構造体のポインタ型を返す
	user7 := new(User)
	fmt.Println(user7)

	//この方法（アドレス演算子）でもポインタ型を定義できる
	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	//user8はUpdateUser2でポインタ型を渡しているので値が変更されている
	fmt.Println(user8)
}
