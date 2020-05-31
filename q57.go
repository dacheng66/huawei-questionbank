/**
题目描述
任意一个偶数（大于2）都可以由2个素数组成，组成偶数的2个素数有很多种情况，本题目要求输出组成指定偶数的两个素数差值最小的素数对
*/

package main
 
import (
    "fmt"
)
 
func isPrime(n int) bool {
    for i := 2; i < n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
 
func main() {
    var n int
     
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
         
        half := n / 2
        for i := 0; i < half-1; i++ {
            v1, v2 := half - i, half + i
            if v1%2 == 0 || v2%2 == 0 {
                continue
            }
            if isPrime(v1) && isPrime(v2) {
                fmt.Printf("%d\n%d\n", v1, v2)
                break
            }
        }
    }
}