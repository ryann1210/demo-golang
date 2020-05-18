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

```go
func main() {
   arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

   // 1.切片是对数组的视图 修改视图会修改原数组
   s1 := arr[:]
   fmt.Println("arr: ", arr)
   fmt.Println("s1: ", s1)
   updateSlice(s1)
   fmt.Println("s1: ", s1)
   fmt.Println("arr: ", arr)

   // 3.修改s2 s1和arr都会改变
   s2 := s1[2:]
   fmt.Println("arr: ", arr)
   fmt.Println("s1: ", s1)
   fmt.Println("s2: ", s2)
   updateSlice(s2)
   fmt.Println("s2: ", s2)
   fmt.Println("s1: ", s1)
   fmt.Println("arr: ", arr)
}

// 2.视图的参数不需要写长度
func updateSlice(s []int) {
   s[0] = 100
}
```

```go
// 1.slice可以向后扩展 不可以向前扩展
// 数据结构包括：
// ptr 指向第一个元素
// len slice的长度
// cap ptr到数组最后一个元素的长度
// 总结：s[i]不可以超越len(s) 向后扩展不可以超过底层数组cap(s)
arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s3 := arr2[2:6] // [2 3 4 5]
s4 := s3[3:5]   // [5 6]
fmt.Println(s3, s4)
// fmt.Println(s3[4]) 报错index out of range
fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n", s3, len(s3), cap(s3)) // s3=[2 3 4 5], len(s3)=4, cap(s3)=6
fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n", s4, len(s4), cap(s4)) // s4=[5 6], len(s4)=2, cap(s4)=3
```

## 切片的操作

```go
   arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
   s1 := arr[2:6]
   s2 := s1[3:5]
   s3 := append(s2, 10)
   // 添加元素时如果超越cap 系统会重新分配更大的底层数组
   s4 := append(s3, 11)
   s5 := append(s4, 12)
   fmt.Println(arr) // [0 1 2 3 4 5 6 10]
   fmt.Println(s1)  // [2 3 4 5]
   fmt.Println(s2)  // [5 6]
   fmt.Println(s3)  // [5 6 10]
   fmt.Println(s4)  // [5 6 10 11]
   fmt.Println(s5)  // [5 6 10 11 12]
```

```go
// cap翻倍的方式扩容
var s6 []int // s6 == nil
for i := 0; i < 10; i++ {
   s6 = append(s6, 2*i+1)
   fmt.Printf("s6=%v, len(s6)=%d, cap(s6)=%d\n", s6, len(s6), cap(s6))
}
/*
s6=[1], len(s6)=1, cap(s6)=1
s6=[1 3], len(s6)=2, cap(s6)=2
s6=[1 3 5], len(s6)=3, cap(s6)=4
s6=[1 3 5 7], len(s6)=4, cap(s6)=4
s6=[1 3 5 7 9], len(s6)=5, cap(s6)=8
s6=[1 3 5 7 9 11], len(s6)=6, cap(s6)=8
s6=[1 3 5 7 9 11 13], len(s6)=7, cap(s6)=8
s6=[1 3 5 7 9 11 13 15], len(s6)=8, cap(s6)=8
s6=[1 3 5 7 9 11 13 15 17], len(s6)=9, cap(s6)=16
s6=[1 3 5 7 9 11 13 15 17 19], len(s6)=10, cap(s6)=16
*/
```

```go
s7 := []int{2, 4, 6, 8}
// make创建slice 可指定len和cap
s8 := make([]int, 16)
s9 := make([]int, 16, 32)
fmt.Printf("s7=%v, len(s7)=%d, cap(s7)=%d\n", s7, len(s7), cap(s7))
fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))
fmt.Printf("s9=%v, len(s9)=%d, cap(s9)=%d\n", s9, len(s9), cap(s9))
/*
s7=[2 4 6 8], len(s7)=4, cap(s7)=4
s8=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len(s8)=16, cap(s8)=16
s9=[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len(s9)=16, cap(s9)=32
*/
```

```go
// s7拷贝到s8
copy(s8, s7)
fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))
// s8=[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len(s8)=16, cap(s8)=16
```

