/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 *
 * https://leetcode.cn/problems/minimum-path-sum/description/
 *
 * algorithms
 * Medium (69.93%)
 * Likes:    1631
 * Dislikes: 0
 * Total Accepted:    552.6K
 * Total Submissions: 790.2K
 * Testcase Example:  '[[1,3,1],[1,5,1],[4,2,1]]'
 *
 * 给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 * 
 * 说明：每次只能向下或者向右移动一步。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
 * 输出：7
 * 解释：因为路径 1→3→1→1→1 的总和最小。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：grid = [[1,2,3],[4,5,6]]
 * 输出：12
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * m == grid.length
 * n == grid[i].length
 * 1 <= m, n <= 200
 * 0 <= grid[i][j] <= 200
 * 
 * 
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = grid[0][i] + dp[i-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			top := math.MaxInt
			left := math.MaxInt
			if i > 0 {
				top = min(top, dp[j])
			}
			if j > 0 {
				left = min(left, dp[j-1])
			}
			dp[j] = min(top, left) + grid[i][j]
		}
		fmt.Println(dp)
	}
	return dp[len(dp)-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
// @lc code=end

