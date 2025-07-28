FROM golang:1.24.3-alpine

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o stress_test main.go

ENTRYPOINT ["/app/stress_test"]