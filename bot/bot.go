package bot

import (
	"log"

	"gitlab.com/telegram_bot/config"
	"gitlab.com/telegram_bot/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotHandler struct {
	cfg     config.Config
	storage storage.StorageI
	bot     *tgbotapi.BotAPI
}

func New(cfg config.Config, strg storage.StorageI) BotHandler {
	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	return BotHandler{
		cfg:     cfg,
		storage: strg,
		bot:     bot,
	}
}

func (h *BotHandler) Start() {
	log.Printf("Authorized on account %s", h.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {
		go h.HandleBot(update)
	}
}

func (h *BotHandler) HandleBot(update tgbotapi.Update) {
	user, err := h.storage.GetOrCreate(update.Message.From.ID, update.Message.From.FirstName)
	if err != nil {
		log.Println("failed to call storage.GetOrCreate: ", err)
		h.SendMessage(user, "error happened")
	}

	if update.Message.Command() == "start" {
		err = h.DisplayEnterFullnameMenu(user)
	} else if update.Message.Text != "" {
		switch user.Step {
		case storage.EnterFullnameStep:
			err = h.HandleEnterFullname(user, update.Message.Text)
		case storage.EnterPhoneNumberStep:
			err = h.HandleEnterPhoneNumber(user, update.Message.Text)
		case storage.CoursesStep:
			err = h.HandleEnterCourses(user, update.Message.Text)
		default:
			h.SendMessage(user, "Iltimos bot qaytadan ishga tushiring /start")
		}
	} else if update.Message.Contact != nil {
		err = h.HandleEnterPhoneNumber(user, update.Message.Contact.PhoneNumber)
	}

	if err != nil {
		log.Println("failed to handle message: ", err)
		h.SendMessage(user, "error happened")
	}
}

func (h *BotHandler) SendMessage(user *storage.User, message string) {
	msg := tgbotapi.NewMessage(user.TgID, message)
	if _, err := h.bot.Send(msg); err != nil {
		log.Println(err)
	}
}
