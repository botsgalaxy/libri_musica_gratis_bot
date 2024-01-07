package modules

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"os"
	"strconv"
)

var IntChatId int64

func init() {
	chatId := os.Getenv("CHAT_ID")
	if chatId == "" {
		panic("CHAT_ID environment variable is empty")
	}

	var err error
	IntChatId, err = strconv.ParseInt(chatId, 10, 64)
	if err != nil {
		panic("Error parsing environment variable:" + err.Error())
	}
}

func verify(b *gotgbot.Bot, ctx *ext.Context) error {
	query := ctx.Update.CallbackQuery
	query.Answer(b, nil)
	userId := ctx.EffectiveUser.Id
	chatId := IntChatId
	chatMember, err := b.GetChatMember(chatId, userId, nil)
	if err != nil {
		_, err := b.SendMessage(userId, StartMessage, &gotgbot.SendMessageOpts{
			ParseMode:   "html",
			ReplyMarkup: ReplyMarkup,
		})
		if err != nil {
			return fmt.Errorf("failed to send start message: %w", err)
		}
	}
	chatMemberStatus := chatMember.GetStatus()
	if chatMemberStatus == "creator" || chatMemberStatus == "administrator" || chatMemberStatus == "member" || chatMemberStatus == "restricted" {
		text := `
<b>Grazie per il tuo supporto ❤️ </b>
	 
	 
<b>👇 CLICCA QUI 👇</b>
	
`
		var replyMarkup = gotgbot.InlineKeyboardMarkup{
			InlineKeyboard: [][]gotgbot.InlineKeyboardButton{
				{gotgbot.InlineKeyboardButton{
					Text: "🔑 VAI AL CONTENUTO SBLOCCATO 🔑",
					Url:  "https://t.me/pazzoscontoitalia/4898",
				}},
			},
		}
		_, err = ctx.EffectiveMessage.Reply(
			b,
			text,
			&gotgbot.SendMessageOpts{
				ParseMode:   "html",
				ReplyMarkup: replyMarkup,
			},
		)
		return err
	} else {
		_, err := b.SendMessage(userId, StartMessage, &gotgbot.SendMessageOpts{
			ParseMode:   "html",
			ReplyMarkup: ReplyMarkup,
		})
		if err != nil {
			return fmt.Errorf("failed to send start message: %w", err)
		}

	}
	return nil

}
