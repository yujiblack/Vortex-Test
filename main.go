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
	if user != nil {
		fmt.Println(user.Name)
	}
}
