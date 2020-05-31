//•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
//•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。
package main
import (
   "bufio"
   "fmt"
   "os"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   for i:=0;i<2;i++{
        data,_,_ := reader.ReadLine()
		s := string(data)
		l := len(s)
		flag := 0
		for l > 8{
		    fmt.Println(s[flag:flag+8])
			l -= 8
			flag += 8
		}
		
		num := s[flag:]
		
		for j:=0;j<8-l;j++{
		  num += "0"
		}
        fmt.Println(num)
   
   }
}