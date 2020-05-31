/**
题目描述
Catcher是MCA国的情报员，他工作时发现敌国会用一些对称的密码进行通信，
比如像这些ABBA，ABA，A，123321，但是他们有时会在开始或结束时加入一些无关的字符以防止别国破解。
比如进行下列变化 ABBA->12ABBA,ABA->ABAKK,123321->51233214　。
因为截获的串太长了，而且存在多种可能的情况（abaaab可看作是aba,或baaab的加密形式），
Cathcer的工作量实在是太大了，他只能向电脑高手求助，你能帮Catcher找出最长的有效密码串吗？
*/
package main
   
import (
    "fmt"
    "bufio"
    "os"
)
   
func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        input,_,_ := reader.ReadLine()
        if len(input) == 0 {
            break
        }
        data := string(input)
        arr := []rune(data)
        max := 1
        for i := 0; i < len(arr); i++ {
            left := i-1
            right := i+1
            count := 1
            for left >= 0 && right < len(arr) {
                if arr[left] != arr[i] && arr[right] != arr[i] {
                    break
                }
                if arr[left] == arr[i] {
                    left--
                }
                if arr[right] == arr[i] {
                    right++
                }
                   
            }
            count = right-left-1
            if max < count {
                max = count
            }
            for left >= 0 && right < len(arr) {
                if arr[left] == arr[right] {
                    left--
                    right++
                    count = count+2
                    if count > max {
                        max = count
                    }
                } else {
                    break
                }
                   
            }
               
        }
        fmt.Println(max)
    }
}