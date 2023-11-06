package main

import (

	"fmt"

	"github.com/wataruFujita/golang_practice/todo_app/app/models"
)

type Person struct {
	Name string
	Age int
}

func main() {
	// fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test3"
	u.Email = "test3@example.com"
	u.Password = "testtest"
	fmt.Println(u)

	u.CreateUser()



	// user, _ := models.GetUser(1)
	// user.CreateTodo("First Todo")

	// u.Name = "Test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// var Db *sql.DB

	// Db, _ = sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "hanako", 19)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd := "SELECT * FROM persons where age = ?"
	// row := Db.QueryRow(cmd, 20)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Third Todo")

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// user2, _ := models.GetUser(2)
	// todos2, _ := user2.GetTodosByUser()
	// for _, v := range todos2 {
	// 	fmt.Println(v)
	// }

	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
