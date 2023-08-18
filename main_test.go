package main

import "testing"

func TestRomanNumeralValues(t *testing.T) {
	romanValues := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100,
		'D': 500, 'M': 1000, 'i': 1, 'v': 5, 'x': 10, 'l': 50, 'c': 100,
		'd': 500, 'm': 1000}

	if romanValues['I'] != 1 {
		t.Error("Expected 1, got ", romanValues['I'])
	}
	if romanValues['V'] != 5 {
		t.Error("Expected 5, got ", romanValues['V'])
	}
	if romanValues['X'] != 10 {
		t.Error("Expected 10, got ", romanValues['X'])
	}
	if romanValues['L'] != 50 {
		t.Error("Expected 50, got ", romanValues['L'])
	}
	if romanValues['C'] != 100 {
		t.Error("Expected 100, got ", romanValues['C'])
	}
	if romanValues['D'] != 500 {
		t.Error("Expected 500, got ", romanValues['D'])
	}
	if romanValues['M'] != 1000 {
		t.Error("Expected 1000, got ", romanValues['M'])
	}
	if romanValues['i'] != 1 {
		t.Error("Expected 1, got ", romanValues['i'])
	}
	if romanValues['v'] != 5 {
		t.Error("Expected 5, got ", romanValues['v'])
	}
	if romanValues['x'] != 10 {
		t.Error("Expected 10, got ", romanValues['x'])
	}
	if romanValues['l'] != 50 {
		t.Error("Expected 50, got ", romanValues['l'])
	}
	if romanValues['c'] != 100 {
		t.Error("Expected 100, got ", romanValues['c'])
	}
	if romanValues['d'] != 500 {
		t.Error("Expected 500, got ", romanValues['d'])
	}
	if romanValues['m'] != 1000 {
		t.Error("Expected 1000, got ", romanValues['m'])
	}
}

// code from main.go copied here for testing
func RomanToInteger(roman string) int {
	// Convert the Roman numeral string to an integer
	romanValues := map[rune]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000, 'i': 1, 'v': 5, 'x': 10, 'l': 50, 'c': 100, 'd': 500, 'm': 1000}
	total := 0
	prevValue := 0
	for _, c := range roman {
		value := romanValues[c]
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}
	return total
}

func TestRomanToInteger(t *testing.T) {
	tests := []struct {
		roman    string
		expected int
	}{
		{"I", 1},
		{"V", 5},
		{"X", 10},
		{"L", 50},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
		{"i", 1},
		{"v", 5},
		{"x", 10},
		{"l", 50},
		{"c", 100},
		{"d", 500},
		{"m", 1000},
		{"II", 2},
	}

	for _, test := range tests {
		result := RomanToInteger(test.roman)
		if result != test.expected {
			t.Error("Expected ", test.expected, ", got ", result)
		}
	}
}
