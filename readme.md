## 规则引擎

## 运行方式

-   api方式，参照 cmd/api/api.go中所示
    - 提供接口添加规则
    - 提供接口输入参数并进行计算输出结果
-  内嵌到程序中

### api方式
```shell
#运行程序
go run cmd/api/api.go

#添加规则
curl -XPOST 'HTTP://127.0.0.1/add?name=compare' \
-H 'content-Type:text/plain' \
-d 'a=3;b=4;a>b'

#调用规则
 curl -XPOST 'HTTP://127.0.0.1/run' \
 -H 'content-type:application/json' \
 -d '{"name":"compare","data":{}}'
```

### 参数
```json
{
  "name":"compare", //name参数为添加的规则名
  "data":{} //data可添加规则运行需要的额外的参数
}

### 运算
- [x] 算术运算 加减乘除
  - [x] float64 op float64
- [x] 比较操作 ``` >、>=、==、<=、<```
- [x] 逻辑运算 ```&&、||、! ```

### 语句
- [x] if、elsif、else
- [ ] return
    - [ ]作用域
### 数据类型
- [x] bool
- [x] string
- [ ] number
  - [x] float64
- [ ] struct
- [x] map
- [x] array

## 开发工具
- goland
- go
- antlr