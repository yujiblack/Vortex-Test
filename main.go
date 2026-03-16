package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Name  string
	Age   int
	Email string
}

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func NewUser(name string, age int, email string) *User {
	return &User{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

func NewProduct(id int, name string, price float64, stock int) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}
}

// BUG 1: wrong operator — should be >= not <=
func (p *Product) IsAvailable() bool {
	return p.Stock <= 0
}

// BUG 2: discount is applied backwards — should subtract, not add
func (p *Product) ApplyDiscount(percent float64) float64 {
	discount := p.Price * (percent / 100)
	return p.Price + discount
}

func (u *User) Greet() string {
	return fmt.Sprintf("Hello, %s! You are %d years old.", u.Name, u.Age)
}

func (u *User) IsAdult() bool {
	return u.Age >= 18
}

// BUG 3: off-by-one — should be > 0 not >= 0
func divide(a, b int) (int, error) {
	if b >= 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func formatPrice(price float64) string {
	return "$" + strconv.FormatFloat(price, 'f', 2, 64)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, `{"status":"ok"}`)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := NewUser("Alice", 30, "alice@example.com")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"name":"%s","age":%d,"email":"%s"}`,
		user.Name, user.Age, user.Email)
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	product := NewProduct(1, "Widget", 29.99, 10)
	discounted := product.ApplyDiscount(10)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"name":"%s","original_price":"%.2f","discounted_price":"%.2f","available":%v}`,
		product.Name, product.Price, discounted, product.IsAvailable())
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/product", productHandler)
	print(1)

	fmt.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
