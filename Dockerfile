### Usually multi-stage build process would be used together with 'stretch' image containing app binary as output,
### but as this is not a standalone application and godog is used for tests - regular golang image will do
FROM golang:1.12.9-alpine3.10

RUN apk add git --no-cache

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /go/src/github.com/r0kas/form3-accountapi-client
COPY . .
RUN go get -d -v && go get github.com/DATA-DOG/godog/cmd/godog

WORKDIR /go/src/github.com/r0kas/form3-accountapi-client/test
ENTRYPOINT ["sh", "-c"]
CMD ["godog"]
