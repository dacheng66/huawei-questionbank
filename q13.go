package main
import (
   "bufio"
    "fmt"
    "os"
    "strings"
)

func main(){
    reader := bufio.NewReader(os.Stdin)
    data,_,_ := reader.ReadLine()
    sList := strings.Split(string(data)," ")
    var s string
    for i:=len(sList)-1;i>=0;i--{
        s += sList[i]
        s += " "
    }
    fmt.Println(s)
    
}