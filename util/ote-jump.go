/*
 * @lc app=leetcode.cn id=815 lang=golang
 *
 * [815] 公交路线
 *
 * https://leetcode.cn/problems/bus-routes/description/
 *
 * algorithms
 * Hard (44.43%)
 * Likes:    365
 * Dislikes: 0
 * Total Accepted:    39.2K
 * Total Submissions: 88.3K
 * Testcase Example:  '[[1,2,7],[3,6,7]]\n1\n6'
 *
 * 给你一个数组 routes ，表示一系列公交线路，其中每个 routes[i] 表示一条公交线路，第 i 辆公交车将会在上面循环行驶。
 *
 *
 * 例如，路线 routes[0] = [1, 5, 7] 表示第 0 辆公交车会一直按序列 1 -> 5 -> 7 -> 1 -> 5 -> 7 -> 1
 * -> ... 这样的车站路线行驶。
 *
 *
 * 现在从 source 车站出发（初始时不在公交车上），要前往 target 车站。 期间仅可乘坐公交车。
 *
 * 求出 最少乘坐的公交车数量 。如果不可能到达终点车站，返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：routes = [[1,2,7],[3,6,7]], source = 1, target = 6
 * 输出：2
 * 解释：最优策略是先乘坐第一辆公交车到达车站 7 , 然后换乘第二辆公交车到车站 6 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：routes = [[7,12],[4,5,15],[6],[15,19],[9,12,13]], source = 15, target =
 * 12
 * 输出：-1
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 .
 * 1
 * routes[i] 中的所有值 互不相同
 * sum(routes[i].length)
 * 0
 * 0
 *
 *
 */

// @lc code=start
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	edge := make([]map[int]bool, len(routes))
	for i := range edge {
		edge[i] = map[int]bool{}
	}

	rec := map[int][]int{}
	for i, route := range routes {
		for _, si := range route {
			lines := rec[si]
			for _, li := range lines {
				edge[i][li] = true
				edge[li][i] = true
			}
			rec[si] = append(rec[si], i)
		}
	}

	var costs []int
	setCost := func() {
		costs = make([]int, len(edge))
		for pi := range edge {
			costs[pi] = math.MaxInt
		}
	}
	starts := rec[source]
	targets := rec[target]
	// fmt.Println(starts)
	// fmt.Println(targets)
	ans := math.MaxInt
	// for _, ai := range edge {
	// 	fmt.Println(ai)
	// }
	for _, si := range starts {
		setCost()
		points := []int{si}
		visited := make([]bool, len(edge))
		if costs[si] == math.MaxInt {
			costs[si] = 1
		}
		for len(points) > 0 { //?
			sourcePoint := points[0]
			points = points[1:]
			if visited[sourcePoint] {
				continue
			}
			visited[sourcePoint] = true
			for targetPoint := range edge[sourcePoint] {
				// fmt.Println(targetPoint)
				costs[targetPoint] = min(costs[targetPoint], costs[sourcePoint]+1)
				points = append(points, targetPoint)
			}
		}
		// fmt.Println(costs)
		for _, target := range targets {
			ans = min(ans, costs[target])
		}
	}

	if ans == math.MaxInt {
		return -1
	}

	return ans
}

// @lc code=end

