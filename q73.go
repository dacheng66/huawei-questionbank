/**
题目描述
验证尼科彻斯定理，即：任何一个整数m的立方都可以写成m个连续奇数之和。
例如：
1^3=1 
2^3=3+5 
3^3=7+9+11 
4^3=13+15+17+19 
*/

package main
 
import(
    "fmt"
    "strconv"
)
func main(){
    for{
        var inputInt int
        _, err := fmt.Scan(&inputInt)
        if err != nil{
            return
        }
        fmt.Println(comF(inputInt))
    }
}
 
//数学公式直接可以推出来首项是m*m+1-m,有m项
func comF(a int) string{
    b := make([]int, a)
    b[0] = a * a - a + 1
    str := strconv.Itoa(b[0])
 
    for i := 1; i < a; i++{
        b[i] = b[i - 1] + 2
        str = str + "+" + strconv.Itoa(b[i])
    }
    return str
}