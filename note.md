# Go 学习

## 常用命令

- go get 获取远程包（需要提前安装 git 或 hg）
- go run 运行程序
- go bulid 测试编译，检查是否有编译错误
- go fmt 测试化源码
- go install 编译包文件并编译整个程序
- go test 运行测试文件
- go doc 查看文档

## Go 编程基础

- Go 程序是通过 packge 来组织的
- 只有 package 名称为 maind 的包可以包含 main 函数
- 一个可执行程序有且仅有一个 main 包
- 通过 import 关键字导入其他非 main 包
- 通过 const 关键字进行常量定义
- 通过在函数体外部使用 var 关键字来进行全局变量的声明与赋值
- 通过 type 关键字来进行结构（struct）或接口（interFace）的声明
- 通过 func 关键字来进行函数的声明

> 关键字

- inport
  - 设置包别名：import asd "fmt"
  - 省略导入：import . "fmt"（不建议使用）

> 既定规则 可见性规则

- 使用大小写来决定该常量、变量、类型、接口、结构或函数是否可以被外部包所调用
- 函数名首字母小写即为 private
- 函数首字母大写即为 public

> 定义简写

- 声明多个常量
  ```Go
  const(
    PI      = 3.14
    contst1 = 2
    contst2 = 3
  )
  ```
- 声明多个全局变量
  ```Go
  var(
    name  = "GSY"
    name1 = 2
    name2 = 3
  )
  ```
- 声明多个一般类型

  ```Go
  type(
    nameType int
    type1 float32
    type2 string
  )
  ```

> Go 基本数据类型

- 布尔型 bool
  - 长度 1 字节
  - 取值范围：true false
  - 注意事项：不可以用数字代替 true 或 false
- 整形 int/uint
  - 根据平台运行可能为 32 或 64 位
- 8 位整形 int8/uint8
  - 长度：1 字节
  - 取值范围：-128 ~127/0 ~ 255
- 16 位整形 int16/uint16
  - 长度：2 字节
  - 取值范围：-32768 ~32767 / 0 ~ 65535
- 32 位整形 int32(rune 别名)/uint32
  - 长度：4 字节
  - 取值范围：-2^32/2 ~ 2^32/2-1/0 ~ 2^32-1
- 64 位整形 int64/uint64
  - 长度：4 字节
  - 取值范围：-2^64/2 ~ 2^64/2-1/0 ~ 2^64-1
- 浮点型 float32/float64
  - 长度：4/8 字节
  - 小数位：精确到 7/15 位小数
- 字节型 byte (uint 别名)
- 复数 complex64/complex128
  - 长度：8/16 字节
- 足够保存指针的 32 位或 64 位整数型 unintptr
- 其他值类型
  - array -->数组
  - struct --> 结构
  - string --> 字符串
- 引用类型
  - slice
  - map
  - chan
- 接口类型
  - interFace
- 函数类型
  - func

> 类型零值

- 零值不等于空值，而是当变量被声明为某种类型后的默认的值，通常情况下值类型的默认值为 0，bool 为 false，string 为空字符串

  ```Go
    var a int
    fmt.Println(a)
    //  a = 0;
    var a bool
    fmt.Println(a)
    //  a = false;
    var a string
    fmt.Println(a)
    //  a = '';
  ```

> 设置类型别名

```Go
type(
  byte int8
  rune int32
  文本   string
)
```

> 单个变量的声明与赋值

- 变量的声明格式：var <变量名称> <变量类型>
  ```Go
  var name string
  ```
- 变量的赋值：<变量名称> = <表达式>
  ```Go
  name  = "GSY"
  ```
- 声明的同时赋值：var <变量名称> [变量类型] = <表达式>
  ```Go
  var name string  = "GSY"
  ```
- 最简声明赋值(系统会根据表达式结果自动推断合适的变量类型)

  ```Go
  name := "GSY"
  ```

> 多个变量的声明与赋值

- 全局变量的声明可使用 var()方式进行简写
  ```Go
  var(
    aaa = "hello"
    sss,bbb = 1, 2
  )
  ```
- 全局变量的声明不可以省略 var，但可使用并行方式
- 所有变量都可使用类型推断
- 局部变量不可以使用 var()方式简写，只能使用并行方式

  ```Go
  var a,b,c,d int = 1,2,3,4
  ```

> 变量类型的转换

- Go 中不存在隐式转换，所有类型必须显式声明
- 转换只能发生在两种相互兼容的类型之间
- 类型转化格式为：

  ```Go
    <ValueA> [:]= <TypeOfValueA> (<ValueB>)
  ```

- 转换示例：
  ```Go
    var a float32 = 1.34
    // var b int
    b := int(a)
    // b = int(a)
    fmt.Println(b)
  ```
