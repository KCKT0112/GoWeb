FROM golang:1.23.
WORKDIR /data/golang/GoWeb
COPY . .
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go mod download \
    && go mod tidy \
    && go build -o main .

FROM alpine:latest
LABEL MAINTAINER="foxiti.a@gmail.com"

WORKDIR /data/golang/GoWeb
COPY --from=0 /data/golang/GoWeb/main .
EXPOSE 8083
ENTRYPOINT ./main