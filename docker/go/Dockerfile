FROM golang:1.20.5-alpine

WORKDIR /app
COPY . /app

RUN apk update && \
    apk add --no-cache \
    alpine-sdk \
    git
RUN go install github.com/cosmtrek/air@latest
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]
