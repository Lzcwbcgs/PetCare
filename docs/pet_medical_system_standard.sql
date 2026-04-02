-- 宠物医疗预约与档案管理系统
-- 标准版数据库建表脚本（含第五阶段 AI/RAG/操作日志能力）
-- 适用数据库：MySQL 8.0+
-- 字符集：utf8mb4

CREATE DATABASE IF NOT EXISTS `pet_medical_system`
DEFAULT CHARACTER SET utf8mb4
DEFAULT COLLATE utf8mb4_unicode_ci;

USE `pet_medical_system`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `operation_log`;
DROP TABLE IF EXISTS `notification`;
DROP TABLE IF EXISTS `knowledge_job`;
DROP TABLE IF EXISTS `knowledge_chunk`;
DROP TABLE IF EXISTS `knowledge_document`;
DROP TABLE IF EXISTS `ai_analysis_record`;
DROP TABLE IF EXISTS `ai_message`;
DROP TABLE IF EXISTS `ai_session`;
DROP TABLE IF EXISTS `medical_report`;
DROP TABLE IF EXISTS `medical_record`;
DROP TABLE IF EXISTS `appointment`;
DROP TABLE IF EXISTS `pet_allergy_record`;
DROP TABLE IF EXISTS `pet_vaccination_record`;
DROP TABLE IF EXISTS `pet_medical_history`;
DROP TABLE IF EXISTS `pet`;
DROP TABLE IF EXISTS `doctor`;
DROP TABLE IF EXISTS `hospital`;
DROP TABLE IF EXISTS `admin`;
DROP TABLE IF EXISTS `user`;

SET FOREIGN_KEY_CHECKS = 1;

-- 1. 宠物主人用户表
CREATE TABLE `user` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '用户主键ID',
    `username` VARCHAR(50) NOT NULL COMMENT '登录用户名',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
    `nickname` VARCHAR(50) DEFAULT NULL COMMENT '昵称',
    `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
    `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
    `avatar_url` VARCHAR(255) DEFAULT NULL COMMENT '用户头像URL',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常，0禁用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_user_username` (`username`),
    KEY `idx_user_phone` (`phone`),
    KEY `idx_user_email` (`email`),
    KEY `idx_user_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物主人用户表';

-- 2. 管理员表
CREATE TABLE `admin` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '管理员主键ID',
    `username` VARCHAR(50) NOT NULL COMMENT '管理员账号',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
    `real_name` VARCHAR(50) DEFAULT NULL COMMENT '管理员姓名',
    `phone` VARCHAR(20) DEFAULT NULL COMMENT '联系电话',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常，0禁用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_admin_username` (`username`),
    KEY `idx_admin_phone` (`phone`),
    KEY `idx_admin_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表';

-- 3. 宠物医院表
CREATE TABLE `hospital` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '医院主键ID',
    `hospital_name` VARCHAR(100) NOT NULL COMMENT '医院名称',
    `address` VARCHAR(255) DEFAULT NULL COMMENT '医院地址',
    `phone` VARCHAR(20) DEFAULT NULL COMMENT '医院联系电话',
    `description` TEXT DEFAULT NULL COMMENT '医院简介',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1启用，0停用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_hospital_name` (`hospital_name`),
    KEY `idx_hospital_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物医院表';

-- 4. 医生表
CREATE TABLE `doctor` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '医生主键ID',
    `hospital_id` BIGINT NOT NULL COMMENT '所属医院ID',
    `username` VARCHAR(50) NOT NULL COMMENT '医生登录账号',
    `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
    `doctor_name` VARCHAR(50) NOT NULL COMMENT '医生姓名',
    `gender` TINYINT DEFAULT NULL COMMENT '性别：1男，2女，0未知',
    `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
    `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
    `title` VARCHAR(50) DEFAULT NULL COMMENT '职称',
    `specialty` VARCHAR(255) DEFAULT NULL COMMENT '擅长领域',
    `avatar_url` VARCHAR(255) DEFAULT NULL COMMENT '头像URL',
    `intro` TEXT DEFAULT NULL COMMENT '医生简介',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1在职可接诊，0停用',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_doctor_username` (`username`),
    KEY `idx_doctor_hospital_id` (`hospital_id`),
    KEY `idx_doctor_name` (`doctor_name`),
    KEY `idx_doctor_status` (`status`),
    CONSTRAINT `fk_doctor_hospital`
        FOREIGN KEY (`hospital_id`) REFERENCES `hospital`(`id`)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医生表';

