/**
题目描述
把M个同样的苹果放在N个同样的盘子里，允许有的盘子空着不放，问共有多少种不同的分法？（用K表示）5，1，1和1，5，1 是同一种分法。
*/
package main
  
import(
    "fmt"
)
  
func demo(m,n int) int{
    if m==0||n==1{
        return 1
    } 
    if m<n{
        return demo(m,m)
    }
    return demo(m,n-1)+demo(m-n,n)
}
  
func main(){
    for{
        var m,n int
        _,err:=fmt.Scan(&m)
        if err!=nil{
            return
        }
        fmt.Scan(&n)
        fmt.Println(demo(m,n))
    }
}