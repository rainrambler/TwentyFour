package main

import (
	"fmt"
	"testing"
)

func TestEnumPossibleValues1(t *testing.T) {
	i := 2
	arr := enumRestTargetValues(i, TARGET_RESULT, true, nil)

	fmt.Printf("%+v\n", arr)

	if len(arr) != 2 {
		t.Errorf("TestEnumPossibleValues1 failed: %v, want: 2", len(arr))
	}
}

func TestEnumPossibleValues3(t *testing.T) {
	i := 3
	arr := enumRestTargetValues(i, TARGET_RESULT, true, nil)

	fmt.Printf("%+v\n", arr)

	if len(arr) != 2 {
		t.Errorf("TestEnumPossibleValues3 failed: %v, want: 2", len(arr))
	}
}

func TestEnumRestValues3(t *testing.T) {
	i := 5
	arr := enumRestTargetValues(i, 4, true, nil)

	for _, v := range arr {
		fmt.Printf("%+v\n", v)
	}

	if len(arr) != 1 {
		t.Errorf("TestEnumPossibleValues3 failed: %v, want: 1", len(arr))
	}
}

func TestEnumRestTargetValues3(t *testing.T) {
	i := 5
	arr := enumRestTargetValues(i, 4, true, nil)

	for _, v := range arr {
		fmt.Printf("%+v\n", v)
	}

	if len(arr) != 1 {
		t.Errorf("TestEnumPossibleValues3 failed: %v, want: 1", len(arr))
	}
}

func TestEnumRestTargetValues1(t *testing.T) {
	i := 5
	arr := enumRestTargetValues(i, 20, true, nil)

	for _, v := range arr {
		fmt.Printf("%+v\n", v)
	}

	if len(arr) != 2 {
		t.Errorf("TestEnumPossibleValues3 failed: %v, want: 2", len(arr))
	}
}

func TestEnumRestTargetValues2(t *testing.T) {
	i := 9
	arr := enumRestTargetValues(i, 3, true, nil)

	for _, v := range arr {
		fmt.Printf("%+v\n", v)
	}

	if len(arr) != 2 {
		t.Errorf("TestEnumPossibleValues3 failed: %v, want: 2", len(arr))
	}
}

func TestMatchTargetValue1(t *testing.T) {
	arr := []int{8, 8, 4, 1}

	e := new(Element)
	e.ParentElement = nil
	e.Result = TARGET_RESULT
	e.LeftVal = TARGET_RESULT
	e.isRoot = true
	matchTargetValue(arr, 0, TARGET_RESULT, e)

	if len(arr) != 4 {
		t.Errorf("TestMatchTargetValue1 failed: %v, want: 2", len(arr))
	}
}
