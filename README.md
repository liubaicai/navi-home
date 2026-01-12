<div align="center">

# 🏠 Navi-Home

**一个简洁、美观、可自定义的浏览器导航主页**

[![License](https://img.shields.io/badge/license-WTFPL-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-Available-2496ED?logo=docker)](Dockerfile)

[English](#english) | [中文](#中文)

</div>

---

## 📸 截图 / Screenshot

![Demo Screenshot](https://raw.githubusercontent.com/liubaicai/navi-home/master/demo.png)

---

<a name="中文"></a>
## 🌟 项目简介

Navi-Home 是一个适合作为浏览器主页的自定义导航页面，使用 Go 语言重构，提供简洁美观的界面和丰富的个性化功能。

### ✨ 功能特点

- 🔐 **用户系统** - 支持用户注册、登录、注销
- 📂 **分类管理** - 自定义导航分类和链接
- 🖱️ **拖拽排序** - 支持拖拽调整分类和链接顺序
- 🌐 **网站图标** - 自动获取网站 Favicon
- 🔍 **多搜索引擎** - 支持百度、必应、谷歌等搜索引擎
- 📱 **响应式设计** - 适配多种设备屏幕

### 🛠️ 技术栈

| 类型 | 技术 |
|------|------|
| **后端** | Go + [Gin](https://github.com/gin-gonic/gin) |
| **数据库** | SQLite |
| **前端** | Vue.js 2 + Bootstrap 4 + Vuedraggable |
| **容器化** | Docker + Docker Compose |

---

## 🚀 快速开始

### 方式一：Docker 运行（推荐）

```bash
# 克隆项目
git clone https://github.com/liubaicai/navi-home.git
cd navi-home

# 使用 Docker Compose 启动
docker-compose up -d
```

访问 http://localhost:9011

### 方式二：本地编译运行

**前置要求：**
- Go 1.21 或更高版本
- GCC（用于 SQLite CGO 编译）

```bash
# 克隆项目
git clone https://github.com/liubaicai/navi-home.git
cd navi-home

# 下载依赖
go mod tidy

# 编译
go build -o navi-home

# 运行
./navi-home
```

访问 http://localhost:3000

### 方式三：使用预编译二进制

从 [Releases](https://github.com/liubaicai/navi-home/releases) 页面下载对应平台的预编译版本：

- `navi-home-linux-amd64` - Linux x86_64
- `navi-home-linux-arm64` - Linux ARM64
- `navi-home-windows-amd64.exe` - Windows x86_64

---

## ⚙️ 配置说明

### 环境变量

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| `PORT` | 服务监听端口 | `3000` |
| `SECRET_KEY_BASE` | Session 加密密钥 | `default-secret-key-for-development` |
| `GIN_MODE` | 运行模式 (`debug` / `release`) | `debug` |
| `DATA_DIR` | 数据存储目录 | `data` |

### Docker Compose 配置示例

```yaml
services:
  web:
    build: .
    volumes:
      - ./data:/app/data          # 数据持久化
      - ./public/icons:/app/public/icons  # 图标缓存
    ports:
      - "9011:3000"
    restart: always
    environment:
      - SECRET_KEY_BASE=your-secret-key-here
      - GIN_MODE=release
      - PORT=3000
```

---

## 📁 项目结构

```
navi-home/
├── main.go              # 应用入口
├── handlers/            # HTTP 处理器
│   ├── auth.go          # 认证相关
│   ├── home.go          # 主页相关
│   └── session.go       # Session 处理
├── models/              # 数据模型
│   ├── db.go            # 数据库初始化
│   ├── user.go          # 用户模型
│   ├── catalog.go       # 分类模型
│   └── link.go          # 链接模型
├── templates/           # HTML 模板
├── static/              # 静态资源
├── public/              # 公共文件
├── data/                # 数据存储目录
├── Dockerfile           # Docker 镜像构建
├── docker-compose.yml   # Docker Compose 配置
├── go.mod               # Go 模块定义
└── go.sum               # 依赖校验
```

---

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 提交 Pull Request

---

## 📄 许可证

本项目采用 [WTFPL](LICENSE) 许可证 - 你可以随意使用。

---

<a name="english"></a>
## 🌟 Introduction

Navi-Home is a customizable browser homepage navigation page, rebuilt with Go, providing a clean and beautiful interface with rich personalization features.

### ✨ Features

- 🔐 **User System** - User registration, login, and logout
- 📂 **Category Management** - Customize navigation categories and links
- 🖱️ **Drag & Drop** - Drag to reorder categories and links
- 🌐 **Website Icons** - Automatic favicon fetching
- 🔍 **Multi-Search Engines** - Support for Baidu, Bing, Google, etc.
- 📱 **Responsive Design** - Works on various screen sizes

### 🛠️ Tech Stack

| Type | Technology |
|------|------------|
| **Backend** | Go + [Gin](https://github.com/gin-gonic/gin) |
| **Database** | SQLite |
| **Frontend** | Vue.js 2 + Bootstrap 4 + Vuedraggable |
| **Container** | Docker + Docker Compose |

---

## 🚀 Quick Start

### Option 1: Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/liubaicai/navi-home.git
cd navi-home

# Start with Docker Compose
docker-compose up -d
```

Visit http://localhost:9011

### Option 2: Build from Source

**Prerequisites:**
- Go 1.21 or higher
- GCC (for SQLite CGO compilation)

```bash
# Clone the repository
git clone https://github.com/liubaicai/navi-home.git
cd navi-home

# Download dependencies
go mod tidy

# Build
go build -o navi-home

# Run
./navi-home
```

Visit http://localhost:3000

### Option 3: Pre-built Binaries

Download pre-built binaries from the [Releases](https://github.com/liubaicai/navi-home/releases) page:

- `navi-home-linux-amd64` - Linux x86_64
- `navi-home-linux-arm64` - Linux ARM64
- `navi-home-windows-amd64.exe` - Windows x86_64

---

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server listening port | `3000` |
| `SECRET_KEY_BASE` | Session encryption key | `default-secret-key-for-development` |
| `GIN_MODE` | Run mode (`debug` / `release`) | `debug` |
| `DATA_DIR` | Data storage directory | `data` |

---

## 📄 License

This project is licensed under the [WTFPL](LICENSE) License - do whatever you want.