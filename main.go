package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func getUser() *User {
	return nil // ← returns nil
}

func main() {
	user := getUser()
	asdf
	fmt.Println(user.Name) // ← nil pointer dereference
}
