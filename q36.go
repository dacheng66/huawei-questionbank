/**
题目描述
有一只兔子，从出生后第3个月起每个月都生一只兔子，小兔子长到第三个月后每个月又生一只兔子，
假如兔子都不死，问每个月的兔子总数为多少？
*/
package main
import (
    "fmt"
)
 
func main() {
    num := 0
    for {
        n,_ := fmt.Scan(&num)
        if n == 0 {
            break
        }
        arr := make([]int, num+1)
        arr[0] = 1
        arr[1] = 1
        arr[2] = 1
        for i := 3; i < num+1; i++ {
            arr[i] = arr[i-1] + arr[i-2]
        }
        fmt.Println(arr[num])
    }
}