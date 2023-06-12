FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN apk add git

RUN apk --no-cache add ca-certificates

RUN go get -d -v

RUN CGO_ENABLED=0 go build -o /bin/app

FROM alpine:latest

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /bin/app /app

COPY --from=builder /app/cert /cert

CMD ["/app"]