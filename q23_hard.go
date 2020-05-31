//计算最少出列多少位同学，使得剩下的同学排成合唱队形
//说明：
/*
N位同学站成一排，音乐老师要请其中的(N-K)位同学出列，使得剩下的K位同学排成合唱队形。
合唱队形是指这样的一种队形：设K位同学从左到右依次编号为1，2…，K，他们的身高分别为T1，T2，…，TK，   则他们的身高满足存在i（1<=i<=K）使得T1<T2<......<Ti-1<Ti>Ti+1>......>TK。
你的任务是，已知所有N位同学的身高，计算最少需要几位同学出列，可以使得剩下的同学排成合唱队形。
请注意处理多组输入输出！
*/

package main

import(
  "fmt" 
)

func main(){
   total := 0
   for {
     _,err := fmt.Scanf("%d",&total)
     if err != nil{
        break
     }
	 
	 height := make([]int,total)
	 for i:=0;i<total;i++{
	     fmt.Scanf("%d",&height[i])
	 }
	 
	 //顺序递增子序列
	 inc := make([]int,total)
	 //逆序递增子序列
	 dec := make([]int,total)
	 
	 //0开始
	 for i:=1;i<total;i++{
	    for j:=i-1;j>=0;j--{
		     if height[i] > height[j] && inc[i]<inc[j]+1{
			     inc[i] = inc[j] + 1
			 }
		}
	 }
       
     for i:= total-2;i>=0;i--{
	    for j:= i+1;j<total;j++{
		   if height[i] > height[j] && dec[i] < dec[j]+1{
		     dec[i] = dec[j]+1
		   }
		}
	 }
	 
	 max := 0
	 for i:=0;i<total;i++{
	    if max < inc[i] + dec[i]{
		    max = inc[i] + dec[i]
		
		}
	 }
	 //子序列-2，每个数在所在队列的人数+1（自己在递增和递减中被重复计算）
	 fmt.Println(total-max-1)
   }
}