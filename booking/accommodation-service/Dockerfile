FROM golang:alpine as build_container
WORKDIR app
RUN apk update \
    && apk add sqlite \
    && apk add socat \
    && apk add build-base
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o accommodation-service
EXPOSE 8091
CMD ["./accommodation-service"]