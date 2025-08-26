# Go Blog

基于 Gin + GORM + MySQL + JWT 的简易博客系统。

## 特性
- 用户注册、登录（JWT）
- 文章 CRUD（作者本人才能修改/删除）
- 评论创建、列表、删除（删除仅作者）
- 统一返回格式

## 目录结构
```
task4/
    config/
        config.yaml
    internal/
        config/
            config.go
        database/
            mysql.go
        handler/
            auth.go
            post.go
            comment.go
        middleware/
            auth.go
            requestLogger.go
        model/
            comment.go
            post.go
            user.go
        router/
            router.go
        utils/
            jwt.go
            response.go
    go.mod
    main.go
    README.md
```

## 环境要求
- Go 1.20+
- MySQL 8.x

## 快速开始
1. 创建数据库：
```sql
CREATE DATABASE blog DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

```

### 配置文件（必需）
- 位置：`task4/config/config.yaml`
- 配置项：
    - server: `port`
    - database: `host`/`port`/`user`/`password`/`name`/`charset`
    - jwt: `secret`、`expire`


### 快速启动
```
cd task4
# 修改 config.yaml 后执行：
go mod tidy
go run main.go

## 接口与测试用例
### 1) 注册
```http
POST http://localhost:8080/register
Content-Type: application/json

{
  "username": "zhangsan",
  "password": "123456",
  "email": "zhangsan@163.com"
}
```

### 2) 登录（获取 token）
```http
POST http://localhost:8080/login
Content-Type: application/json

{
  "username": "zhangsan",
  "password": "123456"
}
```
> 响应 `data.token` 复制到下面的 `{{token}}` 位置。

### 3) 创建文章（鉴权）
```http
POST http://localhost:8080/posts
Authorization: Bearer {{token}}
Content-Type: application/json

{
  "title": "第一篇文章",
  "content": "我的第一篇文章!"
}
```

### 4) 文章列表（公开）
```http
GET http://localhost:8080/posts
```

### 5) 文章详情（公开）
```http
GET http://localhost:8080/posts/1
```

### 6) 更新文章（作者本人）
```http
PUT http://localhost:8080/posts/1
Authorization: Bearer {{token}}
Content-Type: application/json

{
  "title": "更新后的标题",
  "content": "更新后的内容"
}
```

### 7) 删除文章（作者本人）
```http
DELETE http://localhost:8080/posts/1
Authorization: Bearer {{token}}
```

### 8) 发表评论（鉴权）
```http
POST http://localhost:8080/comments
Authorization: Bearer {{token}}
Content-Type: application/json

{
  "content": "写得真不错！",
  "post_id": 1
}
```

### 9) 获取文章评论列表（公开）
```http
GET http://localhost:8080/posts/1/comments
```

### 10) 删除评论（仅评论作者）
```http
DELETE http://localhost:8080/comments/1
Authorization: Bearer {{token}}
```