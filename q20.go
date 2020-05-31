//密码要求:
//1.长度超过8位
//2.包括大小写字母.数字.其它符号,以上四种至少三种
//3.不能有相同长度超2的子串重复
//说明:长度超过2的子串
package main
import(
  "bufio"
  "fmt"
  "os"
  "strings"
)


func main(){
   reader := bufio.NewReader(os.Stdin)
   for {
     data,_,_ := reader.ReadLine()
	 if len(data) == 0{
	   break
	 }
	 
     s := string(data)
     if len(s) <= 8{
	    fmt.Println("NG")
		continue
	 }
	 
	 var up,low,num,ord int
	 
	 for _,i :=range s{
	    if i >= 'a' && i <= 'z'{
		   low =1
		   continue
		}
		 if i >= 'A' && i <= 'Z'{
		   up =1
		   continue
		}
		 if i >= '0' && i <= '9'{
		   num =1
		   continue
		}
        ord = 1			      
	 }
     if up+low+num+ord < 3{
        fmt.Println("NG")
		continue
     }

	 var b bool
     for i:=0;i<len(s)-3;i++{
        if strings.Index(s[i+3:],s[i:i+3]) != -1{
			b = true
			break
        }		

      }	 
	 
	 if b {
	  fmt.Println("NG")
	  continue
	 }
    fmt.Println("OK")
   }
}