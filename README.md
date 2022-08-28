<!--
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-20 10:48:40
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-22 10:32:29
 * @FilePath: 
 * @Description: read me
-->



## golang微服务架构(基于gin)包含以下基础功能
| 序号       | 功能点                    |
| ------------ | ----------------------- |
| `1`        | 动态加载配置文件(viper功能)       |
| `2`        | 支持ZIP日志              |
| `3`        | 服务注册和服务发现（nacos）     |
| `4`        | 集成Redis              |
| `5`        | 集成Mysql             |


## server项目结构（MVC架构）

```shell
├── api
│   └── v1
├── config
├── docs
├── global
├── initialize
├── middleware
├── model
│   ├── orm
│   ├── request
│   └── response
├── router
├── service
└── utils
    ├── nacos
    └── stores
```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `api`        | api层                   | api层 |
| `--v1`       | v1版本接口              | v1版本接口                  |
| `config`     | 配置包                  | config.yaml对应的配置结构体 |
| `global`     | 全局对象                | 全局对象 |
| `initialize` | 初始化 | gin router,redis,gorm,validator, timer的初始化 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码 |
| `model`      | 模型层                  | 模型对应数据表              |
| `--orm`      | 结构体                 | 映射到数据库对应表的各个字段      |
| `--request`  | 入参结构体              | 接收前端发送到后端的数据。  |
| `--response` | 出参结构体              | 返回给前端的数据结构体      |
| `router`     | 路由层                  | 路由层 |
| `service`    | service层               | 存放业务逻辑问题 |
| `utils`      | 工具包                  | 工具函数封装            |
| `--nacos`    | nacos | 注册中心接口封装 |
| `--stores`   | redis                  | redis接口封装        |

