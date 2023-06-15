package utils

import (
	"testing"
)

func TestGetNumber(t *testing.T) {
	for _, number := range []struct {
		input          string
		expectedOutput string
	}{
		{"+000", "0"},
		{"+010", "10"},
		{"+100", "100"},
		{"000", "0"},
		{"010", "10"},
		{"100", "100"},
		{"-000", "0"},
		{"-010", "-10"},
		{"-100", "-100"},
	} {
		tempOutput := getNumber(number.input)
		if tempOutput.sign == "-" {
			tempOutput.number = tempOutput.sign + tempOutput.number
		}
		if number.expectedOutput != tempOutput.number {
			t.Errorf("Error: input: %q, expected to get %q, but got %q", number.input, number.expectedOutput, tempOutput.number)
		}
	}
}

func TestComplentary(t *testing.T) {
	var a number = number{
		sign:   "+",
		number: "100",
	}
	var b number = number{
		sign:   "+",
		number: "1235345",
	}
	complentartZero(&a, &b)

	for _, num := range []struct {
		n1       number
		n2       number
		expected number
	}{
		{
			n1: number{
				sign:   "+",
				number: "100",
			},
			n2: number{
				sign:   "+",
				number: "11100",
			},
			expected: number{
				sign:   "+",
				number: "00100",
			},
		}, {
			n1: number{
				sign:   "-",
				number: "1001234",
			},
			n2: number{
				sign:   "+",
				number: "11100",
			},
			expected: number{
				sign:   "+",
				number: "0011100",
			},
		}, {
			n1: number{
				sign:   "-",
				number: "1234",
			},
			n2: number{
				sign:   "+",
				number: "1110",
			},
			expected: number{
				sign:   "+",
				number: "1234",
			},
		}, {
			n1: number{
				sign:   "-",
				number: "1234",
			},
			n2: number{
				sign:   "+",
				number: "1110",
			},
			expected: number{
				sign:   "+",
				number: "1110",
			},
		},
	} {
		complentartZero(&num.n1, &num.n2)
		if (num.n1.number != num.expected.number) && (num.n2.number != num.expected.number) {
			t.Errorf("failed to complentary 0, n1: %s, n2: %s, expected: %s", num.n1.number, num.n2.number, num.expected.number)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, k := range []struct {
		n1 string
		n2 string
		n3 string
	}{
		{"123", "456", "579"},
		{"999", "1", "1000"},
		{"9", "1", "10"},
	} {
		out := add(k.n1, k.n2)
		if out != k.n3 {
			t.Errorf("Add error: %s + %s should be %s, but got %s", k.n1, k.n2, k.n3, out)
		}
	}
}

func TestSub(t *testing.T) {
	for _, k := range []struct {
		n1 string
		n2 string
		n3 string
	}{
		{"456", "123", "333"},
		{"1000", "1", "999"},
		{"3", "5", "-2"},
		{"30", "50", "-20"},
		{"001", "5", "-4"},
	} {
		out := sub(k.n1, k.n2)
		if out != k.n3 {
			t.Errorf("Sub error: %s - %s should be %s, but got %s", k.n1, k.n2, k.n3, out)
		}
	}
}

func TestSumTwoIntergers(t *testing.T) {
	for _, k := range []struct {
		n1 string
		n2 string
		n3 string
	}{
		{"123", "456", "579"},
		{"-123", "-456", "-579"},
		{"999", "1", "1000"},
		{"-999", "-1", "-1000"},
		{"1", "999", "1000"},
		{"-1", "-999", "-1000"},
		{"33", "-44", "-11"},
		{"-44", "33", "-11"},
		{"0", "-9", "-9"},
		{"-9", "0", "-9"},
		{"9", "-10", "-1"},
		{"-10", "9", "-1"},
		{"-98", "9", "-89"},
		{"9", "-98", "-89"},
		{"123", "-456", "-333"},
		{"-456", "123", "-333"},
		{"123", "-4567", "-4444"},
		{"-4567", "123", "-4444"},
		{"999", "-456", "543"},
		{"-456", "999", "543"},
		{"999999", "-1000000", "-1"},
		{"-1000000", "999999", "-1"},
		{"0", "0", "0"},
		{"01", "02", "3"},
		{"0", "-0", "0"},
		{"-0", "0", "0"},
	} {
		out := SumTwoIntergers(k.n1, k.n2)
		if out != k.n3 {
			t.Errorf("Error: the input1: %s, input2: %s, sum: %s, expected: %s", k.n1, k.n2, out, k.n3)
		}
	}
}
