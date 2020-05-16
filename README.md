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

### 变量类型

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

### 强制类型转换

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

## 循环语句

## 函数

## 指针



