/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.53%)
 * Likes:    3486
 * Dislikes: 0
 * Total Accepted:    781.6K
 * Total Submissions: 1M
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */

// @lc code=start
func generateParenthesis(n int) []string {
	var ans []string
	part := make([]byte, n*2)
	var gen func(i, left, right int, part []byte)
	gen = func(i, left, right int, part []byte) {
		if i >= n*2 {
			if left == right {
				ans = append(ans, string(part)) // ? juge?
			}
			return
		}
		part[i] = '('
		gen(i+1, left+1, right, part)
		if left <= right {
			return
		}
		part[i] = ')'
		gen(i+1, left, right+1, part)

	}
	gen(0, 0, 0, part)
	return ans
}

// @lc code=end

// 可以用dfs 来解，没有考虑到解空间是固定的，搜索解空间即可。