-- 5. 宠物档案表
CREATE TABLE `pet` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '宠物主键ID',
    `user_id` BIGINT NOT NULL COMMENT '所属用户ID',
    `pet_name` VARCHAR(50) NOT NULL COMMENT '宠物名字',
    `pet_type` VARCHAR(20) NOT NULL DEFAULT '猫' COMMENT '宠物类型，当前默认猫',
    `avatar_url` VARCHAR(255) DEFAULT NULL COMMENT '宠物头像URL',
    `gender` TINYINT DEFAULT NULL COMMENT '性别：1公，2母，0未知',
    `age` INT DEFAULT NULL COMMENT '年龄',
    `age_unit` VARCHAR(10) DEFAULT NULL COMMENT '年龄单位：month/月，year/岁',
    `breed` VARCHAR(50) DEFAULT NULL COMMENT '品种',
    `weight` DECIMAL(5,2) DEFAULT NULL COMMENT '体重（kg）',
    `sterilized` TINYINT DEFAULT NULL COMMENT '是否绝育：1是，0否',
    `remark` TEXT DEFAULT NULL COMMENT '备注',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常，0停用/删除',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_pet_user_id` (`user_id`),
    KEY `idx_pet_name` (`pet_name`),
    KEY `idx_pet_status` (`status`),
    CONSTRAINT `fk_pet_user`
        FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物档案表';

