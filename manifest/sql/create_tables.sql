-- PetCare Database Tables Creation Script
-- Generated based on entity structures

USE pet_medical_system;

-- User table
CREATE TABLE `user` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户主键ID',
  `username` VARCHAR(50) NOT NULL COMMENT '登录用户名',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
  `nickname` VARCHAR(50) NOT NULL COMMENT '昵称',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
  `avatar_url` VARCHAR(255) DEFAULT NULL COMMENT '用户头像URL',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常，0禁用',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- Admin table
CREATE TABLE `admin` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '管理员主键ID',
  `username` VARCHAR(50) NOT NULL COMMENT '管理员账号',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
  `real_name` VARCHAR(50) NOT NULL COMMENT '管理员姓名',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '联系电话',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常，0禁用',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='管理员表';

-- Hospital table
CREATE TABLE `hospital` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '医院主键ID',
  `hospital_name` VARCHAR(100) NOT NULL COMMENT '医院名称',
  `address` VARCHAR(255) DEFAULT NULL COMMENT '医院地址',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '医院电话',
  `description` TEXT COMMENT '医院描述',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常，0停用',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医院表';

-- Doctor table
CREATE TABLE `doctor` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '医生主键ID',
  `hospital_id` BIGINT UNSIGNED NOT NULL COMMENT '所属医院ID',
  `username` VARCHAR(50) NOT NULL COMMENT '医生登录账号',
  `password_hash` VARCHAR(255) NOT NULL COMMENT '加密后的密码',
  `doctor_name` VARCHAR(50) NOT NULL COMMENT '医生姓名',
  `gender` TINYINT DEFAULT 0 COMMENT '性别：1男，2女，0未知',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '手机号',
  `email` VARCHAR(100) DEFAULT NULL COMMENT '邮箱',
  `title` VARCHAR(50) DEFAULT NULL COMMENT '职称',
  `specialty` VARCHAR(100) DEFAULT NULL COMMENT '擅长领域',
  `avatar_url` VARCHAR(255) DEFAULT NULL COMMENT '头像URL',
  `intro` TEXT COMMENT '医生简介',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1在职可接诊，0停用',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`),
  KEY `idx_hospital_id` (`hospital_id`),
  KEY `idx_status` (`status`),
  CONSTRAINT `fk_doctor_hospital` FOREIGN KEY (`hospital_id`) REFERENCES `hospital` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医生表';

