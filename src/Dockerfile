FROM golang:latest as build-stage

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN mkdir -p /go/src/myapp
COPY . /go/src/myapp

WORKDIR /go/src/myapp

ENV GOPATH $GOPATH:/go/src/myapp

RUN go get -d -v

RUN go build -o /go/bin/myapp

#Distroless Build
FROM gcr.io/distroless/base

COPY --from=build-stage /go/bin/myapp /

EXPOSE 8080

CMD ["/myapp"]
