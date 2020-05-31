/**
题目描述
有6条配置命令，它们执行的结果分别是：

命   令	执   行
reset	reset what
reset board	board fault
board add	where to add
board delet	no board at all
reboot backplane	impossible
backplane abort	install first
he he	unkown command
注意：he he不是命令。

为了简化输入，方便用户，以“最短唯一匹配原则”匹配：
1、若只输入一字串，则只匹配一个关键字的命令行。例如输入：r，根据该规则，匹配命令reset，执行结果为：reset what；输入：res，根据该规则，匹配命令reset，执行结果为：reset what；
2、若只输入一字串，但本条命令有两个关键字，则匹配失败。例如输入：reb，可以找到命令reboot backpalne，但是该命令有两个关键词，所有匹配失败，执行结果为：unkown command
3、若输入两字串，则先匹配第一关键字，如果有匹配但不唯一，继续匹配第二关键字，如果仍不唯一，匹配失败。例如输入：r b，找到匹配命令reset board 和 reboot backplane，执行结果为：unkown command。

4、若输入两字串，则先匹配第一关键字，如果有匹配但不唯一，继续匹配第二关键字，如果唯一，匹配成功。例如输入：b a，无法确定是命令board add还是backplane abort，匹配失败。
5、若输入两字串，第一关键字匹配成功，则匹配第二关键字，若无匹配，失败。例如输入：bo a，确定是命令board add，匹配成功。
6、若匹配失败，打印“unkown command”
*/
package main
 
import (
    "fmt"
    "bufio"
    "strings"
    "os"
)
var commands=[6][2]string{
    {"reset"},
    {"reset","board"},
    {"board","add"},
    {"board","delet"},
    {"reboot","backplane"},
    {"backplane","abort"},
}
func main(){
     reader := bufio.NewReader(os.Stdin)
      
    for {
        input, _, _ := reader.ReadLine()
        if len(input) == 0 {
            break
        }
        str := string(input)
        i  := compare(strings.Fields(str))
        if i == -1 {
            fmt.Println("unkown command")
        } else {
            fmt.Println(actions[i])
        }
    }
}
func compare(str []string)int{
    l:=len(str)
    index:=-1
    for ix,val:=range commands{
        if l!=clength[ix]{
            continue//命令长度不等返回index=-1
        }
        if l==1&&strings.Contains(val[0],str[0]){
            if index == -1 {
                index = ix//找到索引
            } else {
                index = -1
                break
            }
        }
        if l==2&&strings.Contains(val[0], str[0]) && strings.Contains(val[1], str[1]){
            if index == -1 {
                index = ix
            } else {
                index = -1
                break
            }
        }
    }
    return index
}
var clength = []int {1, 2, 2, 2, 2, 2}
  
var actions = []string {
    "reset what",
    "board fault",
    "where to add",
    "no board at all",
    "impossible",
    "install first",
}