-- Pet table
CREATE TABLE `pet` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '宠物主键ID',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '所属用户ID',
  `pet_name` VARCHAR(50) NOT NULL COMMENT '宠物名字',
  `pet_type` VARCHAR(20) NOT NULL DEFAULT 'cat' COMMENT '宠物类型，当前默认猫',
  `avatar_url` VARCHAR(255) DEFAULT NULL COMMENT '宠物头像URL',
  `gender` TINYINT DEFAULT 0 COMMENT '性别：1公，2母，0未知',
  `age` INT DEFAULT NULL COMMENT '年龄',
  `age_unit` VARCHAR(10) DEFAULT 'year' COMMENT '年龄单位：month/月，year/岁',
  `breed` VARCHAR(50) DEFAULT NULL COMMENT '品种',
  `weight` DECIMAL(5,2) DEFAULT NULL COMMENT '体重（kg）',
  `sterilized` TINYINT DEFAULT 0 COMMENT '是否绝育：1是，0否',
  `remark` TEXT COMMENT '备注',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1正常，0停用/删除',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`),
  CONSTRAINT `fk_pet_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物表';

-- Appointment table
CREATE TABLE `appointment` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '预约主键ID',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
  `doctor_id` BIGINT UNSIGNED NOT NULL COMMENT '医生ID',
  `pet_id` BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
  `appointment_time` DATETIME NOT NULL COMMENT '预约时间',
  `status` TINYINT NOT NULL DEFAULT 0 COMMENT '预约状态：0待确认，1已确认，2已完成，3已取消',
  `symptoms` TEXT COMMENT '症状描述',
  `notes` TEXT COMMENT '预约备注',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_doctor_id` (`doctor_id`),
  KEY `idx_pet_id` (`pet_id`),
  KEY `idx_appointment_time` (`appointment_time`),
  KEY `idx_status` (`status`),
  CONSTRAINT `fk_appointment_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_appointment_doctor` FOREIGN KEY (`doctor_id`) REFERENCES `doctor` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_appointment_pet` FOREIGN KEY (`pet_id`) REFERENCES `pet` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='预约表';

-- Medical Record table
CREATE TABLE `medical_record` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '医疗记录主键ID',
  `appointment_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '关联预约ID',
  `pet_id` BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
  `doctor_id` BIGINT UNSIGNED NOT NULL COMMENT '医生ID',
  `diagnosis` TEXT COMMENT '诊断结果',
  `treatment` TEXT COMMENT '治疗方案',
  `prescription` TEXT COMMENT '处方',
  `notes` TEXT COMMENT '备注',
  `record_date` DATETIME NOT NULL COMMENT '记录日期',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_appointment_id` (`appointment_id`),
  KEY `idx_pet_id` (`pet_id`),
  KEY `idx_doctor_id` (`doctor_id`),
  KEY `idx_record_date` (`record_date`),
  CONSTRAINT `fk_medical_record_appointment` FOREIGN KEY (`appointment_id`) REFERENCES `appointment` (`id`) ON DELETE SET NULL,
  CONSTRAINT `fk_medical_record_pet` FOREIGN KEY (`pet_id`) REFERENCES `pet` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_medical_record_doctor` FOREIGN KEY (`doctor_id`) REFERENCES `doctor` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医疗记录表';

-- Medical Report table
CREATE TABLE `medical_report` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '医疗报告主键ID',
  `medical_record_id` BIGINT UNSIGNED NOT NULL COMMENT '医疗记录ID',
  `report_type` VARCHAR(50) NOT NULL COMMENT '报告类型',
  `report_content` TEXT COMMENT '报告内容',
  `file_url` VARCHAR(255) DEFAULT NULL COMMENT '报告文件URL',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_medical_record_id` (`medical_record_id`),
  CONSTRAINT `fk_medical_report_record` FOREIGN KEY (`medical_record_id`) REFERENCES `medical_record` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='医疗报告表';

-- Notification table
CREATE TABLE `notification` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '通知主键ID',
  `user_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '接收用户ID',
  `doctor_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '接收医生ID',
  `admin_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '接收管理员ID',
  `title` VARCHAR(100) NOT NULL COMMENT '通知标题',
  `content` TEXT NOT NULL COMMENT '通知内容',
  `type` VARCHAR(20) NOT NULL COMMENT '通知类型',
  `is_read` TINYINT NOT NULL DEFAULT 0 COMMENT '是否已读：1已读，0未读',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_doctor_id` (`doctor_id`),
  KEY `idx_admin_id` (`admin_id`),
  KEY `idx_type` (`type`),
  KEY `idx_is_read` (`is_read`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='通知表';

-- Operation Log table
CREATE TABLE `operation_log` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '操作日志主键ID',
  `user_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '操作用户ID',
  `doctor_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '操作医生ID',
  `admin_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '操作管理员ID',
  `operation_type` VARCHAR(50) NOT NULL COMMENT '操作类型',
  `operation_desc` VARCHAR(255) NOT NULL COMMENT '操作描述',
  `ip_address` VARCHAR(45) DEFAULT NULL COMMENT 'IP地址',
  `user_agent` VARCHAR(255) DEFAULT NULL COMMENT '用户代理',
  `created_at` DATETIME NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_doctor_id` (`doctor_id`),
  KEY `idx_admin_id` (`admin_id`),
  KEY `idx_operation_type` (`operation_type`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志表';

-- AI Session table
CREATE TABLE `ai_session` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'AI会话主键ID',
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
  `session_title` VARCHAR(100) DEFAULT NULL COMMENT '会话标题',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态：1活跃，0结束',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='AI会话表';

-- AI Message table
CREATE TABLE `ai_message` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'AI消息主键ID',
  `session_id` BIGINT UNSIGNED NOT NULL COMMENT '会话ID',
  `role` VARCHAR(20) NOT NULL COMMENT '消息角色：user/assistant',
  `content` TEXT NOT NULL COMMENT '消息内容',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_session_id` (`session_id`),
  CONSTRAINT `fk_ai_message_session` FOREIGN KEY (`session_id`) REFERENCES `ai_session` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='AI消息表';

