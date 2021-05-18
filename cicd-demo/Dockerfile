FROM golang:latest as build-env
WORKDIR /go/src/helloworld
ADD . /go/src/helloworld
RUN CGO_ENABLED=0  GOFLAGS=-mod=vendor go build -ldflags '-w -s' -o /go/app
#go build -w -s -o /go/app

FROM golang:latest as prod-env
WORKDIR /go/src/helloworld
COPY --from=build-env /go/app .
EXPOSE 8080
CMD ["./app"]