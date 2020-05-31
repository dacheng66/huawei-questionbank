/**
题目描述
查找和排序
题目：输入任意（用户，成绩）序列，可以获得成绩从高到低或从低到高的排列,相同成绩
都按先录入排列在前的规则处理。
例示：
jack      70
peter     96
Tom       70
smith     67
从高到低  成绩
peter     96
jack      70
Tom       70
smith     67

从低到高
smith     67
Tom       70
jack      70
peter     96

注：0代表从高到低，1代表从低到高
*/

package main
 
import (
    "fmt"
)
 
type Pack struct {
    Name string
    Score int
}
 
func main() {
    for {
        var num,order int
        n,err:=fmt.Scan(&num,&order)
        if n==0 || err!=nil {
            return
        }
        data:=make([]*Pack,num)
        for i:=0;i<num;i++{
            p:=&Pack{}
            fmt.Scan(&p.Name,&p.Score)
            data[i]=p
        }
        sort(data,order)
         
        for i:=0;i<num;i++ {
            fmt.Printf("%s %d\n",data[i].Name,data[i].Score)
        }
         
    }
}
 
func sort(data []*Pack,order int) {
    length:=len(data)
    if order==1 {
        for i:=0;i<length-1;i++ {
            for j:=0;j<length-1-i;j++{
                if data[j].Score>data[j+1].Score {
                    data[j],data[j+1]=data[j+1],data[j]
                }
            }
        }
    }else {
         for i:=0;i<length-1;i++ {
            for j:=0;j<length-1-i;j++{
                if data[j].Score<data[j+1].Score {
                    data[j],data[j+1]=data[j+1],data[j]
                }
            }
        }
    }
}