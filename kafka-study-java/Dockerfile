# 基础镜像
FROM java:8
# 作者
MAINTAINER conor
# 把可执行jar包复制到基础镜像的根目录下
ADD kafka-study-java.jar /kafka-study-java.jar
# 镜像要暴露的端口
EXPOSE 8090
# 在镜像运行为容器后执行的命令
ENTRYPOINT ["java","-jar","/kafka-study-java.jar"]