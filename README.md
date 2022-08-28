<!--
 * @Author: stefan1240 lishaoshu1240@gmail.com
 * @Date: 2022-08-28 12:52:59
 * @LastEditors: stefan1240 lishaoshu1240@gmail.com
 * @LastEditTime: 2022-08-28 12:54:03
 * @FilePath: /gin_micro_service/README.md
 * @Description: readme
-->


# gin_micro_service

**gin_micro_service** 是一个基于golang(gin)的微服务基础框架。

目的是提供一个文件结构清晰、有使用入门例子的微服务架构，主要功能包含：动态加载配置文件、日志、服务注册和服务发现接入（nacos）、链接数据库等功能。

## 特性

你可以把 gin_micro_service。

- **全平台**
  - 运行环境: 支持：linux、windows、mac。
  - 部署方式: docker部署、直接编译运行。
  
- **基础功能**
  - 动态加载配置文件: 基于viper模块。
  - 日志: 基于ZIP模块。
  - 服务注册和服务发现: 基于阿里开源nacos。
  - 数据库连接: 可接入Redis、Mysql, 此外，可自行拓展别数据库。


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

## 立刻开始

### 运行、编译、打包成docker

TODO

## 贡献

我们欢迎来自开源社区、个人和合作伙伴的各种贡献。

- [贡献指南](CONTRIBUTING.md)

## 协议

[Apache 2.0 License](LICENSE)
