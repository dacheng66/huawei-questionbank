/**
题目描述
若两个正整数的和为素数，则这两个正整数称之为“素数伴侣”，如2和5、6和13，它们能应用于通信加密。
现在密码学会请你设计一个程序，从已有的N（N为偶数）个正整数中挑选出若干对组成“素数伴侣”，
挑选方案多种多样，例如有4个正整数：2，5，6，13，如果将5和6分为一组中只能得到一组“素数伴侣”，
而将2和5、6和13编组将得到两组“素数伴侣”，能组成“素数伴侣”最多的方案称为“最佳方案”，当然密码学会希望你寻找出“最佳方案”。

输入:
有一个正偶数N（N≤100），表示待挑选的自然数的个数。后面给出具体的数字，范围为[2,30000]。
输出:
输出一个整数K，表示你求得的“最佳方案”组成“素数伴侣”的对数。
*/
package main

import (
  "fmt"
)

var (
   matchs []int
   used []bool
   odd []int
   even []int
   primeMap map[int]int
)

func main(){
   primeMap=make(map[int]int)
   for {
      var num int
	  x,_ := fmt.Scanln(&num)
	  if x == 0{
	    return
	  }
	  var all = make([]int,num+1)
	  matchs = make([]int,num+1)
	  for i:=1;i<=num;i++{
	     fmt.Scan(&all[i])
	  }
	  odd = []int{}
	  even = []int{}
	  odd = append(odd,0)
	  for i:=1;i<=num;i++{
	     if all[i]%2 == 1{
		   odd = append(odd,all[i])
		 }else{
		   even = append(even,all[i])
		 }
	  }
	  
	  if len(odd)*len(even)==0{
	     fmt.Println(0)
		 return
	  }

      var result int
      for i:=1;i<len(odd);i++{
          used = make([]bool,len(even)+1)
		  if find(i){
		     result++
		  }
      }	 
       fmt.Println(result)	  
   }
}

func find(n int)bool{
   for i:=0;i<len(even);i++{
      if isPrime(odd[n],even[i]) && !used[i]{
	     used[i] = true
		 if matchs[i] == 0 || find(matchs[i]){
		   matchs[i] = n
	       return true
		 }
	  }
   }
   return false 
}

func isPrime(x,y int)bool{
  z := x+y
  if primeMap[z]==1{
     return false
  }else if primeMap[z] == 2{
     return true
  }
  
  for i:=2;i*i<=z;i++{
    if z%i == 0{
	  primeMap[z]=1
	  return false
	}
  }
  primeMap[z] = 2
  return true
}