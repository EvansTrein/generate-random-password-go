FROM golang:1.23-alpine as builder
WORKDIR /app
COPY . /app
RUN go mod download && go build -o password_generator main.go

# further we will use only the compiled file go
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/password_generator .
COPY --from=builder /app/static ./static
EXPOSE 8800
CMD ["./password_generator"]

