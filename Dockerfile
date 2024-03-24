FROM golang:1.22-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/debabky/voting-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/voting-svc /go/src/github.com/debabky/voting-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/voting-svc /usr/local/bin/voting-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["voting-svc"]
