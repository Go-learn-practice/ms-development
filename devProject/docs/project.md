# Project 首页数据库表设计

**菜单表 menus**

- 设置 `SET NAMES utf8mb4;` 作用是设置数据库编码为 utf8mb4，以支持存储表情符号等特殊字符
- 设置 `SET FOREIGN_KEY_CHECKS = 0;` 作用是禁用外键约束检查，以避免在创建表时出现错误

```sql
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
    
DROP TABLE IF EXISTS `ms_project_menus`;
CREATE TABLE `ms_project_menus` (
    `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '菜单id',
    `pid` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父级id',
    `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '名称',
    `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
    `url` varchar(400) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单链接',
    `file_path` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '文件路径',
    `params` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '连接参数',
    `node` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '#' COMMENT '权限节点',
    `sort` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '菜单排序',
    `status` tinyint(1) UNSIGNED NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
    `create_by` bigint(0) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建人',
    `is_inner` tinyint(1) NULL DEFAULLT 0 COMMENT '是否内页',
    `values` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '参数默认值',
    `show_slider` tinyint(1) NULL DEFAULT 1 COMMENT '是否显示侧边栏',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目菜单表' ROW_FORMAT = DYNAMIC;
```

**添加数据项**

```sql
INSERT INTO `ms_project_menus` VALUES (120, 0, '工作台', 'appstore-o', 'home', 'home', ':org', '#', 0, 1, 0, 0, '', 0);
INSERT INTO `ms_project_menus` VALUES (121, 0, '项目管理', 'pro', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (122, 121, '项目列表', 'branches', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (124, 0, '系统设置', 'setting', '#', '#', '', '#', 100, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (125, 124, '成员管理', 'unlock', '#', '#', '', '#', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (126, 125, '账号列表', '', 'system/account', 'system/account', '', 'pro/account/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (127, 122, '我的组织', '', 'organization', 'organization', '', 'pro/organization/index', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (130, 125, '访问授权', '', 'system/account/auth', 'system/account/auth', '', 'pro/auth/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (131, 125, '授权页面', '', 'system/account/apply', 'system/account/apply', ':id', 'pro/auth/apply', 30, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menus` VALUES (138, 121, '消息提醒', 'info-circle-o', '#', '#', '', '#', 30, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (139, 138, '站内消息', '', 'notify/notice', 'notify/notice', '', 'pro/notify/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (140, 138, '系统公告', '', 'notify/system', 'notify/system', '', 'pro/notify/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (143, 124, '系统管理', 'appstore', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (144, 143, '菜单路由', '', 'system/config/menu', 'system/config/menu', '', 'pro/menu/menuadd', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (145, 143, '访问节点', '', 'system/config/node', 'system/config/node', '', 'pro/node/save', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (148, 124, '个人管理', 'user', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (149, 148, '个人设置', '', 'account/setting/base', 'account/setting/base', '', 'pro/index/editpersonal', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (150, 148, '安全设置', '', 'account/setting/security', 'account/setting/security', '', 'pro/index/editpersonal', 0, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menus` VALUES (151, 122, '我的项目', '', 'pro/list', 'pro/list', ':type', 'pro/pro/index', 0, 1, 0, 0, 'my', 1);
INSERT INTO `ms_project_menus` VALUES (152, 122, '回收站', '', 'pro/recycle', 'pro/recycle', '', 'pro/pro/index', 20, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (153, 121, '项目空间', 'heat-map', 'pro/space/task', 'pro/space/task', ':code', '#', 20, 1, 0, 1, '', 1);
INSERT INTO `ms_project_menus` VALUES (154, 153, '任务详情', '', 'pro/space/task/:code/detail', 'pro/space/taskdetail', ':code', 'pro/task/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (155, 122, '我的收藏', '', 'pro/list', 'pro/list', ':type', 'pro/pro/index', 10, 1, 0, 0, 'collect', 1);
INSERT INTO `ms_project_menus` VALUES (156, 121, '基础设置', 'experiment', '#', '#', '', '#', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (157, 156, '项目模板', '', 'pro/template', 'pro/template', '', 'pro/project_template/index', 0, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (158, 156, '项目列表模板', '', 'pro/template/taskStages', 'pro/template/taskStages', ':code', 'pro/task_stages_template/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (159, 122, '已归档项目', '', 'pro/archive', 'pro/archive', '', 'pro/pro/index', 10, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (160, 0, '团队成员', 'team', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (161, 153, '项目概况', '', 'pro/space/overview', 'pro/space/overview', ':code', 'pro/index/info', 20, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (162, 153, '项目文件', '', 'pro/space/files', 'pro/space/files', ':code', 'pro/index/info', 10, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (163, 122, '项目分析', '', 'pro/analysis', 'pro/analysis', '', 'pro/index/info', 5, 1, 0, 0, '', 1);
INSERT INTO `ms_project_menus` VALUES (164, 160, '团队成员', '', '#', '#', '', '#', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (166, 164, '团队成员', '', 'members', 'members', '', 'pro/department/index', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (167, 164, '成员信息', '', 'members/profile', 'members/profile', ':code', 'pro/department/read', 0, 1, 0, 1, '', 0);
INSERT INTO `ms_project_menus` VALUES (168, 153, '版本管理', '', 'pro/space/features', 'pro/space/features', ':code', 'pro/index/info', 20, 1, 0, 1, '', 0);
```

```sql
SET FOREIGN_KEY_CHECKS = 1;
```

** 项目表 Project 表**
```sql
CREATE TABLE `ms_projects`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `cover` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '封面',
  `name` varchar(90) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '名称',
  `description` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '描述',
  `access_control_type` tinyint(0) NULL DEFAULT 0 COMMENT '访问控制l类型',
  `white_list` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '可以访问项目的权限组（白名单）',
  `order` int(0) UNSIGNED NULL DEFAULT 0 COMMENT '排序',
  `deleted` tinyint(1) NULL DEFAULT 0 COMMENT '删除标记',
  `template_code` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '项目类型',
  `schedule` double(5, 2) NULL DEFAULT 0.00 COMMENT '进度',
  `create_time` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '创建时间',
  `organization_code` bigint(0) NULL DEFAULT NULL COMMENT '组织id',
  `deleted_time` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '删除时间',
  `private` tinyint(1) NULL DEFAULT 1 COMMENT '是否私有',
  `prefix` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '项目前缀',
  `open_prefix` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启项目前缀',
  `archive` tinyint(1) NULL DEFAULT 0 COMMENT '是否归档',
  `archive_time` bigint(0) NULL DEFAULT NULL COMMENT '归档时间',
  `open_begin_time` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启任务开始时间',
  `open_task_private` tinyint(1) NULL DEFAULT 0 COMMENT '是否开启新任务默认开启隐私模式',
  `task_board_theme` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'default' COMMENT '看板风格',
  `begin_time` bigint(0) NULL DEFAULT NULL COMMENT '项目开始日期',
  `end_time` bigint(0) NULL DEFAULT NULL COMMENT '项目截止日期',
  `auto_update_schedule` tinyint(1) NULL DEFAULT 0 COMMENT '自动更新项目进度',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `project`(`order`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13043 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目表' ROW_FORMAT = COMPACT;
```

**项目成员表**
```sql
CREATE TABLE `ms_project_members`  (
  `id` bigint(0) NOT NULL AUTO_INCREMENT,
  `project_code` bigint(0) NULL DEFAULT NULL COMMENT '项目id',
  `member_code` bigint(0) NULL DEFAULT NULL COMMENT '成员id',
  `join_time` bigint(0) NULL DEFAULT NULL COMMENT '加入时间',
  `is_owner` bigint(0) NULL DEFAULT 0 COMMENT '拥有者',
  `authorize` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique`(`project_code`, `member_code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 37 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '项目-成员表' ROW_FORMAT = COMPACT;
```