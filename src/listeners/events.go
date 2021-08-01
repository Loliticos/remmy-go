package events

import (
  "github.com/bwmarrin/discordgo"
)

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
  if message.Author.ID == session.State.User.ID {
    return
  }

  if message.Content == ".ping" {
    _, err := session.ChannelMessageSend(message.ChannelID, "Olá")

    if err != nil {
      session.ChannelMessageSend(message.ChannelID, "Falhou")
    }
  }
}

func InteractionCreate(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
  if interaction.Type == discordgo.InteractionApplicationCommand {
    answerCommand(session, interaction)
    return
  }

  if interaction.Type == discordgo.InteractionMessageComponent {
    content := "Obrigado pelo feedback "

    switch interaction.MessageComponentData().CustomID {
      case "yes_btn":
        content += "(yes)"
      case "no_btn":
        content += "(no)"
    }

    session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
      Type: discordgo.InteractionResponseUpdateMessage,
      Data: &discordgo.InteractionResponseData{
        Content: content,
        Components: []discordgo.MessageComponent{},
      },
    })
  }
}

func answerCommand(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
  yes_button := discordgo.Button{
    Label: "Sim",
    Style: discordgo.SuccessButton,
    Disabled: false,
    CustomID: "yes_btn",
  }

  no_button := discordgo.Button{
    Label: "Não",
    Style: discordgo.DangerButton,
    Disabled: false,
    CustomID: "no_btn",
  }

  components := []discordgo.MessageComponent{ yes_button, no_button }

  err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
    Type: discordgo.InteractionResponseChannelMessageWithSource,
    Data: &discordgo.InteractionResponseData{
      Content: "Você está satisfeito com os botões?",
      Components: []discordgo.MessageComponent{
        discordgo.ActionsRow{ Components: components },
      },
    },
  })

  if err != nil {
    panic(err)
  }
}
