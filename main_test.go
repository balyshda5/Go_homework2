package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "positive size", size: 100},
		{name: "zero size", size: 0},
		{name: "negative size", size: -5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if tt.size <= 0 {
				if len(got) != 0 {
					t.Errorf("ожидался пустой слайс, получено %d элементов", len(got))
				}
				return
			}

			if len(got) != tt.size {
				t.Fatalf("ожидалась длина %d, получено %d", tt.size, len(got))
			}

			for i, value := range got {
				if value < 1 || value > 1_000_000 {
					t.Errorf("элемент [%d] вне диапазона: %d", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "maximum at beginning", data: []int{10, 3, 5}, want: 10},
		{name: "maximum at end", data: []int{2, 4, 9}, want: 9},
		{name: "maximum in middle", data: []int{2, 12, 9}, want: 12},
		{name: "equal elements", data: []int{5, 5, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
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
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{42},
			want: 42,
		},
		{
			name: "fewer elements than chunks",
			data: []int{3, 9, 2},
			want: 9,
		},
		{
			name: "exactly eight elements",
			data: []int{1, 2, 3, 4, 5, 6, 7, 20},
			want: 20,
		},
		{
			name: "not divisible by eight",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 99, 10},
			want: 99,
		},
		{
			name: "maximum at the beginning",
			data: []int{100, 2, 3, 4, 5, 6, 7, 8, 9},
			want: 100,
		},
		{
			name: "maximum at the end",
			data: []int{1, 2, 3, 4, 5, 6, 7, 8, 100},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}
