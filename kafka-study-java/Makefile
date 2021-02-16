
BUILD_TAG := $(shell date +'%Y%m%d%H%M')

build-maven:
	@echo maven cliean project ## 执行maven 清理
	mvn clean
	@echo maven package project # 执行maven 打包
	mvn package
	@echo remove old jar and copy new jar # 先删除掉老的jar包，在拷贝新的
	rm -f kafka-study-java.jar
	cp $(PWD)/target/kafka-study-java-0.1.jar kafka-study-java.jar

build-docker:build-maven
	@echo package images, use default DockerFile # 打包镜像，默认使用当前DockerFile
	docker build . -t duxinglangzi/kafka-study-java:$(BUILD_TAG)
	@echo docker push image # 将本地镜像推送到远程仓库
	docker image push duxinglangzi/kafka-study-java:$(BUILD_TAG)
	@echo clean old jar and maven clean # 清理使用过的jar  以及 maven
	rm -f kafka-study-java.jar
	mvn clean


build-docker-start-test:
	docker run -i --rm --name kafka-study-java -p 8090:8090 duxinglangzi/kafka-study-java:$(BUILD_TAG)

build-test:
	@echo 'neira 好 好好'