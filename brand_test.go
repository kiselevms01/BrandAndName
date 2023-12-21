package main

import "testing"

func TestCheck(t *testing.T) {
	data := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "valid",
			input:    "Nike | Кроссовки",
			expected: "Ok",
			err:      nil,
		},
		{
			name:     "valid",
			input:    "Nike | Кроссовки Nike",
			expected: "Ok",
			err:      nil,
		},
		{
			name:     "valid",
			input:    "Nike | ",
			expected: "Ok",
			err:      nil,
		},
		{
			name:     "valid",
			input:    "| Кроссовки",
			expected: "Ok",
			err:      nil,
		},
		{
			name:     "valid",
			input:    "Feishon chic | Кроссовки Feishon chic",
			expected: "Ok",
			err:      nil,
		},
		{
			name:     "invalid",
			input:    "Nike | Кроссовки Adidas",
			expected: "",
			err:      ErrSpam,
		},
		{
			name:     "invalid",
			input:    "Nike кроссовки | Кроссовки Nike",
			expected: "",
			err:      ErrSpam,
		},
		{
			name:     "invalid",
			input:    "Nike Кроссовки Nike",
			expected: "",
			err:      ErrInvalidString,
		},
		{
			name:     "invalid",
			input:    "Nike | Кроссовки | Nike",
			expected: "",
			err:      ErrInvalidString,
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := CheckSpam(d.input)

			if result != d.expected {
				t.Errorf("Expected %s, got %s", d.expected, result)
			}

			if err != d.err {
				t.Errorf("Expected %s, got %s", d.err.Error(), err.Error())
			}
		})
	}
}
