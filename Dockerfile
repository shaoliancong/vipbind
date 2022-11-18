FROM golang:1.17.3 as builder

WORKDIR /app

COPY . .

RUN GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0  go build -o vipbind main.go

FROM alpine:3.15.4

WORKDIR /app

COPY --from=builder /app/vipbind .

CMD ["./vipbind"]