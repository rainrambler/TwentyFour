package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CalcAll(s string) {
	srcarr := strings.Split(s, " ")

	if len(srcarr) != 4 {
		fmt.Printf("WARN: Format error: %s\n", s)
		return
	}

	arrint := []int{}

	for _, s := range srcarr {
		val, _ := strconv.Atoi(s)

		arrint = append(arrint, val)
	}

	//fmt.Printf("%+v\n", arrint)
	permutation(arrint, 0)
}

func Calc2(s string) {
	srcarr := strings.Split(s, " ")

	if len(srcarr) != 4 {
		fmt.Printf("WARN: Format error: %s\n", s)
		return
	}

	arrint := []int{}

	for _, s := range srcarr {
		val, _ := strconv.Atoi(s)

		arrint = append(arrint, val)
	}

	fmt.Printf("Array: %+v\n", arrint)
	permutationCalc(arrint, 0)
}