- int 65 转换为 string 类型

  ```Go
    var a int = 65
    b := string(a)
    fmt.Println(b)
    // 结果为 A
    // 因为string()表示将数据转换为文本格式，因为计算机中存储任何东西本质上都是数字，因此此函数自然地认为我们需要的用数字65表示的文本A
    // ----------------------------------------------------------
    // 引入包 strconv
    var a int = 65
    b := strconv.Itoa(a)
    fmt.Println(b)
    // 结果为 65
  ```

> 常量的定义

- 常量的值在编译时就已经确定
- 常量的定义格式与变量基本相同
- 等号右侧必须是常量或常量表达式
- 常量表达式中函数必须是内置函数
- 声明方式

  ```Go
    const a int = 1
    const b  = "A"
    // 同时定义多个变量
    const(
      text = "123"
      length = len(text)
      num = b = 20
    )
    const c,d,e  =  1,2,3
    const text2,num3,num4  =  "345",12,false

  ```

> 运算符

- Go 中运算符均是从左至右结合
- 优先级：从高到低
- 一元运算符
  - (^) (!)
- 二元运算符

  - (\*) (/) (%) (<< 左移) (>> 右移) (&) (&^)

    - & 运算符
    - 计算方式：二进制同位都为 1 才为 1，不然为 0

    ```GO
      // & 二进制运算
      var a = 6
      // 6 转换为二进制为：0110
      var b = 11
      // 11 转换为二进制为：1011
      // 0110
      // 1011
      //   =
      // 0010 转为十进制为 2
      fmt.Println(a & b)
      // 结果 2
    ```

    - | 运算符
    - 计算方式：二进制同位只要有一个为 1 结果就为 1，不然为 0

    ```GO
      // | 二进制运算
      var a = 6
      // 6 转换为二进制为：0110
      var b = 11
      // 11 转换为二进制为：1011
      // 0110
      // 1011
      //   =
      // 1111 转为十进制为 15
      fmt.Println(a | b)
      // 结果 15
    ```

    - ^ 运算符
    - 计算方式：二进制同位只有一个为 1 结果才为 1，不然为 0

    ```GO
      // ^ 二进制运算
      var a = 6
      // 6 转换为二进制为：0110
      var b = 11
      // 11 转换为二进制为：1011
      // 0110
      // 1011
      //   =
      // 1101 转为十进制为 13
      fmt.Println(a ^ b)
      // 结果 13
    ```

    - &^ 运算符
    - 计算方式：二进制第二位为 1 结果为 0，不然为 1

    ```GO
      // &^ 二进制运算
      var a = 6
      // 6 转换为二进制为：0110
      var b = 11
      // 11 转换为二进制为：1011
      // 0110
      // 1011
      //   =
      // 0100 转为十进制为 4
      fmt.Println(a &^ b)
      // 结果 4
    ```

  - (+) (-) (|) (^)
  - (==) (!=) (<) (<=) (>=) (>)
  - (&&)
  - 计算方式：与 & 一样，只不过前面不成立后面不会执行

    ```Go
      a := 0
      if a > 0 && (10/a) > 1 {
        fmt.Println("大于")
      }
      fmt.Println("小于等于0")
    ```

  - (||)

> 结合常量的 itoa 与<<运算符实现计算机存储单位的枚举

```Go
const (
  B float64 = 1 << (iota * 10)
  KB
  MB
  GB
  TB
  PB
)
func main() {
fmt.Println(B)
fmt.Println(KB)
fmt.Println(MB)
fmt.Println(GB)
fmt.Println(TB)
fmt.Println(PB)
}
```

> 指针

- Go 虽然保留了指针，但是与其他编程语言不同的是，在 Go 中不支持指针运算以及"->"运算符，而直接采"."选择符来操作指针目标对象的成员
- 操作符"&"取变量地址，使用"\*"通过指针间接访问对象

```Go
  	a := 1
    var p *int = &a
    fmt.Println(*p)
    // 结果为 1
```

- 默认值为 nil 而非 null
- 递增递减语句
  - 在 Go 当中，++与--是作为语句，不是表达式

## 控制语句

> 判断语句 if

- 提交表达式没有括号

  ```GO
    a := 2
    if a > 1{}else{}
  ```

- 支持一个初始化表达式
  ```GO
    if a,b := 2,4; a > 1{}else{}
    // 在if语句中初始化变量，它的作用域仅在if中
  ```
- 左大括号必须和条件语句或 else 在同一行
- 支持单行模式
- 初始化语句中变量为 block 级别，同时隐藏外部同名变量
- 1.0.3 版本中编译器 BUG

> 判断语句 for

- Go 中只有 for 一个循环语句关键字，但支持三种形式
- 初始化和步进式可以是多个值
- 条件语句每次循环都会重新检查，因此不建议在条条件语句中使用函数，尽量提前计算好条件并以变量或常量代替
- 左大括号必须和条件语句在同一行

