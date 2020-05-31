/**
题目描述
给定一个正整数N代表火车数量，0<N<10，接下来输入火车入站的序列，一共N辆火车，每辆火车以数字1-9编号。要求以字典序排序输出火车出站的序列号。
*/
package main
  
import (
    "fmt"
    "sort"
    "strings"
)
  
type stack []string
  
func (s *stack) push(val string) {
    *s = append(*s, val)
}
  
// 栈顶
func (s *stack) pop() string {
    l := len(*s)
    tail := (*s)[l-1]
    *s = (*s)[0:l-1]
    return tail
}
  
// 栈底
func (s *stack) unshift() string {
    h := (*s)[0]
    *s = (*s)[1:]
    return h
}
  
func (s *stack) isEmpty() bool {
    return len(*s) == 0
}
  
var result []string
  
func handle(pre, in, after stack) {
    // fmt.Println("pre, in, after: ", pre, in, after)
    if pre.isEmpty() && in.isEmpty() {
        str := strings.Join(after, " ")
        // fmt.Println("str: ", str)
        result = append(result, str)
    } else {
        if !pre.isEmpty() { // 进站
            // 弹出栈底的元素
            prec := make([]string, len(pre))
            inc := make([]string, len(in))
            afc := make([]string, len(after))
            copy(prec, pre)
            copy(inc, in)
            copy(afc, after)
            ps := stack(prec)
            is := stack(inc)
            as := stack(after)
              
            h := ps.unshift()
            is.push(h)
            handle(ps, is, as)
        }
        if !in.isEmpty() { // 出站
            prec := make([]string, len(pre))
            inc := make([]string, len(in))
            afc := make([]string, len(after))
            copy(prec, pre)
            copy(inc, in)
            copy(afc, after)
            ps := stack(prec)
            is := stack(inc)
            as := stack(after)
              
            t := is.pop()
            as.push(t)
            handle(ps, is, as)
        }
    }
}
  
func main() {
    var (
        n int
    )
      
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        train := make([]string, n)
        for i := 0; i < n; i++ {
            fmt.Scan(&train[i])
        }
        // fmt.Println("train: ", train)
          
        result = []string{}
        pre := stack(train)
        in := stack{}
        after := stack{}
        handle(pre, in, after)
          
        sort.Strings(result)
        for _, s := range result {
            fmt.Println(s)
        }
    }
}