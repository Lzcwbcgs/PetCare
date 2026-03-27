# PetCare

宠物医疗协同管理平台后端项目，基于 GoFrame v2 开发。

当前仓库已经完成了基础工程搭建、数据库模型生成、认证模块增强，以及第二阶段的用户/医生/管理员资料模块。现阶段代码可以正常构建和测试，后续将继续按接口文档推进第三阶段及之后的业务模块。

## 技术栈

- Go 1.25
- GoFrame v2.10.0
- MySQL 8.0
- GoFrame CLI `gf`

## 当前开发进度

### 1. 基础工程整理

- 已接入 GoFrame 标准启动方式
- 已固定读取业务配置文件 [manifest/config/config.yaml](./manifest/config/config.yaml)
- 已启用 Swagger / OpenAPI
- 已统一接口响应结构为 `code / message / data`

### 2. 数据模型生成

已基于真实数据库通过 `gf gen dao` 生成标准模型层：

- [internal/dao](./internal/dao)
- [internal/model/do](./internal/model/do)
- [internal/model/entity](./internal/model/entity)

当前已生成的主要表模型包括：

- `user`
- `doctor`
- `admin`
- `hospital`
- `pet`
- `appointment`
- `medical_record`
- `medical_report`
- `notification`
- `ai_session`
- `ai_message`
- `ai_analysis_record`

### 3. 第一阶段：认证模块与测试基建

当前已完成接口：

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/auth/me`
- `POST /api/auth/logout`

当前认证能力包括：

- 用户注册
- 用户 / 医生 / 管理员统一登录
- JWT 令牌签发与解析
- Bearer Token 登录态校验
- 角色鉴权中间件
- 获取当前登录信息
- 注销后当前 Token 立即失效
- 账号被禁用后，已登录接口立即拒绝访问
- 更清晰的错误码与异常定义
- 认证相关单元测试、控制器测试、中间件测试

认证与测试基建相关核心文件：

- API 定义: [api/auth/v1/auth.go](./api/auth/v1/auth.go)
- 路由入口: [internal/cmd/cmd.go](./internal/cmd/cmd.go)
- 控制器: [internal/controller/auth/auth.go](./internal/controller/auth/auth.go)
- 认证服务接口: [internal/service/auth.go](./internal/service/auth.go)
- 认证服务实现: [internal/service/auth_impl.go](./internal/service/auth_impl.go)
- Token 处理: [internal/service/auth_token.go](./internal/service/auth_token.go)
- 角色能力: [internal/service/auth_role.go](./internal/service/auth_role.go)
- 配置读取: [internal/service/auth_config.go](./internal/service/auth_config.go)
- 密码处理: [internal/service/auth_password.go](./internal/service/auth_password.go)
- 鉴权中间件: [internal/middleware/auth.go](./internal/middleware/auth.go)
- 角色中间件: [internal/middleware/role.go](./internal/middleware/role.go)
- 统一响应中间件: [internal/middleware/response.go](./internal/middleware/response.go)
- 测试配置: [utility/test_config.go](./utility/test_config.go)
- 认证控制器测试: [internal/controller/auth/auth_test.go](./internal/controller/auth/auth_test.go)
- 中间件测试: [internal/middleware/auth_test.go](./internal/middleware/auth_test.go)
- Token 测试: [internal/service/auth_token_test.go](./internal/service/auth_token_test.go)
- 认证服务测试: [internal/service/auth_impl_test.go](./internal/service/auth_impl_test.go)
- 角色测试: [internal/service/auth_role_test.go](./internal/service/auth_role_test.go)

### 4. 第二阶段：资料模块

当前已完成接口：

- `GET /api/users/profile`
- `PUT /api/users/profile`
- `PUT /api/users/password`
- `GET /api/doctors/profile`
- `PUT /api/doctors/profile`
- `PUT /api/doctors/password`
- `GET /api/admin/profile`

当前资料模块能力包括：

- 用户个人资料查询、编辑、密码修改
- 医生个人资料查询、编辑、密码修改
- 管理员个人资料查询
- 资料接口按角色路由隔离
- 资料模块控制器级测试覆盖

资料模块相关核心文件：

- 用户 API: [api/auth/v1/profile_user.go](./api/auth/v1/profile_user.go)
- 医生 API: [api/auth/v1/profile_doctor.go](./api/auth/v1/profile_doctor.go)
- 管理员 API: [api/auth/v1/profile_admin.go](./api/auth/v1/profile_admin.go)
- 用户控制器: [internal/controller/auth/profile_user.go](./internal/controller/auth/profile_user.go)
- 医生控制器: [internal/controller/auth/profile_doctor.go](./internal/controller/auth/profile_doctor.go)
- 管理员控制器: [internal/controller/auth/profile_admin.go](./internal/controller/auth/profile_admin.go)
- 用户服务: [internal/service/profile_user.go](./internal/service/profile_user.go)
- 医生服务: [internal/service/profile_doctor.go](./internal/service/profile_doctor.go)
- 管理员服务: [internal/service/profile_admin.go](./internal/service/profile_admin.go)
- 资料模块测试: [internal/controller/auth/profile_test.go](./internal/controller/auth/profile_test.go)

## 当前已实现接口清单

### 认证模块

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/auth/me`
- `POST /api/auth/logout`

