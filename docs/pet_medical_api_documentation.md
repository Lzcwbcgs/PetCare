
---
title: 宠物医疗协同管理平台 RESTful API 接口文档
author: OpenAI
> 适用技术栈：Go（Gin / GoFrame）+ MySQL 8.0 + Vue  
> 文档用途：前后端分离开发与联调  
> 接口风格：RESTful  
> 认证方式：JWT Bearer Token

---

# 1. 文档说明

## 1.1 基础约定

- 基础路径前缀：`/api`
- 接口风格：RESTful
- 数据格式：`application/json`
- 文件上传：`multipart/form-data`
- 时间格式：`YYYY-MM-DD HH:MM:SS`
- 分页参数：`page`、`page_size`
- 所有路径命名使用小写复数形式

## 1.2 统一响应结构

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

## 1.3 通用状态码

| code | 含义 |
|---|---|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未登录或 token 无效 |
| 403 | 无权限访问 |
| 404 | 资源不存在 |
| 409 | 业务冲突 |
| 500 | 服务器内部错误 |

## 1.4 角色定义

| 角色 | 标识 |
|---|---|
| 用户 | `user` |
| 医生 | `doctor` |
| 管理员 | `admin` |

请求头示例：

```http
Authorization: Bearer <token>
```

---

# 2. 枚举定义

## 2.1 预约类型 `appointment_type`

| 值 | 含义 |
|---|---|
| 1 | 体检预约 |
| 2 | 看病预约 |

## 2.2 预约状态 `appointment.status`

| 值 | 含义 |
|---|---|
| 1 | 待就诊 |
| 2 | 已完成 |
| 3 | 已取消 |
| 4 | 已过期 |

## 2.3 预约来源 `appointment.source`

| 值 | 含义 |
|---|---|
| 1 | 用户端预约 |
| 2 | 医生代录入 |
| 3 | 后台创建 |

## 2.4 病历状态 `medical_record.status`

| 值 | 含义 |
|---|---|
| 1 | 已创建 |
| 2 | 已完成 |
| 3 | 已归档 |

## 2.5 AI 会话状态 `ai_session.status`

| 值 | 含义 |
|---|---|
| 1 | 进行中 |
| 2 | 已结束 |
| 3 | 已归档 |

## 2.6 AI 消息发送者 `ai_message.sender_type`

| 值 | 含义 |
|---|---|
| 1 | 用户 |
| 2 | AI |
| 3 | 医生 |
| 4 | 管理员 |

## 2.7 AI 分析类型 `analysis_type`

| 值 | 含义 |
|---|---|
| 1 | 病历总结 |
| 2 | 症状归纳 |
| 3 | 风险提示 |
| 4 | 健康建议 |

## 2.8 风险等级 `risk_level`

| 值 | 含义 |
|---|---|
| 1 | 低 |
| 2 | 中 |
| 3 | 高 |

## 2.9 通知类型 `notification_type`

| 值 | 含义 |
|---|---|
| 1 | 预约提醒 |
| 2 | 系统通知 |
| 3 | AI 分析提醒 |

## 2.10 通知状态 `notification.status`

| 值 | 含义 |
|---|---|
| 0 | 待发送 |
| 1 | 已发送 |
| 2 | 发送失败 |
| 3 | 已读 |

---

# 3. 通用数据结构

## 3.1 分页响应结构

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 100
    }
  }
}
```

## 3.2 用户对象 `UserVO`

```json
{
  "id": 1,
  "username": "linyuan",
  "nickname": "元元",
  "phone": "13800000000",
  "email": "test@example.com",
  "avatar_url": "https://example.com/avatar.jpg",
  "status": 1,
  "created_at": "2026-03-25 10:00:00"
}
```

## 3.3 医生对象 `DoctorVO`

```json
{
  "id": 1,
  "hospital_id": 1,
  "hospital_name": "爱宠动物医院",
  "username": "doctor_zhang",
  "doctor_name": "张医生",
  "gender": 1,
  "phone": "13900000000",
  "email": "doctor@example.com",
  "title": "主治医师",
  "specialty": "猫科内科",
  "avatar_url": "https://example.com/doctor.jpg",
  "intro": "擅长猫科消化系统疾病诊疗",
  "status": 1
}
```

## 3.4 宠物对象 `PetVO`

```json
{
  "id": 1,
  "user_id": 1,
  "pet_name": "小橘",
  "pet_type": "猫",
  "avatar_url": "https://example.com/pet.jpg",
  "gender": 1,
  "age": 2,
  "age_unit": "year",
  "breed": "中华田园猫",
  "weight": "4.50",
  "sterilized": 1,
  "remark": "性格活泼",
  "status": 1,
  "created_at": "2026-03-25 10:00:00",
  "updated_at": "2026-03-25 10:00:00"
}
```

## 3.5 预约对象 `AppointmentVO`

```json
{
  "id": 1001,
  "appointment_no": "APT202603250001",
  "user_id": 1,
  "pet_id": 1,
  "hospital_id": 1,
  "doctor_id": 2,
  "appointment_type": 2,
  "symptom_description": "食欲下降，精神不佳",
  "appointment_time": "2026-03-26 14:00:00",
  "reminder_time": "2026-03-26 13:00:00",
  "status": 1,
  "source": 1,
  "created_at": "2026-03-25 10:00:00",
  "updated_at": "2026-03-25 10:00:00"
}
```

---

# 4. 认证模块

> 支持用户、医生、管理员三类登录主体，统一从认证中心签发 JWT。

## 4.1 用户注册

**接口名称：** 用户注册  
**路径：** `POST /api/auth/register`  
**权限：** 公开

### 请求体

```json
{
  "username": "user001",
  "password": "123456",
  "nickname": "小猫家长",
  "phone": "13800000000",
  "email": "user001@example.com"
}
```

### 响应

```json
{
  "code": 200,
  "message": "注册成功",
  "data": {
    "user_id": 1
  }
}
```

## 4.2 统一登录

**接口名称：** 统一登录  
**路径：** `POST /api/auth/login`  
**权限：** 公开

### 请求体

```json
{
  "username": "doctor_zhang",
  "password": "123456",
  "role": "doctor"
}
```

### 响应

```json
{
  "code": 200,
  "message": "登录成功",
  "data": {
    "token": "jwt_token_string",
    "expire_at": "2026-03-26 10:00:00",
    "user_id": 2,
    "role": "doctor"
  }
}
```

## 4.3 获取当前登录信息

**接口名称：** 获取当前登录信息  
**路径：** `GET /api/auth/me`  
**权限：** 已登录

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 2,
    "username": "doctor_zhang",
    "role": "doctor",
    "doctor_name": "张医生",
    "avatar_url": "https://example.com/doctor.jpg"
  }
}
```

