/**
题目描述
1、对输入的字符串进行加解密，并输出。
2加密方法为：
当内容是英文字母时则用该英文字母的后一个字母替换，同时字母变换大小写,如字母a时则替换为B；字母Z时则替换为a；
当内容是数字时则把该数字加1，如0替换1，1替换2，9替换0；
其他字符不做变化。
3、解密方法为加密的逆过程。
接口描述：
实现接口，每个接口实现1个基本操作：
void Encrypt (char aucPassword[], char aucResult[])：在该函数中实现字符串加密并输出
说明：
1、字符串以\0结尾。
2、字符串最长100个字符。
int unEncrypt (char result[], char password[])：在该函数中实现字符串解密并输出

说明：
1、字符串以\0结尾。
2、字符串最长100个字符。
*/

package main
 
import (
    "bufio"
    "fmt"
    "os"
)
 
func nextChar(a byte) byte {
    if a == '9' {
        return '0'
    } else if a >= '0' && a < '9' {
        return a + 1
    } else if a == 'z' {
        return 'A'
    } else if a == 'Z' {
        return 'a'
    } else if a >= 'a' && a < 'z' {
        return (a + 1) - 32
    } else if a >= 'A' && a < 'Z' {
        return (a + 1) + 32
    }
    return 0
}
 
func prevChar(a byte) byte {
    if a == '0' {
        return '9'
    } else if a > '0' && a <= '9' {
        return a - 1
    } else if a == 'a' {
        return 'Z'
    } else if a == 'A' {
        return 'z'
    } else if a > 'a' && a <= 'z' {
        return (a - 1) - 32
    } else if a > 'A' && a <= 'Z' {
        return (a - 1) + 32
    }
    return 0
}
 
func Decrypt(s []byte) []byte {
    for i, c := range s {
        s[i] = byte(prevChar(c))
    }
    return s
}
 
func Encrypt(s []byte) []byte {
    for i, c := range s {
        s[i] = byte(nextChar(c))
    }
    return s
}
 
func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        s1, _, _ := reader.ReadLine()
        if len(s1) == 0 {
            break
        }
 
        b1 := make([]byte, len(s1))
        copy(b1, s1)
 
        s2, _, _ := reader.ReadLine()
        if len(s2) == 0 {
            break
        }
 
        b2 := make([]byte, len(s2))
        copy(b2, s2)
 
        fmt.Printf("%s\n", Encrypt(b1))
        fmt.Printf("%s\n", Decrypt(b2))
    }
}