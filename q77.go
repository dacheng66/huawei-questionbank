/**
题目描述
题目标题：
将两个整型数组按照升序合并，并且过滤掉重复数组元素[注: 题目更新了。输出之后有换行]
*/
package main
 
import (
    "fmt"
    "sort"
)
 
func main() {
    for {
        var m, n, v int
        var arry1, arry2 []int
        a, err := fmt.Scan(&m)
        if a == 0 || err != nil {
            break
        }
        for i := 0; i < m; i++ {
            a, err = fmt.Scan(&v)
            if a == 0 || err != nil {
                break
            }
            arry1 = append(arry1, v)
        }
        a, err = fmt.Scan(&n)
        if a == 0 || err != nil {
            break
        }
        for i := 0; i < n; i++ {
            a, err = fmt.Scan(&v)
            if a == 0 || err != nil {
                break
            }
            arry2 = append(arry2, v)
        }
        sort.Ints(arry1)
        sort.Ints(arry2)
        CombineBySort(arry1, arry2)
    }
}
 
func CombineBySort(arry1, arry2 []int) {
    //fmt.Printf("arry1: %+v, arry2: %+v", arry1, arry2)
    i := 0
    j := 0
    for i < len(arry1) && j < len(arry2) {
        if arry1[i] == arry2[j] {
            fmt.Print(arry1[i])
            i++
            j++
        } else if arry1[i] < arry2[j] {
            fmt.Print(arry1[i])
            i++
        } else {
            fmt.Print(arry2[j])
            j++
        }
    }
    if i == len(arry1) {
        for j < len(arry2) {
            fmt.Print(arry2[j])
            j++
        }
    } else if j == len(arry2) {
        for i < len(arry1) {
            fmt.Print(arry1[i])
            i++
        }
    }
    fmt.Println()
}
