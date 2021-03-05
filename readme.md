## 规则引擎

## 运行方式

-   api方式，参照 cmd/api/api.go中所示
    - 提供接口添加规则
    - 提供接口输入参数并进行计算输出结果
-  内嵌到程序中

### 运算
- [x] 算术运算 加减乘除
  - [x] float64 op float64
- [x] 比较操作 ``` >、>=、==、<=、<```
- [x] 逻辑运算 ```&&、||、! ```

### 语句
- [x] if、elsif、else
- [x] return

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