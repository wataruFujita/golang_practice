package main

import "fmt"


type T struct {
	//右側のUserは省力できる
	User User
}

type User struct {
	Name string
	Age  int
}

func (u *User) SetName(name string) {
	u.Name = name
}

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Speak() {
	fmt.Printf("%sです。%d歳です。\n", a.Name, a.Age)
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Bark() {
	fmt.Printf("%sです。%d歳です。犬種は%sです。\n", d.Name, d.Age, d.Breed)
}

func main() {
	//構造体の埋め込み
	t := T{User: User{
		Name: "user1",
		Age:  10,
	}}

	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	u := T{User: User{
		Name: "user2",
		Age:  20,
	}}

	u.User.SetName("user3")
	fmt.Println(u.User.Name)

	dog := Dog{
		Animal: Animal{
			Name: "ポチ",
			Age:  10,
		},
		Breed: "Chiwawa",
	}

	fmt.Println("Dog Breed:", dog.Breed)
	dog.Speak()
	dog.Bark()
}
