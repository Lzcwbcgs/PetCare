# PetCare

宠物医疗协同管理平台后端项目，基于 GoFrame v2 开发。

当前仓库已经完成了项目基础脚手架整理，以及认证模块第一版的标准化落地。后续开发将以这一版认证模块作为模板，继续推进用户、宠物、预约、病历、通知、AI 会话等业务模块。

## 技术栈

- Go 1.25
- GoFrame v2.10.0
- MySQL 8.0
- GoFrame CLI `gf`

## 当前已完成

### 1. 基础工程整理

- 项目已接入 GoFrame 标准启动方式
- 项目已固定读取业务配置文件 [manifest/config/config.yaml](D:/code/gocode/pet-hospital/PetCare/manifest/config/config.yaml)
- 已启用 Swagger/OpenAPI 文档
- 已统一接口响应结构为 `code / message / data`

### 2. 数据模型生成

已通过 `gf gen dao` 根据真实数据库生成标准模型层：

- [internal/dao](D:/code/gocode/pet-hospital/PetCare/internal/dao)
- [internal/model/do](D:/code/gocode/pet-hospital/PetCare/internal/model/do)
- [internal/model/entity](D:/code/gocode/pet-hospital/PetCare/internal/model/entity)

当前已识别并生成的主要表包括：

- `user`
- `doctor`
- `admin`
- `pet`
- `appointment`
- `medical_record`
- `medical_report`
- `notification`
- `ai_session`
- `ai_message`
- `ai_analysis_record`

### 3. 认证模块第一版

已完成以下接口：

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/auth/me`
- `POST /api/auth/logout`

实现内容包括：

- 用户注册
- 用户 / 医生 / 管理员统一登录
- JWT 令牌签发与校验
- 登录态校验中间件
- 获取当前登录用户信息
- Swagger 接口文档展示

认证相关核心文件：

- API 定义: [api/auth/v1/auth.go](D:/code/gocode/pet-hospital/PetCare/api/auth/v1/auth.go)
- 路由启动: [internal/cmd/cmd.go](D:/code/gocode/pet-hospital/PetCare/internal/cmd/cmd.go)
- 控制器: [internal/controller/auth/auth.go](D:/code/gocode/pet-hospital/PetCare/internal/controller/auth/auth.go)
- 服务实现: [internal/service/auth.go](D:/code/gocode/pet-hospital/PetCare/internal/service/auth.go)
- 服务实现: [internal/service/auth_impl.go](D:/code/gocode/pet-hospital/PetCare/internal/service/auth_impl.go)
- JWT: [internal/service/auth_token.go](D:/code/gocode/pet-hospital/PetCare/internal/service/auth_token.go)
- 密码处理: [internal/service/auth_password.go](D:/code/gocode/pet-hospital/PetCare/internal/service/auth_password.go)
- 鉴权中间件: [internal/middleware/auth.go](D:/code/gocode/pet-hospital/PetCare/internal/middleware/auth.go)
- 统一响应中间件: [internal/middleware/response.go](D:/code/gocode/pet-hospital/PetCare/internal/middleware/response.go)

## 当前接口文档地址

项目启动后可访问：

- Swagger UI: `http://127.0.0.1:8000/swagger/`
- OpenAPI JSON: `http://127.0.0.1:8000/api.json`

## 本地启动

### 1. 配置数据库

修改业务配置文件：

- [manifest/config/config.yaml](D:/code/gocode/pet-hospital/PetCare/manifest/config/config.yaml)

当前示例配置：

```yaml
database:
  default:
    link: "mysql:root:password@tcp(127.0.0.1:3306)/pet_medical_system"
```

同时如果需要重新生成模型，开发环境配置在：

- [hack/config.yaml](D:/code/gocode/pet-hospital/PetCare/hack/config.yaml)

### 2. 启动项目

```bash
go run .
```

或者：

```bash
gf run main.go
```

默认端口：

```text
http://127.0.0.1:8000
```

## 重新生成数据模型

当数据库表结构发生变化后，可以重新执行：

```bash
gf gen dao
```

当前项目已经将 `gf gen dao` 配置写入：

- [hack/config.yaml](D:/code/gocode/pet-hospital/PetCare/hack/config.yaml)

## 当前目录说明

```text
api/                  接口请求与响应定义
internal/cmd/         服务启动与路由注册
internal/controller/  控制器层
internal/service/     业务服务层
internal/middleware/  中间件
internal/dao/         GoFrame 生成的 DAO
internal/model/do/    GoFrame 生成的 DO 模型
internal/model/entity/GoFrame 生成的实体模型
manifest/config/      业务配置
hack/                 开发工具配置
```

## 已验证内容

当前代码已经验证通过：

- `go build ./...`
- `go test ./...`
- 认证接口真实联调：
  - 注册成功
  - 登录成功
  - 获取当前登录信息成功

## 后续开发计划

接下来将按照更接近商业项目的节奏，逐模块推进：

### 第一阶段

- 继续打磨认证模块
- 增加角色鉴权能力
- 增加更清晰的错误码与异常定义
- 补充认证相关测试

### 第二阶段

- 用户资料模块
- 医生资料模块
- 管理员资料模块

### 第三阶段

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

虽然认证模块已经可用，但还没有完全做到最终商用版，后续仍需继续补强：

- 更完整的错误码体系
- 更完善的日志与审计
- 更严格的参数校验与安全校验
- JWT 刷新/失效策略
- 单元测试与集成测试补全
- 数据库迁移脚本与初始化数据
- Docker / 部署流程完善

## 开发说明

当前项目开发策略是：

- 先用真实数据库生成标准模型
- 再基于 `dao/do/entity` 实现业务
- 每完成一个模块，就同步补齐 Swagger、响应规范、配置和测试

也就是说，后续不会再走“先随手写功能，再回头补结构”的方式，而是按标准工程流程逐步推进。
