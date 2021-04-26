package greedy

// 该包用于存放贪心算法问题

/* 集合覆盖问题
电台覆盖若干地区
@param states []string,所有地区集合
@param station map[string][]string,电台覆盖的地区情况，散列表
@return []string ,电台选择集合
*/
func StationCover(states []string, stations map[string][]string) (res []string) {
	// 基准条件
	if len(states) == 0 {
		return []string{}
	}
	var bestStation string // 最佳电台选择
	var bestCover int      // 最佳电台覆盖数
	for station, statesCover := range stations {
		addCover := insersect(states, statesCover)
		if bestStation == "" || bestCover < len(addCover) {
			bestStation = station
			bestCover = len(addCover)
		}
	}
	// 添加选择，并计算差集
	res = append(res, bestStation)
	states = difference(stations[bestStation], states)
	delete(stations, bestStation)

	// 递归
	next := StationCover(states, stations)

	return append(res, next...)
}

/*
 * 无重复集合计算交集
 * @param s1 []string,无重复字符串切片
 * @param s2 []string,无重复字符串切片
 * @return []string,交集
 * 使用双循环计算差集，O(2n)
 */
func insersect(s1, s2 []string) []string {
	m := make(map[string]struct{})
	res := make([]string, 0)
	for _, v := range s1 {
		m[v] = struct{}{}
	}
	for _, v := range s2 {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
	}
	return res
}

/*
 * 无重复集合计算差集
 * @param s1 []string,无重复字符串切片,全集
 * @param s2 []string,无重复字符串切片,子集
 * @return []string,差集
 * 使用双循环计算差集，O(2n)
 */
func difference(s1, s2 []string) []string {
	m := make(map[string]struct{})
	res := make([]string, 0)
	for _, v := range s1 {
		m[v] = struct{}{}
	}

	for _, v := range s2 {
		// 计算差集
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}
