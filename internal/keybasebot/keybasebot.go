package keybasebot

import (
	"log"
	"os"

	"github.com/keybase/go-keybase-chat-bot/kbchat"
	"github.com/keybase/go-keybase-chat-bot/kbchat/types/chat1"
	"github.com/szeist/home-bot/internal/config"
	"github.com/szeist/home-bot/internal/events"
)

type KeybaseBot struct {
	kbOpts         *kbchat.RunOptions
	kbChatAPI      *kbchat.API
	kbSubscription *kbchat.Subscription
	logger         *log.Logger
}

func New(kbConfig *config.KBConfig) *KeybaseBot {
	kbOpts := &kbchat.RunOptions{
		KeybaseLocation: kbConfig.Location,
		HomeDir:         kbConfig.Home,
		StartService:    false,
		Oneshot: &kbchat.OneshotOptions{
			Username: kbConfig.User,
			PaperKey: kbConfig.PaperKey,
		},
	}

	return &KeybaseBot{
		kbOpts: kbOpts,
		logger: log.New(os.Stderr, "home-bot:keybasebot", log.LstdFlags),
	}
}

func (k *KeybaseBot) Start() error {
	kbc, err := kbchat.Start(*k.kbOpts)
	if err != nil {
		return err
	}
	k.kbOpts = nil
	k.kbChatAPI = kbc

	sub, err := kbc.ListenForNewTextMessages()
	k.kbSubscription = sub

	return err
}

func (k *KeybaseBot) HandleMessages(messageChan chan *events.Message) {
	for {
		msg, err := k.kbSubscription.Read()
		if err != nil {
			k.logger.Printf("failed to read message: %s", err.Error())
		}

		if msg.Message.Content.TypeName != "text" {
			continue
		}

		if msg.Message.Sender.Username == k.kbChatAPI.GetUsername() {
			continue
		}

		cmd := msg.Message.Content.Text.Body
		messageChan <- &events.Message{
			Text:   cmd,
			Sender: msg.Message,
		}
	}

}

func (k *KeybaseBot) HandleResponses(responseChan chan *events.Response) {
	for {
		resp := <-responseChan
		sender := resp.Sender.(chat1.MsgSummary)
		if resp.Error != nil {
			_, err := k.kbChatAPI.SendReply(sender.Channel, &sender.Id, resp.Error.Error())
			if err != nil {
				k.logger.Printf("response failed: %s", err.Error())
			}

		} else {
			_, err := k.kbChatAPI.SendReply(sender.Channel, &sender.Id, resp.Text)
			if err != nil {
				k.logger.Printf("response failed: %s", err.Error())
			}
		}
	}
}
