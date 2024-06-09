-- 设备详情表 t_equipment_detail
CREATE TABLE `t_equipment_detail` (
`id` INT NOT NULL AUTO_INCREMENT,
`class_id` INT NOT NULL COMMENT '分类id',
`name` VARCHAR(255) NOT NULL COMMENT '设备名称',
`type` VARCHAR(50) NOT NULL COMMENT '设备类型',
`specific_model` VARCHAR(255) COMMENT '具体型号',
`unit` VARCHAR(50) COMMENT '单位',
`status` VARCHAR(50) COMMENT '状态',
`image` varchar(512) COMMENT '图片',
`create_user_id` INT COMMENT '创建用户id',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备详情表';

-- 设备分类表 t_equipment_class
CREATE TABLE `t_equipment_class` (
`id` INT NOT NULL AUTO_INCREMENT,
`name` VARCHAR(255) NOT NULL COMMENT '名称',
`code` VARCHAR(100) NOT NULL COMMENT '代码',
`status` TINYINT NOT NULL COMMENT '状态',
`p_id` INT DEFAULT NULL COMMENT '父类id',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='设备分类表';

-- 杆塔设备表 t_tower_equipment
CREATE TABLE `t_tower_equipment` (
`id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
`tower_id` INT UNSIGNED NOT NULL COMMENT '杆塔ID',
`equipment_id` INT UNSIGNED NOT NULL COMMENT '设备管理ID',
`principal_id` INT UNSIGNED NOT NULL COMMENT '负责人ID',
`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='杆塔设备关联表';


-- 杆塔详情表 t_tower_detail
CREATE TABLE `t_tower_detail` (
`id` INT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
`subitem_id` INT NOT NULL COMMENT '子项ID',
`name` VARCHAR(255) NOT NULL COMMENT '基站名称',
`address` VARCHAR(255) NOT NULL COMMENT '详细地区',
`image` TEXT COMMENT '基站图片，考虑是否支持展示多张',
`check_status` VARCHAR(50) COMMENT '检查状态',
`check_time` DATETIME COMMENT '检查时间',
`check_user_id` INT COMMENT '检查人ID',
`principal_id` INT COMMENT '负责人ID',
`plan_time` DATETIME COMMENT '计划检查时间',
`create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
`update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='杆塔详情表';


