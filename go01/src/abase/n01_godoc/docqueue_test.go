package docqueue

import "fmt"

/**
还可以给docqueue的方法写example
Output是fmt.Println的输出结果
*/
func ExampleDocQueue_Pop() {
	q := DocQueue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 2
	// false
}
