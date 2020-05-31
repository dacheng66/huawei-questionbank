/**
题目描述
Lily上课时使用字母数字图片教小朋友们学习英语单词，每次都需要把这些图片按照大小（ASCII码值从小到大）排列收好。
请大家给Lily帮忙，通过C语言解决。
*/

package main
 
import (
    "fmt"
    "sort"
)
 
func main() {
    var s string
    for {
        n,err := fmt.Scan(&s)
        if n == 0 || err != nil {
            return
        }
        arr := []byte(s)
        arr1 := []int{}
        for _, b := range arr {
            arr1 = append(arr1, int(b))
        }
        sort.Ints(arr1)
        for _, b := range arr1 {
            fmt.Print(string(byte(b)))
        }
        fmt.Println()
    }
}