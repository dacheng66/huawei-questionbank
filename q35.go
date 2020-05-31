/**
题目描述
有一种技巧可以对数据进行加密，它使用一个单词作为它的密匙。下面是它的工作原理：首先，选择一个单词作为密匙，如TRAILBLAZERS。如果单词中包含有重复的字母，只保留第1个，其余几个丢弃。现在，修改过的那个单词属于字母表的下面，如下所示：
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z
T R A I L B Z E S C D F G H J K M N O P Q U V W X Y
上面其他用字母表中剩余的字母填充完整。在对信息进行加密时，信息中的每个字母被固定于顶上那行，并用下面那行的对应字母一一取代原文的字母(字母字符的大小写状态应该保留)。因此，使用这个密匙，Attack AT DAWN(黎明时攻击)就会被加密为Tpptad TP ITVH。
请实现下述接口，通过指定的密匙和明文得到密文。
*/

package main
  
import (
    "fmt"
)
  
func main() {
    key := ""
    text := ""
    for {
        n, _ := fmt.Scan(&key, &text)
        if n == 0 {
            break
        }
        arr := []rune(key)
        parr := []rune(text)
          
        m := make(map[rune]int)
        t := 0
        alphabet := make([]rune,26)
        for i := 0; i < len(arr); i++ {
            _, ok := m[arr[i]]
            if ok {
                continue
            } else {
                m[arr[i]] = 1
                alphabet[t] = arr[i]
                t++
            }
        }
          
        for i := 'a'; i <= 'z'; i++ {
            _, ok := m[i]
            if ok {
                continue
            } else {
                m[i] = 1
                alphabet[t] = i
                t++
            }
        }
          
          
        res := ""
        for i := 0; i < len(parr); i++ {
            if parr[i] >= 'A' && parr[i] <= 'Z' {
                tmp := parr[i] - 'A'
                v := alphabet[tmp]
                v = v - 'a' + 'A'
                res += string(v)
            } else {
                tmp := parr[i] - 'a'
                v := alphabet[tmp]
                res += string(v)
            }
        }
        fmt.Println(res)
    }
}