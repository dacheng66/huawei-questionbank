/**
题目描述
编写一个截取字符串的函数，输入为一个字符串和字节数，输出为按字节截取的字符串。但是要保证汉字不被截半个，
如"我ABC"4，应该截为"我AB"，输入"我ABC汉DEF"6，应该输出为"我ABC"而不是"我ABC+汉的半个"。 
*/
package main
import (
    "fmt"
)
func main(){
    var str string
    var num int
    for {
        _,err:=fmt.Scanf("%s",&str)
        if err!=nil{
            return
        }
        _,err=fmt.Scanf("%d",&num)
        if err!=nil{
            return
        }
         
        fmt.Println(str[:num])
    }
}