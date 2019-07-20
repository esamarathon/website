FROM node:8-slim AS builder
WORKDIR /app
COPY . /app
RUN yarn && npm run gulp

FROM golang:1.10.0-alpine AS golang

RUN apk add --no-cache ca-certificates && mkdir -p /go/src/github.com/esamarathon/website/public
WORKDIR /go/src/github.com/esamarathon/website
COPY --from=builder /app .
ENV CGO_ENABLED=0
RUN go build .

FROM alpine:latest as runtime
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=golang /go/src/github.com/esamarathon/website/website /app/
COPY --from=builder /app/templates_minified /app/templates_minified/
COPY --from=builder /app/templates_admin /app/templates_admin/
COPY --from=builder /app/public /app/public/

EXPOSE 3001

CMD ["./website"]
