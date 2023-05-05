FROM golang:1.20 AS builder

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=1 GOOS=linux go build -o telegrambot -a -ldflags '-linkmode external -extldflags "-static"' .

FROM alpine 

WORKDIR /root/

COPY --from=builder /app/telegrambot /app/.env /root/

CMD ./telegrambot