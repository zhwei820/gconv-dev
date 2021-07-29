### gconv 

Created by [郭强](https://goframe.org/display/~john)

## gconv

gconv 可以实现将常用数据类型转换为指定的数据类型，对常用基本数据类型之间的无缝转换，同时也支持任意类型到`struct`对象的转换。由于`gconv`模块内部大量优先使用了断言而非反射，因此执行的效率非常高。

**注意事项：**

`gconv`包的主要目标是提供简便且高效的类型转换功能，开发者应当注意转换的输入参数以及当前使用的业务场景，部分方法如果转换失败，那么方法并不会返回错误原因，也不会产生`panic`，而是直接以转换失败后的数值返回。因此，开发者往往需要结合返回值以及当前使用的业务场景来综合判断结果的正确性。

**使用方式：**

```
import "github.com/zhwei820/gconv"
```

**接口文档**：

<!-- https://godoc.org/github.com/zhwei820/gconv  -->


**单元测试**

```
go test -count=1  ./...
ok      github.com/zhwei820/gconv       3.696s
ok      github.com/zhwei820/gconv/encoding/gbinary      2.199s
ok      github.com/zhwei820/gconv/errors/gerror 3.320s
ok      github.com/zhwei820/gconv/internal/empty        2.914s
ok      github.com/zhwei820/gconv/internal/mutex        3.197s
ok      github.com/zhwei820/gconv/internal/rwmutex      2.856s
ok      github.com/zhwei820/gconv/internal/structs      0.710s
ok      github.com/zhwei820/gconv/internal/utils        1.057s
ok      github.com/zhwei820/gconv/test/gtest    1.506s
```

**基准测试：**

基本类型转换。

```
john@john-B85M:~/Workspace/Go/GOPATH/src/github.com/zhwei820/gconv$ go test *.go -bench=".*" -benchmem
goos: linux
goarch: amd64
BenchmarkString-4               20000000                71.8 ns/op            24 B/op          2 allocs/op
BenchmarkInt-4                  100000000               22.2 ns/op             8 B/op          1 allocs/op
BenchmarkInt8-4                 100000000               24.5 ns/op             8 B/op          1 allocs/op
BenchmarkInt16-4                50000000                23.8 ns/op             8 B/op          1 allocs/op
BenchmarkInt32-4                100000000               24.1 ns/op             8 B/op          1 allocs/op
BenchmarkInt64-4                100000000               21.7 ns/op             8 B/op          1 allocs/op
BenchmarkUint-4                 100000000               22.2 ns/op             8 B/op          1 allocs/op
BenchmarkUint8-4                50000000                25.6 ns/op             8 B/op          1 allocs/op
BenchmarkUint16-4               50000000                32.1 ns/op             8 B/op          1 allocs/op
BenchmarkUint32-4               50000000                27.7 ns/op             8 B/op          1 allocs/op
BenchmarkUint64-4               50000000                28.1 ns/op             8 B/op          1 allocs/op
BenchmarkFloat32-4              10000000               155 ns/op              24 B/op          2 allocs/op
BenchmarkFloat64-4              10000000               177 ns/op              24 B/op          2 allocs/op
BenchmarkTime-4                  5000000               240 ns/op              72 B/op          4 allocs/op
BenchmarkTimeDuration-4         50000000                26.2 ns/op             8 B/op          1 allocs/op
BenchmarkBytes-4                10000000               149 ns/op             128 B/op          3 allocs/op
BenchmarkStrings-4              10000000               223 ns/op              40 B/op          3 allocs/op
BenchmarkInts-4                 20000000                55.0 ns/op            16 B/op          2 allocs/op
BenchmarkFloats-4               10000000               186 ns/op              32 B/op          3 allocs/op
BenchmarkInterfaces-4           20000000                66.6 ns/op            24 B/op          2 allocs/op
PASS
ok      command-line-arguments  35.356s
```


常用基本类型的转换方法比较简单，我们这里使用一个例子来演示转换方法的使用及效果。

# 基本示例

更多的类型转换方法请参考接口文档：
<!-- https://godoc.org/github.com/zhwei820/gconv -->

```
package main

import (
    "fmt"
    "github.com/zhwei820/gconv"
)

func main() {
    i := 123.456
    fmt.Printf("%10s %v\n", "Int:",        gconv.Int(i))
    fmt.Printf("%10s %v\n", "Int8:",       gconv.Int8(i))
    fmt.Printf("%10s %v\n", "Int16:",      gconv.Int16(i))
    fmt.Printf("%10s %v\n", "Int32:",      gconv.Int32(i))
    fmt.Printf("%10s %v\n", "Int64:",      gconv.Int64(i))
    fmt.Printf("%10s %v\n", "Uint:",       gconv.Uint(i))
    fmt.Printf("%10s %v\n", "Uint8:",      gconv.Uint8(i))
    fmt.Printf("%10s %v\n", "Uint16:",     gconv.Uint16(i))
    fmt.Printf("%10s %v\n", "Uint32:",     gconv.Uint32(i))
    fmt.Printf("%10s %v\n", "Uint64:",     gconv.Uint64(i))
    fmt.Printf("%10s %v\n", "Float32:",    gconv.Float32(i))
    fmt.Printf("%10s %v\n", "Float64:",    gconv.Float64(i))
    fmt.Printf("%10s %v\n", "Bool:",       gconv.Bool(i))
    fmt.Printf("%10s %v\n", "String:",     gconv.String(i))
    fmt.Printf("%10s %v\n", "Bytes:",      gconv.Bytes(i))
    fmt.Printf("%10s %v\n", "Strings:",    gconv.Strings(i))
    fmt.Printf("%10s %v\n", "Ints:",       gconv.Ints(i))
    fmt.Printf("%10s %v\n", "Floats:",     gconv.Floats(i))
    fmt.Printf("%10s %v\n", "Interfaces:", gconv.Interfaces(i))
} 
```

执行后，输出结果为：

```
 Int: 123
     Int8: 123
    Int16: 123
    Int32: 123
    Int64: 123
     Uint: 123
    Uint8: 123
   Uint16: 123
   Uint32: 123
   Uint64: 123
  Float32: 123.456
  Float64: 123.456
     Bool: true
   String: 123.456
    Bytes: [119 190 159 26 47 221 94 64]
  Strings: [123.456]
     Ints: [123]
   Floats: [123.456]
Interfaces: [123.456]
```

# 注意事项

数字转换方法例如`gconv.Int/Uint`等等，当给定的转换参数为字符串时，会自动识别十六进制、八进制。

## 八进制转换

`gconv`将前导`0`的数字字符串当做八进制转换。例如，`gconv.Int("010")`将会返回`8`。

## 十六进制转换

`gconv`将`0x`开头的数字字符串当做十六进制转换。例如，`gconv.Int("0xff")`将会返回`255`。

特别注意，`gconv`对于**前导`0`** 开头的字符串处理与标准库的`strconv`包不一样：`gconv`将前导`0`的数字字符串当做**八进制**转换，而`strconv`将会自动去掉前导`0`并按照**十进制**进行转换。

# 类型转换-Map转换

`gconv.Map`支持将任意的`map`或`struct`/`*struct`类型转换为常用的 `map[string]interface{}` 类型。当转换参数为`struct`/`*struct`类型时，支持自动识别`struct`的 `c/gconv/json` 标签，并且可以通过`Map`方法的第二个参数`tags`指定自定义的转换标签，以及多个标签解析的优先级。如果转换失败，返回`nil`。

属性标签：当转换`struct`/`*struct`类型时，支持 `c/gconv/json` 属性标签，也支持 `-`及`omitempty` 标签属性。当使用 `-` 标签属性时，表示该属性不执行转换；当使用 `omitempty` 标签属性时，表示当属性为空时（空指针`nil`, 数字`0`, 字符串`""`, 空数组`[]`等）不执行转换。具体请查看随后示例。

常用转换方法：

```
func Map(value interface{}, tags ...string) map[string]interface{}
func MapDeep(value interface{}, tags ...string) map[string]interface{}
```

其中，`MapDeep`支持递归转换，即会递归转换属性中的`struct`/`*struct`对象。

更多的`map`相关转换方法请参考接口文档：
<!-- https://godoc.org/github.com/zhwei820/gconv -->

# 基本示例

```
package main

import (
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
)

func main() {
    type User struct {
        Uid  int    `c:"uid"`
        Name string `c:"name"`
    }
    // 对象
    pp.Print(gconv.Map(User{
        Uid:  1,
        Name: "john",
    }))
    // 对象指针
    pp.Print(gconv.Map(&User{
        Uid:  1,
        Name: "john",
    }))

    // 任意map类型
    pp.Print(gconv.Map(map[int]int{
        100: 10000,
    }))
}
```

执行后，终端输出：

```
{
    "name": "john",
    "uid": 1
}

{
    "name": "john",
    "uid": 1
}

{
    "100": 10000
}
```

# 属性标签

我们可以通过`c/gconv/json` 标签来自定义转换后的`map`键名，当多个标签存在时，按照`gconv/c/json`的标签顺序进行优先级识别。

```
package main

import (
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
)

func main() {
    type User struct {
        Uid      int
        Name     string `c:"-"`
        NickName string `c:"nickname, omitempty"`
        Pass1    string `c:"password1"`
        Pass2    string `c:"password2"`
    }
    user := User{
        Uid:   100,
        Name:  "john",
        Pass1: "123",
        Pass2: "456",
    }
    pp.Print(gconv.Map(user))
}
```

执行后，终端输出：

```
{
    "Uid": 100,
    "password1": "123",
    "password2": "456"
}
```

# 自定义标签

此外，我们也可以给`struct`的属性自定义自己的标签名称，并在`map`转换时通过第二个参数指定标签优先级。

```
package main

import (
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
)

func main() {
    type User struct {
        Id   int    `c:"uid"`
        Name string `my-tag:"nick-name" c:"name"`
    }
    user := &User{
        Id:   1,
        Name: "john",
    }
    pp.Print(gconv.Map(user, "my-tag"))
}
```

执行后，输出结果为：

```
{
    "nick-name": "john",
    "uid": 1
}
```

# 递归转换

当参数为`map`/`struct`/`*struct`类型时，如果键值/属性为一个对象（或者对象指针）时，并且不是`embedded`结构体且没有任何的别名标签绑定，`Map`方法将会将对象转换为结果的一个键值。我们可以使用`MapDeep`方法递归转换参数的子对象，即把属性也转换为`map`类型。我们来看个例子。

```
package main

import (
    "fmt"
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
    "reflect"
)

func main() {
    type Base struct {
        Id   int    `c:"id"`
        Date string `c:"date"`
    }
    type User struct {
        UserBase Base   `c:"base"`
        Passport string `c:"passport"`
        Password string `c:"password"`
        Nickname string `c:"nickname"`
    }
    user := &User{
        UserBase: Base{
            Id:   1,
            Date: "2019-10-01",
        },
        Passport: "john",
        Password: "123456",
        Nickname: "JohnGuo",
    }
    m1 := gconv.Map(user)
    m2 := gconv.MapDeep(user)
    pp.Print(m1, m2)
    fmt.Println(reflect.TypeOf(m1["base"]))
    fmt.Println(reflect.TypeOf(m2["base"]))
}
```

执行后，终端输出结果为：

```
{
    "base": {
            "Id": 1,
            "Date": "2019-10-01"
    },
    "nickname": "JohnGuo",
    "passport": "john",
    "password": "123456"
}
{
    "base": {
            "date": "2019-10-01",
            "id": 1
    },
    "nickname": "JohnGuo",
    "passport": "john",
    "password": "123456"
}
main.Base
map[string]interface {}
```

看出来差别了吗？

# 类型转换-Struct转换

项目中我们经常会遇到大量`struct`的使用，以及各种数据类型到`struct`的转换/赋值（特别是`json`/`xml`/各种协议编码转换）。为提高编码及项目维护效率，`gconv`模块为各位开发者带来了极大的福利，为数据解析提供了更高的灵活度。

`gconv`模块通过`Struct`转换方法执行`struct`类型转换，其定义如下：

```
// Struct maps the params key-value pairs to the corresponding struct object's attributes.
// The third parameter `mapping` is unnecessary, indicating the mapping rules between the
// custom key name and the attribute name(case sensitive).
//
// Note:
// 1. The `params` can be any type of map/struct, usually a map.
// 2. The `pointer` should be type of *struct/**struct, which is a pointer to struct object
//    or struct pointer.
// 3. Only the public attributes of struct object can be mapped.
// 4. If `params` is a map, the key of the map `params` can be lowercase.
//    It will automatically convert the first letter of the key to uppercase
//    in mapping procedure to do the matching.
//    It ignores the map key, if it does not match.
func Struct(params interface{}, pointer interface{}, mapping ...map[string]string) (err error)
```

其中：

1.  `params`为需要转换到`struct`的变量参数，可以为任意数据类型，常见的数据类型为`map`。
2.  `pointer`为需要执行转的目标`struct`对象，这个参数必须为该`struct`的对象指针，转换成功后该对象的属性将会更新。
3.  `mapping`为自定义的`map键名`到`strcut属性`之间的映射关系，此时`params`参数必须为`map`类型，否则该参数无意义。大部分场景下使用可以不用提供该参数，直接使用默认的转换规则即可。

更多的`struct`相关转换方法请参考接口文档：
<!-- https://godoc.org/github.com/zhwei820/gconv -->

# 转换规则

`gconv`模块的`struct`转换特性非常强大，支持任意数据类型到`struct`属性的映射转换。在没有提供自定义`mapping`转换规则的情况下，默认的转换规则如下：

1.  `struct`中需要匹配的属性必须为 **公开属性** (首字母大写)。
2.  根据`params`类型的不同，逻辑会有不同：
    - `params`参数类型为`map`：键名会自动按照 **不区分大小写** 且 **忽略特殊字符** 的形式与struct属性进行匹配。
    - `params`参数为其他类型：将会把该变量值与`struct`的第一个属性进行匹配。
    - 此外，如果`struct`的属性为复杂数据类型如`slice`,`map`,`strcut`那么会进行递归匹配赋值。
3.  如果匹配成功，那么将键值赋值给属性，如果无法匹配，那么忽略该键值。

以下是几个`map`键名与`struct`属性名称的示例：

```
map键名    struct属性     是否匹配
name       Name           match
Email      Email          match
nickname   NickName       match
NICKNAME   NickName       match
Nick-Name  NickName       match
nick_name  NickName       match
nick name  NickName       match
NickName   Nick_Name      match
Nick-name  Nick_Name      match
nick_name  Nick_Name      match
nick name  Nick_Name      match
```

# 自动创建对象

当给定的`pointer`参数类型为`**struct`时，`Struct`方法内部将会自动创建该`struct`对象，并修改传递变量指向的指针地址。

```
package main

import (
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
)

func main() {
    type User struct {
        Uid  int
        Name string
    }
    params := g.Map{
        "uid":  1,
        "name": "john",
    }
    var user *User
    if err := gconv.Struct(params, &user); err != nil {
        panic(err)
    }
    pp.Print(user)
}
```

执行后，输出结果为：

```
{
    "Name": "john",
    "Uid": 1
}
```

# `Struct`递归转换

递归转换是指当`struct`对象包含子对象时，并且子对象是`embedded`方式定义时，可以将`params`参数数据（第一个参数）同时递归地映射到其子对象上，常用于带有继承对象的`struct`上。

```
package main

import (
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
)

func main() {
    type Ids struct {
        Id         int    `json:"id"`
        Uid        int    `json:"uid"`
    }
    type Base struct {
        Ids
        CreateTime string `json:"create_time"`
    }
    type User struct {
        Base
        Passport   string `json:"passport"`
        Password   string `json:"password"`
        Nickname   string `json:"nickname"`
    }
    data := g.Map{
        "id"          : 1,
        "uid"         : 100,
        "passport"    : "john",
        "password"    : "123456",
        "nickname"    : "John",
        "create_time" : "2019",
    }
    user := new(User)
    gconv.Struct(data, user)
    pp.Print(user)
}
```

执行后，终端输出结果为：

```
{
    "Base": {
        "id": 1,
        "uid": 100,
        "create_time": "2019"
    },
    "nickname": "John",
    "passport": "john",
    "password": "123456"
}
```

# 示例1，基本使用

```
package main

import (
    "github.com/k0kubun/pp"
    "github.com/zhwei820/gconv"
)

type User struct {
    Uid      int
    Name     string
    SiteUrl  string
    NickName string
    Pass1    string `c:"password1"`
    Pass2    string `c:"password2"`
}

func main() {
    var user *User

    // 使用默认映射规则绑定属性值到对象
    user = new(User)
    params1 := g.Map{
        "uid":       1,
        "Name":      "john",
        "site_url":  "https://goframe.org",
        "nick_name": "johng",
        "PASS1":     "123",
        "PASS2":     "456",
    }
    if err := gconv.Struct(params1, user); err == nil {
        pp.Print(user)
    }

    // 使用struct tag映射绑定属性值到对象
    user = new(User)
    params2 := g.Map{
        "uid":       2,
        "name":      "smith",
        "site-url":  "https://goframe.org",
        "nick name": "johng",
        "password1": "111",
        "password2": "222",
    }
    if err := gconv.Struct(params2, user); err == nil {
        pp.Print(user)
    }
}
```

可以看到，我们可以直接通过`Struct`方法将`map`按照默认规则绑定到`struct`上，也可以使用`struct tag`的方式进行灵活的设置。此外，`Struct`方法有第三个`map`参数，用于指定自定义的参数名称到属性名称的映射关系。

执行后，输出结果为：

```
{
    "Uid": 1,
    "Name": "john",
    "SiteUrl": "https://goframe.org",
    "NickName": "johng",
    "Pass1": "123",
    "Pass2": "456"
}

{
    "Uid": 2,
    "Name": "smith",
    "SiteUrl": "https://goframe.org",
    "NickName": "johng",
    "Pass1": "111",
    "Pass2": "222"
}
```

# 示例2，复杂属性类型

属性支持`struct`对象或者`struct`对象指针（目标为指针且未`nil`时，转换时会自动初始化）转换。

```
package main

import (
    "github.com/zhwei820/gconv"
    "github.com/k0kubun/pp"
    "fmt"
)

func main() {
    type Score struct {
        Name   string
        Result int
    }
    type User1 struct {
        Scores Score
    }
    type User2 struct {
        Scores *Score
    }

    user1  := new(User1)
    user2  := new(User2)
    scores := g.Map{
        "Scores": g.Map{
            "Name":   "john",
            "Result": 100,
        },
    }

    if err := gconv.Struct(scores, user1); err != nil {
        fmt.Println(err)
    } else {
        pp.Print(user1)
    }
    if err := gconv.Struct(scores, user2); err != nil {
        fmt.Println(err)
    } else {
        pp.Print(user2)
    }
}
```

执行后，输出结果为：

```
{
    "Scores": {
        "Name": "john",
        "Result": 100
    }
}
{
    "Scores": {
        "Name": "john",
        "Result": 100
    }
}
```