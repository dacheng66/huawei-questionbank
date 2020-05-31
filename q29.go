/**
按照指定规则对输入的字符串进行处理。
详细描述：
将输入的两个字符串合并。
对合并后的字符串进行排序，要求为：下标为奇数的字符和下标为偶数的字符分别从小到大排序。这里的下标意思是字符在字符串中的位置。
对排序后的字符串进行操作，如果字符为‘0’——‘9’或者‘A’——‘F’或者‘a’——‘f’，则对他们所代表的16进制的数进行BIT倒序的操作，并转换为相应的大写字符。如字符为‘4’，为0100b，则翻转后为0010b，也就是2。转换后的字符为‘2’； 如字符为‘7’，为0111b，则翻转后为1110b，也就是e。转换后的字符为大写‘E’。
举例：输入str1为"dec"，str2为"fab"，合并为“decfab”，分别对“dca”和“efb”进行排序，排序后为“abcedf”，转换后为“5D37BF”
接口设计及说明：
/*
功能:字符串处理
输入:两个字符串,需要异常处理
输出:合并处理后的字符串，具体要求参考文档
返回:无

*/

package main
 
import (
    "fmt"
)
 
func SortString(line string) string {
    str := []byte(line)
    for i:=0;i<len(str)/2-1;i++ {
        for j:=0;j<len(str)-2 - i;j+=2 {
            if str[j] > str[j+2] {
                str[j], str[j+2] = str[j+2], str[j]
            }
        }
        for k:=1;k<len(str)-2-i;k+=2 {
            if str[k]>str[k+2] {
                str[k], str[k+2] = str[k+2], str[k]
            }
        }
    }
    //fmt.Println(string(str))
    return string(str)
}
 
func ReverseByte(c byte) byte {
    //fmt.Printf("%c  ",c)
    var num int8
    if c >= '0' && c <= '9' {
        num = int8(c) - '0'
    } else if c >= 'a' && c <= 'f' {
        num = int8(c) - int8('a') + int8(10)
    } else if c >= 'A' && c <= 'F' {
        num = int8(c) - int8('A') + int8(10)
    } else {
        return c
    }
    //fmt.Printf("%d  ", num)
    //分别取出第3,2,1,0位，右移是让数值归1
    bit1 := int8(num)&(1<<3) >> 3
    bit2 := int8(num)&(1<<2) >> 2
    bit3 := int8(num)&(1<<1) >> 1
    bit4 := int8(num)&(1<<0) >> 0
    num = 0
    num |= (bit1 << 0)
    num |= (bit2 << 1)
    num |= (bit3 << 2)
    num |= (bit4 << 3)
    //fmt.Println("num=", num)
    if num >= int8(0) && num <= int8(9) {
        return byte(num+'0')
    } else {
        return byte(num-10+'A')
    }
}
 
func ProcessString(str string) {
    str = SortString(str)
    var res []byte
    for i:=0;i<len(str);i++ {
        c := byte(str[i])
        res =append(res, ReverseByte(c))
    }
    fmt.Println(string(res))
}
 
func main() {
    var str1,str2 string
    for {
        n,_ := fmt.Scan(&str1, &str2)
        if 0 == n {
            break
        }
        str1 = str1 + str2
        ProcessString(str1)
    }
}