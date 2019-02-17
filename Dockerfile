FROM node:lts-alpine AS builder

COPY . app
WORKDIR app
RUN rm -rf node_modules
RUN rm yarn.lock
RUN yarn
RUN npm run gulp
RUN rm -rf node_modules

FROM golang:1.10.0-alpine AS golang

RUN apk add --no-cache ca-certificates && mkdir -p /go/src/github.com/esamarathon/website/public
WORKDIR /go/src/github.com/esamarathon/website
COPY --from=builder /app .
RUN go build .

EXPOSE 3001

CMD ["./website"]
