package main

import (
	"testing"
)

func TestSplitGroup1(t *testing.T) {
	var g Group
	g.numbers = []int{4, 3, 2, 1}

	g1, g2 := g.splitGroup(2)

	if len(g1.numbers) != 2 {
		t.Errorf("Failed: %v, want: 2", len(g1.numbers))
	}

	if len(g2.numbers) != 2 {
		t.Errorf("Failed: %v, want: 2", len(g2.numbers))
	}
}

func TestSplitGroup2(t *testing.T) {
	var g Group
	g.numbers = []int{4, 3, 2, 1, 0}

	g1, g2 := g.splitGroup(2)

	if len(g1.numbers) != 2 {
		t.Errorf("Failed: %v, want: 2", len(g1.numbers))
	}

	if len(g2.numbers) != 3 {
		t.Errorf("Failed: %v, want: 3", len(g2.numbers))
	}
}

func TestGroupDesc1(t *testing.T) {
	var g Group
	g.numbers = []int{4, 3, 2, 1}

	gstr := g.Desc() // [4 3 2 1 ]
	//fmt.Println(gstr)

	if len(gstr) != 8 {
		t.Errorf("Failed: %v, want: 8", len(gstr))
	}
}
