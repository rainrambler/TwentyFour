package main

import (
	"fmt"
)

type Group struct {
	numbers []int
}

func (p *Group) NumCount() int {
	return len(p.numbers)
}

func (p *Group) Desc() string {
	s := ""
	for i := 0; i < p.NumCount(); i++ {
		s += fmt.Sprintf("%d ", p.numbers[i])
	}

	return s
}

type Groups struct {
	group1 Group
	group2 Group
}

type Equation struct {
	EqualDesc string
}

type Equations struct {
	AllEquals []*Equation
}

func (p *Equations) Desc() string {
	s := ""
	for _, eq := range p.AllEquals {
		s += eq.EqualDesc + "\n"
	}

	return s
}

func CreateEquation(leftval, rightval, operdesc int) *Equation {
	var e Equation

	if operdesc == OPER_SELF {
		e.EqualDesc = fmt.Sprintf("%d", leftval)
	} else {
		e.EqualDesc = fmt.Sprintf("%d %s %d", leftval, GetOperDesc(operdesc), rightval)
	}

	return &e
}

func CreateComplexEquation(leftstr, rightstr string, operdesc int) *Equation {
	var e Equation
	left1 := leftstr
	right1 := rightstr

	if len(leftstr) >= 3 {
		left1 = "(" + leftstr + ")"
	}
	if len(rightstr) >= 3 {
		right1 = "(" + rightstr + ")"
	}

	e.EqualDesc = fmt.Sprintf("%s %s %s", left1, GetOperDesc(operdesc), right1)

	return &e
}

func (p *Equations) AddEquation(e *Equation) {
	p.AllEquals = append(p.AllEquals, e)
}

type GroupValues struct {
	Val2Equalations map[int]*Equations
}

func CreateGroupValues() *GroupValues {
	var gv GroupValues
	gv.Val2Equalations = make(map[int]*Equations)
	return &gv
}

func (p *GroupValues) AddEquation(result int, equal *Equation) {
	//fmt.Printf("DBG: Val: %d, Equation: %s\n", result, equal.EqualDesc)
	es, prs := p.Val2Equalations[result]

	if prs {
		es.AddEquation(equal)
	} else {
		// not exist
		es1 := new(Equations)
		es1.AddEquation(equal)

		p.Val2Equalations[result] = es1
	}
}

func (p *GroupValues) CalcEquations(leftval, rightval int) {
	// plus
	added := leftval + rightval
	p.AddEquation(added, CreateEquation(leftval, rightval, OPER_ADD))

	// minus
	if leftval > rightval {
		diff := leftval - rightval
		p.AddEquation(diff, CreateEquation(leftval, rightval, OPER_DEL))
	}

	// mulitply
	p.AddEquation(leftval*rightval, CreateEquation(leftval, rightval, OPER_MULTI))

	// divide
	_, modval := CalcValOper(leftval, rightval, OPER_DIV)
	if modval == 0 {
		p.AddEquation(leftval/rightval, CreateEquation(leftval, rightval, OPER_DIV))
	}
}

func (p *GroupValues) CalcComplexEquations(leftval, rightval int, leftstr, rightstr string) {
	// plus
	added := leftval + rightval
	p.AddEquation(added, CreateComplexEquation(leftstr, rightstr, OPER_ADD))

	// minus
	if leftval > rightval {
		diff := leftval - rightval
		p.AddEquation(diff, CreateComplexEquation(leftstr, rightstr, OPER_DEL))
	}

	// mulitply
	p.AddEquation(leftval*rightval, CreateComplexEquation(leftstr, rightstr, OPER_MULTI))

	// divide
	_, modval := CalcValOper(leftval, rightval, OPER_DIV)
	if modval == 0 {
		p.AddEquation(leftval/rightval, CreateComplexEquation(leftstr, rightstr, OPER_DIV))
	}
}

func CombineGroupValues(leftgroup, rightgroup *GroupValues) *GroupValues {
	gv := CreateGroupValues()

	for leftval1, ve1 := range leftgroup.Val2Equalations {
		for rightval1, ve2 := range rightgroup.Val2Equalations {

			for _, eq1 := range ve1.AllEquals {
				for _, eq2 := range ve2.AllEquals {
					gv.CalcComplexEquations(leftval1, rightval1, eq1.EqualDesc, eq2.EqualDesc)
				}
			}

		}
	}

	return gv
}

