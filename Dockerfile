FROM --platform=$BUILDPLATFORM golang:alpine as builder

WORKDIR /app

ARG BUILDARCH
RUN wget https://github.com/upx/upx/releases/download/v4.2.4/upx-4.2.4-${BUILDARCH}_linux.tar.xz \
    && tar -xvf upx-4.2.4-${BUILDARCH}_linux.tar.xz

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -ldflags '-s -w' -o main main.go
RUN ./upx-4.2.4-${BUILDARCH}_linux/upx -q --lzma -1 main

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/template /app/template

EXPOSE 3000

CMD ["/app/main"]
