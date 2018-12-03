package main

import (
	"burgerking/client"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
	"fmt"
	"strconv"
	"burgerking/utils"
)

func GetValCode(id string) (valCode string) {
	if !utils.ValidateNumber(id){
		return "Invalid code."
	}
	c := client.BurgerkingClient{}
	err := c.StartSurvey(id)
	if err != nil {
		return "Survey starts failed."
	}
	valCode, err = c.DoSurvey()
	if err != nil {
		return "Survey failed."
	}
	return fmt.Sprintf("Survey Code: %s\nValCode: %s", id, valCode)
}

func main()  {
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
		b.Send(m.Sender, GetValCode(m.Text))
	})

	b.Start()

}
