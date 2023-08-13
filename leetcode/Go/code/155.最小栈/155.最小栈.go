// Auxiliary stack
// Use auxiliary stack to store minimum value
// at the each time you push the value into Stack

package main

type stack []int

func (s *stack) Push(val int) {
	*s = append(*s, val)
}

func (s *stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) Top() int {
	lenS := len(*s)
	retV := (*s)[lenS-1]

	return retV
}

// GetMin to get the min value of MinArray.
func (s *stack) GetMin() int {
	lenS := len(*s)
	retV := (*s)[lenS-1]

	return retV
}

type MinStack struct {
	Stack    stack
	MinArray stack
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Stack.Push(val)

	if len(this.MinArray) == 0 || val <= this.MinArray.GetMin() {
		this.MinArray.Push(val)
	}
}

func (this *MinStack) Pop() {
	popV := this.Stack[len(this.Stack)-1]

	this.Stack.Pop()
	if popV == this.MinArray[len(this.MinArray)-1] {
		this.MinArray.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.Stack.Top()
}

func (this *MinStack) GetMin() int {
	return this.MinArray.GetMin()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
