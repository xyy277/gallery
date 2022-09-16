# gallery

模块封装

## 依赖

* go get github.com/xyy277/gallery

* Init初始化

## 代码结构

```shell
├── auth
│   ├── luna
│   └── star
├── cache
├── config
├── constant
├── core
├── global
│   ├── request
│   └── response
├── log
├── oss
├── util
└── util

```

| 文件夹       | 说明                    | 描述                        |
| ------------ | ----------------------- | --------------------------- |
| `auth`        | auth 组件                | auth 接口                 |
| `--luna`      | auth 平台组件            | auth  平台组件接口         |
| `--star`      | auth 业务组件            | auth 业务组件接口          |  
| `cache`       | cache组件                | cache接口                 |
| `config`      | 配置文件                 | 组件配置                   |
| `constant`    | constant常量             | constant常量              |
| `core`        | 核心包                   | 核心包                     |
| `global`      | 全局对象                 | 全局对象                   |
| `--request`   | 入参结构体               | 接收前端发送到后端的数据。   |
| `--response`  | 出参结构体               | 返回给前端的数据结构体       |
| `log`         | 日志组件                 | 日志组件接口                |
| `oss`         | oss组件                  | oss组件接口                 |
| `util`        | 工具包                   | 工具函数封装                |

## ProtoBuf

### 安装

<https://github.com/protocolbuffers/protobuf/releases/>

* linux 安装
* windwos 安装

### protobuf-go

<https://github.com/protocolbuffers/protobuf-go>