### 用户端

- `GET /api/users/profile`
- `PUT /api/users/profile`
- `PUT /api/users/password`

### 医生端

- `GET /api/doctors/profile`
- `PUT /api/doctors/profile`
- `PUT /api/doctors/password`

### 管理端

- `GET /api/admin/profile`

## 接口文档地址

项目启动后可访问：

- Swagger UI: `http://127.0.0.1:8000/swagger/`
- OpenAPI JSON: `http://127.0.0.1:8000/api.json`

接口设计基准文档位于：

- [pet_medical_api_documentation.md](../../api/pet_medical_api_documentation.md)
- [pet_medical_system_standard.sql](../../数据库&关键技术/pet_medical_system_standard.sql)

## 本地启动

### 1. 配置数据库

修改业务配置文件：

- [manifest/config/config.yaml](./manifest/config/config.yaml)

示例：

```yaml
database:
  default:
    link: "mysql:root:password@tcp(127.0.0.1:3306)/pet_medical_system"
```

如需重新生成模型，开发工具配置位于：

- [hack/config.yaml](./hack/config.yaml)

### 2. 启动项目

```bash
go run .
```

或：

```bash
gf run main.go
```

默认地址：

```text
http://127.0.0.1:8000
```

## 重新生成数据模型

当数据库结构发生变更后，可执行：

```bash
gf gen dao
```

对应配置文件：

- [hack/config.yaml](./hack/config.yaml)

## 当前目录说明

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
utility/               测试与通用辅助
```

## 已验证内容

当前代码已经验证通过：

- `go build ./...`
- `go test ./...`

认证与资料模块已完成的验证点包括：

- 注册成功
- 统一登录成功
- 获取当前登录信息成功
- 角色鉴权校验成功
- 注销后 Token 失效
- 禁用账号无法继续访问已登录接口
- 用户/医生/管理员资料接口链路测试通过

## 后续开发计划

后续将继续按接口文档逐阶段推进：

### 第一阶段 `[已完成]`

- 继续打磨认证模块
- 增加角色鉴权能力
- 增加更清晰的错误码与异常定义
- 补充认证相关测试

### 第二阶段 `[已完成]`

- 用户资料模块
- 医生资料模块
- 管理员资料模块

### 第三阶段 `[下一步]`

- 宠物档案模块
- 宠物病史 / 疫苗 / 过敏记录模块

### 第四阶段

- 预约模块
- 病历与报告模块
- 通知模块

### 第五阶段

- AI 会话模块
- AI 分析记录模块
- 操作日志模块

## 当前仍待完善

虽然认证模块和资料模块已经可用，但距离完整商用版还有差距，后续仍需继续补强：

- 第三阶段及后续业务模块尚未落地
- 更完善的日志与审计能力
- 更严格的参数校验与安全校验
- 更完整的 Token 刷新 / 失效策略
- 更多集成测试与真实数据库联调
- 数据库迁移脚本与初始化数据
- Docker / 部署流程完善

## 开发说明

当前项目开发策略是：

- 先用真实数据库生成标准模型
- 再基于 `dao / do / entity` 实现业务
- 不手改生成代码
- 业务逻辑统一放在 `service/`
- 每完成一个模块，同步补齐 Swagger、响应规范、配置和测试

后续将继续严格按照接口文档和数据库结构文档推进，而不是先随手写功能再回头补结构。
