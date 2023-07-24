package substringsearch

func getNext(next []int, s string) {
	j := -1 // j表示 最长相等前后缀长度
	next[0] = j

	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j] // 回退前一位
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j // next[i]是i（包括i）之前的最长相等前后缀长度
	}
}

func getThenext(next []int, s string) {
	// pos 表示与目前比较字母相同的上一个字母的下标 + 1 ====》跳转之后正式开始比较的字母
	pos := 0
	// 前缀表
	next[0] = pos

	// 遍历模式串
	for i := 1; i < len(s); i++ {
		// pos>0指当前位置字母在之前有重复，s[i] != s[pos]指 s[i-1] 与 s[pos-1]是相同的，在前缀表中计算出 s[i-1]的值就是 pos
		// 判断条件是： 只要前面还有重复的字符串，就一直回退
		for pos > 0 && s[i] != s[pos] {
			pos = next[pos-1]
		}
		// == --》模式串内部，第i个位置与第pos位置相同，就继续比较
		if s[i] == s[pos] {
			pos++
		}
		// 经过上述判断，这里有两个情况：
		// 1. 回退过的走for 出来但是没进入if
		// 2. 正常走if++过来的
		next[i] = pos
	}
}

func strStr(haystack string, needle string) int {
	// 由于空字符串是所有字符串的子集，所以返回第一个位置就可i
	if len(needle) == 0 {
		return 0
	}
	// 创建 Next 数据，用以存储模式串 needle 的匹配结果
	next := make([]int, len(needle))
	getNext(next, needle)
	//
	pos := 0
	for i := 0; i < len(haystack); i++ {
		for pos > 0 && haystack[i] != needle[pos] {
			pos = next[pos-1]
		}
		if haystack[i] == needle[pos] {
			pos++
		}
		if pos == len(needle) { // j指向了模式串的末尾
			return i - len(needle) + 1
		}
	}
	return -1
}
