package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	err := loadEnvironment()
	if err != nil {
		panic(err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(getDefaultHandler()),
	}

	b, err := bot.New(os.Getenv("TOKEN"), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}

	b.Start(ctx)
}

func loadEnvironment() error {
	err := godotenv.Load(".env")
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	return err
}

func getDefaultHandler() func(ctx context.Context, b *bot.Bot, update *models.Update) {
	chatId := os.Getenv("CHAT_ID")
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message != nil {
			if update.Message.Photo != nil || update.Message.Video != nil || update.Message.Document != nil {
				fmt.Println("sending photo")
				_, err := b.ForwardMessage(ctx, &bot.ForwardMessageParams{
					ChatID:     chatId,
					FromChatID: update.Message.Chat.ID,
					MessageID:  update.Message.ID,
				})
				if err != nil {
					fmtError := fmt.Errorf("erro sending photo %v", err)
					if fmtError != nil {
						panic(fmtError)
					}
				}
			}
		}
	}
}
