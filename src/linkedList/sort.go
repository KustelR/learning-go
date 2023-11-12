// TODO: make this concurrent

package linkedList

func (list *DoubleLinkedList) Sort(fn func(a, b int) (res bool)) int {
	left := 0
	right := list.Length - 1
	swaps := 0

	list.quickSort(fn, left, right, &swaps)
	return swaps
}

func (list *DoubleLinkedList) quickSort(fn func(a, b int) (res bool), start, end int, swaps *int) {
	if start < end {
		isSorted, p := list.partition(fn, start, end, swaps)
		if isSorted {
			return
		}
		list.quickSort(fn, start, p, swaps)
		list.quickSort(fn, p+1, end, swaps)
	}
}

func (list *DoubleLinkedList) partition(fn func(a, b int) (res bool), start, end int, swaps *int) (bool, int) {
	left := list.findNode(start)
	right := list.findNode(end)
	partition := (end-start)/2 + start
	p := list.findNode(partition).value

	l := start
	r := end
	for l <= r {
		for fn(p, left.value) && l != partition {
			if l == end {
				return true, l
			}
			left = left.next
			l++
		}
		for fn(right.value, p) {
			if r == start {
				return true, r
			}
			right = right.prev
			r--
		}
		if l >= r {
			if (l == end || r == start) && end-start == 1 {
				return true, l
			}
			break
		}
		left = left.next
		right = right.prev
		list.swap(left.prev, right.next)
		l++
		r--
		*swaps++
	}
	return false, r
}
