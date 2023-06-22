# 表示依赖 alpine 最新版
FROM alpine:latest
MAINTAINER 3xxx<504284@qq.com>
ENV VERSION 1.0

ENV GO111MODULE=on
ENV GOROOT=/usr/local/go
ENV GOPATH=/home/gopath
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
ENV GOPROXY=https://goproxy.io
ENV GOOS=linux

RUN go install github.com/astaxie/beego
RUN go install github.com/beego/bee

# 设置固定的项目路径
ENV WORKDIR /root/engineercms

# 在容器根目录 创建一个 engineercms 目录
WORKDIR $WORKDIR
# 挂载容器目录
VOLUME ["/root/engineercms/conf"]
# 拷贝配置文件到容器中
COPY conf/app.conf $WORKDIR/conf/app.conf
COPY conf/app2.conf $WORKDIR/conf/app2.conf

# 添加应用可执行文件，并修改权限
ADD ./engineercms $WORKDIR/engineercms
RUN chmod +x $WORKDIR/engineercms

# 添加静态文件、模版文件、日志目录
ADD static $WORKDIR/static
ADD view $WORKDIR/view
ADD log $WORKDIR/log
 
# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
 
# 设置编码
ENV LANG C.UTF-8

#安装编译需要的环境gcc等
RUN apk add build-base

#交叉编译，需要制定CGO_ENABLED=1，默认是关闭的
RUN  GOOS=linux CGO_ENABLED=1 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o ./bin/localized main.go
#暴露的端口号要与配置文件app.conf中配置的httpport端口号保持一致
EXPOSE 8088 

# 运行golang程序的命令 
ENTRYPOINT ["./engineercms"]
#ENTRYPOINT ["/apps/golang_app"]

#指令编译生成名为main的docker镜像
#docker build -t engineercms .
#docker build -t engineercms:v0.0.1 .
#运行镜像
#docker run -d --name engineercms -p port:port engineercms
#docker run -d --name engineercms -p 8089:8089 engineercms:v0.0.1