/**
题目描述
原理：ip地址的每段可以看成是一个0-255的整数，把每段拆分成一个二进制形式组合起来，然后把这个二进制数转变成
一个长整数。
举例：一个ip地址为10.0.3.193
每段数字             相对应的二进制数
10                   00001010
0                    00000000
3                    00000011
193                  11000001
组合起来即为：00001010 00000000 00000011 11000001,转换为10进制数就是：167773121，即该IP地址转换后的数字就是它了。
的每段可以看成是一个0-255的整数，需要对IP地址进行校验
*/
package main
 
import (
    "fmt"
    "strconv"
    "strings"
)
 
func ipToNum(str string) {
    strs := strings.Split(str, ".")
    a, _ := strconv.Atoi(strs[0])
    b, _ := strconv.Atoi(strs[1])
    c, _ := strconv.Atoi(strs[2])
    d, _ := strconv.Atoi(strs[3])
    fmt.Println((a << 24) + (b << 16) + (c << 8) + d)
}
 
func numToIP(str string) {
    num, _ := strconv.Atoi(str)
    a, b, c, d := num>>24, num>>16&0x00ff, num>>8&0x0000ff, num&0x000000ff
    fmt.Printf("%d.%d.%d.%d\n", a, b, c, d)
}
 
func main() {
    for {
        var str string
        _, err := fmt.Scan(&str)
        if err != nil {
            return
        }
 
        if strings.Index(str, ".") == -1 {
            numToIP(str)
        } else {
            ipToNum(str)
        }
    }
}	