package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Println("Enter some text:")
	fmt.Scanln(&s)
	s1 := make([]rune, len(s)+1)
	for i := 0; i < len(s1)-1; i++ {
		switch s[i] {
		case 'L':
			s1[i] = '1'
			s1[i+1] = '0'
		case 'R':
			s1[i] = '0'
			s1[i+1] = '1'
		case '=':
			if i != 0 {
				s1[i+1] = s1[i]
			} else {
				s1[i+1] = '0'
				s1[i] = '0'
			}
		default:
			fmt.Println("invalid input")
			return
		}
	}
	for i := 1; i < len(s1); i++ {
		cur := s1[i]
		prev := s1[i-1]
		switch s[i-1] {
		case 'L':
			if !(cur < prev) {
				s1[i-1] = s1[i] + 1
			}
		case 'R':
			if !(cur > prev) {
				s1[i] = s1[i-1] + 1
			}
		case '=':
			s1[i] = s1[i-1]
		}
	}

	for i := len(s1) - 1; i > 0; i-- {
		cur := s1[i]
		next := s1[i-1]
		switch s[i-1] {
		case 'L':
			if !(cur < next) {
				s1[i-1] = s1[i] + 1
			}
		}
	}

	sum := 0
	for _, r := range s1 {
		sum += int(r) - '0'
	}
	fmt.Println(string(s1))
	fmt.Println("sum = ", sum)
}
