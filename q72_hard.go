/**
题目标题：
计算两个字符串的最大公共字串的长度，字符不区分大小写
*/
package main
import (
    "fmt"
    "strings"
)
  
func main(){
    for{
        var s1, s2 string
        n, _ := fmt.Scan(&s1,&s2)
        if n != 2 {
            break
        }
  
        var max int
        for i:=0;i<len(s1)-1;i++{
            for j := i+1;j<=len(s1);j++{
                if strings.Contains(s2, s1[i:j]) && max < j-i{
                    max = j-i
                }
            }
        }
        fmt.Println(max)
    }
}