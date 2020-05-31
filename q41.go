/**
Jessi初学英语，为了快速读出一串数字，编写程序将数字转换成英文：
如22：twenty two，123：one hundred and twenty three。
*/

package main
 
import (
    "fmt"
    "strconv"
)
 
var m1 = []string {"","one","two","three","four","five","six","seven","eight","nine","ten","eleven","twelve","thirteen","fourteen","fifteen","sixteen","seventeen","eighteen","nineteen"}
var m2 = []string {"","","twenty","thirty","forty","fifty","sixty","seventy","eighty","ninety"}
 
 
func parse(str string)string{
    var res string
    if len(str)==3{
        n,_:=strconv.Atoi(string(str[0]))
        if n!=0{
            res += m1[n]+" hundred"
        }
        m,_:=strconv.Atoi(string(str[1]))
        j,_:=strconv.Atoi(string(str[2]))
        if m+j!=0 && n!=0{
            res+=" and "
        }
        if m==1{
            res+=m1[j+10]
        }else if m>=2{
            res+=m2[m]
            if j!=0{
                res+=" "+m1[j]
            }
        }else{
            res+=m1[j]
        }
    }
    if len(str)==2{
        m,_:=strconv.Atoi(string(str[0]))
        j,_:=strconv.Atoi(string(str[1]))
        if m==1{
            res+=m1[j+10]
        }else if m>=2{
            res+=m2[m]
            if j!=0{
                res+=" "+m1[j]
            }
        }
    }
    if len(str)==1{
        n,_:=strconv.Atoi(string(str[0]))
        res+=m1[n]
    }
    return res
}
 
func main(){
    //reader := bufio.NewReader(os.Stdin)
    //str, _, _ := reader.ReadLine()
    for{
        //reader := bufio.NewReader(os.Stdin)
        //strr, _, err := reader.ReadLine()
        //str := string(strr)
        var str string
        _,err := fmt.Scanln(&str)
        if err!= nil {
            break
        }
        n,err:=strconv.Atoi(str)
        if err!=nil{
            fmt.Println("error")
            continue
        }
        if n==0{
            fmt.Println("zero")
            continue
        }
        if len(str)>6{
            str1 := parse(str[:len(str)-6])
            str2 := parse(str[len(str)-6:len(str)-3])
            str3:= parse(str[len(str)-3:])
            var res string
            if str1!=""{
                res+=str1+" million"
            }
            if str2!=""{
                res+=" " +str2+" thousand"
            }
            if str3!=""{
                res+=" "+str3
            }
            fmt.Println(res)
        }else if len(str)>3{
            str1:=parse(str[:len(str)-3])
            str2:=parse(str[len(str)-3:])
            var res string
            if str1!=""{
                res += str1+" thousand"
            }
            if str2!=""{
                res +=" "+parse(str[len(str)-3:])
            }
            fmt.Println(res)
        }else{
            fmt.Println(parse(str))
        }
    }
}