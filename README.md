[TOC]

# 第一章 Go语言介绍
## Go简介
- 开发快，运行快

- 没有对象，没有继承多态，没有泛型，没有try/catch

- 有接口，函数式编程，csp并发模型（goroutine + channel）
## 安装与开发环境
**官方下载**：https://golang.org

**国内下载**：https://studygolang.com/dl

**查看当前版本**（镜像功能需要版本在1.13以上）：go version

**当前生效的go语言环境配置**：go env

- GOPROXY 模块镜像地址，proxy.golang.org在大陆无法访问，使用goproxy.cn国内镜像

  go env -w GOPROXY=https://goproxy.cn,direct
  
- GOMUDULE需要打开，注意on必须小写

  go env -w GO111MODULE=on

**安装goimports模块**：go get -v golang.org/x/tools/cmd/goimports

**goland安装file watchers插件**：保存时自动格式化

## 创建项目

- New project
- Go Module(vgo)
- Proxy没有自动带入，复制粘贴https://goproxy.cn,direct
- New->Go File->Simple Application
- Setting->File Watchers->添加goimports的watcher，保存时代码格式化、自动导包等功能
- goland会自动生成go.mod文件，如果用vs需要手动go mod init xxx

 ```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(runtime.GOARCH) // 运行系统架构 输出：amd64
}
 ```



  # 第二章 基本语法

## 变量定义

```go
package main

import "fmt"

func variable() {
	// 1.初始化
    // 变量名在前，变量类型在后
    // 默认初始值int为0，string为""
    // Printf为格式化打印，%q会打印出空字符串
    // 变量定义了必须要用，不然会报红
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b) // 0 ""

	// 2.初始化赋值
    // 初始化可以赋值，可以多个变量同时赋值
	var c, d int = 3, 4
	var e string = "hello"
	fmt.Println(c, d, e)

	// 3. 类型推断
    // 初始化赋值可以不用写类型
	var f, g, h = 3, true, "def"
	fmt.Println(f, g, h)

	// 4. 初始化赋值的简化写法 :=
	i, j, k := 3, false, "hello"
	fmt.Println(i, j, k)
}

func main() {
	variable()
}
```

```go 
package main

import "fmt"

// 5. 方法外变量
// 方法外变量初始化赋值不可以用:=
// 不是全局变量，只在该package下可用
var aa = 3
var bb = false
var cc = "kkk"

// 多个变量可以集中定义，该方式在函数内也可以
var (
	dd = 3
	ee = false
	ff = "ggg"
)

func main() {
   variable()
   fmt.Println(aa, bb, cc)
   fmt.Println(dd, ee, ff)
}
```

## 内建变量类型

**变量类型**

- bool, string

- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr 

  加u就是无符号整数，不加u就是有符号整数

  有符号整数分为固定长度、不固定长度（根据操作系统32还是64位）

  uintptr就是指针

- byte, rune

  rune就是字符型，char只有一字节所以坑多，rune为32位四字节

- float32, float64, complex64, complex128

  complex复数类型，complex64实部虚部分别为32位

  欧拉公式 cmplx.Pow(math.E, 1i*math.Pi) + 1

**强制类型转换**

类型转换只有强制的，没有隐式类型转换

```go
var a, b = 3, 4
var c int
c = int(math.Sqrt(float64(a*a + b*b))) // 如果没有float64 int两次强转会报错，不能隐式互转
fmt.Println(c) // 5
```

## 常量与枚举

```go
func consts() {
   // 1. 常量的数值可以作为各种类型使用
   const fileName = "123.txt"
   const a, b = 3, 4
   var d int = a
   var c = math.Sqrt(a*a + b*b)
   fmt.Println(a, b, c, d, fileName)
   // 2. 常量不必全部大写 首字母大写的变量另有含义
   // 3. 常量也可以集中定义
   const (
      name = "ryan"
      age  = 18
   )
   fmt.Println(name, age)
}
```

```go
func emuns() {
   // 1.通过const自定义枚举类型
   const (
      cpp    = 0
      java   = 1
      python = 2
      golang = 3
   )
   fmt.Println(cpp, java, python, golang) // 0 1 2 3
   // 2. 通过iota自增值定义枚举类型
   // iota表示一组自增值
   // _表示跳过某个自增值
   const (
      beijing = iota
      _
      _
      shanghai
      guangzhou
      shenzhen
   )
   fmt.Println(beijing, shanghai, guangzhou, shenzhen) // 0 3 4 5
   // 3. b kb mb gb tb pb
   const (
      b = 1 << (10 * iota)
      kb
      mb
      gb
      tb
      pb
   )
   fmt.Println(b, kb, mb, gb, tb, pb) // 1 1024 1048576 1073741824...
}
```

