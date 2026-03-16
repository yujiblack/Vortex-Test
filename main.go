package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func getUser() *User {
	return &User{Name: "Test User", Age: 25}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	user := getUser()
	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Name)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
