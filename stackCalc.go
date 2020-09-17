package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type numOp interface{}

// StackCalc is a RPN calculator
type StackCalc struct {
	data []numOp
}

// Push Places a number into the StackCalc or executes the mathematical operation
func (r *StackCalc) Push(a numOp) error {
	switch a.(type) {
	case int: // Store int value
		r.data = append(r.data, a)
	case string: // Check if math operation, then call reduce
		str := a.(string)
		if !isMathOp(str) {
			return errors.New("bad math operation only addition +, subtraction -, division /, and multiplication * supported")
		}
		return r.reduce(str)
	default:
		return errors.New("bad type, Please use int for numbers and string for operations")
	}
	return nil
}

// Value gets the last value in the StackCalc - it returns the result of the previous operation
func (r *StackCalc) Value() interface{} {
	endD := r.Length() - 1
	if endD < 0 {
		return nil
	}
	return r.data[endD]
}

// Length returns the number of elements in the StackCalc
func (r *StackCalc) Length() int {
	return len(r.data)
}

// reduce retrieves the last two elements in the StackCalc
// and tries to execute the operation on them, on success it
// stores the new value
func (r *StackCalc) reduce(operation string) error {
	if len(r.data) >= 2 {
		i := len(r.data)
		previousValue, ok1 := r.data[i-2].(int)
		currentValue, ok2 := r.data[i-1].(int)
		if !ok1 || !ok2 {
			return errors.New("bad math operation, operating on non-integer")
		}
		newElement, ok := doOperation(previousValue, currentValue, operation)
		if ok {
			r.data = append(r.data[:i-2], numOp(newElement))
		} else {
			fmt.Println(r.data)
		}
		return nil
	}
	return errors.New("not enough values to operate on")
}

// doOperation Perform the math
func doOperation(x, y int, op string) (int, bool) {
	var result int
	allGood := false
	switch op {
	case "+":
		result, allGood = x+y, true
	case "-":
		result, allGood = x-y, true
	case "*":
		result, allGood = x*y, true
	case "/":
		result, allGood = x/y, true
	}
	return result, allGood
}

// isMathOp checks for math symbols it uses
func isMathOp(a string) bool {
	// Math Ops are 1 char and contain +-/* only
	if len(a) == 1 && strings.ContainsAny(a, "+-/*") {
		return true
	}
	return false
}

// MakeStackCalc creates a new StackCalc by executing the commands in the computationList string
// according to RPN rules
func MakeStackCalc(computationList string) (*StackCalc, error) {
	tmpStc := StackCalc{}
	t, err := ComputeStackCalc(&tmpStc, computationList)
	return t, err
}

// ComputeStackCalc updates a StackCalc by executing the commands in the computationList string
// according to RPN rules
func ComputeStackCalc(tmpStc *StackCalc, computationList string) (*StackCalc, error) {
	cl := strings.Split(computationList, " ")
	for _, compute := range cl {
		err1 := addInt(compute, tmpStc)
		err2 := tmpStc.Push(compute)
		if err1 != nil && err2 != nil {
			return nil, errors.New(err1.Error() + " " + err2.Error())
		}
	}
	return tmpStc, nil
}

// addInt error or add int to StackCalc
func addInt(compute string, tmpStc *StackCalc) error {
	val, err := strconv.Atoi(compute)
	if err == nil {
		return tmpStc.Push(val)
	}
	return err
}
