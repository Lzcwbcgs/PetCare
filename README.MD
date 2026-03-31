# PetCare

PetCare 是一个基于 GoFrame v2 的宠物医院协同管理平台后端项目，面向用户、医生和管理员三类角色提供认证、档案、预约、病历、报告、通知与后台管理能力。

本文档已按仓库当前代码状态更新，内容对应 **2026-03-31** 的实际实现进度。

## 技术栈

- Go 1.25.0
- GoFrame v2.10.0
- MySQL 8.0
- GoFrame CLI `gf`
- Swagger / OpenAPI
- Docker + Kustomize

## 当前进度

| 模块 | 状态 | 说明 |
| --- | --- | --- |
| 基础工程 | 已完成 | GoFrame 启动、统一响应结构、配置加载、Swagger/OpenAPI、静态文件挂载 |
| 认证模块 | 已完成 | 用户注册、统一登录、获取当前登录信息、退出登录、JWT 与角色鉴权 |
| 个人资料模块 | 已完成 | 用户资料、医生资料、管理员资料查询；用户/医生支持修改资料和密码 |
| 宠物档案模块 | 已完成 | 宠物 CRUD、病史、疫苗、过敏记录；用户写入、医生可读 |
| 预约模块 | 已完成 | 用户创建/查看/取消预约，医生查看接诊详情并更新预约状态 |
| 病历与报告模块 | 已完成 | 医生创建/更新/查看病历，上传/删除报告；用户端可查看病历与报告 |
| 通知模块 | 已完成 | 管理员发布系统通知，医生端通知读取，用户/医生通用通知列表与已读 |
| 管理后台模块 | 已完成 | 医院管理、医生管理、系统通知管理 |
| 公共辅助接口 | 已完成 | 医院下拉、医院医生下拉、通用文件上传 |
| AI 相关能力 | 已预留数据模型 | 已生成 `ai_session`、`ai_message`、`ai_analysis_record` DAO/Model，尚未开放业务接口 |
| 部署资源 | 初步具备 | Dockerfile、Kustomize base/overlays、OpenAPI 清单已在仓库中 |

## 已实现能力概览

### 1. 基础工程

- 服务配置文件固定为 [manifest/config/config.yaml](manifest/config/config.yaml)
- 默认监听地址 `:8000`
- 启动时自动注册 OpenAPI 与 Swagger
- 根路径 `/` 自动跳转到 Swagger 页面
- 统一响应结构为 `code / message / data`
- 已挂载静态资源目录：
  - `/uploads` -> `resource/public/uploads`
  - `/reports` -> `resource/public/uploads/reports`

### 2. 数据模型

项目已基于数据库生成标准的 DAO / DO / Entity：

- [internal/dao](internal/dao)
- [internal/model/do](internal/model/do)
- [internal/model/entity](internal/model/entity)

当前已落库并被代码使用或预留的数据表主要包括：

- `user`
- `doctor`
- `admin`
- `hospital`
- `pet`
- `pet_medical_history`
- `pet_vaccination_record`
- `pet_allergy_record`
- `appointment`
- `medical_record`
- `medical_report`
- `notification`
- `operation_log`
- `ai_session`
- `ai_message`
- `ai_analysis_record`

### 3. 角色与业务边界

- 用户：注册、登录、维护个人资料、管理自己的宠物、发起预约、查看病历与通知
- 医生：登录、维护个人资料、查看本人接诊预约、维护本人病历与报告、查看通知
- 管理员：登录、查看管理员资料、管理医院、管理医生、发布系统通知

## 当前接口清单

