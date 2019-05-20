FROM golang:alpine AS build

WORKDIR /go/src/bitbucket.org/albertogviana/go-microservice/names/
COPY . .

RUN apk update && \
  apk add git &&\
  go get gopkg.in/mgo.v2 && \
  go build -v -o names-demo



FROM alpine

EXPOSE 8080

ENV APP_VERSION 1.0.0

COPY --from=build /go/src/bitbucket.org/albertogviana/go-microservice/names/names-demo /usr/local/bin/names-demo

RUN chmod +x /usr/local/bin/names-demo

HEALTHCHECK --interval=10s CMD wget -qO- localhost:8080/health

CMD ["names-demo"]
