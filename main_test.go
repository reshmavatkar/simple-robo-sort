package main

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		width    int
		height   int
		length   int
		mass     int
		expected string
	}{
		{"Standard Package", 100, 100, 100, 10, "STANDARD"},
		{"Bulky by dimension", 150, 50, 50, 10, "SPECIAL"},
		{"Bulky by volume", 1000, 1000, 2, 10, "SPECIAL"},
		{"Heavy", 100, 100, 100, 25, "SPECIAL"},
		{"Rejected (heavy + bulky)", 200, 200, 200, 25, "REJECTED"},
		{"Edge Case Standard", 149, 149, 149, 19, "STANDARD"},
		{"Edge Case Rejected", 150, 150, 150, 20, "REJECTED"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sort(tt.width, tt.height, tt.length, tt.mass)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
