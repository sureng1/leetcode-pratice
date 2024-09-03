/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 *
 * https://leetcode.cn/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.36%)
 * Likes:    2706
 * Dislikes: 0
 * Total Accepted:    846.4K
 * Total Submissions: 1.2M
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target
 * 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
 * 
 * candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 
 * 
 * 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：candidates = [2,3,6,7], target = 7
 * 输出：[[2,2,3],[7]]
 * 解释：
 * 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
 * 7 也是一个候选， 7 = 7 。
 * 仅有这两种组合。
 * 
 * 示例 2：
 * 
 * 
 * 输入: candidates = [2,3,5], target = 8
 * 输出: [[2,2,2,2],[2,3,3],[3,5]]
 * 
 * 示例 3：
 * 
 * 
 * 输入: candidates = [2], target = 1
 * 输出: []
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= candidates.length <= 30
 * 2 <= candidates[i] <= 40
 * candidates 的所有元素 互不相同
 * 1 <= target <= 40
 * 
 * 
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	choice := make([]int, len(candidates))
	var find func(i int, sum int)
	var part int
	find = func(i int, sum int) {
		if i >= len(candidates) {
			return
		}
		j := 0
		for {
			choice[i] = j
			part = sum + candidates[i]*j
			j++
			if part > target {
				choice[i] = 0
				break
			}
			if part == target {
				var temp []int
				for k, ci := range choice {
					if ci == 0 {
						continue
					}
					for l := 0; l < ci; l++ {
						temp = append(temp, candidates[k])
					}
				}
				if len(temp) > 0 {
					ans = append(ans, temp)
				}
				choice[i] = 0
				break
			}
			find(i+1, part)
		}
	}
	find(0, 0)
	return ans
}
// @lc code=end

