package client

import (
  "log"
  "fmt"
  "os"
  "os/signal"
  "github.com/bwmarrin/discordgo"
  "../listeners"
)

func Start() {
  discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))

  if err != nil {
    log.Fatalf("Invalid bot parameter: %v", err)
    return
  }

  discord.Identify.Intents = discordgo.IntentsGuildMessages

  //_, err = discord.ApplicationCommandCreate("817505090977923093", "800130595794583582",
    //&discordgo.ApplicationCommand{ Name:"feedback", Description: "Give your feedback",},
  //)

  if err != nil {
    log.Fatalf("Could not create slash: %v", err)
    return
  }

  err = discord.Open()
  if err != nil {
    log.Fatalf("error opening connection,", err)
    return
  }

  discord.AddHandler(events.MessageCreate)
  discord.AddHandler(events.InteractionCreate)

  fmt.Println("Bot is now running.")

  defer discord.Close()

  stop := make(chan os.Signal)
  signal.Notify(stop, os.Interrupt)

  <-stop

  log.Println("Gracefully shutdowning")
}
