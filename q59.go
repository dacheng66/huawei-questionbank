/**
输入描述:
输入一个整数

输出描述:
计算整数二进制中1的个数
*/

package main
  
import "fmt"
  
func main(){
    for{
        var a,b int
        _,err:=fmt.Scan(&a)
        if err!=nil{
            return
        }
        for a!=0 {
            if a%2==1 {
                b++
            }
            a = a>>1
        }
        fmt.Println(b)
    }
}