-- 6. 宠物病史表
CREATE TABLE `pet_medical_history` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '病史记录ID',
    `pet_id` BIGINT NOT NULL COMMENT '宠物ID',
    `history_type` VARCHAR(50) DEFAULT NULL COMMENT '病史类型',
    `description` TEXT NOT NULL COMMENT '病史描述',
    `diagnosed_at` DATETIME DEFAULT NULL COMMENT '确诊时间',
    `is_current` TINYINT NOT NULL DEFAULT 0 COMMENT '是否当前仍存在：1是，0否',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_history_pet_id` (`pet_id`),
    KEY `idx_history_is_current` (`is_current`),
    CONSTRAINT `fk_history_pet`
        FOREIGN KEY (`pet_id`) REFERENCES `pet`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物病史表';

-- 7. 宠物疫苗接种记录表
CREATE TABLE `pet_vaccination_record` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '疫苗记录ID',
    `pet_id` BIGINT NOT NULL COMMENT '宠物ID',
    `vaccine_name` VARCHAR(100) NOT NULL COMMENT '疫苗名称',
    `vaccination_date` DATE DEFAULT NULL COMMENT '接种日期',
    `next_due_date` DATE DEFAULT NULL COMMENT '下次应接种日期',
    `hospital_name` VARCHAR(100) DEFAULT NULL COMMENT '接种机构',
    `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_vaccine_pet_id` (`pet_id`),
    KEY `idx_vaccine_name` (`vaccine_name`),
    CONSTRAINT `fk_vaccine_pet`
        FOREIGN KEY (`pet_id`) REFERENCES `pet`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物疫苗接种记录表';

-- 8. 宠物过敏记录表
CREATE TABLE `pet_allergy_record` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '过敏记录ID',
    `pet_id` BIGINT NOT NULL COMMENT '宠物ID',
    `allergen` VARCHAR(100) NOT NULL COMMENT '过敏源',
    `symptom_description` VARCHAR(255) DEFAULT NULL COMMENT '症状描述',
    `severity_level` TINYINT DEFAULT NULL COMMENT '严重程度：1轻微，2中等，3严重',
    `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_allergy_pet_id` (`pet_id`),
    KEY `idx_allergy_severity` (`severity_level`),
    CONSTRAINT `fk_allergy_pet`
        FOREIGN KEY (`pet_id`) REFERENCES `pet`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物过敏记录表';

-- 9. 预约表
CREATE TABLE `appointment` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '预约主键ID',
    `appointment_no` VARCHAR(50) NOT NULL COMMENT '预约单号',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `pet_id` BIGINT NOT NULL COMMENT '宠物ID',
    `hospital_id` BIGINT NOT NULL COMMENT '医院ID',
    `doctor_id` BIGINT DEFAULT NULL COMMENT '医生ID，可为空',
    `appointment_type` TINYINT NOT NULL COMMENT '预约类型：1体检预约，2看病预约',
    `symptom_description` TEXT DEFAULT NULL COMMENT '症状描述（看病预约时填写）',
    `appointment_time` DATETIME NOT NULL COMMENT '预约时间',
    `reminder_time` DATETIME DEFAULT NULL COMMENT '提醒时间，一般为预约前1小时',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1待就诊，2已完成，3已取消，4已过期',
    `source` TINYINT NOT NULL DEFAULT 1 COMMENT '来源：1用户端预约，2医生代录入，3后台创建',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_appointment_no` (`appointment_no`),
    KEY `idx_appointment_user_id` (`user_id`),
    KEY `idx_appointment_pet_id` (`pet_id`),
    KEY `idx_appointment_hospital_id` (`hospital_id`),
    KEY `idx_appointment_doctor_id` (`doctor_id`),
    KEY `idx_appointment_time` (`appointment_time`),
    KEY `idx_appointment_status` (`status`),
    CONSTRAINT `fk_appointment_user`
        FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_appointment_pet`
        FOREIGN KEY (`pet_id`) REFERENCES `pet`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_appointment_hospital`
        FOREIGN KEY (`hospital_id`) REFERENCES `hospital`(`id`)
        ON UPDATE CASCADE ON DELETE RESTRICT,
    CONSTRAINT `fk_appointment_doctor`
        FOREIGN KEY (`doctor_id`) REFERENCES `doctor`(`id`)
        ON UPDATE CASCADE ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='预约表';

