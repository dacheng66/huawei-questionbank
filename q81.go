/**
题目描述
分子为1的分数称为埃及分数。现输入一个真分数(分子比分母小的分数，叫做真分数)，请将该分数分解为埃及分数。如：8/11 = 1/2+1/5+1/55+1/110。
*/
package main
 
import (
    "fmt"
)
 
func main() {
    var num1 int
    var num2 int
    for {
        n, err := fmt.Scanf("%d/%d", &num1, &num2)
        if n == 0 || err != nil {
            //fmt.Println(n)
            // fmt.Println(err)
            break
        }
        ret := make([]int, 0)
        first := true
        for {
            if num2%num1 == 0 && !first {
                ret = append(ret, num2/num1)
                break
            }
            // 类似这种 3/110 = 2/110 + 1/110,直接就可以结束了
            if num2%(num1-1) == 0{
                ret = append(ret,num2/(num1-1))
                ret = append(ret,num2)
                break
            }
             
             
            //新的分母
            base := num2/num1 + 1
            // 1/ (num2/num1+1) 这其实就是最靠近num1/num2，并且分子为1的数
            ret = append(ret, base)
            num1 = num1*base - num2
            num2 = num2 * base
             
            first = false
        }
        for i := 0; i < len(ret)-1; i++ {
            fmt.Printf("1/%d+", ret[i])
        }
 
        fmt.Printf("1/%d", ret[len(ret)-1])
        fmt.Println()
         
    }
}
