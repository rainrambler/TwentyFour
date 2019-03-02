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

type Pair struct {
	LeftVal  int
	RightVal int
}

type Element struct {
	LeftVal       int
	RightVal      int
	OperVal       int
	Result        int
	ResultValid   bool
	isLeftLeaf    bool
	isRightLeaf   bool
	isRoot        bool
	ParentElement *Element
}

func DbgPrintElement(p *Element) {
	//fmt.Println("-->Start")
	curEle := p
	s := ""
	for curEle != nil {
		//fmt.Printf("%+v\n", curEle)
		if IsOperValid(curEle.OperVal) {
			if curEle.isRightLeaf {
				if curEle.isLeftLeaf {
					s += fmt.Sprintf(")%d %s %d(", curEle.RightVal, GetOperDesc(curEle.OperVal), curEle.LeftVal)
				} else {
					s += fmt.Sprintf(")%d %s (", curEle.RightVal, GetOperDesc(curEle.OperVal))
					//s += fmt.Sprintf("%d %s ", curEle.RightVal, GetOperDesc(curEle.OperVal))
				}
			} else {
				if curEle.isLeftLeaf {
					s += fmt.Sprintf(" %s %d(", GetOperDesc(curEle.OperVal), curEle.LeftVal)
					//s += fmt.Sprintf("( %s %d", GetOperDesc(curEle.OperVal), curEle.LeftVal)
				} else {
					operdesc := GetOperDesc(curEle.OperVal)

					switch operdesc {
					case "*", "/":
						s += fmt.Sprintf(" %s (", operdesc)
					default:
						s += fmt.Sprintf(" %s ", operdesc)
					}
				}
			}

		} else {
			s += fmt.Sprintf(")%d", curEle.LeftVal)
			//s += fmt.Sprintf("%d", curEle.LeftVal)
		}

		curEle = curEle.ParentElement

		if curEle.isRoot {
			break
		}
	}

	result := Reverse(s)
	fmt.Println(matchBrackets(result))
	//fmt.Println("<--End")
}

func matchBrackets(s string) string {
	matchcount := 0

	for _, ch := range s {
		if ch == '(' {
			matchcount++
		}

		if ch == ')' {
			matchcount--
		}
	}

	for i := 0; i < matchcount; i++ {
		s += ")"
	}

	return s
}

func createElement(leftval, rightval, operval int, leftleaf, rightleaf bool, parent *Element) *Element {
	var e Element
	e.LeftVal = leftval
	e.RightVal = rightval
	e.OperVal = operval
	e.isLeftLeaf = leftleaf
	e.isRightLeaf = rightleaf
	e.ParentElement = parent

	val, remain := CalcValOper(e.LeftVal, e.RightVal, e.OperVal)

	if e.OperVal == OPER_DIV {
		if remain == 0 {
			e.ResultValid = true
			e.Result = val
		} else {
			e.ResultValid = false
			e.Result = val
		}
	} else {
		e.ResultValid = true
		e.Result = val
	}

	return &e
}

func createPair(v1, v2 int) *Pair {
	var p Pair
	p.LeftVal = v1
	p.RightVal = v2
	return &p
}

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

func matchTargetValue(arr []int, startPos, targetval int, curEle *Element) {
	if startPos == len(arr) {
		return
	}

	findFirstSingle(arr, startPos, targetval, curEle)
	findFirstPair(arr, startPos, targetval, curEle)
}

func findFirstSingle(arr []int, startPos, targetval int, curEle *Element) {
	//fmt.Printf("-->findFirstSingle: startpos: %d, Target: %d\n", startPos, targetval)

	if startPos == len(arr) {
		return
	}

	if (startPos + 1) == len(arr) {
		// last
		if arr[startPos] == targetval {
			//fmt.Printf("Single: %+v\n", arr)
			e := new(Element)
			e.LeftVal = targetval
			e.RightVal = INVALID_VAL
			e.isLeftLeaf = true
			e.isRightLeaf = true
			e.ParentElement = curEle
			DbgPrintElement(e)
		}
		return
	}

	curPos := startPos

	// find the objective value of the rest
	// TODO target = 0?
	arrresults := enumRestTargetValues(arr[curPos], targetval, true, curEle)
	for _, v := range arrresults {
		matchTargetValue(arr, curPos+1, v.RightVal, v)
	}

	//fmt.Println("<--findFirstSingle")
}

