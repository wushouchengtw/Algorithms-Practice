package utils

import (
	"strconv"
	"strings"
)

type number struct {
	sign   string // + or -
	number string
}

func SumTwoIntergers(a, b string) string {
	n1 := getNumber(a)
	n2 := getNumber(b)

	if n1.sign == n2.sign {
		if n1.sign == "-" {
			return "-" + add(a, b)
		} else {
			return add(a, b)
		}
	} else {
		if n1.sign == "-" {
			return sub(b, a)
		} else {
			return sub(a, b)
		}
	}
}

func getNumber(input string) number {
	var n1 number
	leftNumber := input
	if len(input) > 0 {
		switch string(input[0]) {
		case "+":
			n1.sign = "+"
			leftNumber = input[1:]
		case "-":
			n1.sign = "-"
			leftNumber = input[1:]
		default:
			n1.sign = "+"
		}
	}

	if len(leftNumber) > 0 {
		leftNumber = strings.TrimLeft(leftNumber, "0")
	}
	if len(leftNumber) == 0 {
		n1.sign = "+"
		leftNumber = "0"
	}
	n1.number = leftNumber
	return n1
}

func complentartZero(n1, n2 *number) {
	if len(n1.number) >= len(n2.number) {
		iteration := len(n1.number) - len(n2.number)
		for i := 0; i < iteration; i++ {
			n2.number = "0" + n2.number
		}
	} else {
		iteration := len(n2.number) - len(n1.number)
		for i := 0; i < iteration; i++ {
			n1.number = "0" + n1.number
		}
	}
}

func add(a, b string) string {
	n1 := getNumber(a)
	n2 := getNumber(b)
	complentartZero(&n1, &n2)

	sum, carry := 0, 0
	output := ""
	for i := len(n1.number) - 1; i >= 0; i-- {
		j, _ := strconv.Atoi(string(n1.number[i]))
		k, _ := strconv.Atoi(string(n2.number[i]))
		sum = (j + k + carry) % 10
		carry = (j + k + carry) / 10
		output = strconv.Itoa(sum) + output
	}
	if carry == 1 {
		output = "1" + output
	}
	return output
}

func sub(a, b string) string {
	n1 := getNumber(a)
	n2 := getNumber(b)
	complentartZero(&n1, &n2)

	sum, borrow := 0, 0
	output := ""
	var lessThenZero bool
	for i := len(n1.number) - 1; i >= 0; i-- {
		j, _ := strconv.Atoi(string(n1.number[i]))
		k, _ := strconv.Atoi(string(n2.number[i]))
		sum = j - k - borrow
		if sum < 0 {
			borrow = 1
			sum = 10 + sum
		} else {
			borrow = 0
		}
		output = strconv.Itoa(sum) + output
	}
	if borrow == 1 {
		sum, borrow = 0, 0
		output2 := ""
		for i := len(output) - 1; i >= 0; i-- {
			j, _ := strconv.Atoi(string(output[i]))
			sum = 0 - j - borrow
			if sum < 0 {
				borrow = 1
				sum = 10 + sum
			} else {
				borrow = 0
			}
			output2 = strconv.Itoa(sum) + output2
		}
		output = output2
		lessThenZero = true
	}
	output = strings.TrimLeft(output, "0")
	if lessThenZero {
		output = "-" + output
	}
	if output == "" {
		output = "0"
	}
	return output
}
