# PetCare

PetCare 是一个基于 GoFrame v2 的宠物医疗协同管理平台后端项目，面向用户、医生、管理员三类角色，提供认证、宠物档案、预约、病历、报告、通知、AI 问诊、知识库与后台管理能力。

本文档按当前仓库代码状态更新，和现在的实现保持一致。

## 技术栈

- Go 1.25
- GoFrame v2
- MySQL 8.0
- Swagger / OpenAPI
- Docker
- Kustomize
- AI Provider: API / Ollama
- RAG Vector Store: Qdrant

## 当前进度

| 模块 | 状态 | 说明 |
| --- | --- | --- |
| 基础工程 | 已完成 | GoFrame 启动、统一响应结构、配置加载、Swagger/OpenAPI、静态资源挂载 |
| 认证模块 | 已完成 | 用户注册、统一登录、获取当前登录信息、退出登录、JWT 与角色鉴权 |
| 资料模块 | 已完成 | 用户资料、医生资料、管理员资料查询；用户和医生支持更新资料与密码 |
| 宠物档案模块 | 已完成 | 宠物 CRUD、病史、疫苗、过敏记录；用户写入，医生只读 |
| 预约模块 | 已完成 | 用户创建/查看/取消预约，医生查看接诊详情并更新预约状态 |
| 病历与报告模块 | 已完成 | 医生创建/更新/查看病历，上传/删除报告；用户端可查看病历与报告 |
| 通知模块 | 已完成 | 管理员发布系统通知，用户/医生查看通知并标记已读 |
| 管理后台 | 已完成 | 医院管理、医生管理、系统通知管理、操作日志查询 |
| 公共辅助接口 | 已完成 | 医院下拉、医院医生下拉、通用文件上传 |
| AI 问诊模块 | 已完成 | 创建 AI 会话、SSE 流式发送消息、查询会话详情、消息记录、分析结果 |
| 医生 AI 查看 | 已完成 | 医生端 AI 会话列表、消息记录、分析结果查看 |
| 知识库 / RAG | 已完成 | 知识库文件上传、向量化进度查询、RAG 检索接入 AI 问诊 |
| 部署资源 | 已具备 | Dockerfile、Kustomize base/overlay、OpenAPI 清单已在仓库中 |

## 已实现能力

### 基础能力

- 服务配置文件固定为 [`manifest/config/config.yaml`](./manifest/config/config.yaml)
- 默认监听地址 `:8000`
- 根路径 `/` 自动跳转到 Swagger 页面
- 统一响应结构为 `code / message / data`
- 已挂载静态资源目录
  - `/uploads` -> `resource/public/uploads`
  - `/reports` -> `resource/public/uploads/reports`

### 数据模型

当前项目已生成并使用标准 DAO / DO / Entity：

- [`internal/dao`](./internal/dao)
- [`internal/model/do`](./internal/model/do)
- [`internal/model/entity`](./internal/model/entity)

当前核心业务表包括：

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
- `knowledge_document`
- `knowledge_chunk`
- `knowledge_job`

### AI 与知识库

- 支持创建 AI 会话
- 支持通过 SSE 流式发送 AI 消息
- 支持查询 AI 会话详情、消息记录、分析结果
- 支持医生端查看 AI 会话、消息与分析结果
- 支持知识库文档上传
- 支持知识库切片、向量化、进度查询
- 支持 RAG 检索增强问诊
- 支持 API Provider 与本地 Ollama Provider

## 当前接口清单

当前已实现接口共 71 个。

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

### AI 问诊

- `POST /api/ai/sessions`
- `POST /api/ai/sessions/{session_id}/messages`
- `GET /api/ai/sessions/{session_id}`
- `GET /api/ai/sessions/{session_id}/messages`
- `GET /api/ai/sessions/{session_id}/analysis-records`

### 医生 AI 查看

- `GET /api/doctor/ai/sessions`
- `GET /api/doctor/ai/sessions/{session_id}/messages`
- `GET /api/doctor/ai/sessions/{session_id}/analysis-records`

