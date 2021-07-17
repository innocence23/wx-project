FROM golang:alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io

# Get Reflex for live reload in dev
RUN go get github.com/cespare/reflex

COPY . .

RUN go build -o ./run ./cmd/main/

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

#Copy executable from builder
COPY --from=builder /go/src/app/run .

EXPOSE 8080
CMD ["./run"]