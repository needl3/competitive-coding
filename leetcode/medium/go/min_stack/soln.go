package minstack

// started at: 9:15
// end at: 9:51

type MinStack struct {
	data    []int
	minData []int
	top     int
}

func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
		top:     0,
	}
}

func (this *MinStack) Push(val int) {
	if this.top == 0 {
		if len(this.data) == 0 {
			this.data = append(this.data, val)
			this.minData = append(this.minData, val)
		} else {
			this.data[this.top] = val
			this.minData[this.top] = val
		}
	} else {
		if len(this.data) == this.top {
			this.data = append(this.data, val)
			if this.GetMin() > val {
				this.minData = append(this.minData, val)
			} else {
				this.minData = append(this.minData, this.GetMin())
			}
		} else {
			this.data[this.top] = val
			if this.GetMin() > val {
				this.minData[this.top] = val
			} else {
				this.minData[this.top] = this.GetMin()
			}
		}
	}

	this.top++
}

func (this *MinStack) Pop() {
	if this.top == 0 {
		panic("Cannot pop empty stack")
	}
	this.top--
}

func (this *MinStack) Top() int {
	if this.top == 0 {
		panic("Stack is empty.")
	}
	return this.data[this.top-1]
}

func (this *MinStack) GetMin() int {
	if this.top == 0 {
		panic("Stack is empty")
	}
	return this.minData[this.top-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