## 4.4 退出登录

**接口名称：** 退出登录  
**路径：** `POST /api/auth/logout`  
**权限：** 已登录

### 响应

```json
{
  "code": 200,
  "message": "退出成功",
  "data": {}
}
```

---

# 5. 用户端模块

> 用户端包括账号信息、宠物档案、健康记录、预约管理、病历查询、AI 问诊与通知提醒。

## 5.1 用户信息模块

### 5.1.1 获取个人信息

**接口名称：** 获取用户个人信息  
**路径：** `GET /api/users/profile`  
**权限：** 用户

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "user001",
    "nickname": "小猫家长",
    "phone": "13800000000",
    "email": "user001@example.com",
    "avatar_url": "https://example.com/avatar.jpg",
    "status": 1,
    "created_at": "2026-03-25 10:00:00"
  }
}
```

### 5.1.2 更新个人信息

**接口名称：** 更新用户个人信息  
**路径：** `PUT /api/users/profile`  
**权限：** 用户

### 请求体

```json
{
  "nickname": "元元",
  "phone": "13812345678",
  "email": "new@example.com",
  "avatar_url": "https://example.com/new-avatar.jpg"
}
```

### 响应

```json
{
  "code": 200,
  "message": "更新成功",
  "data": {}
}
```

### 5.1.3 修改密码

**接口名称：** 修改用户密码  
**路径：** `PUT /api/users/password`  
**权限：** 用户

### 请求体

```json
{
  "old_password": "123456",
  "new_password": "654321"
}
```

### 响应

```json
{
  "code": 200,
  "message": "密码修改成功",
  "data": {}
}
```

## 5.2 宠物管理模块

### 5.2.1 新增宠物档案

**接口名称：** 新增宠物档案  
**路径：** `POST /api/pets`  
**权限：** 用户

### 请求体

```json
{
  "pet_name": "小橘",
  "pet_type": "猫",
  "avatar_url": "https://example.com/pet.jpg",
  "gender": 1,
  "age": 2,
  "age_unit": "year",
  "breed": "中华田园猫",
  "weight": "4.50",
  "sterilized": 1,
  "remark": "有点挑食"
}
```

### 响应

```json
{
  "code": 200,
  "message": "新增成功",
  "data": {
    "pet_id": 1
  }
}
```

### 5.2.2 获取宠物列表

**接口名称：** 获取宠物列表  
**路径：** `GET /api/pets?page=1&page_size=10`  
**权限：** 用户

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "pet_name": "小橘",
        "pet_type": "猫",
        "gender": 1,
        "age": 2,
        "age_unit": "year",
        "breed": "中华田园猫",
        "weight": "4.50",
        "status": 1
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 5.2.3 获取宠物详情

**接口名称：** 获取宠物详情  
**路径：** `GET /api/pets/{pet_id}`  
**权限：** 用户

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "user_id": 1,
    "pet_name": "小橘",
    "pet_type": "猫",
    "avatar_url": "https://example.com/pet.jpg",
    "gender": 1,
    "age": 2,
    "age_unit": "year",
    "breed": "中华田园猫",
    "weight": "4.50",
    "sterilized": 1,
    "remark": "有点挑食",
    "status": 1,
    "created_at": "2026-03-25 10:00:00",
    "updated_at": "2026-03-25 10:00:00"
  }
}
```

### 5.2.4 修改宠物档案

**接口名称：** 修改宠物档案  
**路径：** `PUT /api/pets/{pet_id}`  
**权限：** 用户

### 请求体

```json
{
  "pet_name": "小橘子",
  "weight": "4.80",
  "remark": "近期食欲正常"
}
```

### 响应

```json
{
  "code": 200,
  "message": "修改成功",
  "data": {}
}
```

### 5.2.5 删除宠物档案

**接口名称：** 删除宠物档案  
**路径：** `DELETE /api/pets/{pet_id}`  
**权限：** 用户

### 响应

```json
{
  "code": 200,
  "message": "删除成功",
  "data": {}
}
```

## 5.3 宠物健康信息模块

### 5.3.1 新增病史记录

**接口名称：** 新增宠物病史  
**路径：** `POST /api/pets/{pet_id}/medical-histories`  
**权限：** 用户

### 请求体

```json
{
  "history_type": "呼吸系统疾病",
  "description": "2025年曾出现上呼吸道感染",
  "diagnosed_at": "2025-06-01 10:00:00",
  "is_current": 0
}
```

### 响应

```json
{
  "code": 200,
  "message": "新增成功",
  "data": {
    "id": 1
  }
}
```

### 5.3.2 获取病史列表

**接口名称：** 获取宠物病史列表  
**路径：** `GET /api/pets/{pet_id}/medical-histories?page=1&page_size=10`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "history_type": "呼吸系统疾病",
        "description": "2025年曾出现上呼吸道感染",
        "diagnosed_at": "2025-06-01 10:00:00",
        "is_current": 0
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 5.3.3 新增疫苗记录

**接口名称：** 新增疫苗接种记录  
**路径：** `POST /api/pets/{pet_id}/vaccinations`  
**权限：** 用户

### 请求体

```json
{
  "vaccine_name": "猫三联",
  "vaccination_date": "2026-03-01",
  "next_due_date": "2027-03-01",
  "hospital_name": "爱宠动物医院",
  "remark": "无不良反应"
}
```

