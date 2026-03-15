package main

import "testing"

func TestGetUser(t *testing.T) {
	user := getUser()
	if user == nil {
		t.Fatal("expected non-nil user, got nil")
	}
	if user.Name == "" {
		t.Errorf("expected non-empty name, got empty")
	}
}

func TestUserAge(t *testing.T) {
	user := getUser()
	if user == nil {
		t.Fatal("expected non-nil user, got nil")
	}
	if user.Age < 0 {
		t.Errorf("expected non-negative age, got %d", user.Age)
	}
}
