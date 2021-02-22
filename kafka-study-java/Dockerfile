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

# 启动参数配置信息，做参考使用
# nohup java -server -Xms1024M -Xmx1536M -Xmn768M -Xss256k  -XX:+UseBiasedLocking -XX:+DisableExplicitGC
# -XX:MaxTenuringThreshold=15 -XX:+UseConcMarkSweepGC -XX:+UseParNewGC  -XX:+CMSParallelRemarkEnabled
# -XX:LargePageSizeInBytes=128m  -XX:+UseFastAccessorMethods -XX:+UseCMSInitiatingOccupancyOnly
# -Djava.awt.headless=true -Duser.timezone=US/Pacific -jar /kafka-study-java.jar >> logs/output-console.log 2>&1 &