```go
// 删除第三个元素
s8 = append(s8[:3], s8[4:]...)
fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))

// 删除头部元素
s8 = s8[1:]
fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))

// 删除末尾元素
s8 = s8[:len(s8)-1]
fmt.Printf("s8=%v, len(s8)=%d, cap(s8)=%d\n", s8, len(s8), cap(s8))
```

## Map

**Key的条件**

- map使用哈希表 必须可以比较相等
- 除了slice、map、function的内建类型都可以作为key
- Struct类型不包括上述字段，也可作为key

```go
// map初始化
m1 := map[string]string{"name": "ryan", "sex": "male"} // 空map
m2 := make(map[string]string)                          // 空map
var m3 map[string]string                               // nil
```

```go
// range方式遍历 无序输出
for k, v := range m1 { // 这里如果用m2m3遍历也不会出错
   fmt.Println(k, v)
}
```

```go
// 通过key取值 如果key不存在返回空字符串
name := m1["name"]
fmt.Println(name)
// 接受两个参数判断是否存在
if sex, exist := m1["sex"]; exist {
    fmt.Println(sex)
} else {
    fmt.Println("Key not exist.")
}

// 通过key删除键值对
delete(m1, "sex")
```

## Map例题

```go
package main

import "fmt"

/*
寻找最长不含有重复字符的子串
*/
func main() {
   result := lengthOfLongestSubstring("helloworld")
   fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
   // 记录每个字符最后出现的位置
   lastOccurred := make(map[byte]int)
   // 有效子串的起始
   start := 0
   // 最大长度
   maxLength := 0
   // 遍历字符串
   for i, ch := range []byte(s) {
      // 当前字符已出现过 并且在有效子串内
      if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
         // 有效子串的起始值+1
         start = lastI + 1
      }
      // 重新计算最大长度
      if i-start+1 > maxLength {
         maxLength = i - start + 1
      }
      // 更新当前字符最后出现的位置
      lastOccurred[ch] = i
   }
   return maxLength
}
```

## 字符和字符串处理

```go
func main() {
   s := "Yes我爱天安门!" // 字符串初始化是utf8编码

   fmt.Println(len(s))                    // 19 中文占3字节
   fmt.Println(utf8.RuneCountInString(s)) // 9 以rune为单位返回字符串长度

   for i, ch := range s { // ch就是rune 类型为int32 固定4字节
      fmt.Printf("(%d %X)", i, ch) // (0 59)(1 65)(2 73)(3 6211)(6 7231)(9 5929)(12 5B89)(15 95E8)(18 21)
   }
   fmt.Println()
   for i, ch := range []rune(s) { // 转成rune数组再输出
      fmt.Printf("(%d %c)", i, ch) // (0 Y)(1 e)(2 s)(3 我)(4 爱)(5 天)(6 安)(7 门)(8 !)
   }
   fmt.Println()

   bytes := []byte(s)
   fmt.Printf("%x\n", []byte(s)) // 596573e68891e788b1e5a4a9e5ae89e997a821
   for len(bytes) > 0 {
      ch, size := utf8.DecodeRune(bytes)
      bytes = bytes[size:]
      fmt.Printf("%c ", ch) // Y e s 我 爱 天 安 门 !
   }
   fmt.Println()
}
```

```go
// 完善上个章节例题对中文的支持
// strings包里有很多对字符串的操作
result := lengthOfLongestSubstringRune("一二三而一")
fmt.Println(result)

func lengthOfLongestSubstringRune(s string) int {
   // 记录每个字符最后出现的位置
   lastOccurred := make(map[rune]int)
   // 有效子串的起始
   start := 0
   // 最大长度
   maxLength := 0
   // 遍历字符串
   for i, ch := range []rune(s) {
      // 当前字符已出现过 并且在有效子串内
      if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
         // 有效子串的起始值+1
         start = lastI + 1
      }
      // 重新计算最大长度
      if i-start+1 > maxLength {
         maxLength = i - start + 1
      }
      // 更新当前字符最后出现的位置
      lastOccurred[ch] = i
   }
   return maxLength
}
```

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

