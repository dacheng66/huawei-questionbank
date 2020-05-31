/**
题目描述
输入n个整数，输出其中最小的k个。
*/
package main
 
import (
    "fmt"
    "sort"
)
 
func main() {
    for {
        var n, m, t int
        var arr []int
        _, err := fmt.Scan(&n, &m)
        if err != nil {
            return
        }
 
        for i := 0; i < n; i++ {
            fmt.Scan(&t)
            arr = append(arr, t)
        }
 
        sort.Ints(arr)
        r := fmt.Sprintf("%v\n", arr[:m])
        fmt.Println(r[1 : len(r)-2])
    }
}