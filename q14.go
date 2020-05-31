package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sort"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   data,_,_ := reader.ReadLine()
   n,_ := strconv.Atoi(string(data))
   var strList []string
   for i:=0;i<n;i++{
     strData,_,_ := reader.ReadLine()
     strList = append(strList,string(strData))	 
   }
   
   sort.Strings(strList)
   
   for _,str := range strList{
      fmt.Println(str)
   }
}