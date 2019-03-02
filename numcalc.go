package main

import (
	"fmt"
)

const (
	OPER_INVALID  = 0
	OPER_ADD      = 1
	OPER_DEL      = 2
	OPER_MULTI    = 3
	OPER_DIV      = 4
	OPER_UNKNOWN  = 5
	OPER_SELF     = 6
	INVALID_VAL   = -1
	TARGET_RESULT = 24
)

func CalcValOper(v1, v2, oper int) (int, int) {
	switch oper {
	case OPER_ADD:
		return v1 + v2, 0
	case OPER_DEL:
		return v1 - v2, 0
	case OPER_MULTI:
		return v1 * v2, 0
	case OPER_DIV:
		return v1 / v2, v1 % v2
	default:
		fmt.Printf("WARN: Cannot calc: %d %s %d\n", v1, GetOperDesc(oper), v2)
		return INVALID_VAL, INVALID_VAL
	}
}

func IsOperValid(oper int) bool {
	switch oper {
	case OPER_ADD, OPER_DEL, OPER_DIV, OPER_MULTI:
		return true
	default:
		return false
	}
}

func GetOperDesc(oper int) string {
	switch oper {
	case OPER_ADD:
		return "+"
	case OPER_DEL:
		return "-"
	case OPER_MULTI:
		return "*"
	case OPER_DIV:
		return "/"
	default:
		fmt.Printf("WARN: Unknown oper: %d\n", oper)
		return "?"
	}
}

func isSwap(arr []int, arrlen int, idx int) bool {
	for i := idx + 1; i < arrlen; i++ {
		if arr[idx] == arr[i] {
			return false
		}
	}

	return true
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
