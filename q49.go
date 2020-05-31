/**
题目描述
输入一个单向链表，输出该链表中倒数第k个结点，链表的倒数第1个结点为链表的尾指针。
*/
package main
 
import "fmt"
 
type node struct {
    val int
    next *node
}
 
func main() {
    var (
        n int
        last int
    )
     
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            //fmt.Println(err.Error)
            break
        }
 
        root := &node{}
        fmt.Scan(&root.val)
 
        ne := root
        for i:=0;i<n-1;i++ {
            ne.next = &node{}
            fmt.Scan(&ne.next.val)
            ne = ne.next
        }
 
        fmt.Scan(&last)
        if last == 0 {
            fmt.Println(0)
            continue
        }
         
        ne = root
        for i:=0;i<n-last;i++{
            ne = ne.next
        }
        fmt.Println(ne.val)
         
    }
 
}
