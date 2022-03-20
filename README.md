## 绿帽子工具 （希望广大程序员能解放双手。当个勤劳的小农民，😁）
```bazaar
├── README.md
├── config
│   ├── config.go   基础配置
│   ├── push.go     push日志过滤清洗脚本
│   └── switch.go   选择执行过滤脚本
├── go.mod
├── go.sum
├── main.go
├── manager
│   └── logfilter.go 基础文件过模块
├── model
│   └── model.go    
└── utils
    ├── cpu.go
    └── file.go   文件操作工具
```

### 新增过滤模块 请在config。switch 中增加选项
###     在config 文件夹 增加对应的过滤 流程即可

本地编译 linux   
终端执行
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build   