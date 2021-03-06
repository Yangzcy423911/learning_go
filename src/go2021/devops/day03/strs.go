package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"unsafe"
)

func main1() {
	var ch byte = 'A'

	fmt.Println(ch)              //  65
	fmt.Printf("%c, %d", ch, ch) // A, 65

}

func main2() {
	fmt.Println("Good")
	fmt.Printf("%c", '\n')
	fmt.Print("Work")
}

func main3() {
	var a rune = '我'             // rune用于汉字, byte use A
	fmt.Printf("%T, %v\n", a, a) // int32, 25105
	fmt.Printf("%c", a)          // 我
}

func main4() {
	var cmd string = "date"
	// fmt.Scanf("%s", &cmd) // input cmd
	fmt.Println("you will exec command:", cmd)
	exec.Command(cmd).Run()
}

// func main() {
// 	var cmd string = "uptime"
// 	fmt.Println("you will exec command:", cmd)

// 	exec.Command(cmd).Run()
// }

func main5() {
	var cmd1 = "cale"
	var cmd2 = "哈哈"

	fmt.Printf("%T %T", cmd1, cmd2)
}

func main6() {
	var mystr01 string = "hello"
	var mystr02 [5]byte = [5]byte{104, 101, 108, 108, 111}
	fmt.Printf("mystr01: %s\n", mystr01)
	fmt.Printf("mystr02: %s", mystr02)
}

func main7() {
	var country string = "hello,中国" // utf-8编写，英文字母和符合占用1个字节,汉字占三个字节
	fmt.Println(len(country))
}

func main8() {
	var ch = 'A'
	var chstr = "A"
	fmt.Printf("%T | %T\n", ch, chstr) // int32 | string
	fmt.Println(unsafe.Sizeof(ch))     // 4 int32
	fmt.Println(unsafe.Sizeof(chstr))  // 16
	fmt.Println(len(chstr))
	fmt.Printf("%c\n", chstr[0]) // A
	// chstr[0] = 'x'  // cannot assign to chstr[0]
	chstr = "XYZ" // 字符串可以修改跳到字符串地址，不可修改字符
	fmt.Println(chstr)
	fmt.Printf("%c %d\n", chstr[0], chstr[0]) // X 88
	fmt.Printf("%c %d\n", chstr[1], chstr[1]) // Y 89
}

func main9() {
	var mystr01 string = "\\r\\n"
	var mystr02 string = `\r\n`  // 无转义字符串
	var mystr03 string = `您好啊!`  // 自带换行符
	fmt.Println(mystr01)         // \r\n
	fmt.Println(mystr02)         // \r\n
	fmt.Printf(" %q\n", mystr01) //  "\\r\\n" 原始字符串
	fmt.Println(mystr03)

}
func main10() {
	var cmd string = "uptime123"
	fmt.Printf("you will exec cmd: %s\n", cmd[0:6])
	exec.Command(cmd[0:6]).Run()
}

func main11() {
	var cmd string = "123uptime"
	fmt.Printf("you will exec cmd: %s\n", cmd[3:])
	exec.Command(cmd[3:]).Run()
}

func main12() {
	var cmd = "uptime"
	for i := 0; i < len(cmd); i++ {
		fmt.Printf("%c %d\n", cmd[i], cmd[i])
	}
	fmt.Println("--------------------------")
	for i, ch := range cmd { // 循环每一个字符，每个汉字
		fmt.Printf("%c %d\n", ch, i)
	}
}

func main13() {
	var string1, string2, string3, string4 = "12", "12.4", "true", "A"
	var a int
	var b float32
	var c bool
	var d byte
	fmt.Sscanf(string1, "%d", &a)
	fmt.Sscanf(string2, "%f", &b)
	fmt.Sscanf(string3, "%t", &c)
	fmt.Sscanf(string4, "%c", &d)
	fmt.Println(a, b, c, d) // 12 12.4 true 65
}

func main() {
	num := 100
	str := strconv.Itoa(num) // 快速整数转换为字符串
	fmt.Printf("%T value %#v\n", str, str)

	str1 := "110"
	str2 := "s100"

	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Printf("%v 转换失败!", str1)
	} else {
		fmt.Printf("%T value: %#v\n", num1, num1)
	}

	num2, err := strconv.Atoi(str2)

	if err != nil {
		fmt.Printf("%v 转换失败!\n", str2)
	} else {
		fmt.Printf("%T value: %#v\n", num2, num2)
	}

	nums := true

	str5 := strconv.FormatBool(nums)
	fmt.Printf("%T | %#v\n", str5, str5) // string | "true"

}
