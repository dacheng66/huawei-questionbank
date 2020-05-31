//数据表记录包含表索引和数值（int范围的整数），请对表索引相同的记录进行合并，即将相同索引的数值进行求和运算，输出按照key值升序进行输出。
package main

import (
   "bufio"
   "fmt"
   "os"
   "strings"
   "strconv"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  data,_,_ := reader.ReadLine()
  n,_ := strconv.Atoi(string(data))
  
  var aMap = make(map[int]int,n)
  for i:=0;i<n;i++{
     data1,_,_ := reader.ReadLine()
	 s := strings.Split(string(data1)," ")
	 if len(s) == 2{
	    index,_ := strconv.Atoi(s[0])
		value,_ := strconv.Atoi(s[1])
		v := aMap[index]
		v += value
	    aMap[index] = v
	 }
  }
  
  var keys []int
  for k,_ := range aMap{
      keys = append(keys,k)

  }  
  for i:=0;i<len(keys)-1;i++{
     for j:=0;j<len(keys)-i-1;j++{
	    if keys[j] > keys[j+1]{
		   keys[j],keys[j+1] = keys[j+1],keys[j]
		
		}
	 
	 }
  }
  
  for _,v := range keys{
    fmt.Println(v,aMap[v])
  
  }
}