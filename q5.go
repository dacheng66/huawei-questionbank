//写出一个程序，接受一个十六进制的数，输出该数值的十进制表示。（多组同时输入 ）
package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   for {
       data,_,_ := reader.ReadLine()
       if len(data) == 0{
          break
       }
       num,err :=  strconv.ParseInt(string(data)[2:],16,32)
       fmt.Println(num)
       if err != nil{
	      break
       }  	   
   }
}