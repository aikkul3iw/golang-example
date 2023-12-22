FROM golang:1.21.5-alpine3.19 AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
ARG VERSION
ARG BUILD
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-X main.version=${VERSION} -X main.build=${BUILD}" -o golang-example .

FROM alpine:3.19.0
RUN apk --no-cache add ca-certificates
RUN addgroup -S runner && adduser -S runner -G runner
USER runner
WORKDIR /home/main
COPY --from=builder /build/golang-example ./
COPY --from=builder /build/.env ./
ENTRYPOINT ["./golang-example"]