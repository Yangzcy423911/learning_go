```shell
# 运行 go run hello_world.go
hello world
# 构建 go build hello_world.go
./hello_world
hello world
知识点
应用程序入口
必须是main包: package main
必须是main方法： fun main()
文件名不一定是main.go
退出返回值
与其它主要编程语言的差异
1、Go中main函数不支持任何返回值
2、通过os.Exit来返回状态
3、main函数不支持传入参数
4、在程序中直接通过os.Args获取命令行参数
5、GO语言不允许隐式类型转换
6、别名和原有类型也不能进行隐式类型转换
7、不支持指针运算
8、string是值类型，其默认的初始化值为空字符串，而不是nil
9、Go语言仅支持循环关键字for
10、if条件
(1) condition表达式结果必须为布尔值
(2) 支持变量赋值
  if var declaration; condition{

  }
11、 switch条件
(1) 条件表达式不限制为常量或者整数
(2) 单个case中，可以出现多个结果选项，使用逗号分隔
(3) 与C语言等规则相反，GO语言不需要用break来明确退出一个case
(4) 可以不设定switch之后的条件表达式，在此种情况下，整个switch结构与多个
位运算符
&^ 按位置零
1 &^ 0 -- 1
1 &^ 1 -- 0
0 &^ 1 -- 0
0 &^ 0 -- 0
12 Map元素的访问
1、在访问Key不存在时，仍会返回零值，不能通过返回nil来判断元素是否存在
13 字符串
1、string是数据类型，不是引用或指针类型
2、string是只读的byte slice,len函数可以它所包含的byte数
3、string的byte数组可以存放任何数据

13 Unicode UTF8
1、 Unicode是一种字符集(code point)
2、 UTF8是unicode的存储实现(转换为字节序列的规则)
if .... else....的逻辑作用等同
14 Go的错误机制
1、没有异常机制
2、error类型实现了error接口
3、可以通过errors.New来快速创建错误实例
15、 panic用于不可以恢复的错误
    panic退出前会执行defer指定的内容
os.Exit退出时不会调用defer指定的函数
os.Exit退出时不输出当前调用栈信息
一、类型的预定义值
1、 math.MaxInt64
2、 math.MaxFloat64
3、 math.MaxUint32

二、用 == 比较数组
1、相同维数且含有相同个数元素的数组才可以比较
2、每个元素都相同的才相等
三、函数是一等公民
1、可以有多个返回值
2、所有参数都是值传递: slice,map,channel会有传引用的错觉
3、函数可以作为变量的值
4、函数可以作为参数和返回值
# go run hello_world.go Yang
hello world Yang
四、行为(方法)定义
type Employee struct {
	Id   string
	Name string
	Age  int
}
1、第一种定义方式在实例对应方法被调用时，实例的成员会进行复制
func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
2、通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}
五、Go接口
1、 接口为非侵入性，实现不依赖于接口定义
2、 所以接口的定义可以包含在接口使用者包内
六、空接口与断言
1、空接口可以表示任何类型
2、通过断言来将空接口转换为制定类型
v,ok := p.(int) // ok = true时转换成功
七、Go接口最佳实践
1、倾向于使用小的接口定义，很多接口只包含一个方法
2、较大的接口定义，可以由多个小接口定义组合而成
3、只依赖于必要功能的最小接口
八、 package
1、基本复用模块单元
 以首字母大写来表明可被包外代码访问
2、代码的package可以和所在的目录不一致
3、同一目录里的Go代码的package要保持一致
4、 init方法
 (1) 在main被执行前，所有依赖的package的init方法都会被执行
 (2) 不同包的init函数按照包导入依赖关系决定执行顺序
 (3) 每个包可以有多个init函数
 (4) 包的每个源文件也可以有多个init函数，这点比较特殊
5、 通过go get来获取远程依赖
  go get -u 强制从网络更新远程依赖
  注意代码在Github上的组织形式，以适合go get
  直接以代码路径开始，不要有src
九、 vendor路径
查找依赖包路径的解决方案如下:
 1、当前包下的vendor目录
 2、向上级目录查找，直到找到src下的vendor目录
 3、在GOPATH下面查找依赖包
 4、在GOROOT目录下查找
```