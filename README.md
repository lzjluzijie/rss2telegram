# rss2telegram

## Config

Get bot token: https://t.me/BotFather

Get chat id: https://api.telegram.org/bot<token>/getUpdates

app.json
```json
{
    "BotToken": "",
    "ChatID": "",
    "FeedURL": "https://zhuji.lu/topics.rss",
    "Interval": 300,
    "LastPublished": "1911-10-10T18:00:00.3182961+08:00"
}
```

## Set proxy (if telegram is blocked)

```bash
export http_proxy='http://127.0.0.1:1080'
export https_proxy='http://127.0.0.1:1080'
```
