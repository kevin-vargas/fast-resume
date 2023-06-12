FROM node:18 as build
WORKDIR /app
COPY app/package*.json .
RUN npm install
# TODO: remove dockerignore and improve copy
COPY app .
RUN npm run build

FROM golang:alpine AS builder

USER root 

WORKDIR /app

COPY --from=build /app/dist /app/app/dist

COPY . .

RUN apk add git

RUN apk --no-cache add ca-certificates

RUN go get -d -v

RUN CGO_ENABLED=0 go build -o /bin/app

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /bin/app /app

CMD ["/app"]
