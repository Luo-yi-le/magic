# 安装依赖
go mod tidy

# 进入server文件夹
cd server

# 本地运行项目
go run main.go

# 编译打包项目
# 打包amd版linux可执行文件
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o magic-linux-amd64 main.go
# 打包arm版linux可执行文件
CGO_ENABLE=0 GOOS=linux GOARCH=arm64 go build -o magic-linux-arm64 main.go
# 打包windows版可执行文件
CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build -o magic-win main.go

go build  -o magic-win main.go  CGO_ENABLE=0 GOOS=windows GOARCH=amd64
