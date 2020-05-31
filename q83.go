/**
题目描述
找出给定字符串中大写字符(即'A'-'Z')的个数
接口说明
原型：int CalcCapital(String str);
返回值：int
*/
package main
 
import (
    "bufio"
    "fmt"
    "os"
)
 
func CalcCapital(str string) int {
    var counts int
    for i := 0; i < len(str); i++ {
        if str[i] >= 'A' && str[i] <= 'Z' {
            counts += 1
        }
    }
    return counts
}
 
func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        input, _, _ := reader.ReadLine()
        if len(input) == 0 {
            break
        }
        fmt.Println(CalcCapital(string(input)))
    }
}