-- AI Analysis Record table
CREATE TABLE `ai_analysis_record` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'AI分析记录主键ID',
  `pet_id` BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
  `analysis_type` VARCHAR(50) NOT NULL COMMENT '分析类型',
  `analysis_result` TEXT COMMENT '分析结果',
  `confidence` DECIMAL(3,2) DEFAULT NULL COMMENT '置信度',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_pet_id` (`pet_id`),
  KEY `idx_analysis_type` (`analysis_type`),
  CONSTRAINT `fk_ai_analysis_pet` FOREIGN KEY (`pet_id`) REFERENCES `pet` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='AI分析记录表';

-- Pet Allergy Record table
CREATE TABLE `pet_allergy_record` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '过敏记录主键ID',
  `pet_id` BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
  `allergen` VARCHAR(100) NOT NULL COMMENT '过敏原',
  `severity` VARCHAR(20) DEFAULT NULL COMMENT '严重程度',
  `symptoms` TEXT COMMENT '症状描述',
  `diagnosis_date` DATE DEFAULT NULL COMMENT '诊断日期',
  `notes` TEXT COMMENT '备注',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_pet_id` (`pet_id`),
  CONSTRAINT `fk_pet_allergy_pet` FOREIGN KEY (`pet_id`) REFERENCES `pet` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物过敏记录表';

-- Pet Medical History table
CREATE TABLE `pet_medical_history` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '病史记录ID',
  `pet_id` BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
  `history_type` VARCHAR(100) NOT NULL COMMENT '病史类型',
  `description` TEXT NOT NULL COMMENT '病史描述',
  `diagnosed_at` DATETIME NOT NULL COMMENT '确诊时间',
  `is_current` TINYINT NOT NULL DEFAULT 0 COMMENT '是否当前仍存在：1是，0否',
  `treatment` TEXT COMMENT '治疗方案',
  `outcome` VARCHAR(50) DEFAULT NULL COMMENT '治疗结果',
  `doctor_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '诊断医生ID',
  `notes` TEXT COMMENT '备注',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_pet_id` (`pet_id`),
  KEY `idx_doctor_id` (`doctor_id`),
  CONSTRAINT `fk_pet_history_pet` FOREIGN KEY (`pet_id`) REFERENCES `pet` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_pet_history_doctor` FOREIGN KEY (`doctor_id`) REFERENCES `doctor` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物医疗历史表';

-- Pet Vaccination Record table
CREATE TABLE `pet_vaccination_record` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '疫苗记录主键ID',
  `pet_id` BIGINT UNSIGNED NOT NULL COMMENT '宠物ID',
  `vaccine_name` VARCHAR(100) NOT NULL COMMENT '疫苗名称',
  `vaccination_date` DATE NOT NULL COMMENT '接种日期',
  `next_due_date` DATE DEFAULT NULL COMMENT '下次接种日期',
  `batch_number` VARCHAR(50) DEFAULT NULL COMMENT '疫苗批号',
  `doctor_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '接种医生ID',
  `notes` TEXT COMMENT '备注',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_pet_id` (`pet_id`),
  KEY `idx_doctor_id` (`doctor_id`),
  KEY `idx_vaccination_date` (`vaccination_date`),
  KEY `idx_next_due_date` (`next_due_date`),
  CONSTRAINT `fk_pet_vaccination_pet` FOREIGN KEY (`pet_id`) REFERENCES `pet` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_pet_vaccination_doctor` FOREIGN KEY (`doctor_id`) REFERENCES `doctor` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='宠物疫苗记录表';