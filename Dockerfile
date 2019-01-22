FROM golang:latest as builder
COPY . /go/src/terminal-adventure
WORKDIR /go/src/terminal-adventure
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags "-s -w" -installsuffix server -o /go/bin/server

FROM alpine:latest
WORKDIR /
RUN mkdir /app
COPY --from=builder /go/bin/server /app/
CMD ["/app/server"]
