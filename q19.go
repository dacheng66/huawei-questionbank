//开发一个简单错误记录功能小模块，能够记录出错的代码所在的文件名称和行号。
//1、 记录最多8条错误记录，循环记录（或者说最后只输出最后出现的八条错误记录），对相同的错误记录（净文件名（保留最后16位）称和行号完全匹配）只记录一条，错误计数增加；
//2、 超过16个字符的文件名称，只记录文件的最后有效16个字符；
//3、 输入的文件可能带路径，记录文件名称不能带路径。

package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  var sMap = make(map[string]int)
  var fileNameList []string
  
  for {
     data,_,_ :=reader.ReadLine()
	 if len(data) == 0{
	    break
	 }
     sList := strings.Split(string(data)," ")   
	 if len(sList) != 2{
	   continue
	 }
     index := strings.LastIndex(sList[0],"\\")
     fileName := sList[0][index+1:]
     
     lf := len(fileName)	 
     if lf > 16{
        fileName = fileName[lf-16:lf]
     } 	 
     
	 nameKey := fileName+" "+sList[1]
     k := sMap[nameKey]
	 if k == 0{
	    sMap[nameKey] = 1
	    fileNameList = append(fileNameList,nameKey)
	 }else{
	    sMap[nameKey] += 1
	 }
  }
  
  fl := len(fileNameList)
  var st int = 0
  if fl > 8{
    st = fl - 8
  }
  
  for ;st<fl;st++{
    fmt.Printf("%s %d\n",fileNameList[st],sMap[fileNameList[st]])
  }
}