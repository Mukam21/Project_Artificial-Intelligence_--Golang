package bot

import (
	"context"
	"strconv"

	"github.com/Mukam21/ChatGPT-Golang-psql/internal/botkit"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SourceDeleter interface {
	Delete(ctx context.Context, sourceID int64) error
}

func ViewCmdDeleteSource(deleter SourceDeleter) botkit.ViewFunc {
	return func(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
		idStr := update.Message.CommandArguments()

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		if err := deleter.Delete(ctx, id); err != nil {
			return nil
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Источник успешно удален")
		if _, err := bot.Send(msg); err != nil {
			return err
		}

		return nil
	}
}
