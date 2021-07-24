FROM golang

WORKDIR /app
COPY . /app

ARG cron="20 0 * * *"
ARG cfgFile=config/config.yaml
COPY ${cfgFile} config/config.yaml

RUN go env -w GO111MODULE=on && \
 go env -w GOPROXY=https://goproxy.cn,direct
RUN go test ./... && \
 go build -o bin/ ./... && \
 mv bin/src bin/main

RUN apt-get update && \
 apt-get install --no-install-recommends -y cron && \
 rm -rf /var/lib/apt/lists/* && \
 apt-get clean
RUN echo "${cron} cd /app && bin/main >> tmp.log" | crontab -

ENTRYPOINT [ "./init.sh" ]
