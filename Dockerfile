FROM golang:1.12 as builder
LABEL maintainer="Joel Messerli <hi.github@peg.nu>"
WORKDIR /go/src/github.com/jmesserli/vros
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/vros .

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /go/bin/vros .
CMD ["./vros"]