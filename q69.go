/**
题目描述
公元前五世纪，我国古代数学家张丘建在《算经》一书中提出了“百鸡问题”：鸡翁一值钱五，鸡母一值钱三，鸡雏三值钱一。
百钱买百鸡，问鸡翁、鸡母、鸡雏各几何？
*/

package main
 
import (
    "fmt"
)
 
func main() {
    var m int
     
    for {
        _, err := fmt.Scan(&m)
        if err != nil {
            break
        }
        // x + y + z = 100
        // 5x + 3y + z/3 = 100
        // y = (100 - 7x) / 4
        // z = 75 + 3x/4
        for x := 0; x <= 20; x+= 4 {
            y := (100 - 7 * x) / 4
            z := 75 + 3 * x / 4
            if y > 0 && z > 0 && x+y+z == 100 {
                fmt.Println(x, y, z)
            }
        }
    }
}