# PetCare

宠物医疗协同管理平台后端项目，基于 GoFrame v2 开发。

当前仓库已经完成基础工程搭建、数据库模型生成、认证模块、资料模块，以及第三阶段中的宠物档案/病史/疫苗/过敏记录模块。项目当前可以正常构建和测试，后续将继续推进预约、病历报告、通知与 AI 相关能力。

## 技术栈

- Go 1.25
- GoFrame v2.10.0
- MySQL 8.0
- GoFrame CLI `gf`

## 当前开发进度

### 1. 基础工程

- 已接入 GoFrame 标准启动方式
- 固定读取业务配置文件 [manifest/config/config.yaml](/d:/code/gocode/review/PetCare/manifest/config/config.yaml)
- 已启用 Swagger / OpenAPI
- 已统一接口响应结构为 `code / message / data`
- 已接入 MySQL 驱动和资源打包入口

### 2. 数据模型

项目已基于真实数据库通过 `gf gen dao` 生成标准模型层：

- [internal/dao](/d:/code/gocode/review/PetCare/internal/dao)
- [internal/model/do](/d:/code/gocode/review/PetCare/internal/model/do)
- [internal/model/entity](/d:/code/gocode/review/PetCare/internal/model/entity)

当前已生成并被业务使用的主要数据表包括：

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
- `ai_session`
- `ai_message`
- `ai_analysis_record`

### 3. 第一阶段：认证模块

已完成接口：

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/auth/me`
- `POST /api/auth/logout`

已具备能力：

- 用户注册
- 用户 / 医生 / 管理员统一登录
- JWT 签发与解析
- Bearer Token 登录态校验
- 角色鉴权中间件
- 获取当前登录信息
- 注销后当前 Token 立即失效
- 账号禁用后已登录接口立即拒绝访问
- 认证相关单元测试、控制器测试、中间件测试

核心文件：

- 路由入口 [internal/cmd/cmd.go](/d:/code/gocode/review/PetCare/internal/cmd/cmd.go)
- API 定义 [api/auth/v1/auth.go](/d:/code/gocode/review/PetCare/api/auth/v1/auth.go)
- 控制器 [internal/controller/auth/auth.go](/d:/code/gocode/review/PetCare/internal/controller/auth/auth.go)
- 认证服务接口 [internal/service/auth.go](/d:/code/gocode/review/PetCare/internal/service/auth.go)
- 认证服务实现 [internal/service/auth_impl.go](/d:/code/gocode/review/PetCare/internal/service/auth_impl.go)
- Token 处理 [internal/service/auth_token.go](/d:/code/gocode/review/PetCare/internal/service/auth_token.go)
- 角色能力 [internal/service/auth_role.go](/d:/code/gocode/review/PetCare/internal/service/auth_role.go)
- 认证中间件 [internal/middleware/auth.go](/d:/code/gocode/review/PetCare/internal/middleware/auth.go)
- 角色中间件 [internal/middleware/role.go](/d:/code/gocode/review/PetCare/internal/middleware/role.go)

### 4. 第二阶段：资料模块

已完成接口：

- `GET /api/users/profile`
- `PUT /api/users/profile`
- `PUT /api/users/password`
- `GET /api/doctors/profile`
- `PUT /api/doctors/profile`
- `PUT /api/doctors/password`
- `GET /api/admin/profile`

已具备能力：

- 用户个人资料查询、编辑、密码修改
- 医生个人资料查询、编辑、密码修改
- 管理员个人资料查询
- 资料接口按角色路由隔离
- 控制器级测试覆盖

核心文件：

- 用户 API [api/auth/v1/profile_user.go](/d:/code/gocode/review/PetCare/api/auth/v1/profile_user.go)
- 医生 API [api/auth/v1/profile_doctor.go](/d:/code/gocode/review/PetCare/api/auth/v1/profile_doctor.go)
- 管理员 API [api/auth/v1/profile_admin.go](/d:/code/gocode/review/PetCare/api/auth/v1/profile_admin.go)
- 资料控制器 [internal/controller/auth](/d:/code/gocode/review/PetCare/internal/controller/auth)
- 资料服务 [internal/service/profile_user.go](/d:/code/gocode/review/PetCare/internal/service/profile_user.go)
- 资料服务 [internal/service/profile_doctor.go](/d:/code/gocode/review/PetCare/internal/service/profile_doctor.go)
- 资料服务 [internal/service/profile_admin.go](/d:/code/gocode/review/PetCare/internal/service/profile_admin.go)

### 5. 第三阶段：宠物档案模块

当前已完成接口：

- `GET /api/pets`
- `POST /api/pets`
- `GET /api/pets/{pet_id}`
- `PUT /api/pets/{pet_id}`
- `DELETE /api/pets/{pet_id}`
- `POST /api/pets/{pet_id}/medical-histories`
- `GET /api/pets/{pet_id}/medical-histories`
- `POST /api/pets/{pet_id}/vaccinations`
- `GET /api/pets/{pet_id}/vaccinations`
- `POST /api/pets/{pet_id}/allergies`
- `GET /api/pets/{pet_id}/allergies`

当前已具备能力：

- 用户维度的宠物档案增删改查
- 宠物病史记录新增与分页查询
- 宠物疫苗记录新增与分页查询
- 宠物过敏记录新增与分页查询
- 用户仅可操作自己的宠物数据
- 医生可读取宠物病史、疫苗、过敏记录
- 宠物列表支持分页、名称模糊查询、类型/状态筛选

核心文件：

- 宠物 API [api/pet/v1/pet.go](/d:/code/gocode/review/PetCare/api/pet/v1/pet.go)
- 病史 API [api/pet/v1/medical_history.go](/d:/code/gocode/review/PetCare/api/pet/v1/medical_history.go)
- 疫苗 API [api/pet/v1/vaccination.go](/d:/code/gocode/review/PetCare/api/pet/v1/vaccination.go)
- 过敏 API [api/pet/v1/allergy.go](/d:/code/gocode/review/PetCare/api/pet/v1/allergy.go)
- 宠物控制器 [internal/controller/pet/pet.go](/d:/code/gocode/review/PetCare/internal/controller/pet/pet.go)
- 病史控制器 [internal/controller/pet/medical_history.go](/d:/code/gocode/review/PetCare/internal/controller/pet/medical_history.go)
- 疫苗控制器 [internal/controller/pet/vaccination.go](/d:/code/gocode/review/PetCare/internal/controller/pet/vaccination.go)
- 过敏控制器 [internal/controller/pet/allergy.go](/d:/code/gocode/review/PetCare/internal/controller/pet/allergy.go)
- 宠物服务 [internal/service/pet.go](/d:/code/gocode/review/PetCare/internal/service/pet.go)
- 病史服务 [internal/service/pet_medical_history.go](/d:/code/gocode/review/PetCare/internal/service/pet_medical_history.go)
- 疫苗服务 [internal/service/pet_vaccination.go](/d:/code/gocode/review/PetCare/internal/service/pet_vaccination.go)
- 过敏服务 [internal/service/pet_allergy.go](/d:/code/gocode/review/PetCare/internal/service/pet_allergy.go)

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

## 接口文档

项目启动后可访问：

- Swagger UI: `http://127.0.0.1:8000/swagger/`
- OpenAPI JSON: `http://127.0.0.1:8000/api.json`

