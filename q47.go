/**
题目描述
问题描述：有4个线程和1个公共的字符数组。线程1的功能就是向数组输出A，线程2的功能就是向字符输出B，
线程3的功能就是向数组输出C，线程4的功能就是向数组输出D。要求按顺序向数组赋值ABCDABCDABCD，ABCD的个数由线程函数1的参数指定。
*/

package main
 
import (
    "fmt"
    "io"
)
 
var str []byte
var count int
var ch1 chan int = make(chan int)
var ch2 chan int = make(chan int)
var ch3 chan int = make(chan int)
var ch4 chan int = make(chan int)
 
var ch5 chan int = make(chan int)
 
func main() {
 
    go addA()
    go addB()
    go addC()
    go addD()
 
    for {
        _, ok := fmt.Scanf("%d\n", &count)
        if ok == io.EOF {
            break
        }
 
        ch1 <- 1
        <-ch5
 
        for i := range str {
            fmt.Printf("%c", str[i])
        }
        fmt.Println()
        str = nil
    }
}
 
func addA() {
    for {
        <-ch1
        if count != 0 {
            str = append(str, 'A')
            ch2 <- 1
        } else {
            ch5 <- 1
        }
        count--
    }
}
 
func addB() {
    for {
        <-ch2
        str = append(str, 'B')
        ch3 <- 1
    }
}
 
func addC() {
    for {
        <-ch3
        str = append(str, 'C')
        ch4 <- 1
    }
}
 
func addD() {
    for {
        <-ch4
        str = append(str, 'D')
        ch1 <- 1
    }
}