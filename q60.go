/**
题目描述
一个DNA序列由A/C/G/T四个字母的排列组合组成。G和C的比例（定义为GC-Ratio）是序列中G和C两个字母的总的出现次数除以总的字母数目
（也就是序列长度）。在基因工程中，这个比例非常重要。因为高的GC-Ratio可能是基因的起始点。
给定一个很长的DNA序列，以及要求的最小子序列长度，研究人员经常会需要在其中找出GC-Ratio最高的子序列。
*/
package main
 
import (
    "fmt"
)
 
func main () {
    for {
        str := ""
        num := 0
        n,_ := fmt.Scan(&str, &num)
        if n == 0 {
            break
        }
        arr := []rune(str)
        max := 0
        index := 0
        for i := 0; i < len(str)-num; i++ {
            count := 0
            for j := i; j < i+num; j++ {
                if arr[j] == 'G' || arr[j] == 'C' {
                    count++
                }
            }
            if count > max {
                max = count
                index = i
            }
        }
        fmt.Println(string(arr[index:index+num]))
    }
}