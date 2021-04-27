package simhash

import "github.com/yanyiwu/gojieba"

/*
 * WordWeight
 * 加入词典模式进行切分
 * @params s string,待分词原句
 * @return []string,分词结果
 */
func WordWeight(s string) []string {
	x := gojieba.NewJieba()
	defer x.Free()
	return x.Cut(s, true)
}