### 通知中心

- `GET /api/notifications`
- `GET /api/notifications/{notification_id}`
- `PUT /api/notifications/{notification_id}/read`
- `GET /api/doctor/notifications`
- `PUT /api/doctor/notifications/{notification_id}/read`
- `POST /api/admin/notifications`
- `GET /api/admin/notifications`
- `DELETE /api/admin/notifications/{notification_id}`

### 管理后台

- `POST /api/admin/hospitals`
- `GET /api/admin/hospitals`
- `GET /api/admin/hospitals/{hospital_id}`
- `PUT /api/admin/hospitals/{hospital_id}`
- `DELETE /api/admin/hospitals/{hospital_id}`
- `POST /api/admin/doctors`
- `GET /api/admin/doctors`
- `GET /api/admin/doctors/{doctor_id}`
- `PUT /api/admin/doctors/{doctor_id}`
- `DELETE /api/admin/doctors/{doctor_id}`
- `GET /api/admin/operation-logs`
- `GET /api/admin/operation-logs/{log_id}`

### 知识库 / RAG

- `POST /api/admin/knowledge/upload`
- `GET /api/admin/knowledge/status`

### 公共辅助接口

- `GET /api/common/hospitals/options`
- `GET /api/common/hospitals/{hospital_id}/doctors/options`
- `POST /api/common/upload`

## 本地启动

### 1. 准备环境

- 安装 Go 1.25+
- 准备 MySQL 8.0
- 创建数据库 `pet_medical_system`
- 按需安装 GoFrame CLI

```bash
go install github.com/gogf/gf/cmd/gf/v2@latest
```

业务配置位于：

- [`manifest/config/config.yaml`](./manifest/config/config.yaml)

配置内容包括：

- 服务监听地址
- Swagger / OpenAPI 路径
- 数据库连接
- JWT 配置
- AI Provider 配置
- RAG / Qdrant 配置

DAO 生成配置位于：

- [`hack/config.yaml`](./hack/config.yaml)

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

OpenAPI 增强配置位于 [`internal/cmd/cmd.go`](./internal/cmd/cmd.go)。

## 常用开发命令

- `make build`：构建二进制
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
manifest/api/          OpenAPI 清单
resource/public/       静态资源与上传目录
utility/               测试辅助代码
```

## 测试与验证

当前仓库状态下已验证：

- `go test ./...`

当前已有测试主要集中在：

- [`internal/controller/auth/auth_test.go`](./internal/controller/auth/auth_test.go)
- [`internal/controller/auth/profile_test.go`](./internal/controller/auth/profile_test.go)
- [`internal/middleware/auth_test.go`](./internal/middleware/auth_test.go)
- [`internal/service/auth_impl_test.go`](./internal/service/auth_impl_test.go)
- [`internal/service/auth_token_test.go`](./internal/service/auth_token_test.go)
- [`internal/service/auth_role_test.go`](./internal/service/auth_role_test.go)

说明：

- 认证、资料、中间件与鉴权基础能力已有自动化测试覆盖
- 业务模块当前主要通过编译与接口联调验证

## 部署资源

仓库中已包含基础部署资源：

- Dockerfile: [`manifest/docker/Dockerfile`](./manifest/docker/Dockerfile)
- Kustomize base: [`manifest/deploy/kustomize/base`](./manifest/deploy/kustomize/base)
- Kustomize develop overlay: [`manifest/deploy/kustomize/overlays/develop`](./manifest/deploy/kustomize/overlays/develop)
- OpenAPI 清单: [`manifest/api/openapi.yaml`](./manifest/api/openapi.yaml)

## 后续建议

- 为预约、宠物、病历、通知、AI、知识库模块补齐更细粒度自动化测试
- 增加数据库迁移、初始化脚本与种子数据
- 把配置中的敏感信息迁移为环境变量或外部密钥管理
- 增加 AI / RAG 相关链路的观测、限流和失败兜底
- 补齐部署文档与生产环境配置说明
