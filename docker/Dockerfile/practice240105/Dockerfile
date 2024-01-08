FROM golang:1.20-alpine3.18 as builder
LABEL maintainer="maxwellke <jeremyke@qq.com>" \
    description="go peoject for reading xls excel"
ARG COMPILER='go build -a'
COPY / /app
WORKDIR /app
ENV GO111MODULE=on GOPROXY="https://goproxy.cn,direct" CGO_ENABLED=0 GOSUMDB=off GOOS=linux
RUN echo use COMPILER=$COMPILER
RUN go mod tidy
RUN $COMPILER -o read-excel main.go

FROM alpine:3.18
RUN apk add --upgrade curl
RUN apk update && apk upgrade && apk add -U tzdata && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /app/read-excel .
COPY --from=builder /app/reviews_1.xls .
CMD ["/read-excel"]