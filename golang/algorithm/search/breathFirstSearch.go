package search

import "learn/golang/algorithm/dataStruct/queue"

// 图
type Graph map[string][]string

/*广度优先搜索
Desc:广度优先搜索，按照关系近远，先判断一度关系，再进行二度关系判断
In:
	g	*Graph,底层为map[string][]string,该方法的所属对象
	v	string,查询值
Return:
	bool,返回寻值结果，存在则返回true ,反之返回false
*/
func (g Graph) BreathFirstSearch(v string) bool {
	q := &queue.Queue{}
	var search []string
	// 一度关系填充
	for k := range g {
		q.Push(k)
	}
	for q.Len() > 0 {
		if eStr := q.Pop().(string); !exists(eStr, search) {
			if eStr == v {
				return true
			}
			// 当未匹配上是，加入已搜索列表
			search = append(search, eStr)
			// 将相邻数据加入待搜索队列
			for _, v := range g[eStr] {
				q.Push(v)
			}
		}
	}
	return false
}

/*判断值在数组中是否存在
Desc:遍历数组查看值是否存在
In:
	v string,查找值
	l []string,归属于字符串数组，搜索数组
Return:
	bool,返回是否存在，是则返回true，反之返回false
*/
func exists(v string, l []string) bool {
	for _, r := range l {
		if r == v {
			return true
		}
	}
	return false
}
