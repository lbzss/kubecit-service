FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

RUN curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip && unzip protoc-3.15.8-linux-x86_64.zip -d /usr/local 

RUN  make init && make all && make build

FROM debian:stable-slim



COPY --from=builder /src/bin /app
COPY --from=builder /src/configs /app/configs

WORKDIR /app



EXPOSE 8000
#EXPOSE 9000
#VOLUME /data/conf

CMD ["./kubecit-service", "-conf", "/app/configs/qa_config.yaml"]
