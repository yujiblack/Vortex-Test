package main

import (
	"testing"
)

func TestIsAvailable_InStock(t *testing.T) {
	p := NewProduct(1, "Widget", 29.99, 10)
	// Stock is 10, product should be available
	if !p.IsAvailable() {
		t.Errorf("expected product with stock=10 to be available, got unavailable")
	}
}

func TestIsAvailable_OutOfStock(t *testing.T) {
	p := NewProduct(2, "Gadget", 9.99, 0)
	// Stock is 0, product should NOT be available
	if p.IsAvailable() {
		t.Errorf("expected product with stock=0 to be unavailable, got available")
	}
}

func TestApplyDiscount(t *testing.T) {
	p := NewProduct(1, "Widget", 100.00, 5)
	// 10% off $100 should be $90, not $110
	result := p.ApplyDiscount(10)
	if result != 90.0 {
		t.Errorf("expected discounted price to be 90.00, got %.2f", result)
	}
}

func TestDivide_Normal(t *testing.T) {
	result, err := divide(10, 2)
	if err != nil {
		t.Errorf("expected no error for divide(10,2), got: %v", err)
	}
	if result != 5 {
		t.Errorf("expected 10/2=5, got %d", result)
	}
}

func TestDivide_ByZero(t *testing.T) {
	_, err := divide(10, 0)
	if err == nil {
		t.Errorf("expected error when dividing by zero, got nil")
	}
}

func TestDivide_Negative(t *testing.T) {
	// Dividing by a negative number is valid — should NOT return an error
	result, err := divide(10, -2)
	if err != nil {
		t.Errorf("expected no error for divide(10,-2), got: %v", err)
	}
	if result != -5 {
		t.Errorf("expected 10/-2=-5, got %d", result)
	}
}

func TestNewUser(t *testing.T) {
	u := NewUser("Bob", 25, "bob@example.com")
	if u == nil {
		t.Fatal("NewUser returned nil")
	}
	if u.Name != "Bob" {
		t.Errorf("expected Name=Bob, got %s", u.Name)
	}
	if u.Age != 25 {
		t.Errorf("expected Age=25, got %d", u.Age)
	}
}

func TestIsAdult(t *testing.T) {
	adult := NewUser("Alice", 20, "a@example.com")
	minor := NewUser("Tim", 16, "t@example.com")

	if !adult.IsAdult() {
		t.Errorf("expected age 20 to be adult")
	}
	if minor.IsAdult() {
		t.Errorf("expected age 16 to not be adult")
	}
}

func TestGreet(t *testing.T) {
	u := NewUser("Alice", 30, "alice@example.com")
	got := u.Greet()
	expected := "Hello, Alice! You are 30 years old."
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
	print(1)
	print(2)

}
