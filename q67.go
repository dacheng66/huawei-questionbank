/**
题目描述
矩阵乘法的运算量与矩阵乘法的顺序强相关。
例如：
A是一个50×10的矩阵，B是10×20的矩阵，C是20×5的矩阵
计算A*B*C有两种顺序：（（AB）C）或者（A（BC）），前者需要计算15000次乘法，后者只需要3500次。
编写程序计算不同的计算顺序需要进行的乘法次数
*/
package main
 
import (
    "fmt"
)
 
// 矩阵的行和列
type matrix struct {
    row int
    col int
}
 
func (m matrix) String() string {
    return fmt.Sprintf("row: %d, col: %d\n", m.row, m.col)
}
 
type stack1 []matrix
 
func (s *stack1) push(m matrix) {
    (*s) = append(*s, m)
}
 
func (s *stack1) pop() (matrix, bool) {
    l := len(*s)
    if l < 1 {
        return matrix{0, 0}, false
    }
    tail := (*s)[l-1]
    (*s) = (*s)[0 : l-1]
    return tail, true
}
 
func main() {
    var (
        n   int
        str string // 矩阵的计算顺序
    )
 
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        arr := make([][]int, n)
        for ix := range arr {
            arr[ix] = make([]int, 2)
            fmt.Scanf("%d %d", &arr[ix][0], &arr[ix][1])
        }
        fmt.Scan(&str)
        // fmt.Println("arr: ", arr)
 
        s := &stack1{}
        res := 0
        ix := 0
        for _, ch := range str {
            if ch == ')' {
                m1, _ := s.pop()
                m2, _ := s.pop()
                // fmt.Println("m1: ", m1)
                // fmt.Println("m2: ", m2)
                // 后一个矩阵乘以前一个
                res += m2.row * m2.col * m1.col
                s.push(matrix{m2.row, m1.col})
            } else if ch >= 'A' && ch <= 'Z' {
                cr := arr[ix]
                ix++
                s.push(matrix{row: cr[0], col: cr[1]})
            }
        }
        fmt.Println(res)
    }
}