package circularqueue

type MyCircularQueue struct {
	queue []int
	head  int
	tail  int
	size  int
	load  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	queue := MyCircularQueue{}
	queue.queue = make([]int, k)
	queue.size = k
	queue.load = 0
	queue.head = 0
	queue.tail = 0
	return queue
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.load < this.size {
		if this.tail == this.size {
			this.tail = 0
		}
		this.queue[this.tail] = value
		this.tail++
		this.load++
		return true
	}
	return false
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.load > 0 {
		element := this.queue[this.head]
		this.queue[this.head] = 0
		this.head++
		this.load--
		return true
	}
	return false
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.load > 0 {
		return this.queue[this.head]
	}
	return -1
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.load > 0 {
		return this.queue[this.tail-1]
	}
	return -1
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.load == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.load == this.size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
