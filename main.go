package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func getUser() *User {
	return nil
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
	print(user.Age)
	print(user.Name)
	fmt.Println(user.Name)
	print(user.Age)
}
