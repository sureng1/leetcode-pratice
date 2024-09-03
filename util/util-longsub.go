/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 *
 * https://leetcode.cn/problems/longest-palindromic-substring/description/
 *
 * algorithms
 * Medium (37.93%)
 * Likes:    7025
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 4.2M
 * Testcase Example:  '"babad"'
 *
 * 给你一个字符串 s，找到 s 中最长的回文子串。
 *
 * 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "babad"
 * 输出："bab"
 * 解释："aba" 同样是符合题意的答案。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "cbbd"
 * 输出："bb"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 仅由数字和英文字母组成
 *
 *
 */

// @lc code=start
func longestPalindrome(s string) string {
	ansi := 0
	ansj := 0
	isPa := make([][]bool, len(s))
	for i := range isPa {
		isPa[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		isPa[i][i] = true
		// if i+1 < len(s) {
		// 	isPa[i][i+1] = true
		// }
	}

	for l := 1; l < len(s); l++ { //j?
		for i := 0; i+l < len(s); i++ {
			j := i + l
			// if l == 4 {
			// 	fmt.Println(s[i : j+1])
			// }
			if s[i] == s[j] && (j-i == 1 || isPa[i+1][j-1]) {
				isPa[i][j] = true
				// fmt.Println(s[i : j+1])
				if l > (ansj - ansi) { // ?
					ansi = i
					ansj = j
				}
			}
		}
	}

	return s[ansi : ansj+1] // j?
}

// @lc code=end

