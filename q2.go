//写出一个程序，接受一个由字母和数字组成的字符串，和一个字符，然后输出输入字符串中含有该字符的个数。不区分大小写。
package main

import(
   "bufio"
   "fmt"
   "os"
   "strings"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   data,_,_ := reader.ReadLine()
   sChar,_,_ := reader.ReadLine()
   fmt.Println(strings.Count(strings.ToLower(string(data)),strings.ToLower(string(sChar))))
}