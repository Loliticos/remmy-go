package commands

import (
  "github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
  if message.Author.ID == session.State.User.ID {
    return
  }

  if message.Content == ".ping" {
    _, err := session.ChannelMessageSend(message.ChannelID, "Pong")

    if err != nil {
      session.ChannelMessageSend(message.ChannelID, "Failed")
    }
  }
}
