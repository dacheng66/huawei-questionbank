/**
假设一个球从任意高度自由落下，每次落地后反跳回原高度的一半; 再落下, 求它在第5次落地时，
共经历多少米?第5次反弹多高？
最后的误差判断是小数点6位
*/

package main
 
import "fmt"
 
func main() {
    for {
        var height float64
        n, err := fmt.Scanln(&height)
        if n == 0 || err != nil || height <= 0 {
            break
        }
        // ↓
        // ↓ ↑ ↓
        // ↓ ↑ ↓ ↑ ↓
        // ↓ ↑ ↓ ↑ ↓ ↑ ↓
        // ↓ ↑ ↓ ↑ ↓ ↑ ↓ ↑ ↓
        // ↓ ↑ ↓ ↑ ↓ ↑ ↓ ↑ ↓ ↑
        //  1   2   3   4   5
 
        // 总路程     s = （1 + 1/2*2 + 1/4*2 + 1/8*2 + 1/16*2）*h0
        // 第5次高度 h5 = height/32
        base := 1.0
 
        fmt.Printf("%.6g\n%.6g\n", (base+base/2*2+base/4*2+base/8*2+base/16*2)*height, height/32)
    }
}