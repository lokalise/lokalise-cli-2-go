FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM golang:1.13.4-alpine3.10 as builder

RUN mkdir -p /build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /bin/lokalise2

FROM scratch
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /bin/lokalise2 /bin/lokalise2


CMD ["/bin/lokalise2"]