### 响应

```json
{
  "code": 200,
  "message": "新增成功",
  "data": {
    "id": 1
  }
}
```

### 5.3.4 获取疫苗记录列表

**接口名称：** 获取疫苗接种记录列表  
**路径：** `GET /api/pets/{pet_id}/vaccinations?page=1&page_size=10`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "vaccine_name": "猫三联",
        "vaccination_date": "2026-03-01",
        "next_due_date": "2027-03-01",
        "hospital_name": "爱宠动物医院"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 5.3.5 新增过敏记录

**接口名称：** 新增过敏记录  
**路径：** `POST /api/pets/{pet_id}/allergies`  
**权限：** 用户

### 请求体

```json
{
  "allergen": "海鲜类罐头",
  "symptom_description": "进食后呕吐",
  "severity_level": 2,
  "remark": "建议避免相关食物"
}
```

### 响应

```json
{
  "code": 200,
  "message": "新增成功",
  "data": {
    "id": 1
  }
}
```

### 5.3.6 获取过敏记录列表

**接口名称：** 获取过敏记录列表  
**路径：** `GET /api/pets/{pet_id}/allergies?page=1&page_size=10`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "allergen": "海鲜类罐头",
        "symptom_description": "进食后呕吐",
        "severity_level": 2,
        "remark": "建议避免相关食物"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

## 5.4 预约管理模块

### 5.4.1 创建预约

**接口名称：** 创建预约  
**路径：** `POST /api/appointments`  
**权限：** 用户

### 请求体

```json
{
  "pet_id": 1,
  "hospital_id": 1,
  "doctor_id": 2,
  "appointment_type": 2,
  "symptom_description": "近两天食欲差，精神萎靡",
  "appointment_time": "2026-03-26 14:00:00"
}
```

### 响应

```json
{
  "code": 200,
  "message": "预约成功",
  "data": {
    "appointment_id": 1001,
    "appointment_no": "APT202603250001",
    "status": 1
  }
}
```

### 5.4.2 获取预约列表

**接口名称：** 获取我的预约列表  
**路径：** `GET /api/appointments?page=1&page_size=10&status=1&appointment_type=2`  
**权限：** 用户

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1001,
        "appointment_no": "APT202603250001",
        "pet_id": 1,
        "pet_name": "小橘",
        "hospital_id": 1,
        "hospital_name": "爱宠动物医院",
        "doctor_id": 2,
        "doctor_name": "张医生",
        "appointment_type": 2,
        "appointment_time": "2026-03-26 14:00:00",
        "status": 1
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 5.4.3 获取预约详情

**接口名称：** 获取预约详情  
**路径：** `GET /api/appointments/{appointment_id}`  
**权限：** 用户 / 医生 / 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1001,
    "appointment_no": "APT202603250001",
    "user_id": 1,
    "user_nickname": "小猫家长",
    "pet_id": 1,
    "pet_name": "小橘",
    "hospital_id": 1,
    "hospital_name": "爱宠动物医院",
    "doctor_id": 2,
    "doctor_name": "张医生",
    "appointment_type": 2,
    "symptom_description": "近两天食欲差，精神萎靡",
    "appointment_time": "2026-03-26 14:00:00",
    "reminder_time": "2026-03-26 13:00:00",
    "status": 1,
    "source": 1,
    "created_at": "2026-03-25 10:00:00"
  }
}
```

### 5.4.4 取消预约

**接口名称：** 取消预约  
**路径：** `PUT /api/appointments/{appointment_id}/cancel`  
**权限：** 用户

### 请求体

```json
{
  "cancel_reason": "时间冲突"
}
```

### 响应

```json
{
  "code": 200,
  "message": "取消成功",
  "data": {}
}
```

## 5.5 病历与报告查看模块

### 5.5.1 获取病历列表

**接口名称：** 获取我的宠物病历列表  
**路径：** `GET /api/medical-records?page=1&page_size=10&pet_id=1`  
**权限：** 用户

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "appointment_id": 1001,
        "pet_id": 1,
        "doctor_id": 2,
        "doctor_name": "张医生",
        "preliminary_diagnosis": "疑似胃肠炎",
        "visit_time": "2026-03-26 14:30:00",
        "status": 2
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 5.5.2 获取病历详情

**接口名称：** 获取病历详情  
**路径：** `GET /api/medical-records/{medical_record_id}`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "appointment_id": 1001,
    "pet_id": 1,
    "user_id": 1,
    "doctor_id": 2,
    "chief_complaint": "食欲下降两天",
    "present_history": "精神略差，无明显腹泻",
    "physical_examination": "体温正常，触诊轻微腹部敏感",
    "preliminary_diagnosis": "疑似胃肠炎",
    "treatment_plan": "对症支持治疗",
    "prescription": "益生菌、肠胃处方粮",
    "doctor_advice": "观察48小时，如持续呕吐需复诊",
    "visit_time": "2026-03-26 14:30:00",
    "status": 2
  }
}
```

### 5.5.3 获取报告列表

**接口名称：** 获取病历报告列表  
**路径：** `GET /api/medical-records/{medical_record_id}/reports`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "report_title": "血常规检查报告",
        "report_type": "检验报告",
        "file_url": "https://example.com/reports/1.pdf",
        "report_content": "白细胞轻度升高",
        "uploaded_at": "2026-03-26 15:20:00"
      }
    ]
  }
}
```

## 5.6 AI 问诊模块

### 5.6.1 创建 AI 会话

**接口名称：** 创建 AI 问诊会话  
**路径：** `POST /api/ai/sessions`  
**权限：** 用户

### 请求体

```json
{
  "pet_id": 1,
  "hospital_id": 1,
  "doctor_id": 2,
  "model_type": "local",
  "model_name": "pet-med-llm-v1"
}
```

### 响应

```json
{
  "code": 200,
  "message": "会话创建成功",
  "data": {
    "session_id": 1,
    "session_no": "AIS202603250001",
    "status": 1
  }
}
```

### 5.6.2 发送 AI 消息

**接口名称：** 发送 AI 问诊消息（SSE 流式输出）  
**路径：** `POST /api/ai/sessions/{session_id}/messages`  
**权限：** 用户 / 医生

### 请求头

```http
Authorization: Bearer <token>
Content-Type: application/json
Accept: text/event-stream
```

**响应类型：** `text/event-stream`

### 请求体

```json
{
  "message_content": "我家猫最近没精神，还不怎么吃东西，应该怎么办？",
  "message_type": 1
}
```

### SSE 事件说明

| event | 含义 |
|---|---|
| `user_message` | 用户消息已写入会话 |
| `message` | AI 增量回复分片 |
| `done` | AI 回复完成，返回最终完整消息 |
| `error` | 流式处理失败 |

### 响应示例

```text
event: user_message
data: {"id":1,"sender_type":1,"message_content":"我家猫最近没精神，还不怎么吃东西，应该怎么办？","created_at":"2026-03-25 11:00:00"}