OpenAPI 由 [internal/cmd/cmd.go](/d:/code/gocode/review/PetCare/internal/cmd/cmd.go) 在启动时注册。

## 本地启动

### 1. 准备配置

业务配置文件：

- [manifest/config/config.yaml](/d:/code/gocode/review/PetCare/manifest/config/config.yaml)

当前默认配置包含：

- 服务监听地址 `:8000`
- Swagger 路径 `/swagger`
- OpenAPI 路径 `/api.json`
- 默认 MySQL 连接
- JWT 配置项 `jwtSecret / jwtIssuer / expireHours`

如需重新生成 DAO，开发工具配置位于：

- [hack/config.yaml](/d:/code/gocode/review/PetCare/hack/config.yaml)

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

## 常用开发命令

项目根目录的 [Makefile](/d:/code/gocode/review/PetCare/Makefile) 引入了 [hack/hack.mk](/d:/code/gocode/review/PetCare/hack/hack.mk) 和 [hack/hack-cli.mk](/d:/code/gocode/review/PetCare/hack/hack-cli.mk)，可以直接使用：

- `make build` 使用 GoFrame 构建二进制
- `make dao` 根据数据库重新生成 DAO / DO / Entity
- `make ctrl` 解析 API 生成控制器骨架
- `make service` 生成 Service 接口代码
- `make enums` 生成枚举文件
- `make pb` 生成 protobuf 相关代码
- `make image` 构建 Docker 镜像

如果本机没有安装 `gf`，`make` 会先尝试自动安装 GoFrame CLI。

## 目录说明

```text
api/                   接口请求与响应定义
internal/cmd/          服务启动与路由注册
internal/controller/   控制器层
internal/service/      业务服务层
internal/middleware/   中间件
internal/dao/          GoFrame 生成的 DAO
internal/model/do/     GoFrame 生成的 DO 模型
internal/model/entity/ GoFrame 生成的实体模型
manifest/config/       业务配置
hack/                  开发工具配置
resource/              项目资源文件
utility/               测试与通用辅助
```

## 已验证内容

当前已实际验证通过：

- `go build ./...`
- `go test ./...`

当前测试覆盖情况：

- 认证控制器测试 [internal/controller/auth/auth_test.go](/d:/code/gocode/review/PetCare/internal/controller/auth/auth_test.go)
- 资料控制器测试 [internal/controller/auth/profile_test.go](/d:/code/gocode/review/PetCare/internal/controller/auth/profile_test.go)
- 认证中间件测试 [internal/middleware/auth_test.go](/d:/code/gocode/review/PetCare/internal/middleware/auth_test.go)
- 认证服务测试 [internal/service/auth_impl_test.go](/d:/code/gocode/review/PetCare/internal/service/auth_impl_test.go)
- Token 测试 [internal/service/auth_token_test.go](/d:/code/gocode/review/PetCare/internal/service/auth_token_test.go)
- 角色能力测试 [internal/service/auth_role_test.go](/d:/code/gocode/review/PetCare/internal/service/auth_role_test.go)

说明：

- 当前宠物档案相关模块已完成实现并通过 `go test ./...` 编译级验证
- `internal/controller/pet` 暂无独立测试文件，后续建议补充控制器和服务层测试

## 后续计划

下一阶段建议继续推进：

- 预约模块
- 病历与医疗报告模块
- 通知模块
- AI 会话与分析记录模块
- 操作日志与审计能力
- 更完善的数据库迁移与初始化脚本
- Docker / 部署流程完善

## 开发说明

当前项目开发策略：

- 先基于真实数据库生成标准模型
- 再基于 `dao / do / entity` 编写业务逻辑
- 尽量不手改生成代码
- 业务逻辑统一放在 `service/`
- 每完成一个模块，同步补齐路由、Swagger、响应规范、配置和测试

当前 README 已按仓库截至 2026-03-30 的实际代码状态更新。
