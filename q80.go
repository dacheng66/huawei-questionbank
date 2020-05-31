/**
题目标题：
判断短字符串中的所有字符是否在长字符串中全部出现
*/
package main
 
import (
    "bufio"
    "fmt"
    "os"
)
 
func main() {
    var short string
    var long string
 
    reader := bufio.NewReader(os.Stdin)
 
    for {
        bytes, _, err := reader.ReadLine()
        if err != nil {
            break
        }
 
        short = string(bytes)
 
        bytes, _, _ = reader.ReadLine()
        long  = string(bytes)
 
        maps := map[byte]int{}
        for i := 0; i < len(long) ; i++  {
            maps[long[i]] = 1
        }
 
        flag := true
        for i := 0; i < len(short); i++ {
            _, ok := maps[short[i]]
            if !ok {
                fmt.Println(false)
                flag = false
                break
            }
        }
 
        if flag == true {
            fmt.Println(true)
        }
    }
}
