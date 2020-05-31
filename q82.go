/**
题目描述
有一个数据表格为二维数组（数组元素为int类型），行长度为ROW_LENGTH,列长度为COLUMN_LENGTH。对该表格中数据的操作可以在单个单元内，也可以对一个整行或整列进行操作，操作包括交换两个单元中的数据；插入某些行或列。
请编写程序，片段对表格的各种操作是否合法。
详细要求:
1.数据表规格的表示方式为“行*列”, 数据表元素的位置表示方式为[行,列]，行列均从0开始编号
2.数据表的最大规格为9行*9列，对表格进行操作时遇到超出规格应该返回错误
3.插入操作时，对m*n表格，插入行号只允许0~m，插入列号只允许0~n。超出范围应该返回错误
4.只需记录初始表格中数据的变化轨迹，查询超出初始表格的数据应返回错误
例如:  初始表格为4*4，可查询的元素范围为[0,0]~[3,3]，假设插入了第2行，数组变为5*4，查询元素[4,0]时应该返回错误
*/

package main
  
import (
    "fmt"
)
  
func main() {
    var (
        row int // 原始行列值
        col int
          
        row1 int // 需要交换的行列值
        col1 int
        row2 int
        col2 int
          
        irow int // 插入的行值
        icol int // 插入的列值
          
        grow int // 要获取的单元值
        gcol int
    )
      
    for {
        _, err := fmt.Scan(&row)
        if err != nil {
            break
        }
        fmt.Scan(&col)
          
        fmt.Scan(&row1)
        fmt.Scan(&col1)
        fmt.Scan(&row2)
        fmt.Scan(&col2)
          
        fmt.Scan(&irow)
        fmt.Scan(&icol)
        fmt.Scan(&grow)
        fmt.Scan(&gcol)
          
        res := make([]int, 5)
        if row <= 0 || col <= 0 {
            res[0] = -1
        }
        // 交换
        if row1 < 0 || row1 >= row || col1 < 0 || col1 >= col || row2 < 0 || row2 >= row || col2 < 0 || col2 >= col {
            res[1] = -1
        }
        if irow < 0 || irow >= row {
            res[2] = -1
        }
        //else {
         //   row++
        //}
        if icol < 0 || icol >= col {
            res[3] = -1
        }
        //else {
        //    col++
        //}
        if grow < 0 || grow >= row || gcol < 0 || gcol >= col {
            res[4] = -1
        }
        for _, v := range res {
            fmt.Println(v)
        }
    }
}