func (p *GroupValues) AddGroupValues(another *GroupValues) {
	for result, v := range another.Val2Equalations {
		vorig, exists := p.Val2Equalations[result]

		if exists {
			for _, eq := range v.AllEquals {
				vorig.AddEquation(eq)
			}
		} else {
			vorig1 := new(Equations)
			for _, eq := range v.AllEquals {
				vorig1.AddEquation(eq)
			}
			p.Val2Equalations[result] = vorig1
		}
	}
}

func (p *GroupValues) Desc() string {
	s := ""
	s += fmt.Sprintf("Total values: %d\n", len(p.Val2Equalations))

	for k, v := range p.Val2Equalations {
		for _, eq := range v.AllEquals {
			s += fmt.Sprintf("%d: [%s]\n", k, eq.EqualDesc)
		}
	}

	return s
}

func CalcPossibleValue(g *Group) *GroupValues {
	gv := CreateGroupValues()

	totallen := g.NumCount()

	for i := 1; i < totallen; i++ {
		g1, g2 := g.splitGroup(i)

		g1vals := CreateGroupValues()
		g2vals := CreateGroupValues()

		if g1.NumCount() > 2 {
			g1vals = CalcPossibleValue(g1)
		} else if g1.NumCount() == 2 {
			g1vals.CalcEquations(g1.numbers[0], g1.numbers[1])
		} else if g1.NumCount() == 1 {
			selfval := g1.numbers[0]
			g1vals.AddEquation(selfval, CreateEquation(selfval, 0, OPER_SELF))
		} else if g1.NumCount() == 0 {
			fmt.Printf("WARN: Group invalid, cannot get g1: [%s]\n", g.Desc())
		}

		if g2.NumCount() > 2 {
			g2vals = CalcPossibleValue(g2)
		} else if g2.NumCount() == 2 {
			g2vals.CalcEquations(g2.numbers[0], g2.numbers[1])
		} else if g2.NumCount() == 1 {
			selfval := g2.numbers[0]
			g2vals.AddEquation(selfval, CreateEquation(selfval, 0, OPER_SELF))
		} else if g2.NumCount() == 0 {
			fmt.Printf("WARN: Group invalid, cannot get g2: [%s]\n", g.Desc())
		}

		gsum := CombineGroupValues(g1vals, g2vals)

		gv.AddGroupValues(gsum)
	}

	return gv
}

func (p *Group) findValue(val int) {
	gvs := CalcPossibleValue(p)
	//fmt.Println(gvs.Desc())

	eqs, exists := gvs.Val2Equalations[val]
	if exists {
		fmt.Printf("%s", eqs.Desc())
	} else {
		//fmt.Println("Not Found!")
	}
}

func (p *Group) splitGroup(pos int) (*Group, *Group) {
	if pos < 0 {
		return nil, nil
	}

	if pos > p.NumCount() {
		return nil, nil
	}

	g1 := new(Group)
	g1.numbers = make([]int, pos)
	g2 := new(Group)
	g2.numbers = make([]int, p.NumCount()-pos)

	for i := 0; i < pos; i++ {
		g1.numbers[i] = p.numbers[i]
	}

	j := 0
	for i := pos; i < len(p.numbers); i++ {
		g2.numbers[j] = p.numbers[i]
		j++
	}

	return g1, g2
}

// https://blog.csdn.net/k346k346/article/details/51154786
func permutationCalc(arr []int, idx int) {
	arrlen := len(arr)

	if idx == arrlen {
		//fmt.Printf("finish arr: %+v\n", arr)
		// finished permutation

		var grp1 Group
		grp1.numbers = append(grp1.numbers, arr...)
		grp1.findValue(TARGET_RESULT)
	} else {
		for i := idx; i < arrlen; i++ {
			if isSwap(arr, arrlen, i) {
				arr[i], arr[idx] = arr[idx], arr[i]

				permutationCalc(arr, idx+1)

				arr[i], arr[idx] = arr[idx], arr[i]
			}
		}
	}
}
