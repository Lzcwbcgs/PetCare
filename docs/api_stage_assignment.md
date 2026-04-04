# 按 README 开发阶段划分的 API 接口清单

> 说明：
> - 本文依据 README 中的开发阶段划分进行整理：第一阶段到第五阶段。fileciteturn5file1
> - 接口条目的编号**严格保留原 API 文档中的编号**，方便后续回查原文。fileciteturn5file0
> - `[已完成] / [下一步] / [待开发]` 状态依据 README 当前开发进度标注。fileciteturn5file1

---

## 第一阶段：认证模块与测试基建 `[已完成]`

README 中第一阶段对应“认证模块与测试基建”，当前已完成接口为 `POST /api/auth/register`、`POST /api/auth/login`、`GET /api/auth/me`、`POST /api/auth/logout`。fileciteturn5file1

### 4. 认证模块

- **4.1 用户注册**  
  `POST /api/auth/register`
- **4.2 统一登录**  
  `POST /api/auth/login`
- **4.3 获取当前登录信息**  
  `GET /api/auth/me`
- **4.4 退出登录**  
  `POST /api/auth/logout`

---

## 第二阶段：资料模块 `[已完成]`

README 中第二阶段对应“用户资料模块、医生资料模块、管理员资料模块”，当前已完成接口为用户资料 3 个、医生资料 3 个、管理员资料 1 个。fileciteturn5file1

### 5.1 用户信息模块

- **5.1.1 获取个人信息**  
  `GET /api/users/profile`
- **5.1.2 更新个人信息**  
  `PUT /api/users/profile`
- **5.1.3 修改密码**  
  `PUT /api/users/password`

### 6.1 医生个人信息模块

- **6.1.1 获取医生个人信息**  
  `GET /api/doctors/profile`
- **6.1.2 更新医生个人信息**  
  `PUT /api/doctors/profile`
- **6.1.3 修改医生密码**  
  `PUT /api/doctors/password`

### 7.1 管理员信息模块

- **7.1.1 获取管理员个人信息**  
  `GET /api/admin/profile`

---

## 第三阶段：宠物档案模块、宠物病史 / 疫苗 / 过敏记录模块 `[下一步]`

README 中第三阶段明确包括“宠物档案模块”和“宠物病史 / 疫苗 / 过敏记录模块”。fileciteturn5file1

### 5.2 宠物管理模块

- **5.2.1 新增宠物档案**  
  `POST /api/pets`
- **5.2.2 获取宠物列表**  
  `GET /api/pets?page=1&page_size=10`
- **5.2.3 获取宠物详情**  
  `GET /api/pets/{pet_id}`
- **5.2.4 修改宠物档案**  
  `PUT /api/pets/{pet_id}`
- **5.2.5 删除宠物档案**  
  `DELETE /api/pets/{pet_id}`

### 5.3 宠物健康信息模块

- **5.3.1 新增病史记录**  
  `POST /api/pets/{pet_id}/medical-histories`
- **5.3.2 获取病史列表**  
  `GET /api/pets/{pet_id}/medical-histories?page=1&page_size=10`
- **5.3.3 新增疫苗记录**  
  `POST /api/pets/{pet_id}/vaccinations`
- **5.3.4 获取疫苗记录列表**  
  `GET /api/pets/{pet_id}/vaccinations?page=1&page_size=10`
- **5.3.5 新增过敏记录**  
  `POST /api/pets/{pet_id}/allergies`
- **5.3.6 获取过敏记录列表**  
  `GET /api/pets/{pet_id}/allergies?page=1&page_size=10`

> 补充说明：5.3.2、5.3.4、5.3.6 虽然权限写有“用户 / 医生”，但从 README 的阶段定义看，它们仍属于“宠物病史 / 疫苗 / 过敏记录模块”，因此统一归入第三阶段。fileciteturn5file0turn5file1

---

## 第四阶段：预约模块、病历与报告模块、通知模块 `[已开发]`

