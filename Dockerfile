# Build Server
FROM golang:latest as builder
ENV GOBIN /go/bin
WORKDIR /go/src/github.com/QualiArts/cb-sample-server
COPY / .
RUN cd pkg && go test
RUN cd pkg && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../bin/cb-sample-server

# COPY theserver file to image from builder
FROM alpine:latest
WORKDIR /usr/local/bin/
COPY --from=builder /go/src/github.com/QualiArts/cb-sample-server/bin/cb-sample-server .
# ConfigファイルをCopy
WORKDIR /usr/local/config/
COPY --from=builder /go/src/github.com/QualiArts/cb-sample-server/config.yaml .

EXPOSE 8080
# Configファイルを設定してアプリを起動
CMD ["/usr/local/bin/cb-sample-server", "-config=/usr/local/config/config.yaml"]
