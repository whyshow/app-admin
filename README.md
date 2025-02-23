# app-admin

Go语言编写的后台管理及API处理项目，使用Gin框架搭建

## 项目目录结构

```text
app-admin/
├── app/
│   ├── common/          # 公共组件
│   ├── config/          # 配置模块
│   ├── modules/         # 业务模块
│   │   ├── user/        # 用户模块
│   │   │   ├── router.go
│   │   │   ├── api.go
│   │   │   └── model.go
│   │   ├── order/       # 订单模块
│   │   └── ...
└── main.go              # 入口文件