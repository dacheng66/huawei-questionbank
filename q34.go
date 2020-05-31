/**
题目描述
题目说明
蛇形矩阵是由1开始的自然数依次排列成的一个矩阵上三角形。
样例输入
5
样例输出
1 3 6 10 15
2 5 9 14
4 8 13
7 12
11
*/

package main
 
import "fmt"
 
func main() {
    var n int
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
 
        a := 1
        b:=0
 
        for i := 1; i <= n; i++ {
            a += i - 1
            b = a
            for j := 1; j <= n-i+1; j++ {
                if j == n-i+1 {
                    fmt.Printf("%d", b)
                } else {
                    fmt.Printf("%d ", b)
                }
                b += i + j
            }
            fmt.Println()
        }
    }
}