func findFirstPair(arr []int, startPos, targetval int, curEle *Element) {
	//fmt.Printf("-->findFirstPair: DBG: curpos: %d, Target: %d\n", startPos, targetval)
	if startPos == len(arr) {
		return
	}

	if (startPos + 1) == len(arr) {
		return
	}

	p := createPair(arr[startPos], arr[startPos+1])

	// possible values of the first pair
	possiblevals := findPossibleResults(p, curEle)

	if (startPos + 2) == len(arr) {
		for _, v := range possiblevals {
			if v.Result == targetval {
				//fmt.Printf("Pair: %+v\n", arr)
				DbgPrintElement(v)
			}
		}
		return
	}

	for _, v := range possiblevals {
		// find the objective value of the rest
		arrRightValues := enumRestTargetValues(v.Result, targetval, false, v)

		for _, v2 := range arrRightValues {
			matchTargetValue(arr, startPos+2, v2.RightVal, v2)
		}
	}
}

func findPossibleResults(p *Pair, curEle *Element) []*Element {

	arrEles := []*Element{}

	e := createElement(p.LeftVal, p.RightVal, OPER_ADD, true, true, curEle)
	arrEles = append(arrEles, e)

	e = createElement(p.LeftVal, p.RightVal, OPER_DEL, true, true, curEle)
	arrEles = append(arrEles, e)

	e = createElement(p.LeftVal, p.RightVal, OPER_MULTI, true, true, curEle)
	arrEles = append(arrEles, e)

	_, modval := CalcValOper(p.LeftVal, p.RightVal, OPER_DIV)

	if modval == 0 {
		e = createElement(p.LeftVal, p.RightVal, OPER_DIV, true, true, curEle)
		arrEles = append(arrEles, e)
	}

	return arrEles
}

func enumRestTargetValues(v1, targetval int, isv1leaf bool, curEle *Element) []*Element {
	//fmt.Printf("-->enumRestValues: DBG: Left val: %d, leaf: %v,  Target: %d\n", v1, isv1leaf, targetval)
	arrEles := []*Element{}

	if targetval > v1 {
		e := createElement(v1, targetval-v1, OPER_ADD, isv1leaf, false, curEle)
		arrEles = append(arrEles, e)

		if v1 != 0 {
			divval := targetval / v1
			modval := targetval % v1

			if modval == 0 {
				e = createElement(v1, divval, OPER_MULTI, isv1leaf, false, curEle)
				arrEles = append(arrEles, e)
			}
		}
	} else {
		e := createElement(v1, v1-targetval, OPER_DEL, isv1leaf, false, curEle)
		arrEles = append(arrEles, e)

		// TODO
		if targetval != 0 {
			divval := v1 / targetval
			modval := v1 % targetval

			if modval == 0 {
				e = createElement(v1, divval, OPER_DIV, isv1leaf, false, curEle)
				arrEles = append(arrEles, e)
			}
		}
	}

	//fmt.Println("<--enumRestValues")
	return arrEles
}

// https://blog.csdn.net/k346k346/article/details/51154786
func permutation(arr []int, idx int) {
	arrlen := len(arr)

	if idx == arrlen {
		//fmt.Printf("arr: %+v\n", arr)

		e := new(Element)
		e.ParentElement = nil
		e.Result = TARGET_RESULT
		e.LeftVal = TARGET_RESULT
		e.isRoot = true
		matchTargetValue(arr, 0, TARGET_RESULT, e)
	} else {
		for i := idx; i < arrlen; i++ {
			if isSwap(arr, arrlen, i) {
				arr[i], arr[idx] = arr[idx], arr[i]

				permutation(arr, idx+1)

				arr[i], arr[idx] = arr[idx], arr[i]
			}
		}
	}
}

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
