FROM golang:1.16

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o push-server .

EXPOSE 8080

CMD ["./push-server", "-config", "/app/configs/config.yaml"]