/**
题目描述
在计算机中，由于处理器位宽限制，只能处理有限精度的十进制整数加减法，比如在32位宽处理器计算机中，
参与运算的操作数和结果必须在-2^31~2^31-1之间。如果需要进行更大范围的十进制整数加法，需要使用特殊
的方式实现，比如使用字符串保存操作数和结果，采取逐位运算的方式。如下：
9876543210 + 1234567890 = ?
让字符串 num1="9876543210"，字符串 num2="1234567890"，结果保存在字符串 result = "11111111100"。
-9876543210 + (-1234567890) = ?
让字符串 num1="-9876543210"，字符串 num2="-1234567890"，结果保存在字符串 result = "-11111111100"。
*/
package main
  
import (
    "fmt"
)
  
// 两个字符相加
// 返回和的各位数以及是否进位
func add(ch1, ch2 byte, carry int) (int, int) {
    res := ch1 - '0' + ch2 - '0'
    if carry == 1 {
        res++
    }
    if res >= 10 {
        return int(res-10), 1
    }
    return int(res), 0
}
  
func main() {
    var (
        str1 string
        str2 string
    )
      
    for {
        _, err := fmt.Scan(&str1)
        if err != nil {
            break
        }
        fmt.Scan(&str2)
          
        if (len(str1) > len(str2)) {
            str1, str2 = str2, str1 // 第二个字符串存更长的
        }
        minl := len(str1)
        maxl := len(str2)
        // fmt.Println("str1, str2: ", str1, str2)
   
        result := make([]int, maxl) // 从各位开始存
        carry := 0
        re := 0
        for i := 0; i < minl; i++ {
            ix1 := minl - i - 1 // 第一个字符串的索引
            ix2 := maxl - i - 1
            re, carry = add(str1[ix1], str2[ix2], carry)
            // fmt.Println("re: ", re)
            result[i] = re
        }
        // fmt.Println("carry 1: ", carry)
        for i := minl; i < maxl; i++ {
            ix := maxl - i - 1
            re, carry = add(str2[ix], '0', carry)
            // fmt.Println("ix: ", ix, re)
            result[i] = re
        }
        if carry == 1 {
            fmt.Printf("%d", 1)
        }
          
        for i := maxl-1; i >= 0; i-- {
            fmt.Printf("%d", result[i])
        }
        fmt.Println()
    }
}
