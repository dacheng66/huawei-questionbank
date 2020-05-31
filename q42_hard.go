/**
题目描述
定义一个二维数组N*M（其中2<=N<=10;2<=M<=10），如5 × 5数组下所示： 
int maze[5][5] = {
0, 1, 0, 0, 0,
0, 1, 0, 1, 0,
0, 0, 0, 0, 0,
0, 1, 1, 1, 0,
0, 0, 0, 1, 0,
};

它表示一个迷宫，其中的1表示墙壁，0表示可以走的路，只能横着走或竖着走，不能斜着走，要求编程序找出从左上角到右下角的最短路线。入口点为[0,0],既第一空格是可以走的路。
Input
一个N × M的二维数组，表示一个迷宫。数据保证有唯一解,不考虑有多解的情况，即迷宫只有一条通道。
Output
左上角到右下角的最短路径，格式如样例所示。
Sample Input
0 1 0 0 0
0 1 0 1 0
0 0 0 0 0
0 1 1 1 0
0 0 0 1 0
Sample Output
(0, 0)
(1, 0)
(2, 0)
(2, 1)
(2, 2)
(2, 3)
(2, 4)
(3, 4)
(4, 4)
*/
package main
 
import (
    "fmt"
)
 
type point struct {
    i, j int
}
 
// 移动的方向
var dirs = [4]point{
    {-1, 0}, // 上
    {0, -1}, // 左
    {1, 0}, // 下
    {0, 1}, // 右
}
 
func (p point) add(r point) point {
    return point{p.i + r.i, p.j + r.j}
}
 
func (p point) at(grid [][]int) (int, bool) {
    // 判断越界
    if p.i < 0 || p.i >= len(grid) {
        return 0, false
    }
    if p.j < 0 || p.j >= len(grid[p.i]) {
        return 0, false
    }
    return grid[p.i][p.j], true
}
 
func main() {
    var (
        n int // 行
        m int // 列
    )
    for {
        _, err := fmt.Scan(&n)
        if err != nil {
            break
        }
        fmt.Scan(&m)
         
        maze := make([][]int, n)
        for i := range maze {
            maze[i] = make([]int, m)
            for j := range maze[i] {
                fmt.Scan(&maze[i][j])
            }
        }
         
        // 记录路径
        steps := make([][]int, n)
        for i := range steps {
            steps[i] = make([]int, m)
        }
         
        // 起始点入队
        queue := []point{{0, 0}}
        for len(queue) > 0 {
            // 获取队首元素
            cur := queue[0]
            queue = queue[1:]
             
            // 终点
            if cur.i == n-1 && cur.j == m-1 {
                break
            }
             
            for _, dir := range dirs {
                // 下一个节点的坐标
                next := cur.add(dir)
                // 迷宫可以走
                // 并且没有走过
                // 并且没有回到起点
                val, ok := next.at(maze)
                if !ok || val == 1 {
                    continue
                }
                val, ok = next.at(steps)
                if !ok || val != 0 {
                    continue
                }
                if next.i == 0 && next.j == 0 {
                    continue
                }
                tmp, _ := cur.at(steps)
                steps[next.i][next.j] = tmp + 1
                queue = append(queue, next)
            }
        }
         
        // fmt.Println("steps: ", steps)
         
        // 打印路径
        step := steps[n-1][m-1] // 走到目标点需要的步数
        path := make([]point, step+1)
        path[step] = point{n-1, m-1}
        path[0] = point{0, 0}
        step--
 
        for ; step > 0; step-- {
            for _, dir := range dirs {
                last := path[step+1].add(dir)
                if val, ok := last.at(steps); ok {
                    if val == step {
                        path[step] = last
                        break
                    }
                }
            }
        }
        for _, p := range path {
            fmt.Printf("(%d,%d)\n", p.i, p.j)
        }
    }
}