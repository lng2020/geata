Geata 是一个基于 IEC61850 的 IoT 网关,使用了包括 Vue3、TailwindCSS、Vite、Go、Gin 和 Xorm 在内的现代 Web 技术栈。它具有以下几个主要特性:

1. 将 MQTT 和 ModbusTCP 数据转换为 IEC61850 格式。IEC61850 中的每个数据属性(DA)都可在 Web UI 中进行配置。
2. 使用自编译的客户端二进制文件从服务器获取 IEC61850 数据,该二进制文件是从 [libiec61850](https://github.com/mz-automation/libiec61850) 生成的。客户端源码位于 `internal/app/client/client.c` 文件中。因为作者太懒了,所以没有提供 Linux 二进制文件,你需要自行构建。
3. 解析 ICD/SCD 文件并将数据存储在数据库中。
4. 基本的电站和用户管理功能。

## 使用说明
1. 复制 `config.example.yaml` 到 `config.yaml` 并修改配置。
2. 前端
```bash
cd web
pnpm install
pnpm run dev
```
3. 后端
```bash
go build
./geata start
```
4. 模拟 MQTT 客户端
```bash
go build
./geata mock
```

## 待完成和限制
- [ ] 添加更多 MQTT 和 ModbusTCP 的配置选项,并支持更多数据类型。
- [ ] 目前 `FC`（也就是IEC61850协议中的Functional Constraint）是硬编码在 IEC61850 客户端中的,应该在数据库中检测和记录。
- [ ] 优化 IEC61850 客户端和 Modbus 客户端的性能。目前使用轮询方式,应改为事件驱动。
- [ ] 考虑使用 IEC61850 Go 库替代自编译的客户端二进制文件。但目前还没有 Go 库可用,没有库的情况下,客户端二进制文件是从服务器获取数据的唯一方式。因为有此限制，所以无法将 IE61850 数据提供给客户端。
- [ ] ICD/SCD 文件解析能力有限,这是由于我对 IEC61850 的知识不足造成的,应该进一步改进。

## 项目结构
```
- cmd # main包和模拟命令
- internal
    - app # 业务逻辑
        - client # IEC61850 客户端
        - handler # 协议处理器
        - parser # ICD/SCD 文件解析器
        - model # 数据模型
        - service # 服务层
        - web # Web 和路由
        - logger # 日志记录器
        config.go # 配置
        app.go # 入口点
    - docs # Swaggo 自动生成的文档
- web # 前端
```

## 如何适配更多协议
在 `internal/app/handler` 目录下实现 `Handler` 和 `HandlerConfig` 接口,并在 `internal/app/handler/handler.go` 文件中进行注册。

## 许可证
MIT