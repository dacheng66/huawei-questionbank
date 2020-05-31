/**
题目描述
Catcher 是MCA国的情报员，他工作时发现敌国会用一些对称的密码进行通信，比如像这些ABBA，ABA，A，123321，
但是他们有时会在开始或结束时加入一些无关的字符以防止别国破解。比如进行下列变化 ABBA->12ABBA,ABA->ABAKK,123321->51233214　。
因为截获的串太长了，而且存在多种可能的情况（abaaab可看作是aba,或baaab的加密形式），
Cathcer的工作量实在是太大了，他只能向电脑高手求助，你能帮Catcher找出最长的有效密码串吗？
*/
package main
 
import "fmt"
 
func main()  {
    str := ""
 
    for {
        _, err := fmt.Scan(&str)
        if err != nil {
            break
        }
 
        dp := make([][] bool, len(str))
        for i := 0; i < len(str); i++ {
            row := make([]bool, len(str))
            dp[i] = row
            dp[i][i] = true
        }
 
        max := 1
        for i := len(str) - 2; i >= 0; i-- {
            for j := i + 1; j < len(str); j++ {
                if str[i] == str[j] {
                    if i + 1 < j && dp[i+1][j-1] == true {
                        dp[i][j] = true
                        if j - i > max {
                            max = j - i + 1
                        }
                    } else if j == i + 1 {
                        dp[i][j] = true
                        if j - i > max {
                            max = j - i + 1
                        }
                    }
                }
            }
        }
 
        fmt.Println(max)
 
    }
     
}
