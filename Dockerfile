FROM golang:alpine AS builder

LABEL stage=gobuilder
ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .

RUN swag init -g ./cmd/main.go -o cmd/docs

RUN go build -ldflags="-s -w" -o /app/worker ./cmd/main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo

ENV TZ Asia/Tokyo

WORKDIR /app

COPY --from=builder /app/worker /app/worker

CMD ["./worker"]