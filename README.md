### 适合作为浏览器主页的自定义导航页面

使用 Go 语言重构的版本，功能和 UI 与原 Ruby on Rails 版本一致。

## 功能特点

- 用户注册/登录/注销
- 自定义导航分类和链接
- 拖拽排序
- 自动获取网站图标
- 多搜索引擎支持（百度、必应、谷歌）

## 技术栈

- **后端**: Go + Gin 框架
- **数据库**: SQLite
- **前端**: Vue.js 2 + Bootstrap 4 + Vuedraggable

## 运行

### 本地运行

```bash
go mod tidy
go build -o navi-home
./navi-home
```

访问 http://localhost:3000

### Docker 运行

```bash
docker-compose up -d
```

访问 http://localhost:9011

## 环境变量

- `PORT`: 服务端口，默认 3000
- `SECRET_KEY_BASE`: Session 密钥
- `GIN_MODE`: 运行模式（debug/release）

运行截图:
![demo](https://raw.githubusercontent.com/liubaicai/navi-home/master/demo.png)