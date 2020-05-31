//假设渊子原来一个BBS上的密码为zvbo9441987,为了方便记忆，他通过一种算法把这个密码变换成YUANzhi1987，这个密码是他的名字和出生年份，怎么忘都忘不了，而且可以明目张胆地放在显眼的地方而不被别人知道真正的密码。
//他是这么变换的，大家都知道手机上的字母： 1--1， abc--2, def--3, ghi--4, jkl--5, mno--6, pqrs--7, tuv--8 wxyz--9, 0--0,就这么简单，渊子把密码中出现的小写字母都变成对应的数字，数字和其他的符号都不做变换，
//声明：密码中没有空格，而密码中出现的大写字母则变成小写之后往后移一位，如：X，先变成小写，再往后移一位，不就是y了嘛，简单吧。记住，z往后移是a哦。

package main

import(
   "bufio"
   "fmt"
   "os"
   "strings"
)

func main(){
   reader := bufio.NewReader(os.Stdin)
   var sMap = make(map[string]string)
   sMap["abc"] = "2"
   sMap["def"] = "3"
   sMap["ghi"] = "4"
   sMap["jkl"] = "5"
   sMap["mno"] = "6"
   sMap["pqrs"] = "7"
   sMap["tuv"] = "8"
   sMap["wxyz"] = "9"
   letter := "abcdefghijklmnopqrstuvwxyz"
   var rs string
   
   for {
      data,_,_ := reader.ReadLine()
	  if len(data) == 0 {
	    break
	  }
	  s := string(data)
	  for _,i := range s{
	     if i>='a'&&i<='z'{
		    for k,v := range sMap{
			   if strings.Contains(k,string(i)){
			       rs += v
				   break
			   }
			   
			}
		 }
		 
		 if i>='A' && i<='Z'{
		    si := strings.ToLower(string(i))
			x := strings.Index(letter,si)
			if x != len(letter)-1{
			   rs += string(letter[x+1]) //注意
			   continue
			}else{
			   rs += "a"
			   continue
			}
		 }
		 
		 if i>='0' && i <='9'{
		     rs += string(i)
			 continue
		 }
	  }
	  
	  fmt.Println(rs)	 
   }
}