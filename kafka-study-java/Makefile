
BUILD_TAG := $(shell date +'%Y%m%d%H%M')

build-maven:
	@echo maven cliean project ## ִ��maven ����
	mvn clean
	@echo maven package project # ִ��maven ���
	mvn package
	@echo remove old jar and copy new jar # ��ɾ�����ϵ�jar�����ڿ����µ�
	rm -f kafka-study-java.jar
	cp $(PWD)/target/kafka-study-java-0.0.1-SNAPSHOT.jar kafka-study-java.jar

build-docker:build-maven
	@echo package images, use default DockerFile # �������Ĭ��ʹ�õ�ǰDockerFile
	docker build . -t duxinglangzi/kafka-study-java:$(BUILD_TAG)
	@echo docker push image # �����ؾ������͵�Զ�ֿ̲�
	docker image push duxinglangzi/kafka-study-java:$(BUILD_TAG)
	@echo clean old jar and maven clean # ����ʹ�ù���jar  �Լ� maven
	rm -f kafka-study-java.jar
	mvn clean


build-docker-start-test:
	docker run -i --rm --name kafka-study-java -p 8090:8090 duxinglangzi/kafka-study-java:$(BUILD_TAG)