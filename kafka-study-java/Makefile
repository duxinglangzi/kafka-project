BUILD_TAG := $(shell date +'%Y%m%d%H%M')

build-maven: ## maven 清理、打包、拷贝
	@echo maven cliean project 执行maven 清理
	mvn clean
	@echo maven package project 执行maven 打包
	mvn package
	@echo remove old jar and copy new jar  先删除掉老的jar包,再拷贝新的
	rm -f kafka-study-java.jar
	cp $(PWD)/target/kafka-study-java-0.1.jar kafka-study-java.jar

build-docker:build-maven ## 生成docker image 并推送至远程仓库
	@echo package images, use default DockerFile 打包镜像默认使用当前DockerFile
	docker build . -t duxinglangzi/kafka-study-java:$(BUILD_TAG)
	@echo docker push image 将本地镜像推送到远程仓库
	docker image push duxinglangzi/kafka-study-java:$(BUILD_TAG)
	@echo clean old jar and maven clean 清理使用过的jar  以及 maven
	rm -f kafka-study-java.jar
	mvn clean

stop-server: ## 停止服务
	@echo stop server 停止服务 ## awk 命令简单教程 https://www.runoob.com/linux/linux-comm-awk.html
	@for PID in $$(ps -ef | grep "java" | grep "kafka-study-java" | awk '{print $$2}'); do \
  		echo $$PID; \
  		kill -9 $$PID; \
  		echo 已停止PID: $$PID; \
	done

build-docker-start-test:
	docker run -i --rm --name kafka-study-java -p 8090:8090 duxinglangzi/kafka-study-java:$(BUILD_TAG)

help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' ./Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
