package main

import (
	"unicode"
)

const (
	MAX_INT = 1<<31 - 1
)

// Automaton 参考官方解答写的状态机: https://leetcode.cn/problems/string-to-integer-atoi/solutions/183164/zi-fu-chuan-zhuan-huan-zheng-shu-atoi-by-leetcode-/
type Automaton struct {
	state string
	sign  int
	ans   int
	table map[string][]string
}

func NewAutomation() *Automaton {
	return &Automaton{
		state: "start",
		sign:  1,
		table: map[string][]string{
			"start":      {"start", "signed", "int_number", "end"},
			"signed":     {"end", "end", "int_number", "end"},
			"int_number": {"end", "end", "int_number", "end"},
			"end":        {"end", "end", "end", "end"},
		},
	}
}

func (a *Automaton) Receive(v rune) {
	a.state = a.table[a.state][getCol(v)]
	if a.state == "signed" {
		if v == '-' {
			a.sign = -1
		}
	} else if a.state == "int_number" {
		a.ans = 10*a.ans + int(v-'0')
		a.clamp() // 提前裁剪, 防止溢出
	}
}

func (a *Automaton) clamp() {
	if a.sign < 0 && a.ans > MAX_INT+1 {
		a.ans = MAX_INT + 1
	} else if a.sign > 0 && a.ans > MAX_INT {
		a.ans = MAX_INT
	}
}

func (a *Automaton) Result() int {
	return a.sign * a.ans
}
func getCol(v rune) int {
	if v == ' ' {
		return 0
	} else if v == '+' || v == '-' {
		return 1
	} else if unicode.IsDigit(v) {
		return 2
	} else {
		return 3
	}
}

func myAtoi(s string) int {
	automaton := NewAutomation()
	for _, v := range s {
		automaton.Receive(v)
		if automaton.state == "end" {
			break
		}
	}
	return automaton.Result()
}
