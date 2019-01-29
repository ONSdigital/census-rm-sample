FROM golang:alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN apk add --no-cache git \
    && go get github.com/gomodule/redigo/redis \
    && go build -o main .
EXPOSE 8125
CMD ["./main"]