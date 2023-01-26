package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var enterPhoneNumberMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButtonContact("📱 Telefon raqamni jo'natish"),
		tgbotapi.NewKeyboardButton("⬅️ Orqaga"),
	),
)

var enterCourseMenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Courses"),
	),
)


var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonURL("", "http://1.com"),
        tgbotapi.NewInlineKeyboardButtonData("2", "2"),
        tgbotapi.NewInlineKeyboardButtonData("3", "3"),
    ),
    tgbotapi.NewInlineKeyboardRow(
        tgbotapi.NewInlineKeyboardButtonData("4", "4"),
        tgbotapi.NewInlineKeyboardButtonData("5", "5"),
        tgbotapi.NewInlineKeyboardButtonData("6", "6"),
    ),
)
