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
		tgbotapi.NewInlineKeyboardButtonURL("GO Bootcamp", "https://t.me/uyga_vazifagofurov/114"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Foundation", "https://t.me/uyga_vazifagofurov/115"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonURL("Dizayn", "https://t.me/uyga_vazifagofurov/116"),
	),
)
