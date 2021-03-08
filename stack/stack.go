package stack

import (
	"errors"
	"fmt"
)

type MyStack struct {
	MaxTop int
	Top int
	Arr []int
}

//入栈
func (stack *MyStack) Push(num int) (err error) {
	// full stack
	if stack.Top + 1 == stack.MaxTop {
		return errors.New("stack full")
	}
	stack.Top ++
	stack.Arr[stack.Top] = num
	return
}

//出栈
func (stack *MyStack) Pop() (val int, err error) {
	if stack.Top == -1 {
		return -1, errors.New("stack empty")
	}
	val = stack.Arr[stack.Top]
	stack.Top --
	return val, nil
}

func (stack *MyStack) List() {
	//先判断栈是否为空
	if stack.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下:")
	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, stack.Arr[i])
	}
}

// 获取栈顶元素
func (stack *MyStack) Peek() int {
	if stack.Top == -1 {
		return -1
	} else {
		return stack.Arr[stack.Top]
	}
}
