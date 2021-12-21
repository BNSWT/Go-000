# Go入门

## Go基本语法和Web框架起步

main函数要点：

- 无参数、无返回值
- main方法必须要在main包里面
- `go run main.go`就可以执行
- 如果文件不叫`main.go`，则需要`go build`得到可运行的文件，直接运行就可以。

包声明：

- 文件夹名字和包名可以不一样
- 同一个文件夹下包名必须一样
- 引入包语法形式： `import [alias] xxx`
- 如果一个包引入了但是没有使用，会报错
- 匿名引入（引入但不使用，是为了用这个包的初始化方法）：前面多一个下划线

string常量：

- 双引号引起来，则内部双引号需要使用\转义
- 反引号引起来，则内部反（`）引号需要使用\转义。反引号通常用于大量字符串的情况，字符串内换行会自己换行。

string长度计算：

- `len(str)`输出的是字节长度，和编码无关
- 需要其他编码相关的方法,比如`utf8.RuneCountInString`

string拼接：

- string不能和别的类型拼接，只能和string拼接

rune类型：

- 直观理解，就是字符
- 不是byte
- 本质是int32，一个rune四个字节
- rune在很多语言里面是没有的，与之对应的是，golang没有char类型。rune不是数字，也不是char，也不是 byte
- 实际中不太常用

一般变量类型：

- bool:true, false
- int8, int16, int32, int64, int
- uint8, uint16, uint32, uint64, uint
- float32, float64

Byte类型：

- 字节，本质是uint8
- 对应的操作包在bytes上

变量声明：

- var,语法:  `var name type = value`
  - 局部变量
  - 包变量
  - 块声明
- 变量名首字符是否大写控制了访问性：大写包外可访问，小写即使是子包也不能访问。
- golang支持类型推断
- 强类型语言，不会作任何隐式类型转换
- `:=`只能用于局部变量声明，即方法内部
- golang使用类型推断来推断类型。数字会被理解为int或者float64。（所以要其他类型的数字，就得用var来声明）
- 变量声明后必须使用
- 同作用域下，一个变量只能使用一次。
- 常量声明只需把`var`改为`const`即可


方法声明：

- `func name(name type)(name type, name type){}`
- 方法的首字符是否大写同样决定了作用域
- 若返回值只有类型，没有名字，可以return a,b 
- 若返回值有名字，可以给其名字赋值后直接返回
- 用到不定参数时，不定参数要放在最后面
- 使用_忽略返回值

数组：

- 声明： `var name [cap]type`
- 初始化时要指定长度
- len(name)和cap(name)操作用于获取数组长度


切片：

- 声明： `var name []type`
- len指切片中已经有多少个元素了，cap指切片中能放多少个元素
- make([]type, len, cap)
- append可能触发扩容，长度1000以下的情况会扩两倍容量
- 不支持随机增删，只支持append

子切片：

- 数组和切片都可以有子切片
- `arr[start:end]`获得[start, end)之间的元素
- `arr[:end]`获得[0,end)之间的元素
- `arr[start:]`获得[start,len(arr))之间的元素
- 子切片和切片共享数组
- 子切片一般只读，不易出错

循环：

- golang没有while
- for后没有括号
- `for {}` 无限循环
- `for i :=0; i< x; i++` 按下标循环
- `for index, value := range arr` range遍历
- break continue和其他语言一样

条件：

- 没有括号
- 条件中可以有两句（声明+条件），分别用分号表示结束
- 条件中声明的变量，出了条件语句就出作用域了
- switch不需要break，不需要default
- switch中，一个case后可以跟多个常量，用都好隔开
- 可以switch结构体，但是一般不会

## type定义与Server抽象


