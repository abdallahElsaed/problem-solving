package main
type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	
	// Push onto the min stack if it's empty or val is less than or equal to current min
	if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	// Only pop if the stack is not empty
	if len(this.data) == 0 {
		return
	}

	// If the popped value is the minimum, pop it from the min stack as well
	if this.data[len(this.data)-1] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}

	// Pop from the main data stack
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	// Return the top element
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	// Return the minimum element
	return this.min[len(this.min)-1]
}