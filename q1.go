//计算字符串最后一个单词的长度，单词以空格隔开。

package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
)   
   
func main(){
   reader := bufio.NewReader(os.Stdin)
   data,_,_ := reader.ReadLine()
   s := string(data)
   sList := strings.Split(s," ")
   
   l := len(sList)
   if l == 0{
      fmt.Println(0)
   }else{
      fmt.Println(len(sList[l-1]))
   }
}