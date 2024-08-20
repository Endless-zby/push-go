FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o push-server .

EXPOSE 10002

CMD ["./push-server", "-config", "/app/configs/config.yaml"]