README 中第四阶段明确包括“预约模块、病历与报告模块、通知模块”。fileciteturn5file1

### 5.4 预约管理模块

- **5.4.1 创建预约**  
  `POST /api/appointments`
- **5.4.2 获取预约列表**  
  `GET /api/appointments?page=1&page_size=10&status=1&appointment_type=2`
- **5.4.3 获取预约详情**  
  `GET /api/appointments/{appointment_id}`
- **5.4.4 取消预约**  
  `PUT /api/appointments/{appointment_id}/cancel`

### 5.5 病历与报告查看模块

- **5.5.1 获取病历列表**  
  `GET /api/medical-records?page=1&page_size=10&pet_id=1`
- **5.5.2 获取病历详情**  
  `GET /api/medical-records/{medical_record_id}`
- **5.5.3 获取报告列表**  
  `GET /api/medical-records/{medical_record_id}/reports`

### 5.7 通知模块

- **5.7.1 获取通知列表**  
  `GET /api/notifications?page=1&page_size=10&status=1`
- **5.7.2 获取通知详情**  
  `GET /api/notifications/{notification_id}`
- **5.7.3 标记通知已读**  
  `PUT /api/notifications/{notification_id}/read`

### 6.2 医生预约管理模块

- **6.2.1 获取医生预约列表**  
  `GET /api/doctor/appointments?page=1&page_size=10&status=1`
- **6.2.2 获取预约接诊详情**  
  `GET /api/doctor/appointments/{appointment_id}/detail`
- **6.2.3 更新预约状态**  
  `PUT /api/doctor/appointments/{appointment_id}/status`

### 6.3 病历管理模块

- **6.3.1 创建病历记录**  
  `POST /api/doctor/medical-records`
- **6.3.2 更新病历记录**  
  `PUT /api/doctor/medical-records/{medical_record_id}`
- **6.3.3 获取病历详情**  
  `GET /api/doctor/medical-records/{medical_record_id}`

### 6.4 医疗报告模块

- **6.4.1 上传医疗报告**  
  `POST /api/doctor/medical-records/{medical_record_id}/reports`
- **6.4.2 获取报告列表**  
  `GET /api/doctor/medical-records/{medical_record_id}/reports`
- **6.4.3 删除医疗报告**  
  `DELETE /api/doctor/reports/{report_id}`

### 6.6 医生通知模块

- **6.6.1 获取医生通知列表**  
  `GET /api/doctor/notifications?page=1&page_size=10`
- **6.6.2 标记医生通知已读**  
  `PUT /api/doctor/notifications/{notification_id}/read`

### 7.2 医生管理模块

- **7.2.1 新增医生**  
  `POST /api/admin/doctors`
- **7.2.2 获取医生列表**  
  `GET /api/admin/doctors?page=1&page_size=10&hospital_id=1&status=1&keyword=张`
- **7.2.3 获取医生详情**  
  `GET /api/admin/doctors/{doctor_id}`
- **7.2.4 修改医生信息**  
  `PUT /api/admin/doctors/{doctor_id}`
- **7.2.5 删除医生**  
  `DELETE /api/admin/doctors/{doctor_id}`

### 7.3 医院管理模块

- **7.3.1 新增医院**  
  `POST /api/admin/hospitals`
- **7.3.2 获取医院列表**  
  `GET /api/admin/hospitals?page=1&page_size=10&status=1&keyword=爱宠`
- **7.3.3 获取医院详情**  
  `GET /api/admin/hospitals/{hospital_id}`
- **7.3.4 修改医院**  
  `PUT /api/admin/hospitals/{hospital_id}`
- **7.3.5 删除医院**  
  `DELETE /api/admin/hospitals/{hospital_id}`

### 7.4 系统通知管理模块

- **7.4.1 发布系统通知**  
  `POST /api/admin/notifications`
- **7.4.2 获取系统通知列表**  
  `GET /api/admin/notifications?page=1&page_size=10&notification_type=2`
- **7.4.3 撤回系统通知**  
  `DELETE /api/admin/notifications/{notification_id}`

