package array

import "testing"

func TestSafeIndexAt(t *testing.T) {
	type input struct {
		arr []string
		ind int
	}
	tests := []struct {
		name     string
		input    input
		expected string
	}{
		{
			name: "return element in array when index in range",
			input: input{
				arr: []string{"0", "1"},
				ind: 1,
			},
			expected: "1",
		},
		{
			name: "return empty string when index is negative",
			input: input{
				arr: []string{"0", "1"},
				ind: -1,
			},
			expected: "",
		},
		{
			name: "return empty string when index out of range",
			input: input{
				arr: []string{"0", "1"},
				ind: 2,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := SafeIndexAt(tt.input.arr, tt.input.ind)
			if actual != tt.expected {
				t.Errorf("expect %s but got %s", tt.expected, actual)
			}
		})
	}
}

func TestStringArrayEqual(t *testing.T) {
	firstArr := []string{"0", "1"}
	secondArray := []string{"1", "2"}
	tests := []struct {
		name     string
		input    [2][]string
		expected bool
	}{
		{
			name:     "array equal",
			input:    [2][]string{firstArr, firstArr},
			expected: true,
		},
		{
			name:     "array not equal",
			input:    [2][]string{firstArr, secondArray},
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := StringArrayEqual(tt.input[0], tt.input[1])
			if actual != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, actual)
			}
		})
	}
}
