package main

import "fmt"
import "unsafe"
import "strconv"


// 1) Go源文件以"go"为扩展名。
// 2) Go应用程序的执行入口是main(方法。
// 3) Go语言严格区分大小写。
// 4) Go方法由一条条语句构成，每个语句后不需要分号(Go语言会在每行后自动加分号),这也体现出Golang的简洁性。
// 5) Go编译器是一行行进行编译的，因此我们一-行就写一条语句，不能把多条语句写在同一个，否则报错
// 6) go语言定义的变量或者import的包如果没有使用到，代码不能编译通过。
// 7)大括号都是成对出现的，缺一不可。

func main()  {
	fmt.Println("printhello");
	const age = 1
	//age = 10
	num := 1
	var xx float32 = 10.10
	fmt.Println(num)
	fmt.Println(age)
	fmt.Println(xx)
	var addr byte = 'A'
	fmt.Println(addr)
	fmt.Printf("%c",addr)

	var flag =  true
	fmt.Println(flag)
	fmt.Println(unsafe.Sizeof(flag))
	var tmp  = "123";
	var tmp1,_ = strconv.ParseInt(tmp,10,64)
	fmt.Println(int32(tmp))
}