event: message
data: {"content":"建议先观察是否伴随呕吐、腹泻、发热等症状，"}

event: message
data: {"content":"如持续超过24小时请尽快就医。"}

event: done
data: {"ai_message":{"id":2,"sender_type":2,"message_content":"建议先观察是否伴随呕吐、腹泻、发热等症状，如持续超过24小时请尽快就医。","created_at":"2026-03-25 11:00:02"}}
```

### 说明

- 该接口使用 SSE 长连接流式返回，不再使用统一 JSON 包装结构。
- 前端应按事件顺序消费 `data` 内容并实时拼接 AI 回复。
- 完整消息记录仍通过 `5.6.4 获取 AI 消息记录` 接口查询。

### 5.6.3 获取 AI 会话详情

**接口名称：** 获取 AI 会话详情  
**路径：** `GET /api/ai/sessions/{session_id}`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "session_no": "AIS202603250001",
    "pet_id": 1,
    "source_type": 1,
    "model_type": "local",
    "model_name": "pet-med-llm-v1",
    "session_summary": "围绕食欲下降与精神不振进行了初步问诊",
    "status": 1,
    "created_at": "2026-03-25 10:55:00",
    "updated_at": "2026-03-25 11:00:02"
  }
}
```

### 5.6.4 获取 AI 消息记录

**接口名称：** 获取 AI 会话消息记录  
**路径：** `GET /api/ai/sessions/{session_id}/messages?page=1&page_size=20`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "sender_type": 1,
        "sender_id": 1,
        "message_content": "我家猫最近没精神，还不怎么吃东西，应该怎么办？",
        "message_type": 1,
        "created_at": "2026-03-25 11:00:00"
      },
      {
        "id": 2,
        "sender_type": 2,
        "sender_id": null,
        "message_content": "建议先观察是否伴随呕吐、腹泻、发热等症状，如持续超过24小时请尽快就医。",
        "message_type": 1,
        "created_at": "2026-03-25 11:00:02"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 20,
      "total": 2
    }
  }
}
```

### 5.6.5 获取 AI 分析结果

**接口名称：** 获取 AI 分析结果  
**路径：** `GET /api/ai/sessions/{session_id}/analysis-records`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "analysis_type": 3,
        "input_source": 1,
        "analysis_result": "存在轻中度消化道不适风险",
        "risk_level": 2,
        "reviewed_by_doctor": 0,
        "created_at": "2026-03-25 11:01:00"
      }
    ]
  }
}
```

## 5.7 通知模块

### 5.7.1 获取通知列表

**接口名称：** 获取通知列表  
**路径：** `GET /api/notifications?page=1&page_size=10&status=1`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "notification_type": 1,
        "title": "预约提醒",
        "content": "您于 2026-03-26 14:00:00 的预约即将开始",
        "appointment_id": 1001,
        "send_time": "2026-03-26 13:00:00",
        "status": 1
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 5.7.2 获取通知详情

**接口名称：** 获取通知详情  
**路径：** `GET /api/notifications/{notification_id}`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "notification_type": 1,
    "title": "预约提醒",
    "content": "您于 2026-03-26 14:00:00 的预约即将开始",
    "appointment_id": 1001,
    "send_time": "2026-03-26 13:00:00",
    "status": 1
  }
}
```

### 5.7.3 标记通知已读

**接口名称：** 标记通知已读  
**路径：** `PUT /api/notifications/{notification_id}/read`  
**权限：** 用户 / 医生

### 响应

```json
{
  "code": 200,
  "message": "已标记为已读",
  "data": {}
}
```

---

# 6. 医生端模块

> 医生端负责查看预约、接诊前查阅宠物档案、创建与更新病历、上传医疗报告，并查看 AI 会话与提醒。

## 6.1 医生个人信息模块

### 6.1.1 获取医生个人信息

**接口名称：** 获取医生个人信息  
**路径：** `GET /api/doctors/profile`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 2,
    "hospital_id": 1,
    "hospital_name": "爱宠动物医院",
    "username": "doctor_zhang",
    "doctor_name": "张医生",
    "gender": 1,
    "phone": "13900000000",
    "email": "doctor@example.com",
    "title": "主治医师",
    "specialty": "猫科内科",
    "avatar_url": "https://example.com/doctor.jpg",
    "intro": "擅长猫科消化系统疾病诊疗",
    "status": 1
  }
}
```

### 6.1.2 更新医生个人信息

**接口名称：** 更新医生个人信息  
**路径：** `PUT /api/doctors/profile`  
**权限：** 医生

### 请求体

```json
{
  "phone": "13911112222",
  "email": "doctor_new@example.com",
  "avatar_url": "https://example.com/new-doctor.jpg",
  "intro": "擅长猫科消化系统与呼吸系统疾病"
}
```

### 响应

```json
{
  "code": 200,
  "message": "更新成功",
  "data": {}
}
```

### 6.1.3 修改医生密码

**接口名称：** 修改医生密码  
**路径：** `PUT /api/doctors/password`  
**权限：** 医生

### 请求体

```json
{
  "old_password": "123456",
  "new_password": "654321"
}
```

### 响应

