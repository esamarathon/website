FROM node:8-slim AS npm-builder
WORKDIR /app
COPY package.json package-lock.json gulpfile.js ./
RUN npm ci
COPY scss scss
COPY scripts scripts
COPY templates templates
COPY templates_admin templates_admin
RUN npm run gulp

FROM golang:1.13.0-alpine AS golang
RUN apk add --no-cache ca-certificates
WORKDIR /go/src/github.com/esamarathon/website
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build .

FROM alpine:latest as runtime
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY public public
COPY --from=golang /go/src/github.com/esamarathon/website/website /app/
COPY --from=npm-builder /app/templates_minified /app/templates_minified/
COPY --from=npm-builder /app/templates_admin /app/templates_admin/
COPY --from=npm-builder /app/public /app/public/
EXPOSE 3001

CMD ["./website"]
