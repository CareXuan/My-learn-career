package main //定义一个包

import "fmt" //通过import导入一个外部包fmt，可以使标准库包，也可以是自定义包或第三方包

func main(){ //main函数是go程序的入口
	fmt.Printf("Hello World!");
}


/**
通过go bulid demo1.go可在当前目录下生成一个可执行文件demo1

之后在当前目录执行./demo1即可运行编译好的文件
 */