```json
{
  "code": 200,
  "message": "密码修改成功",
  "data": {}
}
```

## 6.2 医生预约管理模块

### 6.2.1 获取医生预约列表

**接口名称：** 获取分配给我的预约列表  
**路径：** `GET /api/doctor/appointments?page=1&page_size=10&status=1`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1001,
        "appointment_no": "APT202603250001",
        "pet_id": 1,
        "pet_name": "小橘",
        "user_id": 1,
        "user_nickname": "小猫家长",
        "appointment_type": 2,
        "appointment_time": "2026-03-26 14:00:00",
        "status": 1
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 6.2.2 获取预约接诊详情

**接口名称：** 获取预约接诊详情  
**路径：** `GET /api/doctor/appointments/{appointment_id}/detail`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "appointment": {
      "id": 1001,
      "appointment_no": "APT202603250001",
      "appointment_type": 2,
      "symptom_description": "近两天食欲差，精神萎靡",
      "appointment_time": "2026-03-26 14:00:00",
      "status": 1
    },
    "pet": {
      "id": 1,
      "pet_name": "小橘",
      "pet_type": "猫",
      "gender": 1,
      "age": 2,
      "age_unit": "year",
      "breed": "中华田园猫",
      "weight": "4.50",
      "sterilized": 1,
      "remark": "有点挑食"
    },
    "medical_histories": [
      {
        "id": 1,
        "history_type": "呼吸系统疾病",
        "description": "2025年曾出现上呼吸道感染",
        "diagnosed_at": "2025-06-01 10:00:00",
        "is_current": 0
      }
    ],
    "vaccinations": [
      {
        "id": 1,
        "vaccine_name": "猫三联",
        "vaccination_date": "2026-03-01",
        "next_due_date": "2027-03-01"
      }
    ],
    "allergies": [
      {
        "id": 1,
        "allergen": "海鲜类罐头",
        "symptom_description": "进食后呕吐",
        "severity_level": 2
      }
    ]
  }
}
```

### 6.2.3 更新预约状态

**接口名称：** 更新预约状态  
**路径：** `PUT /api/doctor/appointments/{appointment_id}/status`  
**权限：** 医生

### 请求体

```json
{
  "status": 2
}
```

### 响应

```json
{
  "code": 200,
  "message": "状态更新成功",
  "data": {}
}
```

## 6.3 病历管理模块

### 6.3.1 创建病历记录

**接口名称：** 创建病历记录  
**路径：** `POST /api/doctor/medical-records`  
**权限：** 医生

### 请求体

```json
{
  "appointment_id": 1001,
  "pet_id": 1,
  "user_id": 1,
  "chief_complaint": "食欲下降两天",
  "present_history": "精神略差，无明显腹泻",
  "physical_examination": "体温正常，腹部轻微压痛",
  "preliminary_diagnosis": "疑似胃肠炎",
  "treatment_plan": "对症支持治疗",
  "prescription": "益生菌、胃肠处方粮",
  "doctor_advice": "观察48小时，如症状加重及时复诊",
  "visit_time": "2026-03-26 14:30:00",
  "status": 1
}
```

### 响应

```json
{
  "code": 200,
  "message": "病历创建成功",
  "data": {
    "medical_record_id": 1
  }
}
```

### 6.3.2 更新病历记录

**接口名称：** 更新病历记录  
**路径：** `PUT /api/doctor/medical-records/{medical_record_id}`  
**权限：** 医生

### 请求体

```json
{
  "preliminary_diagnosis": "急性胃肠炎",
  "treatment_plan": "继续口服药物并复查",
  "doctor_advice": "饮食清淡，观察排便",
  "status": 2
}
```

### 响应

```json
{
  "code": 200,
  "message": "病历更新成功",
  "data": {}
}
```

### 6.3.3 获取病历详情

**接口名称：** 获取病历详情  
**路径：** `GET /api/doctor/medical-records/{medical_record_id}`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "appointment_id": 1001,
    "pet_id": 1,
    "user_id": 1,
    "doctor_id": 2,
    "chief_complaint": "食欲下降两天",
    "present_history": "精神略差，无明显腹泻",
    "physical_examination": "体温正常，腹部轻微压痛",
    "preliminary_diagnosis": "急性胃肠炎",
    "treatment_plan": "继续口服药物并复查",
    "prescription": "益生菌、胃肠处方粮",
    "doctor_advice": "饮食清淡，观察排便",
    "visit_time": "2026-03-26 14:30:00",
    "status": 2
  }
}
```

## 6.4 医疗报告模块

### 6.4.1 上传医疗报告

**接口名称：** 上传医疗报告  
**路径：** `POST /api/doctor/medical-records/{medical_record_id}/reports`  
**权限：** 医生  
**Content-Type：** `multipart/form-data`

### 表单参数

| 字段 | 类型 | 必填 | 说明 |
|---|---|---:|---|
| report_title | string | 是 | 报告标题 |
| report_type | string | 否 | 报告类型 |
| report_content | string | 否 | 文字内容 |
| file | file | 否 | 报告附件 |

### 响应

```json
{
  "code": 200,
  "message": "上传成功",
  "data": {
    "report_id": 1,
    "file_url": "https://example.com/reports/report-1.pdf"
  }
}
```

### 6.4.2 获取报告列表

**接口名称：** 获取病历报告列表  
**路径：** `GET /api/doctor/medical-records/{medical_record_id}/reports`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "report_title": "血常规检查报告",
        "report_type": "检验报告",
        "file_url": "https://example.com/reports/report-1.pdf",
        "report_content": "白细胞轻度升高",
        "uploaded_at": "2026-03-26 15:20:00"
      }
    ]
  }
}
```

### 6.4.3 删除医疗报告

**接口名称：** 删除医疗报告  
**路径：** `DELETE /api/doctor/reports/{report_id}`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "删除成功",
  "data": {}
}
```

## 6.5 医疗知识库管理（RAG 支撑）

### 6.5.1 上传知识库文档

**接口名称：** 上传知识库文档  
**路径：** `POST /api/admin/knowledge/upload`  
**权限：** 管理员  
**Content-Type：** `multipart/form-data`

