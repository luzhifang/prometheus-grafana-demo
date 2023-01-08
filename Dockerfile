#源镜像
FROM alpine:latest
#设置工作目录
WORKDIR /
#将服务器的go工程代码加入到docker容器中
ADD prometheus-grafana-demo /prometheus-grafana-demo
#暴露端口
EXPOSE 5000
#最终运行docker的命令
ENTRYPOINT  ["/prometheus-grafana-demo"]