### 认证模块

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/auth/me`
- `POST /api/auth/logout`

### 用户资料

- `GET /api/users/profile`
- `PUT /api/users/profile`
- `PUT /api/users/password`

### 医生资料

- `GET /api/doctors/profile`
- `PUT /api/doctors/profile`
- `PUT /api/doctors/password`

### 管理员资料

- `GET /api/admin/profile`

### 管理员后台

- 医院管理
  - `POST /api/admin/hospitals`
  - `GET /api/admin/hospitals`
  - `GET /api/admin/hospitals/{hospital_id}`
  - `PUT /api/admin/hospitals/{hospital_id}`
  - `DELETE /api/admin/hospitals/{hospital_id}`
- 医生管理
  - `POST /api/admin/doctors`
  - `GET /api/admin/doctors`
  - `GET /api/admin/doctors/{doctor_id}`
  - `PUT /api/admin/doctors/{doctor_id}`
  - `DELETE /api/admin/doctors/{doctor_id}`
- 系统通知管理
  - `POST /api/admin/notifications`
  - `GET /api/admin/notifications`
  - `DELETE /api/admin/notifications/{notification_id}`

### 宠物档案

- `GET /api/pets`
- `POST /api/pets`
- `GET /api/pets/{pet_id}`
- `PUT /api/pets/{pet_id}`
- `DELETE /api/pets/{pet_id}`

### 宠物病史

- `POST /api/pets/{pet_id}/medical-histories`
- `GET /api/pets/{pet_id}/medical-histories`

### 宠物疫苗

- `POST /api/pets/{pet_id}/vaccinations`
- `GET /api/pets/{pet_id}/vaccinations`

### 宠物过敏

- `POST /api/pets/{pet_id}/allergies`
- `GET /api/pets/{pet_id}/allergies`

### 用户预约

- `POST /api/appointments`
- `GET /api/appointments`
- `GET /api/appointments/{appointment_id}`
- `PUT /api/appointments/{appointment_id}/cancel`

### 医生预约工作台

- `GET /api/doctor/appointments`
- `GET /api/doctor/appointments/{appointment_id}/detail`
- `PUT /api/doctor/appointments/{appointment_id}/status`

### 医生病历与报告

- `POST /api/doctor/medical-records`
- `PUT /api/doctor/medical-records/{medical_record_id}`
- `GET /api/doctor/medical-records/{medical_record_id}`
- `POST /api/doctor/medical-records/{medical_record_id}/reports`
- `GET /api/doctor/medical-records/{medical_record_id}/reports`
- `DELETE /api/doctor/reports/{report_id}`

### 用户侧病历查看

- `GET /api/medical-records`
- `GET /api/medical-records/{medical_record_id}`
- `GET /api/medical-records/{medical_record_id}/reports`

### 通知中心

- 通用通知
  - `GET /api/notifications`
  - `GET /api/notifications/{notification_id}`
  - `PUT /api/notifications/{notification_id}/read`
- 医生专属通知
  - `GET /api/doctor/notifications`
  - `PUT /api/doctor/notifications/{notification_id}/read`

### 公共辅助接口

- `GET /api/common/hospitals/options`
- `GET /api/common/hospitals/{hospital_id}/doctors/options`
- `POST /api/common/upload`

## 本地启动

### 1. 准备环境

- 安装 Go 1.25+
- 准备 MySQL 8.0
- 创建数据库：`pet_medical_system`
- 按需安装 GoFrame CLI：`go install github.com/gogf/gf/cmd/gf/v2@latest`

默认业务配置位于：

- [manifest/config/config.yaml](manifest/config/config.yaml)

当前默认配置包括：

- 服务监听地址：`:8000`
- Swagger 路径：`/swagger`
- OpenAPI 路径：`/api.json`
- 默认数据库连接：`mysql:root:123456@tcp(127.0.0.1:3306)/pet_medical_system`
- JWT 配置：`jwtSecret`、`jwtIssuer`、`expireHours`

DAO 生成配置位于：

- [hack/config.yaml](hack/config.yaml)

### 2. 启动项目

```bash
go run .
```

或：

```bash
gf run main.go
```

默认访问地址：

```text
http://127.0.0.1:8000
```

### 3. 接口文档

项目启动后可访问：

- Swagger UI: `http://127.0.0.1:8000/swagger/`
- OpenAPI JSON: `http://127.0.0.1:8000/api.json`

OpenAPI 增强配置位于 [internal/cmd/cmd.go](internal/cmd/cmd.go)。

## 常用开发命令

[Makefile](Makefile) 引入了 [hack/hack.mk](hack/hack.mk) 和 [hack/hack-cli.mk](hack/hack-cli.mk)，常用命令如下：

- `make build`：使用 GoFrame 构建二进制
- `make dao`：重新生成 DAO / DO / Entity
- `make ctrl`：根据 API 生成控制器骨架
- `make service`：生成 Service 接口
- `make enums`：生成枚举代码
- `make pb`：生成 protobuf 代码
- `make image`：构建 Docker 镜像

## 目录说明

```text
api/                   接口请求与响应定义
internal/cmd/          服务启动与路由注册
internal/controller/   控制器层
internal/service/      业务服务层
internal/middleware/   认证与统一响应中间件
internal/dao/          GoFrame 生成的 DAO
internal/model/do/     GoFrame 生成的 DO 模型
internal/model/entity/ GoFrame 生成的 Entity 模型
manifest/config/       业务配置
manifest/docker/       Docker 构建资源
manifest/deploy/       Kustomize 部署清单
resource/public/       静态资源与上传目录
utility/               测试辅助代码
```

## 测试与验证

已在当前仓库状态下实际通过：

- `go test ./...`
- `go build ./...`

当前已有测试主要集中在以下基础能力：

- [internal/controller/auth/auth_test.go](internal/controller/auth/auth_test.go)
- [internal/controller/auth/profile_test.go](internal/controller/auth/profile_test.go)
- [internal/middleware/auth_test.go](internal/middleware/auth_test.go)
- [internal/service/auth_impl_test.go](internal/service/auth_impl_test.go)
- [internal/service/auth_token_test.go](internal/service/auth_token_test.go)
- [internal/service/auth_role_test.go](internal/service/auth_role_test.go)

说明：

- 认证、资料、中间件与鉴权基础能力已有自动化测试覆盖
- 预约、医生工作台、管理员后台、宠物档案、通知等业务模块目前主要通过编译与联调验证

## 部署资源

仓库中已经包含基础部署资源：

- Dockerfile: [manifest/docker/Dockerfile](manifest/docker/Dockerfile)
- Kustomize base: [manifest/deploy/kustomize/base](manifest/deploy/kustomize/base)
- Kustomize develop overlay: [manifest/deploy/kustomize/overlays/develop](manifest/deploy/kustomize/overlays/develop)
- OpenAPI 清单: [manifest/api/openapi.yaml](manifest/api/openapi.yaml)

## 后续建议

按当前仓库状态，下一阶段更适合继续推进以下事项：

- 补齐预约、病历、通知、后台模块的独立自动化测试
- 完善数据库迁移、初始化脚本与种子数据
- 接入操作日志与审计链路
- 推进 AI 会话、消息、分析记录相关业务接口
- 完善 Docker / K8s 环境变量化与部署说明
