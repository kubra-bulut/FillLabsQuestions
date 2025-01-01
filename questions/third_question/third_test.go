package third_question

import "testing"

func TestMostRepeated(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"apple", "pie", "apple", "red", "red", "red"}, "red"},
		{[]string{"blue", "blue", "green", "green", "green", "yellow"}, "green"},
		{[]string{"one", "two", "three", "one", "three", "three"}, "three"},
		{[]string{"a", "b", "c", "a", "b", "a"}, "a"},
		{[]string{"x"}, "x"},
	}

	for _, test := range tests {
		result := MostRepeated(test.input)
		if result != test.expected {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
