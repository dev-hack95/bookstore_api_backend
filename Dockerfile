FROM golang:alpine as builder
LABEL maintainer="dev-hack95"
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build cmd/main/main.go
EXPOSE 9010
CMD ["./main"]