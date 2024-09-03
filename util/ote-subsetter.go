/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 *
 * https://leetcode.cn/problems/subsets/description/
 *
 * algorithms
 * Medium (81.19%)
 * Likes:    2233
 * Dislikes: 0
 * Total Accepted:    722.3K
 * Total Submissions: 889.5K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
 *
 * 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,2,3]
 * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [0]
 * 输出：[[],[0]]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * -10
 * nums 中的所有元素 互不相同
 *
 *
 */

// @lc code=start
func subsets(nums []int) [][]int {
	// * 输入：nums = [1,2,3]
	// * 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
	ans := [][]int{}
	temp := []int{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx >= len(nums) {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		dfs(idx + 1)
		temp = append(temp, nums[idx])
		dfs(idx + 1)
		temp = temp[:len(temp)-1]
	}
	dfs(0)
	return ans
}

// @lc code=end

