package main

import (
	"fmt"
	"go_test/stack"
	"strconv"
)

func isOpera(val int) bool {
	//+,-,*,/
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	}
	return false
}

func Cal(opera int, num1 int, num2 int) int {
	res := 0
	switch opera {
	case 42 :
		res = num2 * num1
	case 43 :
		res = num2 + num1
	case 45 :
		res = num2 - num1
	case 47 :
		res = num2 / num1
	default :
		fmt.Println("运算符错误.")
	}
	return res
}

func Priority(opera int) int {
	res := 0
	if opera == 42 || opera == 47 {
		res = 1
	} else if opera == 43 || opera == 45 {
		res = 0
	}
	return res
}

// 四则运算
func FourOperations(str string) int {
	stackNum := stack.MyStack{MaxTop: 20, Top: -1, Arr: make([]int, 20)}
	stackOpera := stack.MyStack{MaxTop: 20, Top: -1, Arr: make([]int, 20)}
	keepNum := ""
	index := 0
	num1 := 0
	num2 := 0
	for {
		ch := str[index:index+1]
		tmp := int([]byte(ch)[0])
		if isOpera(tmp) {
			num := stackOpera.Peek()
			if num == -1 {
				_ = stackOpera.Push(tmp)
			} else {
				if Priority(stackOpera.Peek()) >= Priority(tmp) {
					num1, _ = stackNum.Pop()
					num2, _ = stackNum.Pop()
					opera, _ := stackOpera.Pop()
					_ = stackNum.Push(Cal(opera, num1, num2))
				}
				_ = stackOpera.Push(tmp)
			}
		} else {
			keepNum += ch
			if index+1 == len(str) {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				_ = stackNum.Push(int(val))
				break
			} else {
				if isOpera(int([]byte(str[index+1:index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					_ = stackNum.Push(int(val))
					keepNum = ""
				}
			}
		}
		if index+1 == len(str) {
			break
		}
		index ++
	}

	for {
		if stackOpera.Peek() == -1 {
			break
		}
		num1, _ = stackNum.Pop()
		num2, _ = stackNum.Pop()
		opera, _ := stackOpera.Pop()
		_ = stackNum.Push(Cal(opera, num1, num2))
	}
	//stackNum.List()
	return stackNum.Peek()
}

func main() {
	Exp := "5+5*10-4-7-8+5"
	fmt.Println(FourOperations(Exp))
}
