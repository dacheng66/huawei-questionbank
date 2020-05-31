/*
先输入字典中单词的个数，再输入n个单词作为字典单词。
输入一个单词，查找其在字典中兄弟单词的个数
再输入数字n
根据输入，输出查找到的兄弟单词的个数
*/
package main

import(
  "fmt"
  "strings"
  "sort"
)

func main(){
   var (
      num int
	  word string
	  n int   
   )
   
   for{
     _,err := fmt.Scan(&num)
	 if err != nil{
        break
	 }
	 
	 words := make([]string,num)
	 siblings := make([]string,num)
	 total := 0 //兄弟单词的个数
	 
	 for i:=0;i<num;i++{
	   fmt.Scan(&words[i])
	 }
	 
	 fmt.Scan(&word)
	 fmt.Scan(&n)
	 
	 tmpArr := strings.Split(word,"")
	 sort.Strings(tmpArr)
	 wordSorted := strings.Join(tmpArr,"")
	 for i:=0;i<num;i++{
	    //查找兄弟单词
	    if words[i] == word || len(words[i]) != len(word){
		   continue
		}
		
		tmp2Arr := strings.Split(words[i],"")
		sort.Strings(tmp2Arr)
		tmp3 := strings.Join(tmp2Arr,"")
		if wordSorted == tmp3{
		   siblings[total] = words[i]
		   total++
		}
	 }
	 
	 //排序
	 slice := siblings[0:total]
	 sort.Strings(slice)
	 
	 fmt.Println(total)
	 if total >= n{
	    fmt.Println(slice[n-1])
	 }
   }
}
