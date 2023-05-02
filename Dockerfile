FROM golang:1.20-alpine

RUN mkdir /app

# WORKDIR /app

COPY . /app

RUN go mod tidy

RUN source .env

WORKDIR /app/cmd/telegrmbot

RUN go build -o telegrambot 

CMD [ "/telegrambot" ]

