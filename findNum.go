package main

import "fmt"

// 数组中有一个元素只出现一次，其他元素出现两次，找出出现一次的元素
func findOnceInArr(arr []int) int {
    var t int
    for i := 0; i < len(arr); i++ {
        t ^= arr[i]
    }
    return t
}

// 数组中有一个元素出现两次，其他元素只出现一次，找出出现两次的元素
func findTwiceArr(arr []int) int {
    var t int
    for i := 0; i < len(arr); i++ {
        t ^= arr[i]
    }
    for i := 0; i < len(arr); i++ {
        t ^= i
    }
    return t

}

// 找出数组中只出现一次的两个数
func findTwoNumTwice(arr []int) {
    var t int
    for i := 0; i < len(arr); i++ {
        t ^= arr[i]
    }
    // 根据异或计算规则，t肯定是数组中两个只出现一次的数的异或值
    // 根据该值的二进制的第一位出现1的位置，把原数组分为两组
    // 分别求出现一次的数
    num1, num2 := 0, 0
    index := findFirstBit(t)
    for i := 0; i < len(arr); i++ {
        if IsBit1(arr[i], uint(index)) {
            num1 ^= arr[i]
        } else {
            num2 ^= arr[i]
        }
    }
    fmt.Println(num1, num2)

}

func findFirstBit(num int) int {
    index := 0
    for num&1 == 0 && index < 32 {
        num = num >> 1
        index++
    }
    return index
}

func IsBit1(num int, index uint) bool {
    num = num >> index
    return num&1 == 1

}

func main() {
    a1 := []int{1,2,2,4,5,4,5}
    fmt.Println(findOnceInArr(a1))
    a2 := []int{1,1,2,3,4,5}
    fmt.Println(findTwiceArr(a2))
    a3 := []int{1,2,3,3,4,4,5,5}
    findTwoNumTwice(a3)
}

