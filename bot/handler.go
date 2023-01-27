package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gitlab.com/telegram_bot/storage"
)

func (h *BotHandler) DisplayEnterFullnameMenu(user *storage.User) error {
	err := h.storage.ChangeStep(user.TgID, storage.EnterFullnameStep)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(user.TgID, "Ism familyangizni kiriting: ")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}

	return nil
}

func (h *BotHandler) HandleEnterFullname(user *storage.User, fullname string) error {
	// TODO: validate text
	err := h.storage.ChangeField(user.TgID, "fullname", fullname)
	if err != nil {
		return err
	}

	return h.DisplayEnterPhoneNumberMenu(user)
}

func (h *BotHandler) DisplayEnterPhoneNumberMenu(user *storage.User) error {
	err := h.storage.ChangeStep(user.TgID, storage.EnterPhoneNumberStep)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(user.TgID, "Telefon raqamingizni kiriting: ")
	msg.ReplyMarkup = enterPhoneNumberMenuKeyboard
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}

	return nil
}

func (h *BotHandler) HandleEnterPhoneNumber(user *storage.User, text string) error {
	if text == "⬅️ Orqaga" {
		return h.DisplayEnterFullnameMenu(user)
	}

	// TODO: validate phone number
	err := h.storage.ChangeField(user.TgID, "phone_number", text)
	if err != nil {
		return err
	}

	return h.DisplayRegisteredMenu(user)
}

func (h *BotHandler) DisplayEnterCourseMenu(user *storage.User) error {
	err := h.storage.ChangeCourseStep(user.TgID, storage.CoursesStep)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(user.TgID, "Kurslarni ko'rish uchun bosing: 👇")
	msg.ReplyMarkup = enterCourseMenuKeyboard
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}

	return nil
}
func (h *BotHandler) HandleEnterCourses(user *storage.User, text string) error {
	if text == "Courses" {
		return h.DisplayEnterCourses(user)
	}

	return h.DisplayRegisteredMenu(user)
}

func (h *BotHandler) DisplayEnterCourses(user *storage.User) error {
	err := h.storage.ChangeCourseStep(user.TgID, storage.CoursesStep)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(user.TgID, "Kurslar bilan tanishing: ")
	msg.ReplyMarkup = numericKeyboard
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
	return nil
}

func (h *BotHandler) DisplayRegisteredMenu(user *storage.User) error {
	err := h.storage.ChangeStep(user.TgID, storage.RegisteredStep)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(user.TgID, "Siz muvaffaqiyatli ro'yxatdan o'tdingiz!!!")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}

	return h.DisplayEnterCourseMenu(user)
}
