# 用户登录注册系统

这是一个基于Go后端和Vue3前端的用户登录注册系统，提供用户注册、登录功能，使用JWT进行身份验证。

### 1. 数据库准备

1. 安装并启动MySQL服务
2. 创建数据库：
   ```sql
   CREATE DATABASE user CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
   ```
3. 导入数据库表结构：
   ```bash
   mysql -u root -p user < user.sql
   ```

### 2. 后端启动

1. 进入项目根目录：
   ```bash
   cd loginRegister
   ```

2. 安装Go依赖：
   ```bash
   go mod tidy
   ```

3. 修改数据库配置（如需要）：
   编辑 `config/config.yaml` 文件，修改数据库连接信息：
   ```yaml
   database:
     dsn: root:123456@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local
   ```

4. 启动后端服务：
   ```bash
   go run main.go
   ```

   后端服务将在 `http://localhost:8080` 启动

### 3. 前端启动

1. 进入前端目录：
   ```bash
   cd fronted
   ```

2. 安装依赖：
   ```bash
   npm install
   ```

3. 启动开发服务器：
   ```bash
   npm run dev
   ```

   前端服务将在 `http://localhost:5173` 启动

### 4. 访问应用

- 前端页面：http://localhost:5173
- 后端API：http://localhost:8080

## 

## 前端展示

### 登录界面：

![登录界面](fronted/assets/loginr.png)

### 注册界面：

![注册界面](fronted/assets/registerr.png)

### 登陆后跳转

![登录后跳转](fronted/assets/dashboardr.png)

## 接口文档（用户登录、注册）

### （1） 用户注册

- **接口说明**: 创建新用户并返回登录凭证（JWT）。
- **请求方法**: `POST`
- **请求URL**: `/api/auth/register`

#### 请求参数（JSON）

- **username**: string，必填，唯一
- **password**: string，必填
- **email**: string，必填，唯一
- **phone**: string，必填，唯一
- ps：数据格式后端未作校验

示例：

```json
{
  "username": "testuser",
  "password": "password123",
  "email": "2983764528@qq.com",
  "phone": "13765273445"
}
```

#### 返回参数

- 成功：`200 OK`

```json
{
    "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTgyOTI4NjAsInVzZXJuYW1lIjoidGVzdXNlciJ9.roL4j_qM_FwbmFiWWdt_Ro0U0gb70p6WqVcal2D3MKk"
}
```

- 失败：
  - `400 Bad Request`（参数绑定失败（json格式有误）/用户名已存在/手机号已存在/邮箱已存在）
  - `500 Internal Server Error`（密码加密失败/JWT生成失败/数据库写入失败）
  - 注：重复字段验证顺序为 用户名->手机号->邮箱

```json
{"error": "用户名已存在"}
```

---

### （2） 用户登录

- **接口说明**: 使用用户名与密码登录，返回登录凭证（JWT）。
- **请求方法**: `POST`
- **请求URL**: `/api/auth/login`

#### 请求参数（ JSON）

- **username**: string，必填
- **password**: string，必填

示例：

```json
{
  "username": "testuser",
  "password": "password123"
}
```

#### 返回参数

- 成功：`200 OK`

```json
{
    "token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTgyOTI5MzIsInVzZXJuYW1lIjoidGVzdHVzZXIifQ.Y-A9G_mZVcsegt5jrXLv7h2KfATC8u_3Nj-h0vhJKA8"
}
```

- 失败：
  - `400 Bad Request`（参数绑定失败）
  - `401 Unauthorized`（用户名或密码错误）
  - `500 Internal Server Error`（JWT生成异常）

```json
{"error": "用户名或密码错误"}
```

---

### user表字段与约束

<img width="735" height="243" alt="82ed7ddf-5d96-48b8-8a15-8a8e3cc92028" src="https://github.com/user-attachments/assets/f5e3bab5-8c40-4afb-a5f5-5238a1b4d90e" />


---

### 部署信息

- **基础URL及服务端口**: 默认`http://localhost:8080`可通过 `config/config.yaml` 的 `app` 配置）
- **CORS**: 使用浏览器联调时，允许域名见 `config/config.yaml` → `cors.AllowOrigins`（默认为 `http://localhost:5173`）
- **数据库**: MySQL（连接信息见 `config/config.yaml` → `database.dsn`，修改`name:password`为用户名、密码），运行user.sql初始化表



