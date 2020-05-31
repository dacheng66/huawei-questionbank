/**
子网掩码是用来判断任意两台计算机的IP地址是否属于同一子网络的根据。
子网掩码与IP地址结构相同，是32位二进制数，其中网络号部分全为“1”和主机号部分全为“0”。利用子网掩码可以判断两台主机是否中同一子网中。若两台主机的IP地址分别与它们的子网掩码相“与”后的结果相同，则说明这两台主机在同一子网中。
示例：
I P 地址　 192.168.0.1
子网掩码　 255.255.255.0
转化为二进制进行运算：
I P 地址　11010000.10101000.00000000.00000001
子网掩码　11111111.11111111.11111111.00000000
AND运算
 　　　　11000000.10101000.00000000.00000000
转化为十进制后为：
 　　　　192.168.0.0
I P 地址　 192.168.0.254
子网掩码　 255.255.255.0
转化为二进制进行运算：
I P 地址　11010000.10101000.00000000.11111110
子网掩码　11111111.11111111.11111111.00000000
AND运算
　　　　　11000000.10101000.00000000.00000000

转化为十进制后为：
　　　　　192.168.0.0

通过以上对两台计算机IP地址与子网掩码的AND运算后，我们可以看到它运算结果是一样的。均为192.168.0.0，所以这二台计算机可视为是同一子网络。
*/
package main
import (
"fmt"
    "strconv"
    "strings"
)
func main() {
    for {
    var mask,ip1,ip2 string
        fmt.Scanf("%s",&mask)
       fmt.Scanf("%s",&ip1)
       n,err:=fmt.Scanf("%s",&ip2)
        if n==0||err!=nil {
           //fmt.Println(1)
            break
        }
   
      fmt.Println(checkNetSegment(mask,ip1,ip2))
    
    }
 
}
 
func checkNetSegment(mask,ip1,ip2 string) int {
    maskFields:=strings.Split(mask,".")
    ip1Fields:=strings.Split(ip1,".")
    ip2Fields:=strings.Split(ip2,".")
    if len(maskFields)!=4 || len(ip1Fields)!=4|| len(ip2Fields)!=4 {
        return 1
    }
    var maskNum,ip1Num,ip2Num uint32
    for i,v:=range maskFields {
        num,_:=strconv.Atoi(v)
        if num>255|| num<0 {
            return 1
        }
        maskNum+=uint32(num)<<uint32((4-i-1)*8)
    }
        for i,v:=range ip1Fields {
        num,_:=strconv.Atoi(v)
        if num>255|| num<0 {
            return 1
        }
         ip1Num+=uint32(num)<<uint32((4-i-1)*8)
    }
            for i,v:=range ip2Fields {
        num,_:=strconv.Atoi(v)
        if num>255|| num<0 {
            return 1
        }
                ip2Num+=uint32(num)<<uint32((4-i-1)*8)
    }
    if (ip1Num&maskNum) ==(ip2Num&maskNum) {
        return 0
    }
    return 2
}