-- 10. 病历记录表
CREATE TABLE `medical_record` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '病历主键ID',
    `appointment_id` BIGINT NOT NULL COMMENT '关联预约ID',
    `pet_id` BIGINT NOT NULL COMMENT '宠物ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `doctor_id` BIGINT NOT NULL COMMENT '接诊医生ID',
    `chief_complaint` TEXT DEFAULT NULL COMMENT '主诉',
    `present_history` TEXT DEFAULT NULL COMMENT '现病史',
    `physical_examination` TEXT DEFAULT NULL COMMENT '体格检查结果',
    `preliminary_diagnosis` TEXT DEFAULT NULL COMMENT '初步诊断',
    `treatment_plan` TEXT DEFAULT NULL COMMENT '治疗方案',
    `prescription` TEXT DEFAULT NULL COMMENT '处方建议',
    `doctor_advice` TEXT DEFAULT NULL COMMENT '医嘱',
    `visit_time` DATETIME NOT NULL COMMENT '就诊时间',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1已创建，2已完成，3已归档',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_medical_record_appointment_id` (`appointment_id`),
    KEY `idx_medical_record_pet_id` (`pet_id`),
    KEY `idx_medical_record_user_id` (`user_id`),
    KEY `idx_medical_record_doctor_id` (`doctor_id`),
    KEY `idx_medical_record_visit_time` (`visit_time`),
    CONSTRAINT `fk_record_appointment`
        FOREIGN KEY (`appointment_id`) REFERENCES `appointment`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_record_pet`
        FOREIGN KEY (`pet_id`) REFERENCES `pet`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_record_user`
        FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_record_doctor`
        FOREIGN KEY (`doctor_id`) REFERENCES `doctor`(`id`)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='病历记录表';

-- 11. 医疗报告表
CREATE TABLE `medical_report` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '报告主键ID',
    `medical_record_id` BIGINT NOT NULL COMMENT '病历ID',
    `doctor_id` BIGINT NOT NULL COMMENT '上传医生ID',
    `report_title` VARCHAR(100) NOT NULL COMMENT '报告标题',
    `report_type` VARCHAR(50) DEFAULT NULL COMMENT '报告类型，如检查报告、检验报告、处置报告',
    `file_url` VARCHAR(255) DEFAULT NULL COMMENT '报告文件路径',
    `report_content` TEXT DEFAULT NULL COMMENT '报告文字内容',
    `uploaded_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
    KEY `idx_medical_report_record_id` (`medical_record_id`),
    KEY `idx_medical_report_doctor_id` (`doctor_id`),
    CONSTRAINT `fk_report_record`
        FOREIGN KEY (`medical_record_id`) REFERENCES `medical_record`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_report_doctor`
        FOREIGN KEY (`doctor_id`) REFERENCES `doctor`(`id`)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医疗报告表';

-- 12. AI问诊会话表
CREATE TABLE `ai_session` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'AI会话ID',
    `session_no` VARCHAR(50) NOT NULL COMMENT '会话编号',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `pet_id` BIGINT NOT NULL COMMENT '宠物ID',
    `hospital_id` BIGINT DEFAULT NULL COMMENT '关联医院ID，可为空',
    `doctor_id` BIGINT DEFAULT NULL COMMENT '关联医生ID，可为空',
    `source_type` TINYINT NOT NULL DEFAULT 1 COMMENT '来源：1用户端发起，2医生端发起',
    `model_type` VARCHAR(50) DEFAULT NULL COMMENT '模型类型，如api/local',
    `model_name` VARCHAR(100) DEFAULT NULL COMMENT '模型名称',
    `provider_name` VARCHAR(100) DEFAULT NULL COMMENT '模型提供方名称',
    `session_summary` TEXT DEFAULT NULL COMMENT 'AI会话总结',
    `rag_enabled` TINYINT NOT NULL DEFAULT 1 COMMENT '是否启用RAG：1是，0否',
    `retrieval_count` INT NOT NULL DEFAULT 0 COMMENT '最近一次召回片段数',
    `sync_to_admin` TINYINT NOT NULL DEFAULT 1 COMMENT '是否同步给管理端：1是，0否',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1进行中，2已结束，3已归档',
    `last_message_at` DATETIME DEFAULT NULL COMMENT '最后消息时间',
    `extra_metadata` JSON DEFAULT NULL COMMENT '扩展元数据，如召回来源、上下文摘要等',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_ai_session_session_no` (`session_no`),
    KEY `idx_ai_session_user_id` (`user_id`),
    KEY `idx_ai_session_pet_id` (`pet_id`),
    KEY `idx_ai_session_hospital_id` (`hospital_id`),
    KEY `idx_ai_session_doctor_id` (`doctor_id`),
    KEY `idx_ai_session_status` (`status`),
    KEY `idx_ai_session_last_message_at` (`last_message_at`),
    CONSTRAINT `fk_ai_session_user`
        FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_ai_session_pet`
        FOREIGN KEY (`pet_id`) REFERENCES `pet`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_ai_session_hospital`
        FOREIGN KEY (`hospital_id`) REFERENCES `hospital`(`id`)
        ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT `fk_ai_session_doctor`
        FOREIGN KEY (`doctor_id`) REFERENCES `doctor`(`id`)
        ON UPDATE CASCADE ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='AI问诊会话表';

