package movingaverage

type MovingAverage struct {
	queue     []int
	load      int
	size      int
	actualCap float64
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	moving_average := MovingAverage{}
	moving_average.queue = make([]int, size)
	moving_average.load = 0
	moving_average.actualCap = 0
	moving_average.size = size
	return moving_average
}

func (this *MovingAverage) Next(val int) float64 {
	if this.load == this.size {
		this.load = 0
		this.queue[this.load] = val
		this.load++
	} else {
		this.queue[this.load] = val
		this.load++
		if this.actualCap < float64(this.size) {
			this.actualCap++
		}
	}
	return this.CalcAverage()
}

func (this *MovingAverage) CalcAverage() float64 {
	total := 0.0
	for _, v := range this.queue {
		total += float64(v)
	}
	return float64(total / this.actualCap)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
