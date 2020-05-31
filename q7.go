//写出一个程序，接受一个正浮点数值，输出该数值的近似整数值。如果小数点后数值大于等于5,向上取整；小于5，则向下取整。
package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  data,_,_ := reader.ReadLine()
  f,_ := strconv.ParseFloat(string(data),32)
  fmt.Println(int(f+0.5))
}