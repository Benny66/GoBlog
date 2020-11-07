### GoBlog
#### 环境
- go环境，配置环境变量
- mod依赖包管理工具

#### 使用
- go run main.go

#### 目录结构
- common        公共模块，包括公共函数、返回response等方法
- config        项目配置模块，包括返回code、默认配置
- controller    控制层 
- database      数据库相关
- middleware    中间件
- models        模型层
- public        静态目录
- router        路由层
- runtime       日志模块
- validator     验证层
- .gitignore    项目忽略上传文件
- go.mod        go mod依赖管理工具文件
- main.go       启动程序main
- README.md     项目说明

#### 功能
- 待完善

#### 中间件
- redis缓存和队列
- http请求封装
- jwt请求token校验
- logger日志保存到文件
- SIGN签名校验封装