/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 *
 * https://leetcode.cn/problems/sort-list/description/
 *
 * algorithms
 * Medium (65.51%)
 * Likes:    2196
 * Dislikes: 0
 * Total Accepted:    455.1K
 * Total Submissions: 694.8K
 * Testcase Example:  '[4,2,1,3]'
 *
 * 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
 * 
 * 
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：head = [4,2,1,3]
 * 输出：[1,2,3,4]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：head = [-1,5,3,4,0]
 * 输出：[-1,0,3,4,5]
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：head = []
 * 输出：[]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 链表中节点的数目在范围 [0, 5 * 10^4] 内
 * -10^5 <= Node.val <= 10^5
 * 
 * 
 * 
 * 
 * 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
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

 func merge(part1, part2 *ListNode, subLen int) (*ListNode, *ListNode) {
	var cur = &ListNode{}
	truncate := func(temp *ListNode) {
		for i := 1; i < subLen && temp != nil; i++ {
			temp = temp.Next
		}
		if temp != nil {
			temp.Next = nil
		}
	}
	truncate(part1)
	truncate(part2)
	var temp *ListNode

	root := cur
	for part1 != nil && part2 != nil {
		if part1.Val > part2.Val {
			temp = part2
			part2 = part2.Next
		} else {
			temp = part1
			part1 = part1.Next
		}
		cur.Next = temp
		cur = cur.Next
	}
	addRest := func(node *ListNode) {
		for node != nil {
			cur.Next = node
			cur = cur.Next
			node = node.Next
		}
	}
	addRest(part1)
	addRest(part2)
	return root.Next, cur
}

func sortList(head *ListNode) *ListNode {
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	root := head
	for sublen := 1; sublen < length; sublen *= 2 {
		var partAStart, partBStart, l, r *ListNode
		l, r = root, root
		first := true
		for r != nil {
			partAStart = r
			for i := 0; i < sublen && r != nil; i++ {
				r = r.Next
			}
			partBStart = r
			for i := 0; i < sublen && r != nil; i++ {
				r = r.Next
			}
			partStart, partEnd := merge(partAStart, partBStart, sublen)
			if first {
				root = partStart
				first = false
			} else {
				l.Next = partStart
			}
			partEnd.Next = r
			l = partEnd
			// fmt.Println("sublen", sublen)
			// printN(root)
		}
	}

	return root
}
// @lc code=end

