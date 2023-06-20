FROM golang:1.20 AS builder

LABEL maintainer = "Giwa Oluwatobi, giwaoluwatobi@gmail.com"

WORKDIR /app

COPY . /app

RUN go mod download 

RUN CGO_ENABLED=1 GOOS=linux go build -o telegrambot -a -ldflags '-linkmode external -extldflags "-static"' .

FROM alpine 

WORKDIR /root/

COPY --from=builder /app/telegrambot /root/

CMD ./telegrambot