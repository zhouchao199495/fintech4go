FROM golang:1.18

ENV BUILD_DIR=/usr/local/app
RUN mkdir -p ${BUILD_DIR}
WORKDIR ${BUILD_DIR}

COPY . ${BUILD_DIR}
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"

EXPOSE 9999
RUN go build -o fintech4go .
ENTRYPOINT  ["./fintech4go"]