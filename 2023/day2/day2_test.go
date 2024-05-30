package main

import (
	"reflect"
	"testing"
)

func TestAddRound(t *testing.T) {
	tests := []struct {
		input    string
		expected Colours
		hasError bool
	}{
		{" 1 red, 2 green, 3 blue", Colours{red: 1, green: 2, blue: 3}, false},
		{" 5 red, 6 green, 7 blue", Colours{red: 5, green: 6, blue: 7}, false},
		{" 10 red, 20 green, 30 blue", Colours{red: 10, green: 20, blue: 30}, false},
		{" 1 red,2 green, error blue", Colours{}, true},
		{"", Colours{}, false},
	}

	for _, tt := range tests {
		result, err := AddRound(tt.input)
		if tt.hasError {
			if err == nil {
				t.Errorf("expected an error for input %s but got none", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for input %s: %v", tt.input, err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v for input %s", tt.expected, result, tt.input)
			}
		}
	}
}
