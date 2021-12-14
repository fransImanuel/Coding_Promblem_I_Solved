package main

type MinStack struct {
	val     []int
	min     int
	tempMin []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.val) == 0 {
		this.val = append(this.val, val)
		this.min = val
		this.tempMin = append(this.tempMin, val)
		return
	}

	if this.min >= val {
		this.tempMin = append(this.tempMin, val)
		this.min = val
	}
	this.val = append(this.val, val)
}

func (this *MinStack) Pop() {
	if this.val[len(this.val)-1] == this.min && this.min == this.tempMin[len(this.tempMin)-1] {
		this.tempMin = this.tempMin[:len(this.tempMin)-1]
		if len(this.tempMin)-1 > 0 {
			this.min = this.tempMin[len(this.tempMin)-1]
		}
		this.val = this.val[:len(this.val)-1]
	} else {
		this.val = this.val[:len(this.val)-1]
	}
}

func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.tempMin)-1 < 1 {
		return this.tempMin[len(this.tempMin)-1]
	}
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

}