-- 13. AI消息记录表
CREATE TABLE `ai_message` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'AI消息ID',
    `session_id` BIGINT NOT NULL COMMENT 'AI会话ID',
    `sender_type` TINYINT NOT NULL COMMENT '发送者类型：1用户，2AI，3医生，4管理员',
    `sender_id` BIGINT DEFAULT NULL COMMENT '发送者ID，可为空',
    `message_content` LONGTEXT NOT NULL COMMENT '消息内容',
    `message_type` TINYINT NOT NULL DEFAULT 1 COMMENT '消息类型：1文本，2系统事件，3结构化结果',
    `provider_type` VARCHAR(50) DEFAULT NULL COMMENT '模型类型，如api/local',
    `provider_name` VARCHAR(100) DEFAULT NULL COMMENT '模型提供方名称',
    `prompt_tokens` INT NOT NULL DEFAULT 0 COMMENT '提示词token数',
    `completion_tokens` INT NOT NULL DEFAULT 0 COMMENT '回复token数',
    `finish_reason` VARCHAR(50) DEFAULT NULL COMMENT '停止原因，如stop/length/error',
    `extra_metadata` JSON DEFAULT NULL COMMENT '扩展元数据，如引用片段、SSE状态、错误信息等',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
    KEY `idx_ai_message_session_id` (`session_id`),
    KEY `idx_ai_message_sender_type` (`sender_type`),
    KEY `idx_ai_message_created_at` (`created_at`),
    CONSTRAINT `fk_ai_message_session`
        FOREIGN KEY (`session_id`) REFERENCES `ai_session`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='AI消息记录表';

-- 14. AI智能分析结果表
CREATE TABLE `ai_analysis_record` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'AI分析记录ID',
    `pet_id` BIGINT NOT NULL COMMENT '宠物ID',
    `session_id` BIGINT DEFAULT NULL COMMENT '关联AI会话ID',
    `medical_record_id` BIGINT DEFAULT NULL COMMENT '关联病历ID',
    `analysis_type` TINYINT NOT NULL COMMENT '分析类型：1病历总结，2症状归纳，3风险提示，4健康建议',
    `input_source` TINYINT NOT NULL COMMENT '输入来源：1AI对话，2病历记录，3体检数据，4综合数据',
    `summary_title` VARCHAR(100) DEFAULT NULL COMMENT '分析标题',
    `analysis_result` LONGTEXT NOT NULL COMMENT '分析结果',
    `rule_based_result` LONGTEXT DEFAULT NULL COMMENT '规则分析结果',
    `llm_based_result` LONGTEXT DEFAULT NULL COMMENT '大模型分析结果',
    `risk_level` TINYINT DEFAULT NULL COMMENT '风险等级：1低，2中，3高',
    `confidence_score` DECIMAL(5,2) DEFAULT NULL COMMENT '置信度分数，0-100',
    `reference_chunks` JSON DEFAULT NULL COMMENT '引用知识片段信息',
    `extra_metadata` JSON DEFAULT NULL COMMENT '扩展元数据',
    `reviewed_by_doctor` TINYINT NOT NULL DEFAULT 0 COMMENT '医生是否已审核：1是，0否',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    KEY `idx_ai_analysis_pet_id` (`pet_id`),
    KEY `idx_ai_analysis_session_id` (`session_id`),
    KEY `idx_ai_analysis_medical_record_id` (`medical_record_id`),
    KEY `idx_ai_analysis_type` (`analysis_type`),
    CONSTRAINT `fk_analysis_pet`
        FOREIGN KEY (`pet_id`) REFERENCES `pet`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_analysis_session`
        FOREIGN KEY (`session_id`) REFERENCES `ai_session`(`id`)
        ON UPDATE CASCADE ON DELETE SET NULL,
    CONSTRAINT `fk_analysis_record`
        FOREIGN KEY (`medical_record_id`) REFERENCES `medical_record`(`id`)
        ON UPDATE CASCADE ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='AI智能分析结果表';

-- 15. 知识库文档表
CREATE TABLE `knowledge_document` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '知识文档ID',
    `document_no` VARCHAR(50) NOT NULL COMMENT '文档编号',
    `title` VARCHAR(200) NOT NULL COMMENT '文档标题',
    `document_type` VARCHAR(50) NOT NULL COMMENT '文档类型，如txt/md/pdf',
    `source_type` TINYINT NOT NULL DEFAULT 1 COMMENT '来源类型：1后台上传，2系统导入',
    `file_name` VARCHAR(255) NOT NULL COMMENT '原始文件名',
    `file_url` VARCHAR(255) NOT NULL COMMENT '文件访问路径',
    `file_size` BIGINT NOT NULL DEFAULT 0 COMMENT '文件大小，单位字节',
    `content_text` LONGTEXT DEFAULT NULL COMMENT '解析后的纯文本内容',
    `embedding_provider` VARCHAR(100) DEFAULT NULL COMMENT '向量化提供方',
    `embedding_model` VARCHAR(100) DEFAULT NULL COMMENT '向量化模型',
    `chunk_count` INT NOT NULL DEFAULT 0 COMMENT '切片数量',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1已上传，2处理中，3成功，4失败，5停用',
    `error_message` VARCHAR(500) DEFAULT NULL COMMENT '失败原因',
    `created_by` BIGINT NOT NULL COMMENT '创建管理员ID',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_knowledge_document_no` (`document_no`),
    KEY `idx_knowledge_document_status` (`status`),
    KEY `idx_knowledge_document_created_by` (`created_by`),
    KEY `idx_knowledge_document_type` (`document_type`),
    CONSTRAINT `fk_knowledge_document_admin`
        FOREIGN KEY (`created_by`) REFERENCES `admin`(`id`)
        ON UPDATE CASCADE ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='知识库文档表';

