/**
题目描述
在命令行输入如下命令：
xcopy /s c:\ d:\，
各个参数如下： 
参数1：命令字xcopy 
参数2：字符串/s
参数3：字符串c:\
参数4: 字符串d:\
请编写一个参数解析程序，实现将命令行各个参数解析出来。
*/
package main
 
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)
 
func main(){
    reader := bufio.NewReader(os.Stdin)
    str,err:=reader.ReadString('\n')
    if err!=nil{
        return
    }
    output:=strings.Fields(str)
    fmt.Println(len(output))
    for _,v:=range output{
        if v[0]=='"'{
            fmt.Println(v[1:len(v)-1])
        }else{
            fmt.Println(v)
        }
    }
}
