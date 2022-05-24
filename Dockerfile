FROM golang:alpine as pre-build

WORKDIR /go/src/send2ding
COPY . .
RUN go build -o send2ding cmd/main.go

FROM alpine
WORKDIR /app/

ENV TZ=Asia/Shanghai

COPY --from=pre-build /go/src/send2ding/send2ding /usr/local/bin
RUN echo "http://mirrors.aliyun.com/alpine/v3.10/main/" > /etc/apk/repositories && \
    apk --update --no-cache add ca-certificates