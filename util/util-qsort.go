/*
 * @lc app=leetcode.cn id=912 lang=golang
 *
 * [912] 排序数组
 *
 * https://leetcode.cn/problems/sort-an-array/description/
 *
 * algorithms
 * Medium (49.31%)
 * Likes:    986
 * Dislikes: 0
 * Total Accepted:    640.5K
 * Total Submissions: 1.3M
 * Testcase Example:  '[5,2,3,1]'
 *
 * 给你一个整数数组 nums，请你将该数组升序排列。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [5,2,3,1]
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [5,1,1,2,0,0]
 * 输出：[0,0,1,1,2,5]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 5 * 10^4
 * -5 * 10^4 <= nums[i] <= 5 * 10^4
 *
 *
 */

// @lc code=start

// func sortArray(nums []int) []int {
// 	quickSortHelper(nums, 0, len(nums)-1)
// 	return nums
// }
// func quickSortHelper(arr []int, left, right int) {
//     if left < right {
//         pivot := partition(arr, left, right)
//         quickSortHelper(arr, left, pivot-1)
//         quickSortHelper(arr, pivot+1, right)
//     }
// }

// func partition(arr []int, left, right int) int {
//     pivot := arr[right]
//     i := left

//     for j := left; j < right; j++ {
//         if arr[j] < pivot {
//             arr[i], arr[j] = arr[j], arr[i]
//             i++
//         }
//     }

//     arr[i], arr[right] = arr[right], arr[i]
//     return i
// }

func sortArray(nums []int) []int {
	helper(nums, 0, len(nums)-1)
	return nums
}
func helper(arr []int, left, right int) {
	if left >= right {
		return
	}
	seed := rand.Intn(right - left)
	arr[left+seed], arr[right] = arr[right], arr[left+seed]
	p := arr[right]
	cur := left
	for j := left; j < right; j++ {
		if arr[j] < p {
			arr[cur], arr[j] = arr[j], arr[cur]
			cur++
		}
	}

	arr[cur], arr[right] = arr[right], arr[cur]
	helper(arr, left, cur-1)
	helper(arr, cur+1, right)
}

// @lc code=end