- > > for 用法

  - 无限循环

    ```Go
    var a int
    for {
      a++
      if a > 10 {
        break
      }
      fmt.Println(a)
    }
    ```

  - 先做表达式判断，后循环
    ```Go
    var a int
    for a < 10 {
      a++
      fmt.Println(a)
    }
    ```
  - 经典模式
    ```Go
    for i := 0; i < 10; i++ {
      fmt.Println(i)
    }
    ```

> 选择语句 switch

- 可以使用任何类型或表达式作为条件语句
- 不需要写 break,一但条件符合自动终止
- 如希望执行下一个 case，需使用 fallthrough 语句
- 支持一个初始化表达式（可以是并行方式），右侧需跟括号
- 左大括号必须和条件语句在同一行

- > > switch 用法

  ```Go
    a := 2
    switch a {
    case 0:
      fmt.Println(0)
    case 1:
      fmt.Println(1)
    default:
      fmt.Println("都不等于")
    }
  ```

> 跳转语句

- goto
- break
- continue
- 三个语句都可配合标签使用
- 标签名区分大小写，若不使用会造成编译错误
- break 和 continue 配合标签可用于多层循环的跳出
- goto 是调整执行位置，与其它 2 个语句配合标签结果并不相同
- break 使用->跳出指定标签循环（goto、continue 同理）

  ```Go
  Lable1:
  for {
  	for a := 0; a < 10; a++ {
  		fmt.Println(a)
  		if a > 5 {
  			break Lable1
  		}
  	}
  }
  fmt.Println("结束了")
  ```

## 数组 Array

- 定义数组的格：var <varName> [n] <type> , n >= 0

  ```Go
  var a [2] int
  a := [2] int{1, 1}
  a := [...]int{1, 1, 1, 3, 4, 5, 6}
  ```

- 数组的长度也是类型的一部分，因此具有不同长度的数组为不同类型
- 注意区分指向数组指针和指针数组
- 数组在 Go 中为值类型
- 数组之间可以用 == 或 != 进行比较，但不可以使用 < 或 >
  ```Go
  a := [...]int{1, 1}
  b := [...]int{1, 2}
  fmt.Println(a == b)
  ```
- 可以使用 new 创建数组，此方法返回一个指向数组的指针
  ```Go
  a := new([10]int)
  fmt.Println(a)
  ```
- Go 支持多维数组
  ```Go
  a := [2][3]int{
  {1, 2, 3},
  {4, 5, 6}}
  fmt.Println(a)
  ```
- Go 冒泡排序
  ```Go
  array := [...]int{23, 45, 7, 89, 23, 12, 563, 323, 12, 32, 45}
  fmt.Println(array)
  length := len(array)
  for i := 0; i < length; i++ {
  	for j := i + 1; j < length; j++ {
  		if array[i] < array[j] {
  			temp := array[i]
  			array[i] = array[j]
  			array[j] = temp
  		}
  	}
  }
  fmt.Println(array)
  ```

## 切片 Slice

- 其本身并不是数组，它指向底层的数组
- 作为变长的数组的替代方案，可以关联底层数组的局部或全部
- 为引用类型
- 可以直接创建或从底层数组中获取生成
  ```Go
  // 创建Slice
  // var s1 [3]int
  // 从已有数组中获取
  array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  s2 := array[5:10] // array 5, 6, 7, 8, 9] 包含第10位
  // 一直取到数组尾部
  // 1.
  // s2 := array[5:len(array)]
  // 2.
  s2 := array[5:]
  fmt.Println(s2)
  ```
- 使用 len 获取元素个数，使用 cap 获取容量
- 一般使用 make 创建
- 如果多个 slice 指向相同的底层数组，其中一个的值改变会影响全部
- make([] T , len , cap)
- 其中 cap 可以省略，则和 len 值相同
- len 表示存数的元素个数，cap 表示容量

> Reslice 从一个切片中重新获取一个切片

- Reclice 时索引以被 slice 的切片为准
  ```Go
  a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}
  s1 := a[2:9]
  fmt.Println(string(s1))
  s2 := s1[1:3]
  // cdefghi
  fmt.Println(string(s2))
  // de
  ```
- 索引不可以超过被 slice 的切片容量 cap 值
- 索引越界不会导致底层数组的重新分配，而是引发错误

> Append 在 slice 尾部追加元素

- 可以在 slice 尾部追加元素
- 可以将一个 slice 追加到另一个 slice 尾部
- 如果最终长度未超过追加到 slice 的容量则返回原始 slice
- 如果超过追加到 slice 的容量则将重新分配数组并拷贝原始数据
  ```Go
  s1 := make([]int, 3, 6)
  fmt.Printf("%p\n", s1)
  s1 = append(s1, 1, 3, 4)
  fmt.Printf("%p\n", s1)
  s1 = append(s1, 1, 3, 4)
  fmt.Printf("%p\n", s1)
  ```

