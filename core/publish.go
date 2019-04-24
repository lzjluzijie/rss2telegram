package core

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/pkg/errors"
)

type Message struct {
	Title string
	Link  string
}

func (app *App) Publish() (err error) {
	t := time.Now()

	req, err := http.NewRequest("GET", app.FeedURL, nil)
	if err != nil {
		return
	}

	req.Header.Set("User-Agent", app.UA)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	fp := gofeed.NewParser()
	feed, err := fp.Parse(resp.Body)
	if err != nil {
		return
	}

	ms := make([]*Message, 0)

	//下面是一段很蠢的代码
	for i := 0; feed.Items[i].PublishedParsed.After(app.LastPublished); i++ {
		ms = append(ms, &Message{
			Title: html.UnescapeString(feed.Items[i].Title),
			Link:  feed.Items[i].Link,
		})

		if i+1 >= len(feed.Items) {
			break
		}
	}

	for i := len(ms) - 1; i >= 0; i-- {
		err = app.SendMessage(ms[i])
		if err != nil {
			log.Printf(err.Error())
			return
		}
	}
	app.LastPublished = t
	err = app.Save()
	return
}

func (app *App) SendMessage(m *Message) (err error) {
	text := fmt.Sprintf("[%s](%s)", m.Title, m.Link)
	log.Println(text)

	resp, err := http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", app.BotToken), map[string][]string{
		"chat_id":    {app.ChatID},
		"text":       {text},
		"parse_mode": {"Markdown"},
	})
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("status code: %d", resp.StatusCode))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println(string(data))
	return
}
