FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY . ./

RUN go build -o /rest-api ./src/cmd/main.go


FROM alpine

WORKDIR /

COPY --from=builder /rest-api /rest-api
EXPOSE 8000

CMD ["/rest-api"]