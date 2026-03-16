package main

import "testing"

func TestGetUser(t *testing.T) {
	user := getUser()
	if user == nil {
		t.Fatal("getUser() returned nil — expected a valid User")
	}
	if user.Name == "" {
		t.Errorf("expected non-empty Name, got empty string")
	}
	if user.Age <= 0 {
		t.Errorf("expected positive Age, got %d", user.Age)
	}
}

func TestMinFunction(t *testing.T) {
	if min(3, 5) != 3 {
		t.Errorf("min(3,5) should be 3")
	}
	if min(10, 2) != 2 {
		t.Errorf("min(10,2) should be 2")
	}
}
