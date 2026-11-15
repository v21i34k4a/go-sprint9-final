package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100, 1000} {
		data := generateRandomElements(size)

		expectedSize := size
		if expectedSize < 0 {
			expectedSize = 0
		}

		if len(data) != expectedSize {
			t.Errorf("size %d: got length %d, want %d",
				size, len(data), expectedSize)
		}

		for _, value := range data {
			if value <= 0 {
				t.Errorf("got non-positive value %d", value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{7}, 7},
		{"ascending", []int{1, 2, 3}, 3},
		{"descending", []int{9, 4, 1}, 9},
		{"repeated", []int{5, 5, 5}, 5},
		{"negative", []int{-9, -3, -7}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d",
					tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{7}, 7},
		{"fewer than eight", []int{1, 8, 3}, 8},
		{"eight elements", []int{3, 9, 1, 5, 7, 2, 4, 6}, 9},
		{"uneven chunks", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, 100},
		{"negative", []int{-9, -3, -7}, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d",
					tt.data, got, tt.want)
			}
		})
	}
}
