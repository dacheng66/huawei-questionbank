//输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
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
   
   var num string
   var s = string(data)
   for i:=len(s)-1;i>=0;i--{
      if !strings.Contains(num,string(s[i])){
	     num += string(s[i])
	  }
   }

   fmt.Printf("%s",num)
 }