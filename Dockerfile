FROM golang:alpine AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN apk add --no-cache git
RUN go get github.com/gomodule/redigo/redis
RUN go build -o main .

FROM alpine:latest
COPY --from=build /app/main .
EXPOSE 8125
CMD ["./main"]
