/**
题目描述
对于不同的字符串，我们希望能有办法判断相似程度，我们定义了一套操作方法来把两个不相同的字符串变得相同，具体的操作方法如下：
1 修改一个字符，如把“a”替换为“b”。
2 增加一个字符，如把“abdd”变为“aebdd”。
3 删除一个字符，如把“travelling”变为“traveling”。
比如，对于“abcdefg”和“abcdef”两个字符串来说，我们认为可以通过增加和减少一个“g”的方式来达到目的。上面的两种方案，都只需要一次操作。把这个操作所需要的次数定义为两个字符串的距离，而相似度等于“距离＋1”的倒数。也就是说，“abcdefg”和“abcdef”的距离为1，相似度为1/2=0.5.
给定任意两个字符串，你是否能写出一个算法来计算出它们的相似度呢？
*/
package main
 
import (
    "fmt"
)
 
const ic = 1
const dc = 1
const rc = 1
 
func min(a, b, c int) int {
    re := a
    if b < re {
        re = b
    }
    if c < re {
        re = c
    }
    return re
}
 
func main() {
    var (
        str1 string
        str2 string
    )
     
    for {
        _, err := fmt.Scan(&str1)
        if err != nil {
            break
        }
        fmt.Scan(&str2)
         
        if str1 == str2 {
            fmt.Println(0)
            continue
        }
        if str1 == "" {
            fmt.Println(len(str2) * rc)
            continue
        }
        if str2 == "" {
            fmt.Println(len(str1) * dc)
            continue
        }
         
        l1, l2 := len(str1), len(str2)
        dp := make([][]int, l1+1)
        for ix := range dp {
            dp[ix] = make([]int, l2+1)
        }
         
        for i := 0; i <= l1; i++ {
            // 将字符串 变成 空串
            dp[i][0] = i * dc
        }
        for j := 0; j <= l2; j++ {
            // 将空串 变成 字符串
            dp[0][j] = j * ic
        }
         
        for i := 1; i <= l1; i++ {
            for j := 1; j <= l2; j++ {
                case1 := dp[i-1][j] + dc // 删除一个字符后变成 str2
                case2 := dp[i][j-1] + ic // 插入一个字符后变成 str2
                case3 := 0
                if str1[i-1] == str2[j-1] { // 从空串算起，字符串的索引要减 1
                    case3 = dp[i-1][j-1]
                } else {
                    case3 = dp[i-1][j-1] + rc // 修改一个字符变成 str2
                }
                dp[i][j] = min(case1, case2, case3)
            }
        }
        fmt.Printf("1/%d\n", dp[l1][l2]+1)
    }
}