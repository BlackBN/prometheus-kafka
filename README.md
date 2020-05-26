# prometheus-kafka-adapter

#### 介绍
prometheus kafka 消息队列适配器，用于将prometheus 采集数据写入kafka

#### 软件架构
架构比较简单，就是以插件的方式部署适配器，接收prometheus 产生的数据并通过异步kafka 生产者模式，将数据输送给kafka

![输入图片说明](https://images.gitee.com/uploads/images/2019/0817/101843_77f1d907_5060217.jpeg "QQ截图20190817101836.jpg")


#### 安装教程

1. 参考Dockerfile 打包镜像

```
$ docker build -t registry.cn-hangzhou.aliyuncs.com/test/prometheus-kafka-adapter:1.6 .'
$ docker push registry.cn-hangzhou.aliyuncs.com/test/prometheus-kafka-adapter:1.6 
```


2. 创建服务configmap / deployment / service

```
$ kubectl apply -f config.yaml
$ kubectl apply -f deployment.yaml
$ kubectl apply -f service.yaml
```


#### 使用说明

1.prometheus配置

```
# 加上remote_write 配置
remote_write:
      - url: "http://3.3.3.3:9201/write"
```



2. 启动kafka 消费者验证消息发送是否成功
```
root@kafka-01:/usr/local/kafka/bin#  ./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic prometheus2kafka --max-messages 20
{"value":0,"@timestamp":"2019-08-17T10:05:26.214+08:00","labels":{"__name__":"prometheus_tsdb_compaction_duration_seconds_bucket","instance":"localhost:9090","job":"prometheus","le":"64"}}
{"value":0,"@timestamp":"2019-08-17T10:05:26.214+08:00","labels":{"__name__":"prometheus_tsdb_compaction_duration_seconds_bucket","instance":"localhost:9090","job":"prometheus","le":"128"}}
{"value":0,"@timestamp":"2019-08-17T10:05:26.214+08:00","labels":{"__name__":"prometheus_tsdb_compaction_duration_seconds_bucket","instance":"localhost:9090","job":"prometheus","le":"256"}}
{"value":0,"@timestamp":"2019-08-17T10:05:26.214+08:00","labels":{"__name__":"prometheus_tsdb_compaction_duration_seconds_bucket","instance":"localhost:9090","job":"prometheus","le":"512"}}
{"value":0,"@timestamp":"2019-08-17T10:05:26.214+08:00","labels":{"__name__":"prometheus_tsdb_compaction_duration_seconds_bucket","instance":"localhost:9090","job":"prometheus","le":"+Inf"}}
{"value":0,"@timestamp":"2019-08-17T10:05:26.214+08:00","labels":{"__name__":"prometheus_tsdb_compaction_duration_seconds_sum","instance":"localhost:9090","job":"prometheus"}}
{"value":0,"@timestamp":"2019-08-17T10:05:26.214+08:00","labels":{"__name__":"prometheus_tsdb_compaction_duration_seconds_count","instance":"localhost:9090","job":"prometheus"}}
... ...
```



