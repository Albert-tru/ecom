# 🛒 Ecom - 电商后端 API

一个基于 Go 语言开发的电商后端系统，支持用户认证、产品管理、购物车和订单功能。

## 📚 技术栈

- **语言**: Go 1.20+
- **Web 框架**: gorilla/mux
- **数据库**: MySQL 8.0
- **认证**: JWT (golang-jwt/jwt)
- **密码加密**: bcrypt
- **数据验证**: go-playground/validator
- **数据库迁移**: golang-migrate

## 🚀 功能特性

### ✅ 已实现功能

- **用户管理**
  - 用户注册
  - 用户登录
  - JWT 令牌认证
  - 密码加密存储

- **产品管理**
  - 获取产品列表
  - 根据 ID 批量查询产品

- **购物车 & 订单**
  - 购物车结账
  - 创建订单
  - 创建订单项
  - 库存检查
  - 价格计算

### 🚧 待开发功能

- [ ] 产品 CRUD（创建/更新/删除）
- [ ] 订单查询
- [ ] 用户个人信息管理
- [ ] 库存自动更新
- [ ] 订单状态管理
- [ ] 支付集成

## 📁 项目结构

```
ecom/
├── cmd/
│   ├── main.go              # 应用入口
│   ├── api/
│   │   └── api.go          # API 服务器
│   └── migrate/
│       ├── main.go         # 数据库迁移入口
│       └── migrations/     # 迁移文件
├── config/
│   └── env.go              # 环境配置
├── db/
│   └── db.go               # 数据库连接
├── service/
│   ├── auth/               # 认证服务
│   │   ├── jwt.go         # JWT 实现
│   │   └── password.go    # 密码哈希
│   ├── user/               # 用户服务
│   │   ├── routes.go      # 用户路由
│   │   └── store.go       # 用户数据层
│   ├── product/            # 产品服务
│   │   ├── routes.go      # 产品路由
│   │   └── store.go       # 产品数据层
│   ├── cart/               # 购物车服务
│   │   ├── routes.go      # 购物车路由
│   │   └── service.go     # 购物车业务逻辑
│   └── order/              # 订单服务
│       └── store.go        # 订单数据层
├── types/
│   └── types.go            # 数据类型定义
├── utils/
│   └── utils.go            # 工具函数
├── .env                    # 环境变量
├── Makefile                # 构建脚本
└── *.http                  # REST Client 测试文件
```

## �� 安装与运行

### 1. 克隆项目

```bash
git clone <your-repo-url>
cd ecom
```

### 2. 配置环境变量

创建 `.env` 文件：

```env
# 数据库配置
DB_USER=ecomuser
DB_PASSWORD=Ecom123.
DB_HOST=localhost
DB_PORT=3306
DB_NAME=ecom
DB_NET=tcp

# JWT 配置
JWT_SECRET=your_jwt_secret_key
JWT_EXP=604800

# 服务器配置
PUBLIC_HOST=http://localhost
PORT=8080
```

### 3. 创建数据库

```bash
mysql -u root -p
```

```sql
CREATE DATABASE ecom;
CREATE USER 'ecomuser'@'localhost' IDENTIFIED BY 'Ecom123.';
GRANT ALL PRIVILEGES ON ecom.* TO 'ecomuser'@'localhost';
FLUSH PRIVILEGES;
```

### 4. 执行数据库迁移

```bash
make migrate-up
```

### 5. 运行服务器

```bash
make run
```

服务器将在 `http://localhost:8080` 启动。

## 📝 API 文档

### 用户认证

#### 注册用户

```http
POST /api/v1/register
Content-Type: application/json

{
  "firstname": "John",
  "lastname": "Doe",
  "email": "john@example.com",
  "password": "123456"
}
```

#### 用户登录

```http
POST /api/v1/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "123456"
}
```

**响应：**
```json
{
  "message": "login successful",
  "user_id": "1",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 产品管理

#### 获取产品列表

```http
GET /api/v1/products
```

### 购物车 & 订单

#### 购物车结账

```http
POST /api/v1/cart/checkout
Content-Type: application/json
Authorization: Bearer <your_token>

{
  "items": [
    {
      "productId": 1,
      "quantity": 2
    },
    {
      "productId": 2,
      "quantity": 1
    }
  ]
}
```

**响应：**
```json
{
  "status": "success",
  "orderId": 1,
  "totalPrice": 2299.97
}
```

## 🧪 测试

项目包含 REST Client 测试文件，可在 VS Code 中使用 REST Client 扩展进行测试：

- `log-api-test.http` - 用户认证测试
- `product-api-test.http` - 产品 API 测试
- `cart-api-test.http` - 购物车 API 测试

## 🛠️ 开发命令

```bash
# 编译项目
make build

# 运行项目
make run

# 运行测试
make test

# 创建迁移文件
make migration name=<migration_name>

# 执行迁移（向上）
make migrate-up

# 回滚迁移（向下）
make migrate-down
```

## 📊 数据库表结构

### users 表
- id (主键)
- firstname
- lastname
- email (唯一)
- password (bcrypt 哈希)
- createdat

### products 表
- id (主键)
- name
- description
- image
- price
- quantity
- createdat

### orders 表
- id (主键)
- user_id (外键)
- total
- status
- address
- created_at

### order_items 表
- id (主键)
- order_id (外键)
- product_id (外键)
- quantity
- price

## 📖 学习笔记

这个项目实践了以下 Go 语言开发技能：

- ✅ RESTful API 设计
- ✅ JWT 认证实现
- ✅ 数据库操作 (database/sql)
- ✅ 中间件模式
- ✅ MVC 架构
- ✅ 错误处理
- ✅ 数据验证
- ✅ 测试驱动开发（TDD）

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

MIT License

## 👨‍�� 作者

Albert - 正在学习 Go 语言开发

---

**⭐ 如果这个项目对你有帮助，请给个 Star！**
