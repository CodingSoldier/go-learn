package docqueue

// 给DocQueue写注释
// 一个先入先出的队列
type DocQueue []int

// 给队列添加一个元素
// 		e.g. q.Push(123)
func (q *DocQueue) Push(v int) {
	*q = append(*q, v)
}

// 从队列头弹出一个元素
func (q *DocQueue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

// 队列是否为空
func (q *DocQueue) IsEmpty() bool {
	return len(*q) == 0
}
