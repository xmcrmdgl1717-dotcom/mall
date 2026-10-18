# 构建阶段
FROM golang:1.21-alpine AS builder
WORKDIR /app

# 使用代理加速依赖下载
ENV GOPROXY=https://goproxy.cn,direct

# 复制 go.mod 和 go.sum
COPY go.mod ./
# 如果你本地没有 go.sum，请把下面这行注释掉，并在本地执行 go mod tidy 生成它
# COPY go.sum ./

# 下载依赖
RUN go mod download

# 复制所有代码并编译
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main cmd/api/main.go

# 运行阶段
FROM alpine:latest
WORKDIR /app
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["/app/main"]
