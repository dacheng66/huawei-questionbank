/**
如果A是个x行y列的矩阵，B是个y行z列的矩阵，把A和B相乘，其结果将是另一个x行z列的矩阵C。
这个矩阵的每个元素是由下面的公式决定的

*/
package main
 
import "fmt"
 
func main()  {
 
    for  {
        var (
            x int
            y int
            z int
        )
 
        _, err := fmt.Scan(&x)
        if err != nil {
            break
        }
        fmt.Scan(&y)
        fmt.Scan(&z)
 
        arr1 := make([][]int, x)
        for i := range arr1 {
            arr1[i] = make([]int, y)
            for j := range arr1[i] {
                fmt.Scan(&arr1[i][j])
            }
        }
 
        arr2 := make([][]int, y)
        for i := range arr2 {
            arr2[i] = make([]int, z)
            for j := range arr2[i] {
                fmt.Scan(&arr2[i][j])
            }
        }
 
        res := make([][]int, x)
        for i := 0; i < x; i++ {
            row := make([]int, z)
            res[i] = row
        }
 
        for i := 0; i < x; i++ {
            for j := 0; j < z; j++ {
                res[i][j] = multi(arr1, arr2, i, j)
            }
        }
 
 
        for i := 0; i < x; i++ {
            for j := 0; j < z; j++ {
                fmt.Printf("%d ", res[i][j])
            }
            fmt.Println()
        }
 
 
 
    }
 
}
 
func multi(arr1 [][]int, arr2 [][]int, i int, j int) int {
    v := 0
    for k := 0; k < len(arr1[0]); k++ {
        v = v + arr1[i][k] * arr2[k][j]
    }
 
    return v
}