## 条件语句

```go
// 1.参数列表变量名在前 变量类型在后
// 2.返回值类型在参数列表后
func bounded(v int) int {
   // 3.if判断条件不需要括号
   if v > 100 {
      return 100
   } else if v < 0 {
      return 0
   } else {
      return v
   }
}
```

```go
const filename = "2.4条件语句.txt"

// 1.函数可以有多个返回值
contents, err := ioutil.ReadFile(filename)
// 2.nil有点像java的null
if err != nil {
   fmt.Println(err)
} else {
   // 3.contents为bytes数组，用Printf可以打印出内容
   fmt.Printf("%s\n", contents)
}

// 4.可以在if里写执行和判断的逻辑 这样写contents和err只能在if内访问
if contents, err := ioutil.ReadFile(filename); err != nil {
    fmt.Println(err)
} else {
    fmt.Printf("%s\n", contents)
}
```

```go
func eval(a, b int, op string) int {
    var result int
    // 1.swtich会自动break 除非使用fallthrough
    // 2.switch变量 case变量
    switch op {
    case "+":
       result = a + b
    case "-":
       result = a - b
    case "*":
       result = a * b
    case "/":
       result = a / b
    default:
       // 3.相当于报错 终止程序执行
       panic("不支持的操作符")
    }
	return result
}

func grade(score int) string {
	var result string
    // 3.也可以swtich空 case里面写判断条件
	switch {
	case score < 0 || score > 100:
        // 4.Sprintf返回string
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		result = "不及格"
	case score <= 100:
		result = "及格"
	}
	return result
}
```

## 循环语句

```go
package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
)

func main() {
   sum := 0
   // 1.for同样不需要括号
   for i := 1; i <= 100; i++ {
      sum += i
   }

   fmt.Println(convertToBin(13)) // 十进制转二进制 1101

   readFile("2.5循环.txt")

   forever()
}

func convertToBin(n int) string {
   result := ""
   // 2.可以省略初始条件
   for ; n > 0; n /= 2 {
      lsb := n % 2
      result = strconv.Itoa(lsb) + result
   }
   return result
}

func readFile(filename string) {
   // 3.os.Open读取文件
   if file, err := os.Open(filename); err != nil {
      panic(err)
   } else {
      // 4.构建扫描器
      scanner := bufio.NewScanner(file)
      // 5.省略初始条件递增条件 只有结束条件 类似java的while
      for scanner.Scan() {
         fmt.Println(scanner.Text())
      }
   }
}

func forever() {
   // 6.死循环
   for {
      fmt.Println("abc")
      // 7.跳出循环
      break
   }
}
```

## 函数

```go
// 1.允许多个返回值
func div1(a, b int) (int, int) {
	return a / b, a % b
}
// 2.多返回值另一种写法 可以给返回值变量名
func div2(a, b int) (q, r int) {
	return a / b, a % b
}
// 3.方法内操作返回变量 最后直接return
func div3(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}
// 4.正常接受多参返回
q, r := div1(13, 3)
// 5.如果只需要其中一个参数
q, _ := div1(13, 3)
```

```go
// 1.计算方法可以改造成多返回
func math(a, b int, op string) (int, error) {
   switch op {
   case "+":
      return a + b, nil
   case "-":
      return a - b, nil
   case "*":
      return a * b, nil
   case "/":
      return a / b, nil
   default:
      return 0, fmt.Errorf("操作符错误：" + op)
   }
}
// 2.接受多返回值判断结果
if result, err := math(13, 2, "asdf"); err != nil {
	fmt.Println(err)
} else {
	fmt.Println(result)
}
```

```go
// 类似java函数式接口 把实现逻辑交到调用方
func math2(op func(int, int) int, a, b int) int {
   pointer := reflect.ValueOf(op).Pointer()    // 反射获取指针
   opName := runtime.FuncForPC(pointer).Name() // 获取方法名称
   fmt.Printf("Calling func %s with args (%d, %d)\n", opName, a, b)
   return op(a, b)
}

// Calling func main.main.func1 with args (3, 4)
// 81
fmt.Println(
    math2(func(a int, b int) int {
        return int(math.Pow(float64(a), float64(b)))
    }, 3, 4))
```

