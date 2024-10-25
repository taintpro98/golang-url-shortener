# Build stage
FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Building app
RUN cd cmd/app && CGO_ENABLED=0 GOOS=linux go build -o app

# Final stage
FROM alpine:3.18 as runner

ARG POSTGRES_HOST
ARG POSTGRES_DB
ARG POSTGRES_USER
ARG POSTGRES_PASSWORD
ARG POSTGRES_PORT
ARG APP_DOMAIN

ENV POSTGRES_HOST $POSTGRES_HOST
ENV POSTGRES_DB $POSTGRES_DB
ENV POSTGRES_USER $POSTGRES_USER
ENV POSTGRES_PASSWORD $POSTGRES_PASSWORD
ENV POSTGRES_PORT $POSTGRES_PORT
ENV APP_DOMAIN $APP_DOMAIN

ENV ENV_CONFIG_ONLY=true
ENV GIN_MODE=release
ENV HOST 0.0.0.0
ENV PORT=8080

# Tạo thư mục config nếu chưa tồn tại
RUN mkdir -p /app/config

COPY --from=builder /app/cmd/app/app /app/

WORKDIR /app

EXPOSE 8080

# Run the web service on container startup.
CMD ["/app/app"]