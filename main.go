package main

import (
	"fmt"
	"net/http" // <-- Add this
)

type User struct {
	Name string
	Age  int
}

func getUser() *User {
	return &User{Name: "Gourav", Age: 20}
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

	// Keep the app running so Kubernetes doesn't restart it!
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil
	print(user.Name)
	print(user.
	print(user.Name
	print(user.Name)

}
