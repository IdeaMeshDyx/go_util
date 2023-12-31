# 字串查找匹配

## Problem

解决的问题和关注的点是字串查找匹配，如果有两个（多个字符串）其中一个字符串是另一个（些）字符串的字串，也就是 A 包含了 B 所有的字母

这个问题中可以关注的问题是：
判断是否是： 判断是否存在字串的关系
查找字串位置： 如果有，长的字符当中第几个字母开始是字串
查找最长字串： 判断一系列字串关系，找出其中最长的那个
计算字串数量： 计算存在字串关系的数量
等等

问题的暴力解法是：
双重遍历两个字符串，一个做为外层循环一个内层循环，每匹配到一个字符串就进入到内层的匹配流程

但这样子就需要多次遍历较短的模式字符串，总共遍历的次数是：m*n 有较高的性能消耗，但是仔细观察会发现在被匹配的字符串当中，存在着非常多的重复字串，他们与匹配字符串大可能只相差一个字符或者几个字符
这样遍历的方式实际上是做了非常多的重复性的工作

所以这个问题的优化方向也就在于： **如何利用已经比较过的流程来简化，也就是之前比较过了就不需要再比较一次**

这也是非常典型的问题优化问题，属于对过去问题的问题集划分并找出重复解的部分，重复利用这些重复解来提高工作效率


## KMP 算法

youtube 上的印度人视频很有借鉴意义，具体的地址是在：

该算法的核心是尽可能地扩大先前匹配字符串的结果，如果是暴力解法，他的时间复杂度是 O(M*N) M与N是两个字符串的长度，实际上是M中每有一个匹配的字符，就会遍历一次N，所以实际上匹配重复操作最多的就是对 N 中重复且能够匹配 M 的字符部分。

接下来本篇将以如下顺序来讲解KMP，

### 1. 什么是 KMP

因为是由这三位学者发明的：Knuth，Morris和Pratt，所以取了三位学者名字的首字母。所以叫做KMP

### 2. KMP 可以解决什么问题

KMP主要应用在字符串匹配上。主要思想是当出现字符串不匹配时，可以知道一部分之前已经匹配的文本内容，可以利用这些信息避免从头再去做匹配了。
从而，也说明这个算法的核心是空间换时间，通过某种方式记录先前匹配的一些结果，然后在后面的匹配过程中直接跳过已经匹配过的字符串。要保证这个行为可以执行的话，首先就得是一个合理的数据结构去存储先前匹配的结果，然后这个结果能够控制接下来匹配过程中的遍历 i 或者 j 减少或者跳过。

### 3. 什么是前缀表---》 next 数组

这个前缀表就是之前说到的存储结构，那他的作用其实就如同前文所说，这里的关键是该怎样去获得他呢？

前缀表是一个 [int]byte 类型的数组[方便理解的]，完全也可以就存储一个 [int] 数组然后下标和要去匹配的字符串[模式串]相对应

前缀表的计算方式是核心：[int] 中的元素表示，与其下标相同的模式串字母相同的自模式串头开始的字母所在位置下标。

这样可以做到什么事情呢？就是在匹配过程中，有匹配到相同的字母，且这个字母的下一个不相同，那么就可以直接回溯到字符串的上一个与这个字母相同的位置，就不用再从头开始匹配

【待丰富】

### 4. 计算前缀表
先记住，再体会
```go
func getnext(next []int, s string){
    // pos 表示与目前比较字母相同的上一个字母的下标 + 1 ====》跳转之后正式开始比较的字母
    pos := 0
    // 前缀表
    next[0] = pos

    // 遍历模式串
    for i := 0; i < len(s); i++ {
        // pos>0指当前位置字母在之前有重复，s[i] != s[pos]指 s[i-1] 与 s[pos-1]是相同的，在前缀表中计算出 s[i-1]的值就是 pos
        // 判断条件是： 只要前面还有重复的字符串，就一直回退
        for pos > 0 && s[i] != s[pos]{
            pos = next[pos]
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
```