package main

import "fmt"

// 并查集
// 并查集用一维数组来表示
type Union struct {
    elements []int
}

func (u *Union) initUnion(size int) {
    u.elements = make([]int, size)
    for i := 0; i < size; i++ {
        u.elements[i] = -1
    }
}

// 找到一个数的根
func (u *Union) find(n int) int {
    origin := n
    // 找到根
    for u.elements[n] != -1 {
        n = u.elements[n]
    }
    // 把查找过程经过的节点都连接到根上
    for origin != n {
        tmp := origin
        origin = u.elements[origin]
        u.elements[tmp] = n
    }
    return n
}

func (u *Union) union(x, y int) {
    rootX := u.find(x)
    rootY := u.find(y)
    if rootX != rootY {
        u.elements[rootX] = rootY
    }
}

//统计总共形成了多少个关系网
func (u *Union) count() int {
    count := 0
    for i := 0; i < len(u.elements); i++ {
        if u.elements[i] == -1 {
            count++
        }
    }
    return count
}

func (u *Union) print() {
    for i := 0; i < len(u.elements); i++ {
        fmt.Println(u.elements[i])
    }
}

func main() {
    u := new(Union)
    u.initUnion(10)
    u.union(0, 1);
    u.union(1, 2);
    u.union(2, 3);
    u.union(3, 4);
    u.union(4, 5);
    u.union(5, 6);
    u.union(6, 7);
    u.union(7, 8);
    u.union(0, 9);
    fmt.Println(u.count())
}

