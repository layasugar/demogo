package algorithm

import (
	"fmt"
	"testing"
)

type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

// 注意go中所有参数都是值传递
// 所以要让h的变化在函数外也起作用，此处要传指针
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		r := 2*i + 2 // 右孩子
		if l >= len(h) {
			break // i 已经是叶子节点，退出操作。
		}
		j := l
		if r < len(h) && h.less(r, l) {
			j = r // 右孩子为最小子节点
		}
		// i 与最小子节点进行比较
		if h.less(i, j) {
			break // 如果父节点比子节点小，则不交换
		}
		h.swap(i, j) // 交换父子节点
		i = j        //继续向下比较
	}
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}
	n := len(*h) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]
	*h = (*h)[0:n]
	// 如果当前元素大于父节点，向下筛选
	if i == 0 || (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else { // 当前元素小于父节点，向上筛选
		h.up(i)
	}
	return x, true
}

// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
	x, _ := h.Remove(0)
	return x
}

func BuildHeap(arr []int) Heap {
	h := Heap(arr)
	n := len(h)
	// 从第一个非叶子节点，到根节点
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}

	return h
}

func HeapSort(arr []int) {
	// 创建堆
	heap := BuildHeap(arr)
	var sortedArr []int
	for len(heap) > 0 {
		sortedArr = append(sortedArr, heap.Pop())
	}

	fmt.Println(sortedArr)
}

func TestHeap(t *testing.T) {
	HeapSort([]int{33, 24, 8, 3, 10, 15, 16, 15, 30, 17, 19})
}