-- 16. 知识库文档切片表
CREATE TABLE `knowledge_chunk` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '知识切片ID',
    `document_id` BIGINT NOT NULL COMMENT '所属文档ID',
    `chunk_no` VARCHAR(50) NOT NULL COMMENT '切片编号',
    `chunk_index` INT NOT NULL COMMENT '切片序号',
    `content` LONGTEXT NOT NULL COMMENT '切片内容',
    `token_count` INT NOT NULL DEFAULT 0 COMMENT '估算token数',
    `vector_store_type` VARCHAR(50) DEFAULT NULL COMMENT '向量库类型，如qdrant/milvus',
    `vector_collection` VARCHAR(100) DEFAULT NULL COMMENT '向量库集合名',
    `vector_point_id` VARCHAR(100) DEFAULT NULL COMMENT '向量点ID',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1待写入，2写入成功，3写入失败',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_knowledge_chunk_no` (`chunk_no`),
    UNIQUE KEY `uk_knowledge_chunk_doc_index` (`document_id`, `chunk_index`),
    KEY `idx_knowledge_chunk_document_id` (`document_id`),
    KEY `idx_knowledge_chunk_status` (`status`),
    KEY `idx_knowledge_chunk_point_id` (`vector_point_id`),
    CONSTRAINT `fk_knowledge_chunk_document`
        FOREIGN KEY (`document_id`) REFERENCES `knowledge_document`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='知识库文档切片表';

