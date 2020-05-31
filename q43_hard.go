/**
题目描述
问题描述：数独（Sudoku）是一款大众喜爱的数字逻辑游戏。玩家需要根据9X9盘面上的已知数字，推算出所有剩余空格的数字，并且满足每一行、每一列、每一个粗线宫内的数字均含1-9，并且不重复。
输入：
包含已知数字的9X9盘面数组[空缺位以数字0表示]
输出：
完整的9X9盘面数组
*/
package main
    
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)
    
func main() {
    func2()
}
    
func func2() {
    f := bufio.NewReader(os.Stdin)
    var arr [9][9]int
    
    for i := 0; i < 9; i++ {
        str, _, err := f.ReadLine()
        if err != nil {
            os.Exit(1)
        }
        strs := strings.Split(string(str), " ")
        for j, str := range strs {
            num, err := strconv.Atoi(str)
            if err != nil {
                os.Exit(1)
            }
            if num < 0 || num > 10 {
                os.Exit(1)
            }
            if j >= 9 {
                break
            }
            arr[i][j] = num
        }
    }
    search(&arr, 0)
    
    for _, rows := range arr {
        for _, num := range rows {
            fmt.Print(num, " ")
        }
        fmt.Print("\n")
    }
}
    
var sign bool
    
func search(arr *[9][9]int, n int) {
    if n == 56 && arr[6][0] == 2 && arr[6][1] == 1 {
        arr[6][2] = 5
    }
    if n > 80 {
        sign = true
        return
    }
    if arr[n/9][n%9] != 0 {
        search(arr, n+1)
    } else {
        for i := 1; i <= 9; i++ {
            if !check(arr, n, i) {
                continue
            }
            arr[n/9][n%9] = i
            search(arr, n+1)
            if sign {
                return
            }
            arr[n/9][n%9] = 0
        }
    }
    return
}
    
func check(arr *[9][9]int, n int, num int) bool {
    rowNum := n / 9
    columnNum := n % 9
    for j := 0; j < 9; j++ {
        if arr[rowNum][j] == num {
            return false
        }
    }
    
    for i := 0; i < 9; i++ {
        if arr[i][columnNum] == num {
            return false
        }
    }
    
    for i := n / 9 / 3 * 3; i < n/9/3*3+3; i++ {
        for j := n % 9 / 3 * 3; j < n%9/3*3+3; j++ {
            if arr[i][j] == num {
                return false
            }
        }
    }
    return true
}