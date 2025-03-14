# Mysql

## 一、数据类型
> MySQL 支持多种数据类型，主要分为以下几类：

- **数值类型**
  - 整数类型
    - tinyint  `1 字节`
    - smallint `2 字节`
    - mediumint `3 字节`
    - int `4 字节`
    - bigint `8 字节`
  - 浮动小数类型
    - float `4 字节`
    - double `8 字节`
    - decimal `可变长度`

- **字符串类型**
  - char `固定长度`
  - varchar `可变长度 最大 65,535 字符`
  - text `可变长度 存储文本数据 最大 65,535 字符`
  - blob `可变长度 存储二进制数据 最大 65,535 字节`
  - mediumtext `可变长度 存储文本数据 最大 16,777,215 字符`
  - mediumblob `可变长度 存储二进制数据 最大 16,777,215 字节`
  - longtext `可变长度 存储文本数据 最大 4,294,967,295 字符`

- **日期和时间类型**
  - date `日期 格式：YYYY-MM-DD`
  - datetime `日期和时间 格式：YYYY-MM-DD HH:MM:SS`
  - timestamp `日期和时间 存储时间戳 通常用于记录数据的创建或更新时间 格式：YYYY-MM-DD HH:MM:SS`
  - time `时间 格式：HH:MM:SS`
  - year `年份 格式：YYYY`

- **枚举类型**
  - enum `枚举类型 可以存储一组预定义的值 例如：enum('男','女','未知')`

- **布尔类型**
  - bool或boolean `布尔类型 0 为禁用，1 为启用`

## 二、约束
> 约束用于限制表中列的数据，并确保数据的完整性和有效性。常见的约束包括

- **主键约束**
  - primary key `主键约束 用于唯一标识表中的每一行`

- **自动递增约束**
  - auto_increment `自动递增约束 用于自动生成唯一的整数值`

- **默认约束**
  - default `默认约束 用于在插入新行时为列提供默认值`

- **非空约束**
  - not null `非空约束 用于确保列中的值不为空`

- **唯一约束**
  - unique `唯一约束 用于确保列中的值是唯一的`

- **外键约束**
  - foreign key `外键约束 用于确保列中的值在另一个表中存在`

- **检查约束**
  - check `检查约束 用于确保列中的值满足特定的条件`

- **索引**
  - index `索引 用于提高查询性能`

## 三、