# Development environment
# -----------------------
FROM golang:1.18-alpine
WORKDIR /browser-chat

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x /browser-chat/docker/app/bin/*

EXPOSE 8080