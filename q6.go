//输入一个正整数，按照从小到大的顺序输出它的所有质因子（如180的质因子为2 2 3 3 5 ）
//最后一个数后面也要有空格
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
     num,err := strconv.ParseInt(string(data),10,64)
     if err != nil{
       break
     }
   
     var n int64 
     n = 2
     for num != 1{
       if num % n == 0{
	       fmt.Printf("%d ",n)
		   num = num / n
	   }else{
	       n++
	   }
     }
   }
}