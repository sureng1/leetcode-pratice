/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode.cn/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (55.87%)
 * Likes:    1721
 * Dislikes: 0
 * Total Accepted:    460.8K
 * Total Submissions: 824.4K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left  。请你反转从位置 left 到位置 right 的链表节点，返回
 * 反转后的链表 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], left = 2, right = 4
 * 输出：[1,4,3,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [5], left = 1, right = 1
 * 输出：[5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目为 n
 * 1
 * -500
 * 1
 *
 *
 *
 *
 * 进阶： 你可以使用一趟扫描完成反转吗？
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if right <= left || head == nil || head.Next == nil {
		return head
	}

	root := &ListNode{}
	root.Next = head

	cur := root
	bhead := cur
	for i := 0; i < left; i++ {
		bhead = cur
		cur = cur.Next
	}

	var subHead, pre, next *ListNode
	subHead = cur

	for i := 0; i <= right-left && cur != nil; i++ {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	bhead.Next = pre
	subHead.Next = cur
	return root.Next
}

// @lc code=end

