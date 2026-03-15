package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func getUser() *User {
	return &User{
		Name: "John Doe",
		Age:  30,
	}
}

func main() {
	user := getUser()
	fmt.Println("User Name:", user.Name)
}
