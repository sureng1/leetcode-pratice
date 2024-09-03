/*
 * @lc app=leetcode.cn id=179 lang=golang
 *
 * [179] 最大数
 *
 * https://leetcode.cn/problems/largest-number/description/
 *
 * algorithms
 * Medium (40.97%)
 * Likes:    1229
 * Dislikes: 0
 * Total Accepted:    215.1K
 * Total Submissions: 524.9K
 * Testcase Example:  '[10,2]'
 *
 * 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
 * 
 * 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [10,2]
 * 输出："210"
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,30,34,5,9]
 * 输出："9534330"
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 100
 * 0 <= nums[i] <= 10^9
 * 
 * 
 */

// @lc code=start

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	sort.Slice(nums, func(i, j int) bool {
		ai := fmt.Sprintf("%d%d", nums[i], nums[j])
		bi := fmt.Sprintf("%d%d", nums[j], nums[i])
		for k := 0; k < len(ai) && k < len(bi); k++ {
			if ai[k] == bi[k] {
				continue
			}
			val := ai[k] > bi[k]
			return val
		}
		return false
	})
	if nums[0] == 0 {
		return "0"
	}
	stb := strings.Builder{}
	for _, si := range nums {
		stb.WriteString(fmt.Sprint(si))
	}
	return stb.String()
}
// @lc code=end