```go
// 1.函数的可变参数列表
func sum(numbers ...int) int {
   result := 0
   // 2.range遍历获得下标和值
   for _, v := range numbers {
      result += v
   }
   return result
}

fmt.Println(sum(1, 2, 3, 4, 5)) // 15
```

## 指针

**值传递？引用传递？**

go语言只有值传递一种方式

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b := 1, 2
	fmt.Println(reflect.TypeOf(a))   // int
	fmt.Println(&a)                  // 十六进制内存地址 0xc00000a0b8
	fmt.Println(reflect.TypeOf(&a))  // *int 指针类型
	fmt.Println(reflect.TypeOf(*&a)) // 指针前面加* 得到指针对应的值

	fmt.Println("方法外 起始值", a, b) // 1 2
	valSwap(a, b)
	fmt.Println("方法外 值传递结束", a, b) // 1 2

	a, b = 1, 2
	fmt.Println("方法外 起始值", a, b) // 1 2
	refSwap(&a, &b)
	fmt.Println("方法外 指针传递结束", a, b) // 2 1
}

// 值传递对原本的a b不影响
func valSwap(a, b int) {
	fmt.Println("方法内 值传递开始: ", a, b) // 1 2
	a, b = b, a
	fmt.Println("方法内 值传递结束: ", a, b) // 2 1
}

// 指针传递修改了对应的值
func refSwap(a, b *int) {
	fmt.Println("方法内 指针传递开始: ", a, b) // 0xc00000a0b8 0xc00000a0d0
	*a, *b = *b, *a
	fmt.Println("方法内 指针传递结束: ", a, b) // 0xc00000a0b8 0xc00000a0d0
}
```

# 第三章 内建容器

## 数组

```go
package main

import "fmt"

func main() {
   // 数组初始化及赋值
   var arr1 [5]int
   arr2 := [3]int{1, 3, 5}
   arr3 := [...]int{2, 4, 6}
   fmt.Println(arr1, arr2, arr3)
   var grid [4][5]bool
   fmt.Println(grid)

   // 遍历数组 用range 用i遍历也可以
   // 如果只用v的话 i用_占位
   for i, v := range arr3 {
      fmt.Println(i, v)
   }
}
```

```go
// 1.数组是值类型 传递的时候会进行拷贝
// 2.[10]int和[20]int是不同类型
fmt.Println(arr1) // [0 0 0 0 0]
printArr(arr1)
fmt.Println(arr1) // [0 0 0 0 0]

func printArr(arr [5]int) {
   arr[0] = 100 // [100 0 0 0 0]
   fmt.Println(arr)
   for i, v := range arr {
      fmt.Println(i, v)
   }
}
```

```go
// 1.可以通过指针传递进行操作
fmt.Println(arr1) // [0 0 0 0 0]
printArr2(&arr1)
fmt.Println(arr1) // [100 0 0 0 0]

func printArr2(arr *[5]int) {
   fmt.Println(arr) // &[0 0 0 0 0]
   // 2.可以直接操作指针下标修改数据
   arr[0] = 100
   fmt.Println(arr) // &[100 0 0 0 0]
}
```

## 切片的概念

## 切片的操作

## Map

## Map例题

## 字符和字符串处理

# 第四章 面向“对象”

## 结构体和方法

## 包和封装

## 扩展已有类型

## 使用内嵌来扩展已有类型

# 第五章 Go语言依赖管理

## 依赖管理

## GOPATH和GOVENDOR

## go mod的使用

## 目录的整理

# 第六章 面向接口

## 接口概念

## duck typing概念

## 接口的定义和实现

## 接口的值类型

## 接口的组合

## 常用系统接口

# 第七章 函数式编程

## 函数式编程

## 函数式编程例一

## 函数式编程例二

# 第八章 错误处理和资源管理

## defer调用

## 错误处理概念

## 服务器统一出错处理1

## panic和recover

## 服务器统一出错处理2

# 第九章 测试与性能调优

## 测试

## 代码覆盖率和性能测试

## 使用pprof进行调优

## 测试http服务器1

## 测试http服务器2

## 生成文档和示例代码

## 测试总结

# 第十章 Goroutine

## goroutine

## go语言的调度器

# 第十一章 Channel

## channel

## 使用Channel等待任务结束

## 使用Channel进行树的遍历

##  用select进行调度

## 传统同步机制

# 第十二章 迷宫的广度优先搜索

## 迷宫算法

## 迷宫代码实现

# 第十三章 http及其他标准库

## http标准库

## 其他标准库

## gin框架介绍

## 为gin添加middleware

