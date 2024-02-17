FROM golang:1.20.3-alpine AS builder

COPY . /github.com/Murat993/chat-client/source/
WORKDIR /github.com/Murat993/chat-client/source/

RUN go mod download
RUN go build -o ./bin/client cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Murat993/chat-client/source/bin/client .

CMD ["./client"]