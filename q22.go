/*
“某商店规定：三个空汽水瓶可以换一瓶汽水。小张手上有十个空汽水瓶，她最多可以换多少瓶汽水喝？
”答案是5瓶，方法如下：先用9个空瓶子换3瓶汽水，喝掉3瓶满的，喝完以后4个空瓶子，用3个再换一瓶，
喝掉这瓶满的，这时候剩2个空瓶子。然后你让老板先借给你一瓶汽水，喝掉这瓶满的，
喝完以后用3个空瓶子换一瓶满的还给老板。如果小张手上有n个空汽水瓶，最多可以换多少瓶汽水喝？
*/
package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   for {
      //喝了多少瓶
      var n int
	  data,_,_ := reader.ReadLine()
	  if len(data) == 0{
	     break
	  }
	  
	  s := string(data)
	  x,_ := strconv.Atoi(s)
      if x == 0{
	     break
	  }
	  
	  for {
	     if x < 2{
		    fmt.Println(n)
			break
		 }
	     if x % 3 == 0{
		    n += x / 3
			x = x / 3
		 }
		 
		 if x % 3 == 1{
		    n += x / 3
			x = x / 3 + 1		   
		 }
		 
		 if x % 3 == 2{
		    n += (x+1) / 3
			x = (x-2) / 3
		 }
	  }
   }
}