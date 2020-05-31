/**
题目描述
输出7有关数字的个数，包括7的倍数，还有包含7的数字（如17，27，37...70，71，72，73...）的个数（一组测试用例里可能有多组数据，请注意处理）
*/
package main
 
import (
    "fmt"
    "strconv"
    "strings"
)
 
func main() {
    for {
        var n int
        _, err := fmt.Scan(&n)
        if err != nil {
            return
        }
        var res int
        for i := 1; i <= n; i++ {
            if i%7 == 0 || strings.Index(strconv.Itoa(i), "7") != -1 {
                res++
            }
        }
        fmt.Println(res)
    }
}