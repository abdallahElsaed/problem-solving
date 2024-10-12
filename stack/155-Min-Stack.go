package main
type MinStack struct {
	stack []int
	minStack  []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	
	// Push onto the min stack if it's empty or val is less than or equal to current min
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	// Only pop if the stack is not empty
	if len(this.stack) == 0 {
		return
	}

	// If the popped value is the minimum, pop it from the min stack as well
	if this.stack[len(this.stack)-1] == this.minStack[len(this.minStack)-1] {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}

	// Pop from the main stack stack
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	// Return the top element
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	// Return the minimum element
	return this.minStack[len(this.minStack)-1]
}