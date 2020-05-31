/**
题目描述
给出一个名字，该名字有26个字符串组成，定义这个字符串的“漂亮度”是其所有字母“漂亮度”的总和。
每个字母都有一个“漂亮度”，范围在1到26之间。没有任何两个字母拥有相同的“漂亮度”。字母忽略大小写。
给出多个名字，计算每个名字最大可能的“漂亮度”。
*/
package main
  
import (
    "fmt"
    "strings"
    "sort"
)
  
func main(){
    for{
        var num int
        _,err:=fmt.Scanf("%d",&num)
        if err!=nil{
            return
        }
    for i:=0;i<num;i++{
        var str string
      
        _,err:=fmt.Scanf("%s",&str)
        if err!=nil{
            return
        }
        str=strings.ToLower(str)
        var strMap =make([]int,26)
        for _,value:=range str{
            strMap[value-'a']++
        }
        maxNum:=26
        sum:=0
        sort.Ints(strMap)
        //fmt.Println(strMap)
        for i:=len(strMap)-1;i>=0;i--{
            sum+=maxNum*strMap[i]
            maxNum--
        }
        fmt.Println(sum)
    }
    }
      
}