package search

// 节点
type Point map[string]int

// 图
type Graph map[string]Point

/** 狄克斯特拉算法，加权图算法，计算最低权重的路线
@param g *map[string]map[string]int,加权图
@parpm dsr string,目的地
@result int,最小开销
*/
// func Dijkstra(g *Graph, dsr string) int {
// 	// 最小节点开销
// 	costs := make(map[string]int)

// 	// 父节点散列表
// 	parents := make(map[string]string)

// 	// 记录处理过的节点
// 	var processed []string

// }
