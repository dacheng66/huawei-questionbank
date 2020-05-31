/**
题目描述
输入一行字符，分别统计出包含英文字母、空格、数字和其它字符的个数。
*/
package main
  
import (
    "fmt"
    "bufio"
    "os"
)
  
func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        input, _, _ := reader.ReadLine()
        if len(input) == 0 {
            break
        }
          
        v1, v2, v3, v4 := 0, 0, 0, 0
        for _, ch := range input {
            if ch == ' ' {
                v2++
            } else if ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' {
                v1++
            } else if ch >= '0' && ch <= '9' {
                v3++
            } else {
                v4++
            }
        }
        fmt.Println(v1)
        fmt.Println(v2)
        fmt.Println(v3)
        fmt.Println(v4)
    }
}
