package joinString

import (
	"fmt"
	"math/rand"
	"strings"
)

// 结合两个 string  https://leetcode.cn/problems/merge-strings-alternately/
func mergerAlternately(word1 string, word2 string) string {
	// 按照顺序分别添加字母即可，关键是对对象的操作
	// 长度提前获取，取决于使用频率
	n, m := len(word1), len(word2)
	// 返回结果是 string , 操作细度是字符类型的话，就需要从byte入手
	ans := make([]byte, 0, n+m)

	// 将两个字符合并，同时加上剩余字符==》 并集计算
	for i := 0; i < n || i < m; i++ {
		// 使用if控制逻辑
		if i < n {
			ans = append(ans, word1[i])
		}
		if i < m {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

// 一般形式，但是可以控制的字符串必须得是整串，不能够是但单个字符
func commonMerge(word1 string, word2 string) (string, string) {
	// + 对象必须是两个string
	ans1 := word1 + word2

	// printf 构造字符串格式
	ans2 := fmt.Sprintf("%s%s", word1, word2)

	return ans1, ans2
}

// 依据传入的数字决定要生成多少个不同的字符
func generateRandom(n int, num int) []string {
	const leeterByte = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	ans := make([]string, num)
	// range 遍历
	for j := range ans {
		s := make([]byte, 0)
		for i := range s {
			s[i] = leeterByte[rand.Intn(len(leeterByte))]
		}
		ans[j] = string(s)
	}
	return ans
}

// 性能最好的 strings.Builder
func builderMerge(word1 string, word2 string) []string {
	var builder strings.Builder
	ans := make([]string, 2)
	builder.WriteString(word1)
	builder.WriteString(word2)
	ans[0] = builder.String()

	builder.WriteString(word1[:len(word1)-1])
	builder.WriteString(word2[:len(word2)-1])
	ans[1] = builder.String()
	return ans
}
