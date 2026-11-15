package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{"negative", -5, 0},
		{"zero", 0, 0},
		{"one", 1, 1},
		{"many", 1000, 1000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)

			if len(data) != test.want {
				t.Fatalf(
					"длина: получили %d, ожидали %d",
					len(data),
					test.want,
				)
			}

			for index, value := range data {
				if value <= 0 {
					t.Fatalf(
						"элемент %d должен быть положительным, получили %d",
						index,
						value,
					)
				}
			}
		})
	}
}

func TestMaximumFunctions(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{35}, 35},
		{"maximum first", []int{9, 3, 2}, 9},
		{"maximum middle", []int{2, 9, 3}, 9},
		{"maximum last", []int{2, 3, 9}, 9},
		{"equal", []int{2, 2, 2}, 2},
		{"zeros", []int{0, 0}, 0},
		{"negative", []int{-9, -2, -7}, -2},
	}

	functions := []struct {
		name string
		find func([]int) int
	}{
		{"maximum", maximum},
		{"maxChunks", maxChunks},
	}

	for _, function := range functions {
		for _, test := range tests {
			t.Run(function.name+"/"+test.name, func(t *testing.T) {
				got := function.find(test.data)

				if got != test.want {
					t.Fatalf(
						"получили %d, ожидали %d",
						got,
						test.want,
					)
				}
			})
		}
	}
}

func TestMaxChunksEveryPosition(t *testing.T) {
	for length := 1; length <= 2*CHUNKS+3; length++ {
		for position := 0; position < length; position++ {
			data := make([]int, length)

			for index := range data {
				data[index] = -10
			}

			data[position] = -1

			if got := maxChunks(data); got != -1 {
				t.Fatalf(
					"длина %d, позиция %d: получили %d, ожидали -1",
					length,
					position,
					got,
				)
			}
		}
	}
}
