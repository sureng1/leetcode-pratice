/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 *
 * https://leetcode.cn/problems/merge-k-sorted-lists/description/
 *
 * algorithms
 * Hard (58.75%)
 * Likes:    2738
 * Dislikes: 0
 * Total Accepted:    746.8K
 * Total Submissions: 1.3M
 * Testcase Example:  '[[1,4,5],[1,3,4],[2,6]]'
 *
 * 给你一个链表数组，每个链表都已经按升序排列。
 *
 * 请你将所有链表合并到一个升序链表中，返回合并后的链表。
 *
 *
 *
 * 示例 1：
 *
 * 输入：lists = [[1,4,5],[1,3,4],[2,6]]
 * 输出：[1,1,2,3,4,4,5,6]
 * 解释：链表数组如下：
 * [
 * ⁠ 1->4->5,
 * ⁠ 1->3->4,
 * ⁠ 2->6
 * ]
 * 将它们合并到一个有序链表中得到。
 * 1->1->2->3->4->4->5->6
 *
 *
 * 示例 2：
 *
 * 输入：lists = []
 * 输出：[]
 *
 *
 * 示例 3：
 *
 * 输入：lists = [[]]
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * k == lists.length
 * 0 <= k <= 10^4
 * 0 <= lists[i].length <= 500
 * -10^4 <= lists[i][j] <= 10^4
 * lists[i] 按 升序 排列
 * lists[i].length 的总和不超过 10^4
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type Heap struct {
	arr    []*ListNode
	length int
}
func (h *Heap) Pop() (*ListNode, bool) {
	if h.length <= 0 {
		return nil, false
	}
	ret := h.arr[0]
	h.arr[0] = h.arr[h.length-1]
	h.length--
	idx := 0
	for {
		min := idx
		l, r := idx*2+1, (idx+1)*2
		if l < h.length && h.arr[l].Val < h.arr[min].Val { //?
			min = l
		}
		if r < h.length && h.arr[r].Val < h.arr[min].Val {
			min = r
		}
		if min == idx {
			break
		}
		h.arr[min], h.arr[idx] = h.arr[idx], h.arr[min]
		idx = min
	}
	return ret, true
}
func (h *Heap) Push(a *ListNode) {
	h.arr[h.length] = a
	idx := h.length
	h.length++
	for {
		p := (idx - 1) / 2
		if h.arr[p].Val > a.Val {
			h.arr[p], h.arr[idx] = h.arr[idx], h.arr[p]
			idx = p
		} else {
			break
		}
	}
}
func mergeKLists(lists []*ListNode) *ListNode {
	heap := &Heap{arr: make([]*ListNode, len(lists))}
	for _, li := range lists {
		if li != nil {
			heap.Push(li)
		}
	}
	root := &ListNode{}
	cur := root
	for {
		min, ok := heap.Pop()
		if !ok {
			break
		}
		cur.Next = min
		cur = cur.Next
		if min.Next != nil {
			heap.Push(min.Next)
		}
	}
	return root.Next
}

// @lc code=end

