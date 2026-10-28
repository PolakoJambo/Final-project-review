FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-app main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /todo-app .
COPY --from=builder /app/web ./web

EXPOSE 7540
CMD ["./todo-app"]