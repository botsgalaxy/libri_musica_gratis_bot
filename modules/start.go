package modules

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var StartMessage = `<b>✅ | SBLOCCA L' ACCESSO </b>
 
Per sbloccare l'accesso a libri, audiolibri e musica GRATIS, iscriviti ai canali qui sotto. 
	 
Segui questi 2 step: 
🔸1. Clicca su "ISCRIVITI AL CANALE" ed unisciti ad entrambi i canali. 
🔸2. Ritorna qui e clicca su "SBLOCCA"`

var ReplyMarkup = gotgbot.InlineKeyboardMarkup{
	InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
		{gotgbot.InlineKeyboardButton{
			Text: "🔵 ISCRIVITI AL CANALE 1🔵",
			Url:  "https://t.me/+hkF_zFBk66ZmYTY0",
		}},
		{gotgbot.InlineKeyboardButton{
			Text: "🔵 ISCRIVITI AL CANALE 2🔵",
			Url:  "https://t.me/instagram_followers_likeeee",
		}},
		{
			gotgbot.InlineKeyboardButton{
				Text:         "✅ SBLOCCA ✅",
				CallbackData: "verify",
			},
		},
	},
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, StartMessage, &gotgbot.SendMessageOpts{
		ParseMode:   "html",
		ReplyMarkup: ReplyMarkup,
	})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}
