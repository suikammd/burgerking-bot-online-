package main

import (
	"fmt"
	"strconv"
	"time"

	"burgerking/client"
	"burgerking/utils"
	tb "gopkg.in/tucnak/telebot.v2"
	"strings"
)

func getValCodeOnline(id string) (message string) {
	if !utils.ValidateNumber(id) {
		return "Invalid code."
	}
	c := client.BurgerkingClient{}
	err := c.StartSurvey(id)
	if err != nil {
		return "Online Mode:\nSurvey starts failed."
	}
	message, err = c.DoSurvey()
	if err != nil {
		return "Online Mode:\nSurvey failed."
	}
	return fmt.Sprintf("Online Mode:\nSurvey Code: %s\nValCode: %s", id, message)
}

func getValCodeOffline(id string) (message string) {
	valCode, _ := utils.CalcValCode(id)
	message = fmt.Sprintf("Offline Mode:\nSurvey Code: %s\nValCode: %s", id, valCode)
	return
}

func getValCode(text string) (message string) {
	if strings.HasPrefix(text, "online") || strings.HasPrefix(text, "Online") {
		message = getValCodeOnline(text[6:])
	} else {
		message = getValCodeOffline(text)
	}
	return
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "TOKEN HERE",
		Poller: &tb.LongPoller{Timeout: 1 * time.Second},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	b.Handle("/id", func(m *tb.Message) {
		b.Send(m.Sender, strconv.FormatInt(m.Chat.ID, 10))
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		b.Send(m.Sender, getValCode(m.Text))
	})

	b.Start()
}
