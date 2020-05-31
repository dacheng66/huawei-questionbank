//编写一个函数，计算字符串中含有的不同字符的个数。字符在ACSII码范围内(0~127)，换行表示结束符，不算在字符里。不在范围内的不作统计。
package main
import (
   "bufio"
   "fmt"
   "os"
)
func main(){
   
   reader := bufio.NewReader(os.Stdin)
   data,_,_ := reader.ReadLine()
   var iMap = make(map[byte]bool,len(data))
   for _,b := range data {
       if int(b) < 127 && int(b) > 0 {
           iMap[b] = true
      }	   
   }
   fmt.Println(len(iMap))
}