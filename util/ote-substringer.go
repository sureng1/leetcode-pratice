/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (45.51%)
 * Likes:    2807
 * Dislikes: 0
 * Total Accepted:    507.7K
 * Total Submissions: 1.1M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 * 
 * 
 * 
 * 注意：
 * 
 * 
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 解释：整个字符串 s 是最小覆盖子串。
 * 
 * 
 * 示例 3:
 * 
 * 
 * 输入: s = "a", t = "aa"
 * 输出: ""
 * 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * ^m == s.length
 * ^n == t.length
 * 1 <= m, n <= 10^5
 * s 和 t 由英文字母组成
 * 
 * 
 * 
 * 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
 */

// @lc code=start
// todo 用滑动窗口优化下。


func minWindow(s string, t string) string {
	set := map[byte]int{}
	for i := range t {
		set[t[i]]++
	}
	ansl, ansr := 0, 0
	ansLen := math.MaxInt
	ansArr := []int{} // idx of chars in t
	queues := map[byte]*Queue{}
	for ci := range set {
		queues[ci] = &Queue{}
	}
	hit := false
	for i := range s {
		ch := s[i]
		need := set[ch] != 0
		if !need {
			continue
		}
		// > need,pop, update ans if enough
		if queues[ch].Size() == set[ch] { // todo use arr
			idx := queues[ch].Pop()
			for i, idxi := range ansArr {
				if idxi == idx {
					ansArr = append(ansArr[:i], ansArr[i+1:]...)
					break
				}
			}
		}
		// push
		queues[ch].Push(i)
		ansArr = append(ansArr, i)
		if !hit {
			done := true
			for ch, size := range set {
				if queues[ch].Size() != size {
					done = false
					break
				}
			}
			hit = done
		}
		if hit {
			if ansLen > (ansArr[len(ansArr)-1] - ansArr[0]) {
				ansLen = ansArr[len(ansArr)-1] - ansArr[0]
				ansl = ansArr[0]
				ansr = ansArr[len(ansArr)-1]
			}
		}
	}
	if ansLen == math.MaxInt {
		return ""
	}
	return s[ansl : ansr+1]
}

type Queue []int

func (s *Queue) Push(a int) {
	(*s) = append(*s, a)
}
func (s *Queue) Pop() int {
	val := (*s)[0]
	(*s) = (*s)[1:]
	return val
}
func (s *Queue) Peek() int {
	return (*s)[0]
}
func (s *Queue) Size() int { return len(*s) }
func (s *Queue) Empty() bool {
	return len(*s) == 0
}
// @lc code=end

