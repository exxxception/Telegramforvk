package vk

import (
	"context"
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/SevereCloud/vksdk/v2/events"
	"github.com/SevereCloud/vksdk/v2/longpoll-bot"
	"log"
)

type VkBot struct {
	bot *api.VK
}

func NewVkBot(bot *api.VK) *VkBot {
	return &VkBot{
		bot: bot,
	}
}

func (vk *VkBot) Start() error {
	// Initialize Long Pool
	lp, err := vk.initLongPoll()
	if err != nil {
		return err
	}

	vk.handlerUpdates(lp)

	// Run Bots Long Poll
	log.Println("Start Long Poll")
	if err := lp.Run(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (vk *VkBot) handlerUpdates(longpool *longpoll.LongPoll) {
	// New message event
	longpool.MessageNew(func(_ context.Context, obj events.MessageNewObject) {
		log.Printf("%d: %s", obj.Message.PeerID, obj.Message.Text)

		if obj.Message.Text == "ping" {
			b := params.NewMessagesSendBuilder()
			b.Message("pong")
			b.RandomID(0)
			b.PeerID(obj.Message.PeerID)

			_, err := vk.bot.MessagesSend(b.Params)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func (vk *VkBot) initLongPoll() (*longpoll.LongPoll, error) {
	// get information about the group
	group, err := vk.bot.GroupsGetByID(nil)
	if err != nil {
		return nil, err
	}

	// Initializing Long Poll
	lp, err := longpoll.NewLongPoll(vk.bot, group[0].ID)
	if err != nil {
		return nil, err
	}

	return lp, nil
}
