/**
题目描述
找出字符串中第一个只出现一次的字符
*/
package main
 
import (
    "fmt"
)
 
func main() {
    for {
        var str string
        var a bool
        _, err := fmt.Scan(&str)
        if err != nil {
            return
        }
        m := make(map[rune]int)
        for _, b := range str {
            m[b]++
        }
        for _, b := range str {
            if m[b] == 1 {
                fmt.Println(string(b))
                a = true
                break
            }
        }
        if !a {
            fmt.Println(-1)
        }
    }
}