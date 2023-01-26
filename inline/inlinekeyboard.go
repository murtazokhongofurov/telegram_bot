package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/telegram_bot/storage"
)

type InlineKeyboardsI interface {
	NewCourseKeyboard()
}
type InlineKeyboards struct {
}

func NewInlineKeyboards() InlineKeyboards {
	return InlineKeyboards{}
}

func (InlineKeyboards) NewCourseKeyboard(s []*storage.Course) tgbotapi.InlineKeyboardMarkup {
	markup := tgbotapi.NewInlineKeyboardMarkup()
	for _, val := range s {
		callbackData := fmt.Sprintf("course:%d", val.Id)
		markup.InlineKeyboard = append(markup.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{Text: val.CourseName, CallbackData: &callbackData},
		))
	}
	return markup
}
