FROM golang:1.20

WORKDIR app/

COPY . .

RUN go build -o /app/main ./cmd/service

CMD sleep 20 && /app/main
