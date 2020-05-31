package main
  
import (
    "fmt"
    "strconv"
)
  
type stack []string
  
func (s *stack) push(v string) {
    *s = append(*s, v)
}
  
func (s *stack) top() string {
    l := len(*s)
    return (*s)[l-1]
}
  
func (s *stack) pop() string {
    l := len(*s) - 1
    tail := (*s)[l]
    *s = (*s)[:l]
    return tail
}
  
func (s *stack) shift() string {
    h := (*s)[0]
    if len(*s) > 1 {
       *s = (*s)[1:]
    } else {
        *s = nil
    }
    return h
}
  
func (s *stack) length() int {
    return len(*s)
}
  
func (s *stack) isEmpty() bool {
    return len(*s) == 0
}
  
  
type stackInt []int
  
func (s *stackInt) push(v int) {
    *s = append(*s, v)
}
  
func (s *stackInt) pop() int {
    l := len(*s) - 1
    tail := (*s)[l]
    *s = (*s)[:l]
    return tail
}
  
// 操作符等级
var opRank = map[byte]int{
    '{': 1,
    '[': 1,
    '(': 1,
    '*': 3,
    '/': 3,
    '+': 2,
    '-': 2,
}
// 括号匹配
var braMap = map[byte]byte{
    '}': '{',
    ']': '[',
    ')': '(',
}
  
func main() {
    var expr string
      
    for {
        _, err := fmt.Scan(&expr)
        if err != nil {
            break
        }
        // fmt.Println("expr: ", expr)
          
        // 存储操作数
        s1 := stack{}
        // 存储操作符
        s2 := stack{}
          
        lastIsOp := false // 上一个是否是操作符
        lastLeft := false // 上一个是否是左括号
        for _, ch := range expr {
            chStr := string(ch)
            // fmt.Println("post prefix: ", string(s1))
            // fmt.Println("ch: ", string(ch))
            if ch >= '0' && ch <= '9' {
                if !s1.isEmpty() && !lastIsOp {
                    // 上一个也是数字
                    tmp := s1.pop()
                    s1.push(tmp+chStr)
                } else {
                    s1.push(chStr)
                }
                lastIsOp = false
                lastLeft = false
                continue
            }
              
            if ch == '{' || ch == '[' || ch == '(' {
                s2.push(chStr)
                // fmt.Println("s2~~: ", string(s2))
                lastIsOp = true
                lastLeft = true
                continue
            }
              
            // 出现在左括号后面的才是 负数
            if ch == '-' && lastLeft {
                s1.push("0")
            }
            lastIsOp = true
            lastLeft = false
  
            if ch == '}' || ch == ']' || ch == ')' {
                // fmt.Println("s2: ", string(s2))
                i := s2.pop()
                for {
                    // fmt.Println("ch: ", string(i), string(ch))
                    if s2.isEmpty() || i[0] == braMap[byte(ch)] {
                        break
                    }
                    s1.push(i)
                    i = s2.pop()
                }
                continue
            }
             
            if ch == '*' || ch == '/' || ch == '+' || ch == '-' {
                if !s2.isEmpty() {
                    i := s2.top()
                    for opRank[i[0]] >= opRank[byte(ch)] {
                        // fmt.Println("i: ", string(i))
                        s1.push(i)
                        s2.pop()
                        if s2.isEmpty() {
                            break
                        }
                        i = s2.top()
                    }
                }
                s2.push(chStr)
            }
            // fmt.Println("s1: ", string(s1))
            // fmt.Println("s2: ", string(s2))
        }
        // fmt.Println("s1: ", string(s1))
        // fmt.Println("s2: ", string(s2))
        for !s2.isEmpty() {
            s1.push(s2.pop())
        }
          
        // fmt.Println("post prefix: ", s1)
         
        re := stackInt{}
        for !s1.isEmpty() {
            ch := s1.shift()
            // fmt.Println("ch in post prefix: ", ch, re)
            if v, err := strconv.Atoi(ch); err == nil {
                re.push(v)
            } else {
                v2 := re.pop()
                v1 := re.pop()
                switch ch {
                case "+":
                    re.push(v1 + v2)
                case "-":
                    re.push(v1 - v2)
                case "*":
                    re.push(v1 * v2)
                case "/":
                    re.push(v1 / v2)
                }
            }
  
        }
          
        fmt.Println(re.pop())
    }
}