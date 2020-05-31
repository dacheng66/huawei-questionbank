/**
题目描述
问题描述：给出4个1-10的数字，通过加减乘除，得到数字为24就算胜利
输入：
4个1-10的数字。[数字允许重复，但每个数字仅允许使用一次，测试用例保证无异常数字]
输出：
true or false
*/ 
package main
import (
    "fmt"
)
 
func get_points(A int, res []int) []int{
    points := []int{}
    if len(res) < 1 {
        points = append(points, A)
    }else {
        for _, k := range res {
            points = append(points, k + A)
            points = append(points, k * A)
            points = append(points, k - A)
            if A == 0 {
                points = append(points, 0)
            }else {
                points = append(points, k / A)
            }
        }
    }
    return points
}
 
func game24(A int, N []int) []int {
    res := []int{}
    M := []int{}
    count :=0
    for _,v := range N {
        if v == A && count<1{
            count++
            continue
        }
        M = append(M, v)
    }
 
    if len(N) > 1 {
        for _,V := range M {
            res = append(res, get_points(V, game24(V, M))...)
        }
    }else {
        return get_points(A, []int{})
    }
    return res
}
 
func Error(err string) {
    panic(err)
}
 
func main() {
 
    A := 0
    B := 0
    C := 0
    D := 0
 
    for {
        n, _ := fmt.Scan(&A,&B,&C,&D)
        if n == 0 {
            break
        }else {
            res := []int{}
            res = append(res, game24(A, []int{A, B, C, D})...)
            res = append(res, game24(B, []int{A, B, C, D})...)
            res = append(res, game24(C, []int{A, B, C, D})...)
            res = append(res, game24(D, []int{A, B, C, D})...)
 
            result := false
            for _, v := range res {
                if v == 24 {
                    result = true
                }
            }
            fmt.Println(result)
        }
    }
}