package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type App struct {
	BotToken       string
	ChatID         string
	FeedURL        string
	Interval       time.Duration `json:"-"`
	IntervalSecond int           `json:"Interval"`

	LastPublished time.Time
}

func GetApp() (app *App, err error) {
	data, err := ioutil.ReadFile("app.json")
	if err != nil {
		return
	}

	app = &App{}
	err = json.Unmarshal(data, app)
	if err != nil {
		return
	}

	app.Interval = time.Second * time.Duration(app.IntervalSecond)
	return
}

func (app *App) Run() (err error) {
	err = app.Publish()
	if err != nil {
		log.Println(err.Error())
	}

	for range time.Tick(app.Interval) {
		err := app.Publish()
		if err != nil {
			log.Println(err.Error())
		}
	}
	return
}

func (app *App) Save() (err error) {
	data, err := json.MarshalIndent(app, "", "    ")
	if err != nil {
		return
	}

	err = ioutil.WriteFile("app.json", data, 0600)
	return
}
