//明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，他先用计算机生成了N个1到1000之间的随机整数（N≤1000），
//对于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。
//然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。请你协助明明完成“去重”与“排序”的工作
package main

import (
   "bufio"
   "strconv"
   "fmt"
   "os"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   for{
      data,_,_ := reader.ReadLine()
	   if len(data) <= 0{
	    break
	  }
	  
	  num,_ := strconv.Atoi(string(data))
	 
	  var b [1024]bool
	  
      for i:=0;i<num;i++{
	     data,_,_ := reader.ReadLine()
	     number,_ := strconv.Atoi(string(data))	 
		 b[number] = true
	  }
	  
	  for k,v := range b{
	     if v == true{
		    fmt.Println(k)
		 }
	  }
   }
}