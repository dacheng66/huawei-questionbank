/**
题目描述
根据输入的日期，计算是这一年的第几天。。
详细描述：
输入某年某月某日，判断这一天是这一年的第几天？。
*/
package main
 
import "fmt"
 
var ping = [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var run = [...]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
 
func main() {
    var nian, yue, ri int
    for {
        _, err := fmt.Scan(&nian)
        if err != nil {
            break
        }
        fmt.Scan(&yue)
        fmt.Scan(&ri)
 
        var sum int
        if nian%4 == 0 {
            for i := 0; i < yue-1; i++ {
                sum += run[i]
            }
            fmt.Println(sum + ri)
        } else {
            for i := 0; i < yue-1; i++ {
                sum += ping[i]
            }
            fmt.Println(sum + ri)
        }
    }
}
