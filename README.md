# Stub Sample Service
This repository contains a stub implementation of the [Sample service] it simply provides a REST query front end to sample attributes stored in Redis loaded using (https://github.com/ONSdigital/census-rm-sample-loader). 

## Service Information API
* `GET /info` will return information about this service, collated from when it was last built.
* `GET /samples/<UUID>/attributes` will return the json stored under key "sampleunit:<UUID>" in redis

### Example JSON Response
```json
{
  "name": "samplesvc",
  "version": "1.0.0",
  "origin": "git@github.com:ONSdigital/go-sample.git",
  "commit": "b7ae66fbc54ca0abafe30950f69e48de72f66699",
  "branch": "master",
  "built": "2018-09-05T13:00:00Z"
}
```

## Environment
The following environment variables may be overridden:

| Environment Variable | Purpose            | Default Value |
| :------------------- | :------------------| :-------------|
| PORT                 | HTTP listener port | :8081         |
| REDIS_SERVICE_HOST   | host of redis      | localhost     |
| REDIS_SERVICE_PORT   | redis port         | 6379          |

## Dependencies
```
go get github.com/gomodule/redigo/redis
```

## Load test
You can use Hey to loadtest, your go bin directory must be on your path to run hey after it's installed. 

```
go get -u github.com/rakyll/hey
hey -n 10000 -c 20 http://localhost:8081/
```

## Copyright
Copyright (C) 2018 Crown Copyright (Office for National Statistics)