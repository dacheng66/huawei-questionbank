/**
题目描述
对字符串中的所有单词进行倒排。
说明：
1、构成单词的字符只有26个大写或小写英文字母；
2、非构成单词的字符均视为单词间隔符；
3、要求倒排后的单词间隔符以一个空格表示；如果原字符串中相邻单词间有多个间隔符时，倒排转换后也只允许出现一个空格间隔符；
4、每个单词最长20个字母；
*/
package main
  
import (
    "bufio"
    "os"
    "strings"
    "fmt"
)
  
func main(){
    reader := bufio.NewReader(os.Stdin)
    str, _, _ := reader.ReadLine()
    for i,c :=range str{
        if !((c>='a'&&c<='z')||(c>='A' && c<='Z')){
            str[i]=' '
        }
    }
    ans := strings.Fields(string(str))
  
    for i:=len(ans)-1;i>=0;i--{
        if i==0{
            fmt.Printf("%s",ans[i])
        }else {
            fmt.Printf("%s ",ans[i])
  
        }
    }
}  
  