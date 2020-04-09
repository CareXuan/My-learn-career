#### 标识符
---
标识符分两类，一类是语言设计者确定的，预留的标识符，另一类就是编程者可以自定义的标识符

标识符构成规则是，开头必须是字母或下划线，后边跟上多个字符，数字，下划线，并且区分大小写

- 预声明标识符
1. 关键字
	1. 引导程序整体结构的8个关键字
	```
	package import const var func defer go return
	```
	2. 声明数据结构的4个关键字
	```
	struct interface map chan
	```
	3. 控制程序结构的13个关键字
	```
	if else                                    //if else语句
	for range break continue                   //for循环
	switch select type case default fallthrough//switch和select语句
	goto                                       //goto跳转语句
	```
2. 内置数据类型标识符
	1. 数值
		1. 整型
		```
		byte int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64 uintptr
		```
		2. 浮点型
		```
		float32 float64
		```
		3. 复数
		```
		complex64 complex128
		```
	2. 字符和字符串型
		```
		string rune
		```
	3. 接口型
		```
		error
		```
	4. 布尔型
		```
		bool
		```
3. 内置函数
```
make new len cap append copy delete panic recover close complex real image print println
```
4. 常量值标识符
```
true false   //bool类型的两个常量值
iota             //用在连续的枚举类型的声明
nil               //指针、引用型变量的默认值
```
5. 空白标识符
```
_
```