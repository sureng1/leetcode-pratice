/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 *
 * https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (62.32%)
 * Likes:    2376
 * Dislikes: 0
 * Total Accepted:    969.3K
 * Total Submissions: 1.6M
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
 *
 * 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
 *
 * 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: [3,2,1,5,6,4], k = 2
 * 输出: 5
 *
 *
 * 示例 2:
 *
 *
 * 输入: [3,2,3,1,2,4,5,5,6], k = 4
 * 输出: 4
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= nums.length <= 10^5
 * -10^4 <= nums[i] <= 10^4
 *
 *
 */

// @lc code=start

/* 堆排序 */
// // 注意size的需要报错 i 这个位置
// func adj(as []int, idx, leng int) {
// 	if idx >= leng {
// 		return
// 	}
// 	maxi := idx
// 	l := (idx*2 + 1)
// 	r := (idx + 1) * 2
// 	if l < leng && as[maxi] < as[l] {
// 		maxi = l
// 	}
// 	if r < leng && as[maxi] < as[r] {
// 		maxi = r
// 	}
// 	if maxi == idx {
// 		return
// 	}

// 	as[idx],as[maxi] = as[maxi],as[idx]
// 	adj(as, maxi, leng)
// 	return
// }

// func findKthLargest(nums []int, k int) int {
// 	for i := len(nums) / 2; i >= 0; i-- {
// 		adj(nums, i, len(nums))
// 	}

// 	leng := len(nums)
// 	for i := 1; i < k; i++ {
// 		nums[0], nums[leng-1] = nums[leng-1], nums[0]
// 		leng--
// 		adj(nums, 0, leng)
// 	}
// 	return nums[0]
// }

/* 快排 */

// 要求O(n)，因此需要提前返回，如果k-1是base，name就可以返回了。
func handle(nums []int, l, r,k int) {
	/*
		这是一个降序序列。
		实现时必须从右往左找值。必须使用l作为基准位置。
	*/

	// 考虑的边界条件
	// case1 [3,2] 不需要交换
	// case2 [2,3] 需要交换一次
	// case3 [2,3,4] 需要交换一次，交换完，3不需要交换，用来验证实现有没问题
	if l >= r {
		return
	}

	// 选择nums[l] 作为 base, 记录下nums[l] 的值，腾出位置
	base := nums[l]
	pl, pr := l, r
	for pl < pr {
		// 从右开始遍历找到符合条件的元素（比base大）
		for nums[pr] <= base && pl < pr {
			pr--
		}
		nums[pl] = nums[pr]
		// 现在腾出来的位置在 pr 了，从左开始找元素填充到pr
		for nums[pl] >= base && pl < pr {
			pl++
		}
		nums[pr] = nums[pl]
	}
	// pl即为base的位置，把base放回目标位置
	nums[pl] = base
	if pl == k-1 {
		return
	}
	// 排序剩下的
	handle(nums, l, pl-1,k)
	handle(nums, pl+1, r,k)
}

func findKthLargest(nums []int, k int) int {
	handle(nums, 0, len(nums)-1, k)
	return nums[k-1]
}

// @lc code=end

