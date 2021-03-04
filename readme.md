### Skeleton

```bash
# Export example config
go run main.go config init -o config.example.yaml

# Generate
go generate -x ./...

# Build
go build -o skeleton .

# Format
go fmt -x ./...
```

> 目录结构

```
├── app                     # 主目录
│   ├── commands            # 命令目录: 存放所有的命令定义
│   │   └── version.go      # 显示版本命令 (示例命令)
│   └── module.go           # 模块核心配置, 包括依赖注入的注册等
├── bootstrap               # 启动目录: 项目与 Core 框架核心关联
│   ├── bootstrap.go        # 项目与 Core 的核心关联
│   └── root.go             # 项目根命令
├── config                  # 全局配置目录: 全局模块的注册和管理
│   └── app.go              # 配置模块加载与命令加载等
├── .gitignore              # Git 忽略文件
├── config.example.yaml     # 示例配置文件
├── config.yaml             # 配置文件
├── go.mod                  # go module
├── go.sum                  # go module
├── main.go                 # 项目入口
└── readme.md               # Readme
```
