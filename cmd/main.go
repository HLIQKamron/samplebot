package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nickname76/telegrambot"
)

func main() {
	api, me, err := telegrambot.NewAPI("5850204459:AAGzVr1hz6tBQHYFBcge84qlxjkWSUXwhVk")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	stop := telegrambot.StartReceivingUpdates(api, func(update *telegrambot.Update, err error) {
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		msg := update.Message
		if msg == nil {
			return
		}
		m, _ := json.MarshalIndent(msg, " ", " ")
		fmt.Println("MESSAGE ", string(m))
		_, err = api.SendMessage(&telegrambot.SendMessageParams{
			ChatID: msg.Chat.ID,
			Text:   fmt.Sprintf("What is your name?"),
			// ReplyMarkup: &telegrambot.ReplyKeyboardMarkup{
			// 	Keyboard: [][]*telegrambot.KeyboardButton{{
			// 		{
			// 			Text: "Hello",
			// 		},
			// 	}},
			// 	ResizeKeyboard:  true,
			// 	OneTimeKeyboard: true,
			// 	Selective:       true,
			// },
		})

		if err != nil {
			log.Printf("Error: %v", err)
			return
		}
	})

	log.Printf("Started on %v", me.Username)

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt)

	<-exitCh

	// Waits for all updates handling to complete
	stop()
}
