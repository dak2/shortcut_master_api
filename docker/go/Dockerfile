FROM golang:1.16.6-alpine

WORKDIR /app
COPY . /app

RUN apk update && \
    apk add --no-cache \
    alpine-sdk \
    git
RUN go mod tidy
RUN go get -u github.com/cosmtrek/air
ENV GO111MODULE=on
CMD ["air"]
