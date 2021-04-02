GO=go
GOFLAGS ?= $(GOFLAGS:) -mod=vendor
BUILD_TAG := $(shell date +'%Y%m%d%H%M')
export GOBIN = $(PWD)/kafka-study-go
build-linux: ## 编译
	@echo 编译安装环境
	env GOOS=linux GOARCH=amd64 $(GO) build -o $(GOBIN) $(GOFLAGS) com/conor/kafka/main

build-macOS: ## 编译
	@echo 编译MacOs安装环境
	$(GO) build -o $(GOBIN) $(GOFLAGS) com/conor/kafka/main

build-docker: build-linux ## 打包镜像
	@echo 打包镜像且默认使用当前目录的DockerFile文件
	docker build . -t duxinglangzi/kafka-study-go:$(BUILD_TAG)
	@echo "推送到docker仓库"
	docker image push duxinglangzi/kafka-study-go:$(BUILD_TAG)
	@echo 清理文件
	rm -f kafka-study-go

build-test: ## 本地测试
	docker run -i --rm --name kafka-study-go -p 8091:8091 duxinglangzi/kafka-study-go:202102011657

dev:
	@echo "Start local service ... "
	$(GO) run $(GOFLAGS) com/conor/kafka/main