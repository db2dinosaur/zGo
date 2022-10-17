FROM golang:1.11
USER root
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY ./goweb/goweb.go /go/src/app
COPY ./goweb/go.mod /go/src/app/
EXPOSE 9999
RUN go build
CMD ["./app"]