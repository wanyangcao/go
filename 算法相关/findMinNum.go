package main

import (
    "fmt"
    "strings"
)

// 一个多位数字，删除k位之后使得剩下的数最小
// 常規解法，暴力搜索
// O(nk)
func forceFindMin(num string, k int) string{
    if len(num) == 0 || len(num) < k {
        return "0"
    }
    s := []rune(num)
    hasCut := false
    for i := 0; i < k; i++ {
        for j := 0; j < len(s); j++ {
            // 每次删除第一个出现左边元素大于右边元素的元素
            if s[j] > s[j + 1] {
                s = append(s[:j], s[j + 1 :]...)
                hasCut = true
                break
            }
        }
        // 如果没找到要删除的数，删除最后1位
        if !hasCut {
            s = s[:len(s) - 1]
        }
    }
    if len(s) == 0 {
        return "0"
    }
    return string(s)
}

// 優化方案
// 额外使用一个数组用来存放结果
// O(n)
func findMin(num string, k int) string {
    if len(num) == 0 || len(num) < k {
        return "0"
    }
    s := []rune(num)
    result := make([]string, len(s))
    top := 0
    for i := 0; i < len(s); i++ {
        c := string(s[i])
        for top > 0 && strings.Compare(result[top - 1], c) == 1 && k > 0{
            top -= 1
            k -= 1
        }
        result[top] = c
        top++
    }
    return strings.Join(result, "")
}


func main() {
    s := "56739"
    fmt.Println(forceFindMin(s, 1))
    fmt.Println(findMin(s, 1))
}
