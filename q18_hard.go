//请解析IP地址和对应的掩码，进行分类识别。要求按照A/B/C/D/E类地址归类，不合法的地址和掩码单独归类。
//所有的IP地址划分为 A,B,C,D,E五类
//A类地址1.0.0.0~126.255.255.255;
//B类地址128.0.0.0~191.255.255.255;
//C类地址192.0.0.0~223.255.255.255;
//D类地址224.0.0.0~239.255.255.255；
//E类地址240.0.0.0~255.255.255.255
//私网IP范围是：
//10.0.0.0～10.255.255.255
//172.16.0.0～172.31.255.255
//192.168.0.0～192.168.255.255


package main
 
import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)
 
func validMaskOne(mask int, zero bool) bool {
    judge := mask == 254 || mask == 252 || mask == 248 || mask == 240 || mask == 224 || mask == 192 || mask == 128
    // 是否包含 0
    if zero {
        judge = judge || mask == 0
    }
    return judge
}
 
func validMask(mask []int) bool {
    if mask[0] == 255 {
        if mask[1] == 255 {
            if mask[2] == 255 {
                return validMaskOne(mask[3], true)
            } else {
                return validMaskOne(mask[2], true) && mask[3] == 0
            }
        } else {
            return validMaskOne(mask[1], true) && mask[2] == 0 && mask[3] == 0
        }
    }
    return validMaskOne(mask[0], false) && mask[1] == 0 && mask[2] == 0 && mask[3] == 0
}
 
func classIp(ip []int, iptype []int) {
    if ip[0] >= 1 && ip[0] <= 126 {
        iptype[0] += 1
        if ip[0] == 10 {
            iptype[6] += 1
        }
    } else if ip[0] >= 128 && ip[0] <= 191 {
        iptype[1] += 1
        if ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31 {
            iptype[6] += 1
        }
    } else if ip[0] >= 192 && ip[0] <= 223 {
        iptype[2] +=1
        if ip[0] == 192 && ip[1] == 168 {
            iptype[6] += 1
        }
    } else if ip[0] >= 224 && ip[0] <= 239 {
        iptype[3] += 1
    } else if ip[0] >= 240 && ip[0] <= 255 {
        iptype[4] += 1
    }
    // else {
    //    iptype[5] += 1
    // }
}
 
func main() {
    var (
        // 输入值
        ip [4]int
        mask [4]int
        // ip 分类: A，B，C，D，E，private
        iptype [7]int
    )
     
    reader := bufio.NewReader(os.Stdin)
     
    for {
        input, _, _ := reader.ReadLine()
        if len(input) == 0 {
            break
        }
         
        inputStr := string(input)
        inputArr := strings.Split(inputStr, "~")
        if len(inputArr) != 2 {
            iptype[5] += 1
            continue
        }
        maskArr := strings.Split(inputArr[1], ".")
        if len(maskArr) != 4 {
            iptype[5] += 1
            continue
        }
        ipArr := strings.Split(inputArr[0], ".")
        if len(ipArr) != 4 {
            iptype[5] += 1
            continue
        }
        valid := true
        for i := 0; i < 4; i++ {
            mask[i], _ = strconv.Atoi(maskArr[i])
            ip[i], _ = strconv.Atoi(ipArr[i])
            if (mask[i] < 0 || mask[i] > 255 || ip[i] < 0 || ip[i] > 255) {
                // 错误的 ip
                valid = false
                break
            }
        }
        if !valid {
            iptype[5] += 1
            continue
        }
        // 判断子网掩码
        v := validMask(mask[:])
        if !v {
            iptype[5] += 1
            continue
        }
         
        // 判断 ip 分类
        classIp(ip[:], iptype[:])
    }
     
    for ix, val := range iptype {
        if ix == 0 {
            fmt.Printf("%d", val)
        } else {
            fmt.Printf(" %d", val)
        }
    }
    fmt.Println()
}