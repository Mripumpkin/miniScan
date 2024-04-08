
# 约定:
# 构建与生产使用固定的目录 /app

# 阶段一: build
# docker pull golang:1.17.5-alpine3.15
FROM golang:1.17.5-alpine3.15 AS build
WORKDIR /app
COPY . .
RUN apk add build-base
ENV GORPOXY=https://goproxy.cn,direct
RUN go build -o server .


# 阶段二: 生产环境
FROM alpine:latest as prod
WORKDIR /app
COPY --from=build /app/server .    
# EXPOSE 8000
ENTRYPOINT [ "/app/server" ]
