# SeatFlow

SeatFlow 是一个基于 Go 的高并发票务与智能座位分配系统。第一阶段建立模块化单体工程、本地 MySQL 环境和 HTTP 健康检查。

## 技术栈

- Go 1.26 或以上
- Gin 1.12.0
- GORM 1.31.2，配合 GORM MySQL 适配器 1.6.0 和 Go MySQL 驱动 1.10.2
- MySQL 8.4 LTS，通过 Docker Compose V2 启动

## 启动 HTTP 服务

在仓库根目录执行以下命令。仅验证 HTTP 服务时，可以直接启动 Go 进程。

```sh
go run ./cmd/server
```

另开一个终端请求健康检查接口。

```sh
curl -fsS http://127.0.0.1:8080/healthz
```

预期返回 HTTP 200 和 JSON 响应 `{"status":"ok"}`。可通过进程环境变量 `HTTP_ADDR` 调整监听地址。

## 启动本地 MySQL

安装 Docker 和 Compose V2 后，在仓库根目录复制配置模板。

```sh
cp .env.example .env
```

填写 `.env` 中的 `MYSQL_ROOT_PASSWORD` 和 `MYSQL_PASSWORD`，再执行以下命令。`.env` 供 Compose 读取，Go 服务从进程环境读取配置。

```sh
docker compose --env-file .env -f deployments/compose.yaml up -d --wait
docker compose --env-file .env -f deployments/compose.yaml ps
```

MySQL 默认监听宿主机的 `127.0.0.1:3306`，使用名为 `seatflow` 的数据库和应用用户。可通过 `MYSQL_PORT` 调整宿主机端口。初始化账号和密码仅在数据卷首次初始化时生效。

如需在服务启动时验证数据库连接，可设置 `MYSQL_DSN`。将示例中的占位符换成刚设置的应用用户密码，端口与 `MYSQL_PORT` 保持一致。

```sh
MYSQL_DSN='seatflow:替换为本地应用用户密码@tcp(127.0.0.1:3306)/seatflow?parseTime=true&timeout=5s&readTimeout=5s&writeTimeout=5s' go run ./cmd/server
```

停止容器时执行以下命令。MySQL 数据保存在命名卷中。

```sh
docker compose --env-file .env -f deployments/compose.yaml down
```

## 工程文档

- [开发规范](AGENTS.md)
- [总体架构](docs/architecture.md)
- [第一阶段里程碑](docs/milestones/01-foundation.md)
