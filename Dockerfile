FROM golang:latest  AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV GOPROXY=https://goproxy.cn,direct

# 移动到工作目录 build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译为可执行文件到fwbot
RUN go build -o fwbot .

# 创建一个小镜像
#FROM scratch
FROM busybox

# 从builder镜像中将拉取 build/app 到当前目录
COPY --from=builder /build/fwbot /

EXPOSE 8077
ENTRYPOINT ["/fwbot"]