### 请求体

| 字段 | 类型 | 必填 | 说明 |
|---|---|---|---|
| file | file | 是 | PDF 或 Markdown 医疗指南文件 |
| category | string | 否 | 知识库分类，如 `guide`、`disease`、`drug` |
| title | string | 否 | 文档标题 |

### 响应

```json
{
  "code": 200,
  "message": "上传成功，已进入向量化队列",
  "data": {
    "knowledge_id": 1,
    "file_name": "feline_digestive_guide.pdf",
    "status": "processing"
  }
}
```

### 说明

- 后端接收 PDF/Markdown 文件后，自动执行切片、向量化并写入知识库。
- `status` 可取值：`processing`、`completed`、`failed`。

### 6.5.2 查看向量化进度

**接口名称：** 查看知识库向量化进度  
**路径：** `GET /api/admin/knowledge/status`  
**权限：** 管理员

### 请求参数

| 参数 | 位置 | 必填 | 类型 | 说明 |
|---|---|---|---|---|
| knowledge_id | query | 否 | integer | 指定知识库文档 ID；不传则返回最近任务列表 |

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "knowledge_id": 1,
        "file_name": "feline_digestive_guide.pdf",
        "status": "completed",
        "progress": 100,
        "chunk_count": 86,
        "vector_count": 86,
        "error_message": "",
        "updated_at": "2026-04-02 14:30:00"
      }
    ]
  }
}
```

## 6.6 AI 查看模块

### 6.6.1 获取 AI 会话列表

**接口名称：** 获取 AI 会话列表  
**路径：** `GET /api/doctor/ai/sessions?page=1&page_size=10&pet_id=1`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "session_no": "AIS202603250001",
        "pet_id": 1,
        "pet_name": "小橘",
        "model_name": "pet-med-llm-v1",
        "status": 1,
        "created_at": "2026-03-25 10:55:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 6.6.2 获取 AI 会话消息记录

**接口名称：** 获取 AI 会话消息记录  
**路径：** `GET /api/doctor/ai/sessions/{session_id}/messages?page=1&page_size=20`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "sender_type": 1,
        "sender_id": 1,
        "message_content": "我家猫最近没精神",
        "created_at": "2026-03-25 11:00:00"
      },
      {
        "id": 2,
        "sender_type": 2,
        "sender_id": null,
        "message_content": "建议尽快就诊并观察体温变化",
        "created_at": "2026-03-25 11:00:02"
      }
    ]
  }
}
```

### 6.6.3 获取 AI 分析结果

**接口名称：** 获取 AI 分析结果  
**路径：** `GET /api/doctor/ai/sessions/{session_id}/analysis-records`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "analysis_type": 2,
        "input_source": 1,
        "analysis_result": "症状集中于食欲下降与精神不振",
        "rule_based_result": "建议排查消化系统问题",
        "llm_based_result": "综合风险为中等",
        "risk_level": 2,
        "reviewed_by_doctor": 0,
        "created_at": "2026-03-25 11:01:00"
      }
    ]
  }
}
```

## 6.6 医生通知模块

### 6.6.1 获取医生通知列表

**接口名称：** 获取医生通知列表  
**路径：** `GET /api/doctor/notifications?page=1&page_size=10`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "notification_type": 1,
        "title": "接诊提醒",
        "content": "您在 2026-03-26 14:00:00 有一条新的预约",
        "appointment_id": 1001,
        "send_time": "2026-03-26 13:00:00",
        "status": 1
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 6.6.2 标记医生通知已读

**接口名称：** 标记医生通知已读  
**路径：** `PUT /api/doctor/notifications/{notification_id}/read`  
**权限：** 医生

### 响应

```json
{
  "code": 200,
  "message": "已标记为已读",
  "data": {}
}
```

---

# 7. 管理端模块

> 管理端负责医生信息管理、医院信息管理、系统通知发布与系统日志查看。

## 7.1 管理员信息模块

### 7.1.1 获取管理员个人信息

**接口名称：** 获取管理员个人信息  
**路径：** `GET /api/admin/profile`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "admin001",
    "real_name": "系统管理员",
    "phone": "13700000000",
    "status": 1,
    "created_at": "2026-03-25 09:00:00"
  }
}
```

## 7.2 医生管理模块

### 7.2.1 新增医生

**接口名称：** 新增医生  
**路径：** `POST /api/admin/doctors`  
**权限：** 管理员

### 请求体

```json
{
  "hospital_id": 1,
  "username": "doctor_li",
  "password": "123456",
  "doctor_name": "李医生",
  "gender": 2,
  "phone": "13922223333",
  "email": "li@example.com",
  "title": "执业医师",
  "specialty": "猫科皮肤科",
  "avatar_url": "https://example.com/li.jpg",
  "intro": "擅长常见皮肤病诊疗",
  "status": 1
}
```

### 响应

```json
{
  "code": 200,
  "message": "新增成功",
  "data": {
    "doctor_id": 3
  }
}
```

### 7.2.2 获取医生列表

**接口名称：** 获取医生列表  
**路径：** `GET /api/admin/doctors?page=1&page_size=10&hospital_id=1&status=1&keyword=张`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 2,
        "hospital_id": 1,
        "hospital_name": "爱宠动物医院",
        "username": "doctor_zhang",
        "doctor_name": "张医生",
        "title": "主治医师",
        "specialty": "猫科内科",
        "phone": "13900000000",
        "status": 1,
        "created_at": "2026-03-20 10:00:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 7.2.3 获取医生详情

