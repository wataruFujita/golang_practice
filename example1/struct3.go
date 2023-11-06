package main

import "fmt"

type User struct {
	Name string
	Age  int
}

//構造体のコンストラクタ
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 10)

	fmt.Println(user1)
	fmt.Println(*user1)
}