-- 17. 知识库处理任务表
CREATE TABLE `knowledge_job` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '知识任务ID',
    `job_no` VARCHAR(50) NOT NULL COMMENT '任务编号',
    `document_id` BIGINT NOT NULL COMMENT '文档ID',
    `job_type` TINYINT NOT NULL DEFAULT 1 COMMENT '任务类型：1解析切片，2向量化入库，3重建索引',
    `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1待执行，2处理中，3成功，4失败',
    `progress` INT NOT NULL DEFAULT 0 COMMENT '进度百分比',
    `started_at` DATETIME DEFAULT NULL COMMENT '开始时间',
    `finished_at` DATETIME DEFAULT NULL COMMENT '结束时间',
    `error_message` VARCHAR(500) DEFAULT NULL COMMENT '失败原因',
    `extra_metadata` JSON DEFAULT NULL COMMENT '扩展元数据',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY `uk_knowledge_job_no` (`job_no`),
    KEY `idx_knowledge_job_document_id` (`document_id`),
    KEY `idx_knowledge_job_status` (`status`),
    KEY `idx_knowledge_job_type` (`job_type`),
    CONSTRAINT `fk_knowledge_job_document`
        FOREIGN KEY (`document_id`) REFERENCES `knowledge_document`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='知识库处理任务表';

-- 18. 通知提醒表
CREATE TABLE `notification` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '通知ID',
    `user_id` BIGINT DEFAULT NULL COMMENT '接收用户ID',
    `doctor_id` BIGINT DEFAULT NULL COMMENT '接收医生ID',
    `appointment_id` BIGINT DEFAULT NULL COMMENT '关联预约ID',
    `notification_type` TINYINT NOT NULL COMMENT '通知类型：1预约提醒，2系统通知，3AI分析提醒',
    `title` VARCHAR(100) NOT NULL COMMENT '通知标题',
    `content` TEXT NOT NULL COMMENT '通知内容',
    `send_time` DATETIME NOT NULL COMMENT '计划发送时间',
    `status` TINYINT NOT NULL DEFAULT 0 COMMENT '状态：0待发送，1已发送，2发送失败，3已读',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY `idx_notification_user_id` (`user_id`),
    KEY `idx_notification_doctor_id` (`doctor_id`),
    KEY `idx_notification_appointment_id` (`appointment_id`),
    KEY `idx_notification_send_time` (`send_time`),
    KEY `idx_notification_status` (`status`),
    CONSTRAINT `fk_notification_user`
        FOREIGN KEY (`user_id`) REFERENCES `user`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_notification_doctor`
        FOREIGN KEY (`doctor_id`) REFERENCES `doctor`(`id`)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT `fk_notification_appointment`
        FOREIGN KEY (`appointment_id`) REFERENCES `appointment`(`id`)
        ON UPDATE CASCADE ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='通知提醒表';

-- 19. 系统操作日志表
CREATE TABLE `operation_log` (
    `id` BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT '日志ID',
    `operator_type` TINYINT NOT NULL COMMENT '操作人类型：1管理员，2医生，3用户',
    `operator_id` BIGINT NOT NULL COMMENT '操作人ID',
    `operation_module` VARCHAR(50) NOT NULL COMMENT '操作模块',
    `operation_type` VARCHAR(50) NOT NULL COMMENT '操作类型',
    `target_type` VARCHAR(50) DEFAULT NULL COMMENT '目标对象类型，如knowledge_document/ai_session',
    `target_id` BIGINT DEFAULT NULL COMMENT '目标对象ID',
    `request_method` VARCHAR(10) DEFAULT NULL COMMENT '请求方法',
    `request_path` VARCHAR(255) DEFAULT NULL COMMENT '请求路径',
    `operation_result` TINYINT NOT NULL DEFAULT 1 COMMENT '操作结果：1成功，0失败',
    `operation_desc` TEXT DEFAULT NULL COMMENT '操作描述',
    `ip_address` VARCHAR(50) DEFAULT NULL COMMENT 'IP地址',
    `extra_data` JSON DEFAULT NULL COMMENT '扩展数据，如请求参数、返回摘要、错误信息等',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
    KEY `idx_operation_log_operator` (`operator_type`, `operator_id`),
    KEY `idx_operation_log_module` (`operation_module`),
    KEY `idx_operation_log_target` (`target_type`, `target_id`),
    KEY `idx_operation_log_result` (`operation_result`),
    KEY `idx_operation_log_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统操作日志表';
