package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  data,_,_ := reader.ReadLine()
  num,_ := strconv.ParseInt(string(data),10,64)
  sBin := strconv.FormatInt(num,2)
  var l int
  for _,s := range sBin{
    if s == '1'{
	  l+=1
	}
  }
  fmt.Println(l)
}