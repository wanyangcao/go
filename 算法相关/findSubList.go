package main

import "fmt"

/**
     * T: 和为S的连续正数序列
     *
     * 题目描述
     * 小明很喜欢数学,有一天他在做数学作业时,要求计算出9~16的和,他马上就写出了正确答案是100。
     * 但是他并不满足于此,他在想究竟有多少种连续的正数序列的和为100(至少包括两个数)。
     * 没多久,他就得到另一组连续正数和为100的序列:18,19,20,21,22。
     * 现在把问题交给你,你能不能也很快的找出所有和为S的连续正数序列? Good Luck!
     *
     * 输出描述:
     * 输出所有和为S的连续正数序列。序列内按照从小至大的顺序，序列间按照开始数字从小到大的顺序
     */
// 穷举法，最容易想到，效率最低
// O(n^2)
func findSubArr(arr []int, s int) [][]int {
    var result [][]int
    if len(arr) <= 3 || s == 0 {
        return result
    }
    for i := 0; i < len(arr); i++ {
        num := s
        var tmp []int
        for j := i; j < len(arr); j++ {
            tmp = append(tmp, arr[j])
            num -= arr[j]
            if num == 0 {
                if len(tmp) < 2 {
                    break
                }
                result = append(result, tmp)
                break
            } else {
                continue
            }
        }
    }
    return result
}


// 设置两个指针small和big，分别指向当前序列中最小和最大的数
// 当前序列和如果小于s,则移动big指针
// 当前序列和大于s,则缩小区间范围，移动small指针

func findSubArr2(arr []int, s int) [][]int {
    var result [][]int
    if len(arr) <= 3 || s == 0 {
        return result
    }
    small, big := 1, 2
    total := 0
    for small < big {
        // 等差数列求和 s = (small + big) * (big - small + 1) / 2
        total = (small + big) * (big - small + 1) / 2
        if total == s {
            var tmp []int
            left := small
            for left <= big {
                tmp = append(tmp, left)
                left++
            }
            result = append(result, tmp)
            small++
            big++
        } else if total < s {
            big++
        } else {
            small++
        }
    }
    return result
}



func main() {
    a := []int{1,2,3,4,5}
    fmt.Println(findSubArr(a, 5))
    fmt.Println(findSubArr2(a, 5))
}




