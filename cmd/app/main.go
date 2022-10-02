package main

import (
	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/exxxception/pkg/vk"
)

func main() {
	// use os.Getenv("TOKEN")
	token := "TOKEN"
	bot := api.NewVK(token)

	// Start Bot VK
	vkBot := vk.NewVkBot(bot)
	vkBot.Start()
}
