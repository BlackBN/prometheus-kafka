FROM registry.cn-hangzhou.aliyuncs.com/test/ubuntu_16.04:latest  
LABEL description="Kafka write adapter for Prometheus remote storage."

ENV APP_PATH /data/services/prometheus-kafka

RUN mkdir -p $APP_PATH/config

WORKDIR $APP_PATH

COPY prometheus-kafka ./
COPY config.yaml ./

CMD ["./prometheus-kafka"]

