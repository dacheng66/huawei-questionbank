/***
题目描述
MP3 Player因为屏幕较小，显示歌曲列表的时候每屏只能显示几首歌曲，用户要通过上下键才能浏览所有的歌曲。
为了简化处理，假设每屏只能显示4首歌曲，光标初始的位置为第1首歌。
现在要实现通过上下键控制光标移动来浏览歌曲列表，控制逻辑如下：
歌曲总数<=4的时候，不需要翻页，只是挪动光标位置。
*/
package main
 
import (
    "fmt"
)
 
const page = 4
 
func main() {
    var (
        n int // 歌曲总数
        action string // 操作
    )
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        fmt.Scan(&action)
         
        cursor := 0 // 当前指向的歌曲,从 0 开始计数
        if n <= page {
            for _, ch := range action {
                if ch == 'U' {
                    cursor = (cursor+1) % n
                } else {
                    cursor = (cursor-1) % n
                }
            }
            for i := 0; i < n; i++ {
                fmt.Printf("%d ", i+1)
            }
            fmt.Println()
            fmt.Println(cursor+1)
            continue
        }
         
        header := 0 // 当前列表的头
        for _, ch := range action {
            if ch == 'U' {
                if cursor == 0 {
                    header = n - page
                    cursor = n - 1
                } else if cursor == header {
                    // 光标在当前页的第一首
                    cursor--
                    header--
                } else {
                    cursor--
                }
            } else {
                if cursor == n-1 {
                    header = 0
                    cursor = 0
                } else if cursor == header+page-1 {
                    // 当前光标在当前页的最后一首
                    cursor++
                    header++
                } else {
                    cursor++
                }
            }
        }
         
        for i := 1; i <= page; i++ {
            fmt.Printf("%d ", header+i)
        }
        fmt.Println()
        fmt.Println(cursor+1)
    }
}