package main

import "fmt"

type pack struct{
  v int
  p int
}

func main(){
   var V,m int
   fmt.Scan(&V,&m)
   v /= 10
   
   arr := make([][]pack,m+1)
   var v,p,q int
   
   //建立二维数组存主副  
   for i:=1;i<=m;i++{
      fmt.Scan(&v,&p,&q)
      v /= 10
      p*=v
      if q == 0{
        if arr[i] == nil{
           arr[i] = make([]pack,1)
        }
        arr[i][0] = pack{v,p}   //主物品都是放 arr[i][0]		
      } else{                  //存副物品
         if arr[q] == nil{
		    arr[q] = make([]pack,2)
			arr[q][1] = pack{v,p}
		 }else{
		    arr[q] = append(arr[q],pack{v,p})
		 }
 
      }	 
   }
   
   // 创建一个一维数组，表示价格为i时的最大价值
   dp := make([]int,V+1)
   
   //每次循环表示物品存放的个数
   for i:=1;i<=m;i++{
     if arr[i]== nil{
	   continue
	 }
	 
	 //计算包含主件最大价值，j大于主件物品
	 for j:=V;j>=arr[i][0].v;j--{
	     //ans为最大V时的价值
	     ans := dp[j]
		 //only main
		 //增加一件主件的总价值【选择大的主件】与len(arr[i]) > 1 不冲突 
		 //dp[j-arr[i][0].v]的值为（V-当前主件价格）的最大值，可包含之前主件的值，并非dp[j]的值
		 if ans < dp[j-arr[i][0].v] + arr[i][0].p{
		    ans = dp[j-arr[i][0].v] + arr[i][0].p
		 }
		 //每个主可以有0，1，2个副
		 if len(arr[i]) > 1{
             v := arr[i][0].v + arr[i][1].v // main + sub1
             p := arr[i][0].p + arr[i][1].p
			 //j为多少钱
             if j >= v && ans < dp[j-v]+ p {
                    ans = dp[j-v]+p
             }
			 if len(arr[i]) > 2 {
                    // main + sub1 + sub2
                    v += arr[i][2].v
                    p += arr[i][2].p
                    if j >= v && ans < dp[j-v]+ p {
                        ans = dp[j-v]+p
                    }
                    // main + sub2
                    v -= arr[i][1].v
                    p -= arr[i][1].p
                    if j >= v && ans < dp[j-v]+ p {
                        ans = dp[j-v]+p
                    }
                }
         }		 
		 dp[j] = ans
	 }
   }   
   
   fmt.Println(dp[V]*10)
}
   
      

