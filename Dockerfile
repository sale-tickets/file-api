FROM golang:1.25-alpine AS builder

WORKDIR /build

COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine:latest AS runner

WORKDIR /run

COPY --from=builder /build/ /run/

EXPOSE 8000
EXPOSE 50000

ENTRYPOINT [ "/run/main" ]