**接口名称：** 获取医生详情  
**路径：** `GET /api/admin/doctors/{doctor_id}`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 2,
    "hospital_id": 1,
    "username": "doctor_zhang",
    "doctor_name": "张医生",
    "gender": 1,
    "phone": "13900000000",
    "email": "doctor@example.com",
    "title": "主治医师",
    "specialty": "猫科内科",
    "avatar_url": "https://example.com/doctor.jpg",
    "intro": "擅长猫咪消化系统疾病诊疗",
    "status": 1,
    "created_at": "2026-03-20 10:00:00",
    "updated_at": "2026-03-25 09:00:00"
  }
}
```

### 7.2.4 修改医生信息

**接口名称：** 修改医生信息  
**路径：** `PUT /api/admin/doctors/{doctor_id}`  
**权限：** 管理员

### 请求体

```json
{
  "hospital_id": 1,
  "doctor_name": "张主任",
  "title": "副主任医师",
  "specialty": "猫科内科、老年病",
  "phone": "13999998888",
  "status": 1
}
```

### 响应

```json
{
  "code": 200,
  "message": "修改成功",
  "data": {}
}
```

### 7.2.5 删除医生

**接口名称：** 删除医生  
**路径：** `DELETE /api/admin/doctors/{doctor_id}`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "删除成功",
  "data": {}
}
```

## 7.3 医院管理模块

### 7.3.1 新增医院

**接口名称：** 新增医院  
**路径：** `POST /api/admin/hospitals`  
**权限：** 管理员

### 请求体

```json
{
  "hospital_name": "爱宠动物医院",
  "address": "深圳市南山区科技园一路100号",
  "phone": "0755-88886666",
  "description": "提供猫犬全科与专科医疗服务",
  "status": 1
}
```

### 响应

```json
{
  "code": 200,
  "message": "新增成功",
  "data": {
    "hospital_id": 1
  }
}
```

### 7.3.2 获取医院列表

**接口名称：** 获取医院列表  
**路径：** `GET /api/admin/hospitals?page=1&page_size=10&status=1&keyword=爱宠`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "hospital_name": "爱宠动物医院",
        "address": "深圳市南山区科技园一路100号",
        "phone": "0755-88886666",
        "status": 1,
        "created_at": "2026-03-20 09:00:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 7.3.3 获取医院详情

**接口名称：** 获取医院详情  
**路径：** `GET /api/admin/hospitals/{hospital_id}`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "hospital_name": "爱宠动物医院",
    "address": "深圳市南山区科技园一路100号",
    "phone": "0755-88886666",
    "description": "提供猫犬全科与专科医疗服务",
    "status": 1,
    "created_at": "2026-03-20 09:00:00",
    "updated_at": "2026-03-25 09:00:00"
  }
}
```

### 7.3.4 修改医院

**接口名称：** 修改医院  
**路径：** `PUT /api/admin/hospitals/{hospital_id}`  
**权限：** 管理员

### 请求体

```json
{
  "hospital_name": "爱宠动物专科医院",
  "address": "深圳市南山区科技园一路100号",
  "phone": "0755-88889999",
  "description": "升级为专科医院",
  "status": 1
}
```

### 响应

```json
{
  "code": 200,
  "message": "修改成功",
  "data": {}
}
```

### 7.3.5 删除医院

**接口名称：** 删除医院  
**路径：** `DELETE /api/admin/hospitals/{hospital_id}`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "删除成功",
  "data": {}
}
```

## 7.4 系统通知管理模块

### 7.4.1 发布系统通知

**接口名称：** 发布系统通知  
**路径：** `POST /api/admin/notifications`  
**权限：** 管理员

### 请求体

```json
{
  "receiver_type": "user",
  "receiver_ids": [1, 2, 3],
  "notification_type": 2,
  "title": "系统升级通知",
  "content": "平台将于今晚 23:00 进行系统维护",
  "send_time": "2026-03-25 20:00:00"
}
```

### 响应

```json
{
  "code": 200,
  "message": "发布成功",
  "data": {
    "count": 3
  }
}
```

### 7.4.2 获取系统通知列表

**接口名称：** 获取系统通知列表  
**路径：** `GET /api/admin/notifications?page=1&page_size=10&notification_type=2`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 10,
        "notification_type": 2,
        "title": "系统升级通知",
        "content": "平台将于今晚 23:00 进行系统维护",
        "send_time": "2026-03-25 20:00:00",
        "status": 1,
        "created_at": "2026-03-25 18:00:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 7.4.3 撤回系统通知

**接口名称：** 撤回系统通知  
**路径：** `DELETE /api/admin/notifications/{notification_id}`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "撤回成功",
  "data": {}
}
```

## 7.5 操作日志模块

### 7.5.1 获取操作日志列表

**接口名称：** 获取操作日志列表  
**路径：** `GET /api/admin/operation-logs?page=1&page_size=10&operator_type=2&operation_module=doctor`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "operator_type": 2,
        "operator_id": 2,
        "operation_module": "medical_record",
        "operation_type": "update",
        "operation_desc": "更新病历记录ID=1",
        "ip_address": "127.0.0.1",
        "created_at": "2026-03-26 15:10:00"
      }
    ],
    "pagination": {
      "page": 1,
      "page_size": 10,
      "total": 1
    }
  }
}
```

### 7.5.2 获取操作日志详情

**接口名称：** 获取操作日志详情  
**路径：** `GET /api/admin/operation-logs/{log_id}`  
**权限：** 管理员

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "operator_type": 2,
    "operator_id": 2,
    "operation_module": "medical_record",
    "operation_type": "update",
    "operation_desc": "更新病历记录ID=1",
    "ip_address": "127.0.0.1",
    "created_at": "2026-03-26 15:10:00"
  }
}
```

---

# 8. 公共辅助接口建议

## 8.1 获取启用中的医院下拉数据

**接口名称：** 获取医院下拉数据  
**路径：** `GET /api/common/hospitals/options`  
**权限：** 已登录

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 1,
      "hospital_name": "爱宠动物医院"
    }
  ]
}
```

## 8.2 获取医院医生下拉数据

**接口名称：** 获取医院医生下拉数据  
**路径：** `GET /api/common/hospitals/{hospital_id}/doctors/options`  
**权限：** 已登录

### 响应

```json
{
  "code": 200,
  "message": "success",
  "data": [
    {
      "id": 2,
      "doctor_name": "张医生",
      "title": "主治医师"
    }
  ]
}
```

## 8.3 通用文件上传

