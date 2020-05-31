/**
编写一个程序，将输入字符串中的字符按如下规则排序。
规则 1 ：英文字母从 A 到 Z 排列，不区分大小写。
如，输入： Type 输出： epTy
规则 2 ：同一个英文字母的大小写同时存在时，按照输入顺序排列。
如，输入： BabA 输出： aABb
规则 3 ：非英文字母的其它字符保持原来的位置。
如，输入： By?e 输出： Be?y
*/
package main

import(
  "bufio"
  "fmt"
  "os"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   for {
      strArr,_,_ := reader.ReadLine()
	  if len(strArr) == 0{
	     break
	  }
	  //特殊符号位置记录
      symbols := make(map[int]string)
	  words := make([]byte,0)
	  for i,v := range strArr {
	     if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z'{
		    words = append(words,v)
		 }else{
		    symbols[i] = string(v)
		 }
	  }
	  
	  words = MergeSort(words)
	  
	  newStr := ""
	  for _,v := range words{
	     newStr += string(v)
	  }
       
	  //补上特殊字符
      for i:=0;i<len(strArr);i++{
          if v,ok := symbols[i];ok{
		      if i >= len(newStr){
			      newStr += v
			  }else{
			      newStr = newStr[:i] + v + newStr[i:]
			  }
		  }
  	  }
	  fmt.Println(newStr)   
   }
}

//归并排序，稳定
func MergeSort(slice []byte)(result []byte){
   length := len(slice)
   if length <= 1{
       return slice
   }
   middle := length / 2
   left := MergeSort(slice[:middle])
   right := MergeSort(slice[middle:])
   l,r := 0,0
   for l < len(left) && r < len(right){
      s1,s2 := left[l],right[r]
	  if s1 >= 'A' && s1 <= 'Z'{
	     s1 += 'a' - 'A'
	  }
	  if s2 >='A' && s2 <= 'Z'{
	     s2 += 'a' - 'A'
	  }
	  
	  if s1 <= s2{
	      result = append(result,left[l])
		  l++
	  }else{
	      result = append(result,right[r])
		  r++
	  }
   }
   result = append(result,left[l:]...)
   result = append(result,right[r:]...)
   return
}
