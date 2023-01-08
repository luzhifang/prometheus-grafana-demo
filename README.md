## 介绍

本示例项目包括了 Prometheus、Grafana、AlertManager、node-exporter、docker-compose

## 教程

可参考[个人博客](https://luzhifang.github.io/categories/%E7%9B%91%E6%8E%A7%E5%91%8A%E8%AD%A6/)

## 使用

修改 `hoststats-alert.yml` 和 `prometheus.yml` 文件的需要修改为本机IP的地方。

运行以下命令

```Shell
mkdir -p alertmanager-data
make build
docker-compose up
```

访问 http://localhost:9090/ 可查看 Prometheus 控制台

访问 http://localhost:9093/ 可查看 AlertManager 控制台

访问 http://localhost:3000/ 可查看 Grafana 控制台