**接口名称：** 通用文件上传  
**路径：** `POST /api/common/upload`  
**权限：** 已登录  
**Content-Type：** `multipart/form-data`

### 表单参数

| 字段 | 类型 | 必填 | 说明 |
|---|---|---:|---|
| file | file | 是 | 上传文件 |
| biz_type | string | 否 | 业务类型，如 `avatar`、`report` |

### 响应

```json
{
  "code": 200,
  "message": "上传成功",
  "data": {
    "file_url": "https://example.com/uploads/2026/03/25/abc.jpg",
    "file_name": "abc.jpg"
  }
}
```

---

# 9. 权限矩阵建议

| 模块 | 用户 | 医生 | 管理员 |
|---|---:|---:|---:|
| 用户个人信息 | √ | × | × |
| 宠物档案管理 | √ | 只读 | × |
| 病史/疫苗/过敏管理 | √ | 只读 | × |
| 创建预约/取消预约 | √ | × | 可代创建 |
| 查看本人预约 | √ | × | √ |
| 查看分配预约 | × | √ | √ |
| 创建/更新病历 | × | √ | × |
| 上传报告 | × | √ | × |
| AI 会话发起 | √ | √ | × |
| AI 分析查看 | √ | √ | √ |
| 通知查看 | √ | √ | √ |
| 医生管理 | × | × | √ |
| 医院管理 | × | × | √ |
| 操作日志查看 | × | × | √ |

---

# 10. 命名与实现建议

## 10.1 URL 命名规范

- 资源集合使用复数：`/pets`、`/appointments`
- 资源从属关系清晰表达：`/pets/{pet_id}/vaccinations`
- 动作用子资源表示：`/appointments/{id}/cancel`、`/notifications/{id}/read`

## 10.2 Go 后端建议分层

- `router`：路由注册
- `handler/controller`：参数接收、响应封装
- `service`：业务逻辑
- `repository/dao`：数据库访问
- `model/entity`：表结构实体
- `dto/vo`：请求响应对象
- `middleware`：JWT、角色鉴权、日志记录

## 10.3 中间件建议

- JWT 认证中间件
- RBAC 角色校验中间件
- 操作日志中间件
- 请求参数校验中间件
- 文件上传大小与类型校验中间件

## 10.4 业务落地建议

- `appointment_no`、`session_no` 采用日期 + 流水号生成
- 取消预约前校验预约是否已完成或已过期
- 创建病历前校验预约状态与医生归属
- 删除类接口优先逻辑删除
- 通知建议使用定时任务 + 消息队列/Redis 延迟队列实现

---

# 11. 接口清单总览

## 11.1 认证模块

- `POST /api/auth/register`
- `POST /api/auth/login`
- `GET /api/auth/me`
- `POST /api/auth/logout`

## 11.2 用户端模块

- `GET /api/users/profile`
- `PUT /api/users/profile`
- `PUT /api/users/password`
- `POST /api/pets`
- `GET /api/pets`
- `GET /api/pets/{pet_id}`
- `PUT /api/pets/{pet_id}`
- `DELETE /api/pets/{pet_id}`
- `POST /api/pets/{pet_id}/medical-histories`
- `GET /api/pets/{pet_id}/medical-histories`
- `POST /api/pets/{pet_id}/vaccinations`
- `GET /api/pets/{pet_id}/vaccinations`
- `POST /api/pets/{pet_id}/allergies`
- `GET /api/pets/{pet_id}/allergies`
- `POST /api/appointments`
- `GET /api/appointments`
- `GET /api/appointments/{appointment_id}`
- `PUT /api/appointments/{appointment_id}/cancel`
- `GET /api/medical-records`
- `GET /api/medical-records/{medical_record_id}`
- `GET /api/medical-records/{medical_record_id}/reports`
- `POST /api/ai/sessions`
- `POST /api/ai/sessions/{session_id}/messages`
- `GET /api/ai/sessions/{session_id}`
- `GET /api/ai/sessions/{session_id}/messages`
- `GET /api/ai/sessions/{session_id}/analysis-records`
- `GET /api/notifications`
- `GET /api/notifications/{notification_id}`
- `PUT /api/notifications/{notification_id}/read`

## 11.3 医生端模块

- `GET /api/doctors/profile`
- `PUT /api/doctors/profile`
- `PUT /api/doctors/password`
- `GET /api/doctor/appointments`
- `GET /api/doctor/appointments/{appointment_id}/detail`
- `PUT /api/doctor/appointments/{appointment_id}/status`
- `POST /api/doctor/medical-records`
- `PUT /api/doctor/medical-records/{medical_record_id}`
- `GET /api/doctor/medical-records/{medical_record_id}`
- `POST /api/doctor/medical-records/{medical_record_id}/reports`
- `GET /api/doctor/medical-records/{medical_record_id}/reports`
- `DELETE /api/doctor/reports/{report_id}`
- `GET /api/doctor/ai/sessions`
- `GET /api/doctor/ai/sessions/{session_id}/messages`
- `GET /api/doctor/ai/sessions/{session_id}/analysis-records`
- `GET /api/doctor/notifications`
- `PUT /api/doctor/notifications/{notification_id}/read`

## 11.4 管理端模块

- `GET /api/admin/profile`
- `POST /api/admin/doctors`
- `GET /api/admin/doctors`
- `GET /api/admin/doctors/{doctor_id}`
- `PUT /api/admin/doctors/{doctor_id}`
- `DELETE /api/admin/doctors/{doctor_id}`
- `POST /api/admin/hospitals`
- `GET /api/admin/hospitals`
- `GET /api/admin/hospitals/{hospital_id}`
- `PUT /api/admin/hospitals/{hospital_id}`
- `DELETE /api/admin/hospitals/{hospital_id}`
- `POST /api/admin/notifications`
- `GET /api/admin/notifications`
- `DELETE /api/admin/notifications/{notification_id}`
- `POST /api/admin/knowledge/upload`
- `GET /api/admin/knowledge/status`
- `GET /api/admin/operation-logs`
- `GET /api/admin/operation-logs/{log_id}`

---
