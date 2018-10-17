FROM alpine:latest
RUN apk update && apk upgrade
COPY build/linux-amd64/bin/samplesvc /usr/local/bin/
EXPOSE 8081
ENTRYPOINT ["/usr/local/bin/samplesvc"]
