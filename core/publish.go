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

func (app *App) Publish() (err error) {
	t := time.Now()
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(app.FeedURL)

	fmt.Println(feed)

	for i := 0; feed.Items[i].PublishedParsed.After(app.LastPublished); i++ {
		title := html.UnescapeString(feed.Items[i].Title)
		link := feed.Items[i].Link
		err = app.SendMessage(title, link)
		if err != nil {
			return
		}
	}
	app.LastPublished = t
	err = app.Save()
	return
}

func (app *App) SendMessage(title, link string) (err error) {
	text := fmt.Sprintf("[%s](%s)", title, link)
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
