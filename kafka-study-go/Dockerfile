FROM alpine:3.12.1
MAINTAINER conor

RUN apk add libc6-compat \ 
		tzdata \
		ca-certificates \
		linux-headers \
          netcat-openbsd \
        && rm -rf /tmp/*
# apk add 是在打包时增加软件

# 拷贝代码
COPY ./kafka-study-go /kafka-study-go

# 端口
EXPOSE 8091

ENTRYPOINT ["./kafka-study-go"]

