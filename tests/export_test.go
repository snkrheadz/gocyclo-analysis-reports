package main

import (
	"testing"
)

func TestSimpleFunction(t *testing.T) {
	t.Parallel()
	result := simpleFunction(5)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestModerateFunction(t *testing.T) {
	t.Parallel()
	result := moderateFunction(15)
	if result != 17 {
		t.Errorf("Expected 17, but got %d", result)
	}
}

func TestComplexFunction(t *testing.T) {
	t.Parallel()
	result := complexFunction(25)
	if result != 28 {
		t.Errorf("Expected 28, but got %d", result)
	}
}

func TestVeryComplexFunction(t *testing.T) {
	t.Parallel()
	result := veryComplexFunction(35)
	if result != 43 {
		t.Errorf("Expected 43, but got %d", result)
	}
}
