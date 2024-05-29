FROM golang:alpine as builder
LABEL authors="Yun Seongmin ilcm96@gmail.com"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -ldflags '-s -w' -o main main.go \
    && wget -q https://github.com/upx/upx/releases/download/v4.2.4/upx-4.2.4-amd64_linux.tar.xz \
    && tar -xf upx-4.2.4-amd64_linux.tar.xz \
    && ./upx-4.2.4-amd64_linux/upx -q --lzma -1 main

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/template /app/template

EXPOSE 3000

CMD ["/app/main"]