> 分配说明：虽然 README 第四阶段只写了“预约模块、病历与报告模块、通知模块”，但这些业务落地需要后台侧的医生管理、医院管理与通知发布能力共同支撑，因此 7.2、7.3、7.4 一并归入第四阶段，更符合联调与管理端交付顺序。fileciteturn5file0turn5file1

---

## 第五阶段：AI 会话模块、AI 分析记录模块、操作日志模块、医疗知识库管理（RAG 支撑） `[待开发]`

README 中第五阶段明确包括“AI 会话模块、AI 分析记录模块、操作日志模块”。补充将管理员端知识库管理接口一并纳入第五阶段，用于支撑 AI 问诊的 RAG 能力。

### 5.6 AI 问诊模块

- **5.6.1 创建 AI 会话**  
  `POST /api/ai/sessions`
- **5.6.2 发送 AI 消息（SSE 流式输出）**  
  `POST /api/ai/sessions/{session_id}/messages`
- **5.6.3 获取 AI 会话详情**  
  `GET /api/ai/sessions/{session_id}`
- **5.6.4 获取 AI 消息记录**  
  `GET /api/ai/sessions/{session_id}/messages?page=1&page_size=20`
- **5.6.5 获取 AI 分析结果**  
  `GET /api/ai/sessions/{session_id}/analysis-records`

### 6.5 医疗知识库管理（RAG 支撑）

- **6.5.1 上传知识库文档**  
  `POST /api/admin/knowledge/upload`
- **6.5.2 查看向量化进度**  
  `GET /api/admin/knowledge/status`

### 6.6 AI 查看模块

- **6.6.1 获取 AI 会话列表**  
  `GET /api/doctor/ai/sessions?page=1&page_size=10&pet_id=1`
- **6.6.2 获取 AI 会话消息记录**  
  `GET /api/doctor/ai/sessions/{session_id}/messages?page=1&page_size=20`
- **6.6.3 获取 AI 分析结果**  
  `GET /api/doctor/ai/sessions/{session_id}/analysis-records`

### 7.5 操作日志模块

- **7.5.1 获取操作日志列表**  
  `GET /api/admin/operation-logs?page=1&page_size=10&operator_type=2&operation_module=doctor`
- **7.5.2 获取操作日志详情**  
  `GET /api/admin/operation-logs/{log_id}`

---

## 公共辅助接口的归属建议

API 文档中还包含一组“公共辅助接口建议”，README 没有单独为它们设阶段。为便于实际开发，这里给出建议归属。fileciteturn5file0turn5file1

### 建议归入第四阶段

- **8.1 获取启用中的医院下拉数据**  
  `GET /api/common/hospitals/options`  
  说明：主要服务于预约创建、后台配置等场景。
- **8.2 获取医院医生下拉数据**  
  `GET /api/common/hospitals/{hospital_id}/doctors/options`  
  说明：主要服务于预约创建和医生选择。
- **8.3 通用文件上传**  
  `POST /api/common/upload`  
  说明：主要服务于头像、报告等上传场景，与病历/报告模块联动较强。

---

## 汇总视图

- **第一阶段**：4.1 ~ 4.4
- **第二阶段**：5.1.1 ~ 5.1.3，6.1.1 ~ 6.1.3，7.1.1
- **第三阶段**：5.2.1 ~ 5.2.5，5.3.1 ~ 5.3.6
- **第四阶段**：5.4.1 ~ 5.4.4，5.5.1 ~ 5.5.3，5.7.1 ~ 5.7.3，6.2.1 ~ 6.2.3，6.3.1 ~ 6.3.3，6.4.1 ~ 6.4.3，6.6.1 ~ 6.6.2，7.2.1 ~ 7.2.5，7.3.1 ~ 7.3.5，7.4.1 ~ 7.4.3，8.1 ~ 8.3（建议）
- **第五阶段**：5.6.1 ~ 5.6.5，6.5.1 ~ 6.5.3，7.5.1 ~ 7.5.2
