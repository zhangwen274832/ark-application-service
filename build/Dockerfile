## 设置编译环境
FROM harbor.nuclearport.com/golang/golang:1.16 as builder

MAINTAINER panzhiqiang@fotoable.com

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPRIVATE=gitlab.ftsview.com \
    GOPROXY=https://proxy.golang.org,direct

WORKDIR /app

COPY . .

RUN go build -o ark-application-service .

## 设置运行环境
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ark-application-service .

COPY --from=builder /app/configs configs

COPY --from=builder /app/static static

EXPOSE 40200

ENV AMAN_ADDR "https://aman.akgoo.net/"

ENV AMAN_ENV_ID 1

ENV AMAN_PROJECT_ID "ark-application-service"

ENV VERSION latest

ENTRYPOINT ./ark-application-service "--aman_addr=$AMAN_ADDR" "--aman_env_id=$AMAN_ENV_ID" "--aman_project_id=$AMAN_PROJECT_ID" "--version=$VERSION"
