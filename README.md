## Conversational Telegram Bot
## Setup
Provide an .env file with bot token from BotFather with key "BOT_TOKEN" or edit the main.go
```go
 func getBotToken() (token string) {
	token, ok := os.LookupEnv("BOT_TOKEN")
  ...
```

or supply the bot token directly to
```go
pref := tele.Settings{
		Token: getBotToken(),
	}
 ...
 ```

##  build
To get an executable binary
> go build

To run the bot
> go run main.go

### Features

- Inspirational quote ✅
- Inspiratioal Image ✅
- Random inspiratioal quote ✅
- Cron jobs setup ✅
- Bible verse of the day ✅
- Audio bible message of the day ✅
- Persistent storage ✅
- Career motivational audio ❎
- Conversational AI ❎
- Remainder system ❎
- Game ❎
- Settings ❎
