package modules

import (
	"errors"
	"os"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

var AdminID int64

func init() {
	adminID := os.Getenv("ADMIN_ID")
	if adminID == "" {
		panic("admin id not found on env variable. Exiting...")
	}

	AdminID, _ = strconv.ParseInt(adminID, 10, 64)
}

func broadcast(b *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveUser.Id != AdminID {
		ctx.EffectiveMessage.Reply(b, "You are not allowed to broadcast messages...", nil)
		return errors.New("user not permitted to broadcast")
	}
	if ctx.Message.ReplyToMessage == nil {
		ctx.EffectiveMessage.Reply(b, "You have to reply to a message in order to broadcast...", nil)
		return errors.New("no message to broadcast")
	}
	users, _ := getUsers()

	for _, user := range users {
		ctx.Message.ReplyToMessage.Forward(b, user.UserId, nil)

	}
	return nil

}