> Copy 拷贝数组

```Go
  array := []int{1, 2, 3, 4, 5, 6, 7}
  array1 := []int{9, 9, 9, 0}
  copy(array, array1)
  fmt.Println(array, array1)
  // [9 9 9 0 5 6 7] [9 9 9 0]
```

## Map

- 类似其它编程语言中的哈希表或字典，以 key-value 形式存储数据
- key 必须是支持 ==或 != 比较运算的类型，不可以是函数、map 或 slice
- Map 查找比线性搜索快很多，但比索引访问数据的类型慢 100 倍
- Map 使用 make()创建，支持 := 简写方式
- make ([keyType] valueType , cap)，cap 表示容量，可省略

  ```Go
    // var m map[int]map[int]string
    // m = make(map[int]map[int]string)
    // var m = map[int]string{} //第一种
    // var m = make(map[int]string) //第二种
    m := make(map[int]string) //第三种
  ```

- 超出容量时会自动扩容，但尽量提供一个合理的初始值
- 使用 len 获取元素个数
- 键值对不存在时自动添加，使用 delete()删除某键值对
  ```Go
  delete(m,1)
  ```
- 使用 for range 对 map 和 slice 进行迭代操作

  ```Go
    // 对 map 进行迭代操作
    	for K:v := range map{
        ...
      }
    // 对 slice 进行迭代操作
    	for i:v := range slice{
        ...
      }
    sm := make([]map[int]string, 10)
    // ---------------------------------------
    for i := range sm {
      sm[i] = make(map[int]string, 1)
      sm[i][1] = "OK"
      fmt.Println(sm[i])
    }
    fmt.Println(sm)
  ```

- 对 map 进行排序

  ```Go
  package main

  import (
    "fmt"
    "sort"
  )

  func main() {
    m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
    s := make([]int, len(m))
    i := 0
    for k, \_ := range m {
    s[i] = k
    i++
  }

  sort.Ints(s)
  fmt.Println(s)
  }
  ```

- 对 map 进行键值调换

  ```Go
  m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
  s := make(map[string]int)

  for k, v := range m {
  	s[v] = k
  }

  fmt.Println(s)
  ```

## 函数 Function

- Go 函数不支持嵌套、重载和默认参数
- 但支持以下特性

  - 无须声明原型
  - 不定长度变参

    ```Go
    func A(a ... int) {
      // a 变成了一个slice
      ...
    }
    ```

  - 多返回值
  - 命名返回值参数
  - 匿名函数
    ```Go
    func main() {
    a := func() {
      fmt.Print("OK")
    }
    a()
    }
    ```
  - 闭包

    ```Go
    package main
    import "fmt"

    func main() {
      f := closure(10)
      fmt.Println(f(1))
    }

    func closure(x int) func(int) int {
      return func(i int) int {
        eturn x + i
      }
    }
    ```

- 定义函数使用关键字 func，且左大括号不能另起一行
- 函数也可以作为一种类型使用

## defer

- defer 的执行方式类似其它语言中的析构函数，在函数体执行结束后按照调用顺序的相反顺序逐个执行
- 即使函数发生严重错误也会执行
- 支持匿名函数的调用
- 常用于资源管理、文件关闭、解锁以及记录时间等操作
- 通过与匿名函数配合可在 return 之后修改函数计算结果
- 如果函数体内某个变量作为 defer 是匿名函数的参数，则在定义 defer 时即已获得了拷贝，否则则是引用某个变量的地址
- Go 没有异常机制，但有 panic/reover 模式来处理错误
- Panic 可以在任何地方引发，但 reover 只有在 defer 调用的函数中有效

  ```Go
  // 引发Panic ，随后使用 defer reover 跳过panic ,继续执行代码
  func main() {
  	A()
  	B()
  	C()
  }

  func A() {
    fmt.Println("Func A")
  }

  func B() {
    defer func() {
      if err := recover(); err != nil {
        fmt.Println("Recover in B")
      }
    }()
    panic("Painc in B")
  }

  func C() {
    fmt.Println("Func C")
  }
  ```

- 一般使用

  ```Go
  fmt.Println("a")
  defer fmt.Println("b")
  defer fmt.Println("c")
  // 结果为 a c b ，defer定义的操作从下往上执行
  ```

- 闭包使用

  ```Go
  for i := 0; i < 4; i++ {
  	defer func() {
  		fmt.Println(i)
  	}()
  }
  // 结果为 4 4 4 4
  // 因为匿名函数引用的是i的实际地址
  ```
