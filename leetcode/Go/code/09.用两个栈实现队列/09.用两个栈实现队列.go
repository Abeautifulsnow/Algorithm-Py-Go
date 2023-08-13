package main

type CQueue struct {
	Stack1 []int
	Stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		Stack1: []int{},
		Stack2: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.Stack1 = append(this.Stack1, value)
}

func (this *CQueue) DeleteHead() int {
	stack2Len := len(this.Stack2)
	if stack2Len > 0 {
		retV := this.Stack2[stack2Len-1]
		this.Stack2 = this.Stack2[:stack2Len-1]
		return retV
	}
	if len(this.Stack1) == 0 {
		return -1
	}

	for len(this.Stack1) > 0 {
		this.Stack2 = append(this.Stack2, this.Stack1[len(this.Stack1)-1])
		this.Stack1 = this.Stack1[:len(this.Stack1)-1]
	}

	lenStack2 := len(this.Stack2)
	retV := this.Stack2[lenStack2-1]
	this.Stack2 = this.Stack2[:lenStack2-1]
	return retV
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

// More gopher way

// type stack []int

// func (s *stack) Push(value int) {
// 	*s = append(*s, value)
// }

// func (s *stack) Pop() int {
// 	n := len(*s)
// 	v := (*s)[n-1]
// 	*s = (*s)[:n-1]
// 	return v
// }

// type CQueue struct {
// 	in, out stack
// }

// func Constructor() CQueue {
// 	return CQueue{}
// }

// func (this *CQueue) AppendTail(value int) {
// 	this.in.Push(value)
// }

// func (this *CQueue) DeleteHead() int {
// 	if len(this.out) > 0 {
// 		return this.out.Pop()
// 	} else if len(this.in) > 0 {
// 		for len(this.in) > 0 {
// 			this.out.Push(this.in.Pop())
// 		}
// 		return this.out.Pop()
// 	}
// 	return -1
// }
