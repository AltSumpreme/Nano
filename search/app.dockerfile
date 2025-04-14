FROM golang:1.19.4-alpine3.18 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/AltSumpreme/Nano/
COPY go.mod go.sum ./
COPY vendor vendor
COPY search search
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./search/cmd/search
FROM alpine:3.18
WORKDIR /usr/bin
COPY --from=build /go/bin/ .
EXPOSE 8080
CMD ["app"]

